package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gameserver/internal/objs"
	"time"
)

type KyEvent_startFight struct {
	*kyEventLog.KyEventPropBase
	StageId   int `json:"config_id1"`
	StageType int `json:"type"`
}

type KyEvent_leaveFight struct {
	*kyEventLog.KyEventPropBase
	StageId       int `json:"config_id1"`
	StageType     int `json:"type"`
	Page_staytime int `json:"page_staytime"` //挑战时长(秒)	数值	否
}

type KyEvent_endFight struct {
	*kyEventLog.KyEventPropBase
	Result        int         `json:"status"`          //挑战结果	数值	是	1：成功 2：失败
	StageId       int         `json:"config_id1"`    //关卡id	数值	是
	StageType     int         `json:"type"`   //关卡id	数值	是
	Page_staytime int         `json:"page_staytime"` //挑战时长(秒)	数值	否
	Rewards       map[int]int `json:"info"`       //奖励道具
}

type KyEvent_miningFightStart struct {
	*kyEventLog.KyEventPropBase
	RobType         int    `json:"type"` //掠夺（0）夺回 1
	FightRoleOpenId string `json:"info"`
	FightRoleId     int    `json:"log_event_parm1"`      //掠夺/夺回玩家Id
	FightRoleName   string `json:"log_event_parm2"`   //掠夺/夺回玩家Id
	FightRoleCombat int    `json:"num1"` //掠夺/夺回玩家Id
}

type KyEvent_miningFightEnd struct {
	*kyEventLog.KyEventPropBase
	RobType         int    `json:"rob_type"` //掠夺（0）夺回 1
	Result          int    `json:"type"`     //战斗结果
	FightRoleOpenId string `json:"info"`
	FightRoleId     int    `json:"log_event_parm1"`      //掠夺/夺回玩家Id
	FightRoleName   string `json:"log_event_parm2"`   //掠夺/夺回玩家Id
	FightRoleCombat int    `json:"num1"` //掠夺/夺回玩家Id
	Miner           int    `json:"miner"`             //矿工等级
	RobTimes        int    `json:"rob_times"`         //掠夺次数
}

type KyEvent_areanFightStart struct {
	*kyEventLog.KyEventPropBase
	FightRoleOpenId string `json:"info"`
	FightRoleId     int    `json:"log_event_parm1"`      //掠夺/夺回玩家Id
	FightRoleName   string `json:"log_event_parm2"`   //掠夺/夺回玩家Id
	FightRoleCombat int    `json:"num1"` //掠夺/夺回玩家Id
	FightRoleRank   int    `json:"num2"`
}

type KyEvent_areanFightEnd struct {
	*kyEventLog.KyEventPropBase
	Result int `json:"status"` //战斗结果
	//FightRoleOpenId string `json:"fight_role_open_id"`
	//FightRoleId     int    `json:"role_role_id"`      //掠夺/夺回玩家Id
	//FightRoleName   string `json:"fight_role_name"`   //掠夺/夺回玩家Id
	//FightRoleCombat int    `json:"fight_role_combat"` //掠夺/夺回玩家Id
	Rank  int `json:"num1"`  //新段位
	Score int `json:"num2"` //积分
}

type KyEvent_fieldFightStart struct {
	*kyEventLog.KyEventPropBase
	FightRoleOpenId string `json:"info"`
	FightRoleId     int    `json:"log_event_parm1"`      //掠夺/夺回玩家Id
	FightRoleName   string `json:"log_event_parm2"`   //掠夺/夺回玩家Id
	FightRoleCombat int    `json:"num1"` //掠夺/夺回玩家Id
}

type KyEvent_fieldFightEnd struct {
	*kyEventLog.KyEventPropBase
	FightRoleOpenId string `json:"info"`
	Result          int    `json:"status"`              //战斗结果
	FightRoleId     int    `json:"log_event_parm1"`      //掠夺/夺回玩家Id
	FightRoleName   string `json:"log_event_parm2"`   //掠夺/夺回玩家Id
	FightRoleCombat int    `json:"num1"` //掠夺/夺回玩家Id
}

type KyEvent_startStage struct {
	*kyEventLog.KyEventPropBase
	StageId int `json:"config_id1"`
}

type KyEvent_endStage struct {
	*kyEventLog.KyEventPropBase
	Result        int `json:"type"`          //挑战结果	数值	是	1：成功 2：失败
	StageId       int `json:"config_id1"`    //关卡id	数值	是
	Page_staytime int `json:"page_staytime"` //挑战时长(秒)	数值	否
}

type KyEvent_startTower struct {
	*kyEventLog.KyEventPropBase
	StageId int `json:"config_id1"`
}

type KyEvent_endTower struct {
	*kyEventLog.KyEventPropBase
	Result        int `json:"type"`          //挑战结果	数值	是	1：成功 2：失败
	StageId       int `json:"config_id1"`    //关卡id	数值	是
	Page_staytime int `json:"page_staytime"` //挑战时长(秒)	数值	否
}

func StageStart(user *objs.User, stageId int) {

	stageConf := gamedb.GetStageStageCfg(stageId)
	if stageConf == nil {
		return
	}
	if stageConf.Type == constFight.FIGHT_TYPE_ARENA || stageConf.Type == constFight.FIGHT_TYPE_FIELD || stageConf.Type == constFight.FIGHT_TYPE_MINING {
		return
	}
	fightStart(user, stageConf)
}

func StageEnd(user *objs.User, stageId int, result int, startTime time.Time, reward map[int]int) {
	stageConf := gamedb.GetStageStageCfg(stageId)
	if stageConf == nil {
		return
	}
	fightEnd(user, stageConf, result, startTime, reward)
}

