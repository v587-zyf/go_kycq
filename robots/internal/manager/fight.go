package manager

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"
)

var (
	fightType = flag.Int("fightType", 8, "specify config file")
)

func init() {
	pb.Register(pb.CmdSceneEnterNtfId, func(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

		msg := p.(*pb.SceneEnterNtf)
		for _, v := range msg.Objs {

			sceneObjs[v.ObjId] = v
			if v.ObjType != pb.SCENEOBJTYPE_USER {
				continue
			}
			if v.User.UserId == m.robot.user.Userid {

				fmt.Println("接收到玩家进入场景消息：stageId:", msg.StageId,m.robot.openId)
				userObjs[int(v.User.HeroIndex)] = &RobotSceneInfo{
					sceneObj:   v,
					birthPoint: v.Point,
				}

				m.robot.inFightSceneStatus = 2
			}
		}
		return nil, nil, nil
	})
	pb.Register(pb.CmdSceneMoveNtfId, func(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
		msg := p.(*pb.SceneMoveNtf)
		for _, v := range userObjs {
			if msg.ObjId == v.sceneObj.ObjId {
				v.moveTime = time.Now().Unix()
				v.sceneObj.Point = msg.Point
			}
		}
		for _, v := range sceneObjs {
			if v.ObjId == msg.ObjId {
				v.Point = msg.Point
			}
		}
		return nil, nil, nil
	})

	pb.Register(pb.CmdAttackEffectNtfId, func(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

		msg := p.(*pb.AttackEffectNtf)
		for _, v := range userObjs {
			if v.sceneObj.ObjId == msg.AttackerId {

				for _, heroBaseInfo := range m.robot.user.Heros {
					if heroBaseInfo.Index == v.sceneObj.User.HeroIndex {
						for _, skill := range heroBaseInfo.Skills {
							if skill.SkillId == msg.SkillId {
								skill.EndTime = msg.SkillStopT
								v.userSkillTime = time.Now()
							}
						}
					}
				}
			}
		}
		for _, v := range msg.Hurts {
			if v.IsDeath {
				delete(sceneObjs, v.ObjId)
			}
			for _, userHero := range userObjs {
				if v.ObjId == userHero.sceneObj.ObjId {
					userHero.sceneObj.Hp = v.Hp
				}
			}
		}

		allDie := true
		for _, v := range userObjs {
			if v.sceneObj.Hp > 0 {
				allDie = false
				break
			}
		}
		if allDie {
			logger.Info("准备复活")
			m.robot.initDataPb = append(m.robot.initDataPb, &pb.FightUserReliveReq{
				SafeRelive: true,
			})
		}

		return nil, nil, nil
	})

	pb.Register(pb.CmdSceneLeaveNtfId, func(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
		msg := p.(*pb.SceneLeaveNtf)
		for _, v := range msg.ObjIds {
			delete(sceneObjs, v)
		}
		return nil, nil, nil
	})

	pb.Register(pb.CmdSceneUserReliveNtfId, func(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
		msg := p.(*pb.SceneUserReliveNtf)
		sceneObjs[msg.Obj.ObjId] = msg.Obj
		return nil, nil, nil
	})
}

type RobotSceneInfo struct {
	sceneObj      *pb.SceneObj
	moveTime      int64 //移动时间
	birthPoint    *pb.Point
	userSkillTime time.Time
}

var sceneObjs map[int32]*pb.SceneObj
var userObjs map[int]*RobotSceneInfo

func (this *Robot) inFightSceneTest() {

	var msg nw.ProtoMessage
	stageId := 0
	if *fightType == constFight.FIGHT_TYPE_MAIN_CITY {
		msg = &pb.EnterPublicCopyReq{
			StageId:   int32(constFight.FIGHT_TYPE_MAIN_CITY_STAGE),
			Condition: 1,
		}
		stageId = constFight.FIGHT_TYPE_MAIN_CITY_STAGE
	} else if *fightType == constFight.FIGHT_TYPE_SHABAKE {
		msg = &pb.EnterShaBaKeFightReq{}
		stageId = constFight.FIGHT_TYPE_SHABAKE_STAGE
		//未加入公会
		if this.guildStatue != 2 {
			return
		}
		now := common.GetTimeSeconds(time.Now())
		if now < gamedb.GetConf().ShabakeTime3[0].GetSecondsFromZero() || now > gamedb.GetConf().ShabakeTime3[1].GetSecondsFromZero() {
			logger.Debug("沙巴克还未开放")
			return
		}
	} else if *fightType == constFight.FIGHT_TYPE_STAGE {
		stageId = int(m.robot.user.StageId)
		msg = &pb.StageFightStartReq{}
	}
	if msg == nil {
		return
	}
	stageConf := gamedb.GetStageStageCfg(stageId)
	mapT := gamedb.GetMapMapCfg(stageConf.Mapid)
	this.sceneT = gamedb.GetDb().GetSceneConf(mapT.Resource)

	sceneObjs = make(map[int32]*pb.SceneObj)
	userObjs = make(map[int]*RobotSceneInfo)
	this.inFightSceneStatus = 1
	this.SendMessage(0, msg)
}

