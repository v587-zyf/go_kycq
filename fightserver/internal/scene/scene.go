package scene

import (
	"cqserver/gamelibs/errex"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/protobuf/pb"
	"fmt"
	"math"
	"math/rand"
	"runtime/debug"
	"sync"
	"time"

	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
)

type Scene struct {
	sceneT         *gamedb.SceneConf
	mapT           *gamedb.MapMapCfg
	notifier       Notifier
	objs           map[int]ISceneObj
	playerObjs     map[int]ISceneObj
	itemObjs       map[int]*SceneItem
	buffObjs       map[int]*SceneBuff
	collectionObjs map[int]*SceneCollection
	points         map[int]*Point
	objsMu         sync.RWMutex
	astarManager   *AStarManager
	safeArea       map[int]map[int]bool //安全区

	shieldMove    bool
	posChangeChan chan *posChangeItem
	notifyChan    chan *notifyItem
	done          chan struct{}
}

func NewScene(stageId int) (*Scene, error) {

	mapId := gamedb.GetStageStageCfg(stageId).Mapid
	mapT := gamedb.GetMapMapCfg(mapId)
	if mapT == nil {
		return nil, errex.Create(-1, fmt.Sprintf("get map conf err:%v", mapId))
	}
	sceneT := gamedb.GetDb().GetSceneConf(mapT.Resource)
	if sceneT == nil {
		return nil, errex.Create(-1, fmt.Sprintf("get map json conf err:%v", mapT.Resource))
	}
	notifier := NewNotifier(stageId, sceneT)
	if notifier == nil {
		return nil, errex.Create(-1, fmt.Sprintf("unknown nofifierType:%v", mapId))
	}
	scene := &Scene{
		sceneT:         sceneT,
		mapT:           mapT,
		notifier:       notifier,
		objs:           make(map[int]ISceneObj),
		playerObjs:     make(map[int]ISceneObj),
		itemObjs:       make(map[int]*SceneItem),
		buffObjs:       make(map[int]*SceneBuff),
		collectionObjs: make(map[int]*SceneCollection),
		posChangeChan:  make(chan *posChangeItem, 100),
		notifyChan:     make(chan *notifyItem, 100),
		done:           make(chan struct{}),
		shieldMove:     false,
	}
	scene.initGrids()
	scene.initSafeArea()
	scene.astarManager = NewAStarManager(scene)

	//if sceneT.Id == 1004 {
	//	startPoint := scene.points[gamedb.GetPointIndex(66, 142)]
	//	endPoint := scene.points[gamedb.GetPointIndex(61, 139)]
	//	now := time.Now()
	//
	//	scene.astarManager.FindAStarRoad(startPoint, endPoint)
	//	logger.Info("cost time",time.Now().Sub(now).Milliseconds())
	//}

	go scene.worker()
	return scene, nil
}

func (this *Scene) FindRoad(start, end *Point) []*Point {
	return this.astarManager.FindAStarRoad(start, end)
}

func (this *Scene) initGrids() {
	this.points = make(map[int]*Point, len(this.sceneT.WalkableMap))
	for k, v := range this.sceneT.WalkableMap {
		this.points[int(k)] = NewGird(int(k), int(v.X), int(v.Y), false)
	}

	// 构建八叉树
	for x := 0; x < this.sceneT.ColNum; x++ {
		for y := 0; y < this.sceneT.RowNum; y++ {
			index := gamedb.GetPointIndex(int32(x), int32(y))
			p := this.points[index]
			if p == nil {
				//添加一个不可走空格
				p := NewGird(index, x, y, true)
				this.points[index] = p
			}
			if p == nil {
				continue
			}
			// 遍历该格子周围的格子，进行周围是否能走的判断
			nearGirds := make([]*Point, 8)
			for index := 0; index < 16; index += 2 {
				tx := x + EIGHT_DIR_OFFSET[index]
				ty := y + EIGHT_DIR_OFFSET[index+1]
				if tx < 0 || ty < 0 || tx >= this.sceneT.ColNum || ty >= this.sceneT.RowNum {
					continue
				}
				tindex := gamedb.GetPointIndex(int32(tx), int32(ty))
				tp := this.points[tindex]
				if tp == nil {
					//添加一个不可走格子
					tp = NewGird(tindex, tx, ty, true)
					this.points[tindex] = tp
				}
				dir := GetDir(x, y, tx, ty)
				nearGirds[dir] = tp
			}
			p.setNears(nearGirds)
		}
	}
}

