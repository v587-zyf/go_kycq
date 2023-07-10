package user

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func (this *UserManager) SendDisplay(user *objs.User) {
	pbDisplayNtfs := make(map[int32]*pb.Display)

	for _, v := range user.Heros {
		//更新武将显示信息
		this.updateHeroDisplay(user, v)
		pbDisplayNtf := this.GetHeroDisplay(v)
		pbDisplayNtfs[int32(v.Index)] = pbDisplayNtf
	}
	//推送消息
	this.GetUserManager().SendMessage(user, &pb.DisplayNtf{Display: pbDisplayNtfs}, false)
}

func (this *UserManager) updateHeroDisplay(user *objs.User, hero *objs.Hero) {
	oldDisplay := hero.Display
	hero.Display = &model.Display{
		ClothType:  pb.DISPLAYTYPE_EQUIP,
		WeaponType: pb.DISPLAYTYPE_EQUIP,
	}

	//衣服
	clothItemId := hero.Equips[pb.EQUIPTYPE_CLOTHES].ItemId
	hero.Display.ClothItemId = clothItemId

	//武器
	weaponItemId := hero.Equips[pb.EQUIPTYPE_WEAPON_R].ItemId
	hero.Display.WeaponItemId = weaponItemId

	//判断时装系统
	if hero.Wear.FashionWeaponId > 0 {
		fashion := hero.Fashions[hero.Wear.FashionWeaponId]
		fashionConf := gamedb.GetFashionConf(fashion.Id, fashion.Lv)
		hero.Display.WeaponType = pb.DISPLAYTYPE_FASHION
		hero.Display.WeaponItemId = fashionConf.Id
	}

	if hero.Wear.FashionClothId > 0 {
		fashion := hero.Fashions[hero.Wear.FashionClothId]
		fashionConf := gamedb.GetFashionConf(fashion.Id, fashion.Lv)
		hero.Display.ClothType = pb.DISPLAYTYPE_FASHION
		hero.Display.ClothItemId = fashionConf.Id
	}
	//羽翼
	hero.Display.WingId = hero.Wear.WingId

	//法阵
	hero.Display.MagicCircleLvId = hero.Wear.MagicCircleLvId

	//称号
	hero.Display.TitleId = hero.Wear.TitleId

	//头衔
	if hero.Index == constUser.USER_HERO_MAIN_INDEX {
		hero.Display.LabelId = user.Label.Id
		hero.Display.LabelJob = user.Label.Job
	}

	if hero.Display.ClothItemId != oldDisplay.ClothItemId ||
		hero.Display.WeaponItemId != oldDisplay.WeaponItemId ||
		hero.Display.WingId != oldDisplay.WingId ||
		hero.Display.MagicCircleLvId != oldDisplay.MagicCircleLvId ||
		hero.Display.TitleId != oldDisplay.TitleId ||
		hero.Display.LabelId != oldDisplay.LabelId ||
		hero.Display.LabelJob != oldDisplay.LabelJob {
		//标记通知战斗服 更新玩家数据
		user.UpdateFightUserHeroIndexFun(hero.Index)
	}

}

func (this *UserManager) GetHeroDisplay(hero *objs.Hero) *pb.Display {

	pbDisplayNtf := &pb.Display{
		ClothType:       int32(hero.Display.ClothType),
		ClothItemId:     int32(hero.Display.ClothItemId),
		WeaponType:      int32(hero.Display.WeaponType),
		WeaponItemId:    int32(hero.Display.WeaponItemId),
		WingId:          int32(hero.Display.WingId),
		MagicCircleLvId: int32(hero.Display.MagicCircleLvId),
		TitleId:         int32(hero.Display.TitleId),
		LabelId:         int32(hero.Display.LabelId),
		LabelJob:        int32(hero.Display.LabelJob),
	}
	return pbDisplayNtf
}