func (this *Robot) sceneAction() {

	for _, v := range this.user.Heros {

		heroInfo := userObjs[int(v.Index)]
		if heroInfo == nil {
			continue
		}
		if heroInfo.sceneObj.Hp <= 0 {
			continue
		}
		if this.UseSkill(heroInfo) {
			continue
		}

		if *fightType == constFight.FIGHT_TYPE_MAIN_CITY && v.Index != constUser.USER_HERO_MAIN_INDEX {
			continue
		}
		this.randMove(heroInfo)
	}
}

func (this *Robot) randMove(userHero *RobotSceneInfo) {
	now := time.Now()
	if userHero.moveTime > 0 && (now.Unix()-userHero.moveTime) < 1 {
		return
	}
	for i := 0; i < 50; i++ {
		dir := rand.Intn(8)
		if Distance(int(userHero.sceneObj.Point.X), int(userHero.sceneObj.Point.Y), int(userHero.birthPoint.X), int(userHero.birthPoint.Y)) > 5 {
			dir = GetDir(int(userHero.sceneObj.Point.X), int(userHero.sceneObj.Point.Y), int(userHero.birthPoint.X), int(userHero.birthPoint.Y))
		}
		offsetX, offsetY := GetDirOffset(dir)
		if this.sceneT.Walkable(userHero.sceneObj.Point.X+int32(offsetX), userHero.sceneObj.Point.Y+int32(offsetY)) {
			this.SendMessage(0, &pb.SceneMoveRpt{
				ObjId: userHero.sceneObj.ObjId,
				Point: &pb.Point{
					X: userHero.sceneObj.Point.X + int32(offsetX),
					Y: userHero.sceneObj.Point.Y + int32(offsetY),
				},
			})
			logger.Debug("玩家发送走动：%v,heroIndex:%v", this.openId, userHero.sceneObj.User.HeroIndex)
			userHero.moveTime = time.Now().Unix()
			return
		}
	}
}

func (this *Robot) UseSkill(heroInfo *RobotSceneInfo) bool {

	if *fightType != constFight.FIGHT_TYPE_SHABAKE {
		return false
	}

	//找最近的敌人
	armyObjId, armyDis := findNearestArmy(heroInfo)
	if armyObjId < 0 {
		return false
	}
	//找能攻击的技能
	skillId := this.fineSkill(heroInfo, armyDis)
	//释放技能
	if skillId > 0 {
		army := sceneObjs[armyObjId]
		dir := GetDir(int(heroInfo.sceneObj.Point.X), int(heroInfo.sceneObj.Point.Y), int(army.Point.X), int(army.Point.Y))
		this.SendMessage(0, &pb.AttackRpt{
			SkillId: int32(skillId),
			ObjIds:  []int32{armyObjId},
			Dir:     int32(dir),
			Point:   heroInfo.sceneObj.Point,
			ObjId:   heroInfo.sceneObj.ObjId,
		})
		return true
	}
	return false
}

