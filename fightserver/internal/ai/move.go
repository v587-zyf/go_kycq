package ai

import (
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/scene"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/protobuf/pb"
	"math"
)

func GetActionInterval(speed int) int64 {
	if speed == 0 {
		return math.MaxInt64
	}
	dSpeed := float64(speed)
	//没格大小scene.CellSize=50,为了保持移动的连贯性,改为47,保证每次160ms的心跳都可以移动
	return int64(1e9 * float64(47) / dSpeed)
}

/**
 *  @Description: 目标移动
 *  @param actor 角色
 *  @param point 起点
 *  @param toPoint 终点
 */
func moveToPoint(actor base.Actor, point *scene.Point, toPoint *scene.Point, moveType int) {

	//先查找玩家身上路径
	nextPoint, _ := actor.GetNextPathPoint()
	if nextPoint == nil {
		//按方向走
		nextPoint = GetMoveDirPos(point, toPoint)
	}

	if nextPoint == nil {

		pathSlice := actor.GetScene().FindRoad(point, toPoint)
		//pathSlice := bStar(actor.GetScene(), point, toPoint)
		if pathSlice != nil {
			if moveType == constFight.MOVE_TYPE_CHASE {
				pathSlice = pathSlice[1:len(pathSlice)]
			}

			actor.SetPathSlice(pathSlice)
			nextPoint, _ = actor.GetNextPathPoint()
		}

	}
	if nextPoint != nil && nextPoint.CanStand(actor) {
		actor.MoveTo(nextPoint, pb.MOVETYPE_RUN, false, true)
	} else {
		//清空寻路路径
		actor.SetPathSlice(nil)
		moveByDir(actor, toPoint)
	}
}

func bStar(sceneParam *scene.Scene, point *scene.Point, tPoint *scene.Point) []*scene.Point {
	baseRightPoint, baseLeftPoint := point, point
	Rightstep := make([]*scene.Point, 0)
	Leftstep := make([]*scene.Point, 0)
	for i := 0; i < FIND_PATH_NUM; i++ {

		baseRightPoint = RightSpread(sceneParam, Rightstep, baseRightPoint, tPoint)
		Rightstep = append(Rightstep, baseRightPoint)
		if baseRightPoint == tPoint {
			return Rightstep
		}

		baseLeftPoint = LeftSpread(sceneParam, Leftstep, baseLeftPoint, tPoint)
		Leftstep = append(Leftstep, baseLeftPoint)
		if baseLeftPoint == tPoint {
			return Leftstep
		}
	}
	if len(Rightstep) > 0 {
		return Rightstep
	}
	if len(Leftstep) > 0 {
		return Leftstep
	}
	return []*scene.Point{}
}

func GetMoveDirPos(point, tpoint *scene.Point) *scene.Point {
	if point.X() == tpoint.X() && point.Y() == tpoint.Y() {
		return nil
	}
	moveDir := scene.GetDir(point.X(), point.Y(), tpoint.X(), tpoint.Y())
	nextDir := point.GetNewNearPointByDir(moveDir)
	return nextDir
}

func RightSpread(scene1 *scene.Scene, pathSlice []*scene.Point, basePoint *scene.Point, toPoint *scene.Point) *scene.Point {
	curDir := scene.GetDir(basePoint.X(), basePoint.Y(), toPoint.X(), toPoint.Y())
loop:
	for j := 0; j < 8; j++ {
		newDir := curDir + j
		if newDir > 7 {
			newDir = newDir - 8
		}
		newPoint := basePoint.GetNewNearPointByDir(newDir)
		if newPoint == nil {
			continue
		}
		for _, v := range pathSlice {
			if v == newPoint {
				continue loop
			}
		}

		if !newPoint.IsBlock() {
			//logger.Info("dir:%v,nowDir:%v,starPoint:%v,endPoint:%v", curDir, newDir, basePoint.ToString(), toPoint.ToString())
			return newPoint
		}
	}
	return basePoint
}

func LeftSpread(scene1 *scene.Scene, pathSlice []*scene.Point, basePoint *scene.Point, toPoint *scene.Point) *scene.Point {
	curDir := scene.GetDir(basePoint.X(), basePoint.Y(), toPoint.X(), toPoint.Y())
loop:
	for j := 0; j < 8; j++ {
		newDir := curDir - j
		if newDir < 0 {
			newDir = newDir + 8
		}

		newPoint := basePoint.GetNewNearPointByDir(newDir)
		if newPoint == nil {
			continue
		}
		for _, v := range pathSlice {
			if v == newPoint {
				continue loop
			}
		}
		if !newPoint.IsBlock() {
			return newPoint
		}
	}
	return basePoint
}
