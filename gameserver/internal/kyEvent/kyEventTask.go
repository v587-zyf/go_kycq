package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_task struct {
	*kyEventLog.KyEventPropBase
	Main_type  int `json:"main_type"`   //任务大类	数值	是
	Config_id1 int `json:"config_id1"` //当前任务id	数值	是
	Config_id2 int `json:"config_id2"` //下一个任务id
}

func UserTask(user *objs.User, taskType int, nowTaskId int, preTaskId int) {
	data := &KyEvent_task{
		KyEventPropBase: getEventBaseProp(user),
		Main_type:       taskType,
		Config_id1:      nowTaskId,
		Config_id2:      preTaskId,
	}
	writeUserEvent(user, "task_complete", data)
}