/**
 *  @Description: 初始化安全区坐标属性
 */
func (this *Scene) initSafeArea() {

	for _, v := range this.mapT.SafeRect {

		safeArea := this.sceneT.Rect[v]
		for pointIndex, _ := range safeArea {

			if point, ok := this.points[pointIndex]; ok {
				point.SetIsSafe(true)
			}
		}
	}
}

func (this *Scene) RandomDelivery(obj ISceneObj, safeArea bool) *Point {

	if safeArea {

		areaIndex := rand.Intn(len(this.mapT.SafeRect))
		point, err := this.randPointByRect(areaIndex)
		if err != nil {
			return nil
		}
		return point
	} else {

		l := len(this.sceneT.WalkableMap)
		for i := 0; i < 5; i++ {
			randIndex := rand.Intn(l)
			count := 0
			for k, _ := range this.sceneT.WalkableMap {

				if count == randIndex {
					point := this.points[int(k)]
					if point != nil {
						if point.CanStand(obj) {
							return point
						} else {
							nears := point.GetNearPoints()
							for _, near := range nears {
								if near.CanStand(obj) {
									return near
								}
							}
						}
					}
				}
				count++
			}
		}
	}
	return nil
}

func (this *Scene) GetPointByXY(x, y int) *Point {

	index := gamedb.GetPointIndex(int32(x), int32(y))
	return this.points[index]
}

func (this *Scene) GetUserBirthPoint(index int) (*Point, error) {
	if index == -1 {
		index = rand.Intn(len(this.mapT.Born))
	}
	rectIndex := this.mapT.Born[index]
	return this.randPointByRect(rectIndex)
}

func (this *Scene) GetBirthPointByAreaIndex(rectIndex int) (*Point, error) {
	return this.randPointByRect(rectIndex)
}

func (this *Scene) GetHeroBirthPoint(obj ISceneObj, point *Point, dir int, heroIndex int, birthType int) *Point {

	left, back := constFight.HERO_BIRTH[heroIndex]["left"], constFight.HERO_BIRTH[heroIndex]["back"]

	var newPoint *Point
	offsetX, offsetY := 0, 0
	if birthType == constFight.FIGHT_BIRTH_TYPE_TRIANGLE {
		if len(gamedb.GetConf().HeroBirthOffset) == 2 {
			if heroIndex == constUser.USER_HERO_SECOND_INDEX {
				left, back = gamedb.GetConf().HeroBirthOffset[0][0], gamedb.GetConf().HeroBirthOffset[0][1]

			} else if heroIndex == constUser.USER_HERO_THREE_INDEX {
				left, back = gamedb.GetConf().HeroBirthOffset[1][0], gamedb.GetConf().HeroBirthOffset[1][1]
			}
		}
		offsetX, offsetY = GetLeftBottomOffset(dir, left, back)

	} else {

		if heroIndex == constUser.USER_HERO_SECOND_INDEX {
			leftDir := GetDirLeftDir(dir)
			offsetX, offsetY = GetDirOffset(leftDir)
		} else {
			leftDir := GetDirRightDir(dir)
			offsetX, offsetY = GetDirOffset(leftDir)
		}
	}
	newPointX, newPonitY := point.X()+offsetX, point.Y()+offsetY
	newPoint = this.GetPointByXY(newPointX, newPonitY)

	if newPoint != nil && newPoint.CanStand(obj) {
		return newPoint
	} else {
		return this.GetPointByPointRange(obj, newPointX, newPonitY, point)
	}
}

func (this *Scene) GetPointByPointRange(obj ISceneObj, x, y int, withoutPoint *Point) *Point {
	for i := 1; i <= 10; i++ {
		for j := -1 * i; j <= i; j++ {
			for n := -1 * i; n <= i; n++ {
				if int(math.Abs(float64(j))) == i || int(math.Abs(float64(n))) == i {
					newPoint := this.GetPointByXY(x+j, y+n)
					if newPoint != nil && newPoint.CanStand(obj) {
						if withoutPoint == nil || !newPoint.Equal(withoutPoint) {
							return newPoint
						}
					}
				}
			}
		}
	}
	return nil
}

