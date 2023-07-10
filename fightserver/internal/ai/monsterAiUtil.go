package ai

import (
	"cqserver/fightserver/internal/actorPkg"
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/scene"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/protobuf/pb"
	"math/rand"
)

//怪物回家
func monsterHomeWalk(monster *actorPkg.MonsterActor) {

	point := monster.Point()
	tPoint := monster.BirthPoint()
	moveToPoint(monster, point, tPoint, constFight.MOVE_TYPE_HOME)
}

//怪物随机走动
func monsterRandWalk(monster base.Actor) {

	monster.SetPathSlice(nil)
	for i := 0; i < 8; i++ {

		randDir := rand.Intn(8)
		nextPoint := monster.Point().GetNewNearPointByDir(randDir)
		if nextPoint != nil && nextPoint.CanStand(monster) {
			monster.MoveTo(nextPoint, pb.MOVETYPE_RUN, false, true)
			break
		}
	}
}

func findNearestMonsterEnemy(actor base.Actor, rangeDis int) base.Actor {
	point := actor.Point()
	pointSlice := actor.GetScene().GetPointByRangeDis(point, rangeDis)

	var nearEnemy base.Actor
	for _, v := range pointSlice {

		if nearEnemy != nil && scene.DistanceByPoint(point, nearEnemy.Point()) < scene.DistanceByPoint(point, v) {
			continue
		}

		if v.GetSceneObjsNum() > 0 {
			objs := v.GetAllObject()
			for _, obj := range objs {
				if obj.IsSceneObj() {
					continue
				}
				enemyActor := obj.GetContext().(base.Actor)
				if !enemyActor.CanAttack() {
					continue
				}
				if enemyActor.GetVisible() && actor.IsEnemy(enemyActor) {
					nearEnemy = enemyActor
					break
				}
			}

		}
	}
	return nearEnemy
}

/**
 * 沿着某个方向一定(AI调用,不允许玩家调用)
 *
 * @param map
 * @param monster
 * @param tPoint
 */
func moveByDir(monster base.Actor, tPoint *scene.Point) {

	mPoint := monster.Point()
	if mPoint == nil {
		return
	}
	var nextPoint *scene.Point
	dir := scene.GetFaceDirByPoint(mPoint, tPoint) // 回家的方向

	if dir == scene.DIR_NONE {
		return
	}

	nextPoint = scene.NextDirPoint(mPoint, dir)
	if nextPoint != nil {
		if !nextPoint.CanStand(monster) {
			nextPoint = nil
		}
	}

	if nextPoint == nil { // 朝着回家的方向两侧走（左侧）
		left := scene.GetDirLeftDir(dir)
		nextPoint = scene.NextDirPoint(mPoint, left)
		if nextPoint != nil {
			if !nextPoint.CanStand(monster) {
				nextPoint = nil
			}
		}
	}
	if nextPoint == nil { // 朝着回家的方向两侧走（右侧）
		right := scene.GetDirRightDir(dir)
		nextPoint = scene.NextDirPoint(mPoint, right)
		if nextPoint != nil {
			if !nextPoint.CanStand(monster) {
				nextPoint = nil
			}
		}
	}

	if nextPoint != nil {
		monster.MoveTo(nextPoint, pb.MOVETYPE_RUN, false, true)
	} else {
		//随机走动
		monsterRandWalk(monster)
	}
}