func (this *Robot) fineSkill(hero *RobotSceneInfo, dis int) int {

	var heroBaseInfo *pb.HeroInfo
	for _, v := range this.user.Heros {
		if v.Index == hero.sceneObj.User.HeroIndex {
			heroBaseInfo = v
		}
	}
	if heroBaseInfo == nil {
		return 0
	}

	attackSpeed := heroBaseInfo.GetHeroProp().Props[pb.PROPERTY_ATT_SPEED]
	attackInterval := int64(gamedb.GetAttIntervalByAttSpeed(int(attackSpeed))) - 20
	attackIntervalIn := !hero.userSkillTime.IsZero() && time.Now().Sub(hero.userSkillTime).Milliseconds() < attackInterval

	now := common.GetNowMillisecond()
	for _, v := range heroBaseInfo.SkillBag {
		if v == 0 {
			continue
		}
		var skill *pb.SkillUnit
		for _, skillTemp := range heroBaseInfo.Skills {
			if skillTemp.SkillId == v {
				skill = skillTemp
				break
			}
		}
		if skill.EndTime > now {
			continue
		}
		skillConf := gamedb.GetSkillSkillCfg(int(skill.SkillId))

		if skillConf.Aspd && attackIntervalIn {
			continue
		}

		skillLvConf := gamedb.GetSkillLvConf(int(skill.SkillId), int(skill.Level))
		if skillLvConf == nil {
			continue
		}

		if skillLvConf.Distance > dis {
			skill.EndTime = now + 1000
			hero.userSkillTime = time.Now()
			return int(skill.SkillId)
		}
	}
	return 0
}

func findNearestArmy(hero *RobotSceneInfo) (int32, int) {

	minDis := math.MaxInt32
	var minDisObjId int32 = 0
	for _, v := range sceneObjs {
		if v.ObjType != pb.SCENEOBJTYPE_USER {
			continue
		}

		if v.User.GuildId == hero.sceneObj.User.GuildId {
			continue
		}

		dis := Distance(int(v.Point.X), int(v.Point.Y), int(hero.sceneObj.Point.X), int(hero.sceneObj.Point.Y))
		if dis < minDis {
			minDis = dis
			minDisObjId = v.ObjId
		}
	}
	return minDisObjId, minDis
}

const (
	DIR_NONE         = -1
	DIR_BOTTOM       = pb.SCENEDIR_BOTTOM
	DIR_RIGHT_BOTTOM = pb.SCENEDIR_RIGHT_BOTTOM
	DIR_RIGHT        = pb.SCENEDIR_RIGHT
	DIR_RIGHT_TOP    = pb.SCENEDIR_RIGHT_TOP
	DIR_TOP          = pb.SCENEDIR_TOP
	DIR_LEFT_TOP     = pb.SCENEDIR_LEFT_TOP
	DIR_LEFT         = pb.SCENEDIR_LEFT
	DIR_LEFT_BOTTOM  = pb.SCENEDIR_LEFT_BOTTOM
)

func GetDirOffset(dir int) (int, int) {
	switch dir {
	case DIR_TOP:
		return 0, -1
	case DIR_RIGHT_TOP:
		return 1, -1
	case DIR_RIGHT:
		return 1, 0
	case DIR_RIGHT_BOTTOM:
		return 1, 1
	case DIR_BOTTOM:
		return 0, 1
	case DIR_LEFT_BOTTOM:
		return -1, 1
	case DIR_LEFT:
		return -1, 0
	case DIR_LEFT_TOP:
		return -1, -1
	}
	return 0, 0
}

func Distance(x1, y1, x2, y2 int) int {
	return int(math.Max(math.Abs(float64(x1-x2)), math.Abs(float64(y1-y2))))
}

func GetDir(fx, fy, tx, ty int) int {

	todir := 0
	angle := getangle(fx, fy, tx, ty, false)
	if angle < 0 {
		angle = 360 + angle
	}

	if angle > 335 || angle < 25 {
		todir = pb.SCENEDIR_RIGHT

	} else if angle > 290 {
		todir = pb.SCENEDIR_RIGHT_TOP
	} else if angle > 245 {
		todir = pb.SCENEDIR_TOP
	} else if angle > 200 {
		todir = pb.SCENEDIR_LEFT_TOP
	} else if angle > 155 {
		todir = pb.SCENEDIR_LEFT
	} else if angle > 110 {
		todir = pb.SCENEDIR_LEFT_BOTTOM
	} else if angle > 65 {
		todir = pb.SCENEDIR_BOTTOM
	} else {
		todir = pb.SCENEDIR_RIGHT_BOTTOM
	}
	return todir
}

func getangle(startx int, starty int, endx int, endy int, isradian bool) int {

	disX := endx - startx
	disY := endy - starty
	angle := math.Atan2(float64(disY), float64(disX))
	if !isradian {
		angle = angle * 180 / math.Pi
	}
	return int(angle)
}