func (this *Scene) randPointByRect(rectIndex int) (*Point, error) {
	rect := this.sceneT.Rect[rectIndex]
	rectSlice := make([]int, len(rect))
	i := 0
	for k, _ := range rect {
		rectSlice[i] = k
		i++
	}

	//查找出生点怪物最少的一个格子
	var minPoint *Point
	hasSel := make(map[int]int)
	pointsNum := len(rect)
	for i := 0; i < pointsNum*3; i++ {
		selIndex := rand.Intn(pointsNum)
		if _, ok := hasSel[selIndex]; ok {
			continue
		}
		hasSel[selIndex] = selIndex
		pIndex := rectSlice[selIndex]
		if this.points[pIndex] == nil {
			continue
		}
		if this.points[pIndex].IsBlock() {
			continue
		}
		if minPoint == nil || minPoint.GetSceneObjsNum() > this.points[pIndex].GetSceneObjsNum() {
			minPoint = this.points[pIndex]
		}
		if minPoint.GetSceneObjsNum() == 0 {
			return minPoint, nil
		}
	}
	if minPoint == nil {
		logger.Error("随机区域内坐标异常,地图：%v 区域索引：%v,坐标:%v", this.mapT.Id, rectIndex, rect)
	}
	return minPoint, nil
}

func (this *Scene) GetSceneRectObjs(rectIndex int) []int {

	allObjs := make([]int, 0)
	rect := this.sceneT.Rect[rectIndex]
	for k, _ := range rect {
		p := this.points[k]
		objs := p.GetAllObject()
		for _, v := range objs {
			if v.GetType() == pb.SCENEOBJTYPE_USER || v.GetType() == pb.SCENEOBJTYPE_FIT {
				allObjs = append(allObjs, v.GetObjId())
			}
		}
	}
	return allObjs
}

func (this *Scene) SetSceneBlockByRectIndex(rectIndex int, isBlock bool) {

	rect := this.sceneT.Rect[rectIndex]
	for k, _ := range rect {
		this.points[k].SetIsBlock(isBlock)
	}
}

func (this *Scene) GetPointByDirAndMaxDis(point *Point, dir, maxDis int, isStand bool, obj ISceneObj) *Point {

	offsetX, offsetY := GetDirOffset(dir)
	for i := maxDis; i > 0; i++ {

		tempPoint := this.GetPointByXY(point.X()+offsetX*i, point.Y()+offsetY*i)
		if tempPoint != nil {
			if isStand {
				if tempPoint.CanStand(obj) {
					return tempPoint
				}
			} else {
				return tempPoint
			}
		}
	}
	return nil
}

func (this *Scene) SetShieldMove(flag bool) {
	this.shieldMove = flag
}

func (this *Scene) GetShieldMove() bool {
	return this.shieldMove
}

func (this *Scene) Walkable(x, y int) bool {
	return this.sceneT.Walkable(int32(x), int32(y))
}

func (this *Scene) GetSceneObj(objId int) ISceneObj {
	this.objsMu.RLock()
	obj := this.objs[objId]
	this.objsMu.RUnlock()
	return obj
}

func (this *Scene) GetSceneAllObj() map[int]ISceneObj {
	return this.objs
}

func (this *Scene) GetPlayerObjsNum() int {
	return len(this.playerObjs)
}

func (this *Scene) GetItemObjsNum() int {
	return len(this.itemObjs)
}

func (this *Scene) GetAllItemObjsByPlayer(owner int) []int {

	ids := make([]int, 0)
	for _, v := range this.itemObjs {
		if v.GetContext().(*SceneItem).owner == owner {
			ids = append(ids, v.GetObjId())
		}
	}
	return ids
}

