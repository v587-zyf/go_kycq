package friend

import (
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"encoding/json"
	"fmt"
)

/**
 *  @Description: 记录好友数据
 *  @param user
 */
func (this *FriendManager) WriteFriendUserInfo(user *objs.User) {
	heroEquips := make(map[int]model.Equips)
	heroProp := make(map[int]map[int]int)
	userHero := user.Heros
	for heroIndex, hero := range userHero {
		heroPropInfo := hero.Prop.BuildClient()
		heroEquips[heroIndex] = hero.Equips
		if heroProp[heroIndex] == nil {
			heroProp[heroIndex] = make(map[int]int)
		}
		for pid, pVal := range heroPropInfo {
			heroProp[heroIndex][int(pid)] = int(pVal)
		}
	}
	//组织字符串
	//装备
	equipStr := "["
	if len(heroEquips) > 0 {
		for heroIndex, equips := range heroEquips {
			for pos, equip := range equips {
				randPropStr := "["
				if len(equip.RandProps) > 0 {
					for _, v := range equip.RandProps {
						randPropStr += fmt.Sprintf("[%d,%d,%d],", v.PropId, v.Color, v.Value)
					}
					randPropStr = randPropStr[:len(randPropStr)-1]
				}
				randPropStr += "]"
				equipStr += fmt.Sprintf(`[%d,%d,%d,%d,%s,%d],`, heroIndex, pos, equip.Index, equip.ItemId, randPropStr, equip.Lucky)
			}
		}
		equipStr = equipStr[:len(equipStr)-1]
	}
	equipStr += "]"
	//详细属性
	propStr := "["
	if len(heroProp) > 0 {
		for heroIndex, prop := range heroProp {
			for pid, pVal := range prop {
				propStr += fmt.Sprintf(`[%d,%d,%d],`, heroIndex, pid, pVal)
			}
		}
		propStr = propStr[:len(propStr)-1]
	}
	propStr += "]"
	valStr := fmt.Sprintf(`[%d,%d,%s,%s]`, user.Id, user.Combat, equipStr, propStr)
	rmodel.Friend.SetFriendInfo(user.Id, valStr)
}

/**
 *  @Description: 获取好友数据
 *  @param friendId
 *  @return *pb.FriendUserInfo
 */
func (this *FriendManager) GetFriendUserInfo(friendId int) *pb.FriendUserInfo {
	friendUserInfo := &pb.FriendUserInfo{
		FriendHeroInfo: make(map[int32]*pb.FriendHeroInfo),
	}
	//在线直接取数据
	if user := this.GetUserManager().GetUser(friendId); user != nil {
		userHero := user.Heros
		for heroIndex, hero := range userHero {
			friendUserInfo.FriendHeroInfo[int32(heroIndex)] = &pb.FriendHeroInfo{
				Equips: make(map[int32]*pb.EquipUnit),
				Props:  make(map[int32]int64),
			}
			friendUserInfo.FriendHeroInfo[int32(heroIndex)].Display = this.GetUserManager().GetHeroDisplay(hero)
			friendUserInfo.FriendHeroInfo[int32(heroIndex)].Job = int32(hero.Job)
			friendUserInfo.FriendHeroInfo[int32(heroIndex)].Sex = int32(hero.Sex)
			friendUserInfo.FriendHeroInfo[int32(heroIndex)].Lv = int32(hero.ExpLvl)
			friendUserInfo.FriendHeroInfo[int32(heroIndex)].Name = hero.Name
			friendUserInfo.FriendHeroInfo[int32(heroIndex)].Combat = int64(hero.Combat)
			friendUserInfo.FriendHeroInfo[int32(heroIndex)].Equips = builder.BuildEquiqs(hero)
			friendUserInfo.FriendHeroInfo[int32(heroIndex)].Props = hero.Prop.BuildClient()
		}
		friendUserInfo.UserId = int32(friendId)
		friendUserInfo.Combat = int64(user.Combat)
	} else {
		//不在线取redis数据
		user := this.GetUserManager().GetUserBasicInfo(friendId)
		if len(user.HeroDisplay) > 0 {
			for heroIndex, display := range user.HeroDisplay {
				friendUserInfo.FriendHeroInfo[int32(heroIndex)] = &pb.FriendHeroInfo{
					Equips: make(map[int32]*pb.EquipUnit),
					Props:  make(map[int32]int64),
				}
				friendUserInfo.FriendHeroInfo[int32(heroIndex)].Display = builder.BuildHeroDisplay(display.Display)
				friendUserInfo.FriendHeroInfo[int32(heroIndex)].Job = int32(display.Job)
				friendUserInfo.FriendHeroInfo[int32(heroIndex)].Sex = int32(display.Sex)
				friendUserInfo.FriendHeroInfo[int32(heroIndex)].Lv = int32(display.ExpLvl)
				friendUserInfo.FriendHeroInfo[int32(heroIndex)].Name = display.Name
				friendUserInfo.FriendHeroInfo[int32(heroIndex)].Combat = int64(display.Combat)
			}
		}
		//数据编译为数组
		bytes := rmodel.Friend.GetFriendInfo(friendId)
		data := make([]interface{}, 0)
		if err := json.Unmarshal(bytes, &data); err != nil {
			logger.Error("friend GetFriendUserInfo Unmarshal err:%v", err)
			return friendUserInfo
		}
		//装备信息
		equipInterface := data[2].([]interface{})
		for _, equip := range equipInterface {
			equipArr := equip.([]interface{})
			randPropInterface := equipArr[4].([]interface{})
			randProps := make([]*pb.EquipRandProp, 0)
			for _, randProp := range randPropInterface {
				randPropArr := randProp.([]interface{})
				randProps = append(randProps, &pb.EquipRandProp{
					PropId: int32(randPropArr[0].(float64)),
					Color:  int32(randPropArr[1].(float64)),
					Value:  int32(randPropArr[2].(float64)),
				})
			}
			heroIndex := int(equipArr[0].(float64))
			friendUserInfo.FriendHeroInfo[int32(heroIndex)].Equips[int32(equipArr[1].(float64))] = &pb.EquipUnit{
				ItemId:     int32(equipArr[3].(float64)),
				RandProps:  randProps,
				EquipIndex: int32(equipArr[2].(float64)),
				Lucky:      int32(equipArr[5].(float64)),
			}
		}
		//详细属性
		propInterface := data[3].([]interface{})
		for _, prop := range propInterface {
			propArr := prop.([]interface{})
			heroIndex := int(propArr[0].(float64))
			friendUserInfo.FriendHeroInfo[int32(heroIndex)].Props[int32(propArr[1].(float64))] = int64(propArr[2].(float64))
		}
		friendUserInfo.UserId = int32(data[0].(float64))
		friendUserInfo.Combat = int64(data[1].(float64))
	}
	return friendUserInfo
}