func stageBossStart(user *objs.User, stageId int) {

	data := &KyEvent_startStage{
		KyEventPropBase: getEventBaseProp(user),
		StageId:         stageId,
	}
	writeUserEvent(user, "start_stage", data)
}

func stageBossEnd(user *objs.User, stageConf *gamedb.StageStageCfg, result int, startTime time.Time) {
	data := &KyEvent_endStage{
		KyEventPropBase: getEventBaseProp(user),
		Result:          result,
		StageId:         stageConf.Id,
	}
	if !startTime.IsZero() {
		data.Page_staytime = int(time.Now().Sub(startTime).Seconds())
	}
	writeUserEvent(user, "end_stage", data)
}

func stageTowerStart(user *objs.User, stageId int) {

	data := &KyEvent_startTower{
		KyEventPropBase: getEventBaseProp(user),
		StageId:         stageId,
	}
	writeUserEvent(user, "start_tower", data)
}

func stageTowerEnd(user *objs.User, stageConf *gamedb.StageStageCfg, result int, startTime time.Time) {
	data := &KyEvent_endTower{
		KyEventPropBase: getEventBaseProp(user),
		Result:          result,
		StageId:         stageConf.Id,
	}
	if !startTime.IsZero() {
		data.Page_staytime = int(time.Now().Sub(startTime).Seconds())
	}
	writeUserEvent(user, "end_tower", data)
}

func fightStart(user *objs.User, stageConf *gamedb.StageStageCfg) {

	data := &KyEvent_startFight{
		KyEventPropBase: getEventBaseProp(user),
		StageId:         stageConf.Id,
		StageType:       stageConf.Type,
	}
	writeUserEvent(user, "start_fight", data)
}

func FightLeave(user *objs.User, stageId int, startTime time.Time, ) {
	stageConf := gamedb.GetStageStageCfg(stageId)
	if stageConf == nil {
		return
	}
	if stageConf.Type != constFight.FIGHT_TYPE_PAODIAN && stageConf.Type != constFight.FIGHT_TYPE_GUILD_BONFIRE {
		return
	}

	data := &KyEvent_leaveFight{
		KyEventPropBase: getEventBaseProp(user),
		StageId:         stageConf.Id,
		StageType:       stageConf.Type,
	}
	if !startTime.IsZero() {
		data.Page_staytime = int(time.Now().Sub(startTime).Seconds())
	}
	writeUserEvent(user, "leave_fight", data)
}

func fightEnd(user *objs.User, stageConf *gamedb.StageStageCfg, result int, startTime time.Time, rewards map[int]int) {

	data := &KyEvent_endFight{
		KyEventPropBase: getEventBaseProp(user),
		Result:          result,
		StageId:         stageConf.Id,
		StageType:       stageConf.Type,
	}
	if !startTime.IsZero() {
		data.Page_staytime = int(time.Now().Sub(startTime).Seconds())
	}
	if rewards != nil {
		data.Rewards = rewards
	}
	writeUserEvent(user, "end_fight", data)
}

func MiningFightStart(user *objs.User, isRob bool, roleId int, roleName string, roleCombat int) {
	data := &KyEvent_miningFightStart{
		KyEventPropBase: getEventBaseProp(user),
		RobType:         1, //掠夺（0）夺回 1
		FightRoleId:     roleId,
		FightRoleName:   roleName,
		FightRoleCombat: roleCombat,
	}
	if isRob {
		data.RobType = 0
	}
	writeUserEvent(user, "start_mining", data)
}

func MiningFightEnd(user *objs.User, isRob bool, result int, roleId int, roleName string, roleCombat int, miner int, rob_times int) {

	data := &KyEvent_miningFightEnd{
		KyEventPropBase: getEventBaseProp(user),
		RobType:         1, //掠夺（0）夺回 1
		FightRoleId:     roleId,
		FightRoleName:   roleName,
		FightRoleCombat: roleCombat,
		Result:          result,
		Miner:           miner,
		RobTimes:        rob_times,
	}
	if isRob {
		data.RobType = 0
	}
	writeUserEvent(user, "end_mining", data)
}

func ArenaFightStart(user *objs.User, openID string, roleId int, roleName string, roleCombat int) {
	data := &KyEvent_areanFightStart{
		KyEventPropBase: getEventBaseProp(user),
		FightRoleOpenId: openID,
		FightRoleId:     roleId,
		FightRoleName:   roleName,
		FightRoleCombat: roleCombat,
	}
	writeUserEvent(user, "start_arena", data)
}

func ArenaFightEnd(user *objs.User, result int, rank int, score int) {

	data := &KyEvent_areanFightEnd{
		KyEventPropBase: getEventBaseProp(user),
		//FightRoleOpenId: openId,
		//FightRoleId:     roleId,
		//FightRoleName:   roleName,
		//FightRoleCombat: roleCombat,
		Result: result,
		Rank:   rank,
		Score:  score,
	}
	writeUserEvent(user, "end_arena", data)
}

func FieldFightStart(user *objs.User, openId string, roleId int, roleName string, roleCombat int) {
	data := &KyEvent_fieldFightStart{
		KyEventPropBase: getEventBaseProp(user),
		FightRoleOpenId: openId,
		FightRoleId:     roleId,
		FightRoleName:   roleName,
		FightRoleCombat: roleCombat,
	}
	writeUserEvent(user, "start_field", data)
}

func FieldFightEnd(user *objs.User, result int, openId string, roleId int, roleName string, roleCombat int) {

	data := &KyEvent_fieldFightEnd{
		KyEventPropBase: getEventBaseProp(user),
		FightRoleOpenId: openId,
		FightRoleId:     roleId,
		FightRoleName:   roleName,
		FightRoleCombat: roleCombat,
		Result:          result,
	}
	writeUserEvent(user, "end_field", data)
}