func (this *Scene) AddSceneObj(obj ISceneObj, point *Point) error {
	if !this.sceneT.Walkable(int32(point.X()), int32(point.Y())) {
		return ErrNotWalkable
	}

	this.objsMu.Lock()
	this.addSceneObj(obj, point)
	this.objsMu.Unlock()
	if constFight.SCENE_NOTIFIER_CHAN_TYPE {
		this.posChangeChan <- newPosChangeItem(PosChangeTypeAdd, obj, point)
	} else {
		this.notifier.Add(obj)
	}
	return nil
}

func (this *Scene) addSceneObj(obj ISceneObj, point *Point) {
	obj.SetPoint(point)
	this.objs[obj.GetObjId()] = obj
	if obj.GetType() == pb.SCENEOBJTYPE_USER {
		this.playerObjs[obj.GetObjId()] = obj
	} else if obj.GetType() == pb.SCENEOBJTYPE_ITEM {
		this.itemObjs[obj.GetObjId()] = obj.GetContext().(*SceneItem)
	} else if obj.GetType() == pb.SCENEOBJTYPE_COLLECTION {
		this.collectionObjs[obj.GetObjId()] = obj.GetContext().(*SceneCollection)
	} else if obj.GetType() == pb.SCENEOBJTYPE_BUFF {
		this.buffObjs[obj.GetObjId()] = obj.GetContext().(*SceneBuff)
	}
	point.addObject(obj)
}

func (this *Scene) AddSceneObjs(objs []ISceneObj, points []*Point, enterType int) error {

	if len(objs) != len(points) {
		return ErrInSceneData
	}
	for _, v := range points {
		if !this.sceneT.Walkable(int32(v.X()), int32(v.Y())) {
			return ErrNotWalkable
		}
	}
	this.objsMu.Lock()
	for k, v := range objs {
		this.addSceneObj(v, points[k])
		v.SetScene(this)
	}
	this.objsMu.Unlock()
	if constFight.SCENE_NOTIFIER_CHAN_TYPE {
		this.posChangeChan <- newPosChangeItem(PosChangeTypeAdds, nil, nil, objs, points, enterType)
	} else {
		this.notifier.Adds(objs, points, enterType)
	}
	return nil
}

func (this *Scene) RemoveSceneObj(obj ISceneObj) {
	this.objsMu.Lock()
	delete(this.objs, obj.GetObjId())
	delete(this.playerObjs, obj.GetObjId())
	delete(this.itemObjs, obj.GetObjId())
	delete(this.collectionObjs, obj.GetObjId())
	delete(this.buffObjs, obj.GetObjId())
	oldPoint := obj.Point()
	if oldPoint != nil {
		oldPoint.removeObject(obj)
	}
	this.objsMu.Unlock()
	if constFight.SCENE_NOTIFIER_CHAN_TYPE {
		this.posChangeChan <- newPosChangeItem(PosChangeTypeRemove, obj, nil)
	} else {
		this.notifier.Remove(obj)
	}
}

func (this *Scene) MoveSceneObj(obj ISceneObj, point *Point, moveType int, moveForce bool, sendClient bool) error {
	if !this.sceneT.Walkable(int32(point.X()), int32(point.Y())) || point.IsBlock() {
		return ErrNotWalkable
	}
	oldPoint := obj.Point()
	if oldPoint.Equal(point) {
		return nil
	}
	if oldPoint != nil {
		oldPoint.removeObject(obj)
	}
	obj.SetPoint(point)
	point.addObject(obj)
	if constFight.SCENE_NOTIFIER_CHAN_TYPE {
		this.posChangeChan <- newPosChangeItem(PosChangeTypeMove, obj.GetContext(), oldPoint, moveType, moveForce, sendClient)
	} else {
		this.notifier.Move(obj.GetContext(), oldPoint, moveType, moveForce, sendClient)
	}
	return nil
}

//func (this *Scene) UpdateSceneObj(obj ISceneObj, point *Point) error {
//	if !this.sceneT.Walkable(int32(point.X()), int32(point.Y())) {
//		return ErrNotWalkable
//	}
//	oldPoint := obj.Point()
//	if oldPoint != nil {
//		oldPoint.removeObject(obj)
//	}
//	obj.SetPoint(point)
//	point.addObject(obj)
//	if notifierChanType{
//		this.posChangeChan <- newPosChangeItem(PosChangeTypeUpdate, obj, point)
//	}else{
//
//	}
//	return nil
//}

func (this *Scene) ReliveSceneObj(obj ISceneObj, point *Point, reliveType int) error {
	if !this.sceneT.Walkable(int32(point.X()), int32(point.Y())) {
		return ErrNotWalkable
	}
	oldPoint := obj.Point()
	if oldPoint != nil {
		oldPoint.removeObject(obj)
	}
	obj.SetVisible(true)

	obj.SetPoint(point)
	point.addObject(obj)
	if constFight.SCENE_NOTIFIER_CHAN_TYPE {

		this.posChangeChan <- newPosChangeItem(PosChangeTypeRelive, obj, oldPoint, reliveType)
	} else {
		this.notifier.Relive(obj, oldPoint, reliveType)
	}
	return nil
}

func (this *Scene) NotifyNearby(obj ISceneObj, msg nw.ProtoMessage, excludeSession map[uint32]bool) {
	if constFight.SCENE_NOTIFIER_CHAN_TYPE {
		if this.notifyChan != nil {
			this.notifyChan <- newNotifyItem(false, obj, msg, excludeSession)
		}
	} else {
		this.notifier.NotifyNearby(obj, msg, excludeSession)
	}

}

func (this *Scene) NotifyAll(msg nw.ProtoMessage) {
	if constFight.SCENE_NOTIFIER_CHAN_TYPE {
		if this.notifyChan != nil {
			this.notifyChan <- newNotifyItem(true, nil, msg, nil)
		}
	} else {
		this.notifier.NotifyAll(msg)
	}

}

func (this *Scene) worker() {
	defer func() {
		if r := recover(); r != nil {
			stackBytes := debug.Stack()
			logger.Error("panic when worker:%v,%s,%d", r, stackBytes, this.sceneT.Id)
			//panic(r)
			go this.worker()
		}
	}()
	for {
		select {
		case posChangeItem := <-this.posChangeChan:
			if posChangeItem.typ == PosChangeTypeAdd {
				this.notifier.Add(posChangeItem.obj)
			} else if posChangeItem.typ == PosChangeTypeAdds {
				if len(posChangeItem.arg) == 3 {
					this.notifier.Adds(posChangeItem.arg[0].([]ISceneObj), posChangeItem.arg[1].([]*Point), posChangeItem.arg[2].(int))
				}
			} else if posChangeItem.typ == PosChangeTypeRemove {
				this.notifier.Remove(posChangeItem.obj)
			} else if posChangeItem.typ == PosChangeTypeMove {
				moveFace := false
				moveType := pb.MOVETYPE_WALK
				sendClient := true
				if len(posChangeItem.arg) > 1 {
					moveFace = posChangeItem.arg[1].(bool)
				}
				if len(posChangeItem.arg) > 0 {
					moveType = posChangeItem.arg[0].(int)
				}
				if len(posChangeItem.arg) > 2 {
					sendClient = posChangeItem.arg[2].(bool)
				}
				this.notifier.Move(posChangeItem.obj, posChangeItem.point, moveType, moveFace, sendClient)
			} else if posChangeItem.typ == PosChangeTypeUpdate {
				this.notifier.Update(posChangeItem.obj)
			} else if posChangeItem.typ == PosChangeTypeRelive {
				reliveType := constFight.RELIVE_TYPE_NOMAL
				if len(posChangeItem.arg) > 0 {
					reliveType = posChangeItem.arg[0].(int)
				}
				this.notifier.Relive(posChangeItem.obj, posChangeItem.point, reliveType)
			}
		case notifyItem := <-this.notifyChan:
			if notifyItem.isNotifyAll {
				this.notifier.NotifyAll(notifyItem.msg)
			} else {
				this.notifier.NotifyNearby(notifyItem.obj, notifyItem.msg, notifyItem.excludeSession)
			}
		case <-this.done:
			return
		}
	}
}

/**
*  @Description: 刷新
*  @receiver this
*  @return bool 是否有道具消息
**/
func (this *Scene) UpdateFrame() bool {

	hasItemDisappeared := false
	now := int(time.Now().Unix())
	for _, v := range this.itemObjs {
		if v.disappeared(now) {
			this.RemoveSceneObj(v)
			hasItemDisappeared = true
		}
	}
	return hasItemDisappeared
}

/**
 *  @Description: 采集帧刷新
 *  @return map[int]int
 */
func (this *Scene) UpdateFrameCollection() map[int]int {

	now := int(time.Now().Unix())
	collections := make(map[int]int)

	for _, v := range this.collectionObjs {
		if v.disappeared(now) {
			collections[v.collectionObjId] = v.getCollectionId()
			collectionConf := gamedb.GetCollectionCollectionCfg(v.getCollectionId())
			if collectionConf.Type != constFight.COLLECTION_TYPE_ONE {
				this.RemoveSceneObj(v)
			} else {
				v.SetVisible(false)
				v.Reset(true)
				msg := v.BuildDisappearMessage()
				this.NotifyNearby(v, msg, nil)
			}
		} else if v.ReShow() {
			v.SetVisible(true)
			msg := v.BuildAppearMessage()
			this.NotifyNearby(v, msg, nil)
		}
	}
	return collections
}

func (this *Scene) Destroy() {
	logger.Debug("scene.go:Destroy:", this.sceneT.Id)

	go func() {

		defer func() {
			if r := recover(); r != nil {
				stackBytes := debug.Stack()
				logger.Error("panic Scene Destroy:%v,%s", r, stackBytes)
			}
		}()

		var ticker = time.NewTicker(time.Second * 1)
		for {
			<-ticker.C
			if len(this.posChangeChan) != 0 || len(this.notifyChan) != 0 {
				continue
			}
			close(this.done)
			return
		}
	}()
}

func (this *Scene) GetMapId() int {
	return this.sceneT.Id
}

func (this *Scene) GetSceneGirdCol() int {
	return this.sceneT.ColNum
}

func (this *Scene) GetPointByRangeDis(point *Point, rangeDis int) []*Point {

	dis := rangeDis / 2
	if rangeDis%2 != 0 {
		dis += 1
	}

	startX := point.X() - dis
	if startX < 0 {
		startX = 0
	}
	startY := point.Y() - dis
	if startY < 0 {
		startY = 0
	}
	endX := point.X() + dis
	if endX > this.sceneT.ColNum {
		endX = this.sceneT.ColNum
	}
	endY := point.Y() + dis
	if endY > this.sceneT.RowNum {
		endY = this.sceneT.RowNum
	}
	nearestPoint := make([]*Point, 0)
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			index := gamedb.GetPointIndex(int32(x), int32(y))
			if this.points[index] != nil && !this.points[index].IsBlock() {
				nearestPoint = append(nearestPoint, this.points[index])
			}
		}
	}
	return nearestPoint
}

/**
 *  @Description:	根据坐标点 偏移获取地图点
 *  @param point	原点
 *  @param offset   偏移区域
 *  @return []*Point 返回偏移区域坐标
 *  @return error
 */
func (this *Scene) GetSceneAreaByPointOffset(point *Point, offset []int) ([]*Point, error) {

	points := make([]*Point, 0)
	for i := 0; i < len(offset); i += 2 {

		offsetIndex := gamedb.GetPointIndex(int32(point.X()+offset[i]), int32(point.Y()+offset[i+1]))
		offsetPoint := this.points[offsetIndex]
		//判断坐标格子，及是否安全区
		if offsetPoint != nil && !offsetPoint.IsSafe() {

			points = append(points, offsetPoint)
		}
	}
	return points, nil
}

func (this *Scene) CheckInNpcRange(npcMapId int, point *Point) (bool, error) {

	for k, v := range this.sceneT.Npc {
		if k == npcMapId {
			mapPoint := this.points[v]
			if mapPoint == nil {
				logger.Error("判断是否在npc范围，npc坐标未找到", npcMapId)
				return false, gamedb.ERRPARAM
			} else {
				dis := DistanceByPoint(point, mapPoint)
				if dis > 2 {
					return false, nil
				} else {
					return true, nil
				}
			}
		}
	}
	logger.Error("判断是否在npc范围，npc地图上未找到", npcMapId)
	return false, gamedb.ERRPARAM
}
