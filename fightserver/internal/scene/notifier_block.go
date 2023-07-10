package scene

import (
	"cqserver/fightserver/internal/net"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/common"
	"cqserver/protobuf/pb"
	"math"
	"reflect"

	"cqserver/gamelibs/errex"

	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
)

const (
	BLOCK_SIZE_X = 6
	BLOCK_SIZE_Y = 10
	RANGE        = 3
)

type Block struct {
	blockNotifier *BlockNotifier
	x             int
	y             int
	objs          map[int]ISceneObj
}

type BlockNotifier struct {
	*NotifierBase
	blockss [][]*Block
}

func NewBlock(blockNotifier *BlockNotifier, x, y int) *Block {
	return &Block{
		blockNotifier: blockNotifier,
		x:             x,
		y:             y,
		objs:          make(map[int]ISceneObj),
	}
}

func (this *Block) AddSceneObj(obj ISceneObj) {
	if !reflect.ValueOf(obj).IsNil() {
		obj.SetBlockIndex(this.x*1000 + this.y)
		this.objs[obj.GetObjId()] = obj
	} else {
		err := errex.New("add nil obj")
		logger.Error(err.Error() + err.ErrorStack())
	}
}

func (this *Block) RemoveSceneObj(obj ISceneObj) {
	if _, ok := this.objs[obj.GetObjId()]; ok {
		obj.SetBlockIndex(0)
		delete(this.objs, obj.GetObjId())
	} else {
		logger.Info("移除对象失败，区域未找到该对象", obj.GetObjId())
	}
}

func (this *Block) Notify(msg nw.ProtoMessage) {
	if len(this.objs) == 0 {
		return
	}
	ids := make(map[int]map[int]int)
	for _, obj := range this.objs {
		if obj != nil && obj.SessionId() > 0 {
			hostId := obj.HostId()
			if ids[hostId] == nil {
				ids[hostId] = make(map[int]int)
			}
			ids[hostId][int(obj.SessionId())] = hostId
		}
	}
	if len(ids) > 0 {
		net.GetGateConn().BroadcastToGate(ids, msg)
	}
}

func (this *Block) NotifyWithExclude(msg nw.ProtoMessage, excludeSession map[uint32]bool) {

	if len(this.objs) == 0 {
		return
	}
	ids := make(map[int]map[int]int)
	for _, obj := range this.objs {
		if obj != nil && obj.SessionId() > 0 {
			if excludeSession != nil && excludeSession[obj.SessionId()] {
				continue
			}
			hostId := obj.HostId()
			if ids[hostId] == nil {
				ids[hostId] = make(map[int]int)
			}
			ids[hostId][int(obj.SessionId())] = hostId
		}
	}
	if len(ids) > 0 {
		net.GetGateConn().BroadcastToGate(ids, msg)
	}
}

// sendOthersMsgTo 发送其他人的出现消息
func (this *Block) sendOthersMsgTo(target ISceneObj, typ BatchBuilderType, withoutobjs map[int]bool) {

	hostId := target.HostId()
	if hostId <= 0 {
		return
	}

	var builder = GetBatchBuilder(this.blockNotifier.StageId, typ)
	var notifyCount = 0
	for _, obj := range this.objs {
		if obj == nil {
			continue
		}
		if withoutobjs != nil && withoutobjs[obj.GetObjId()] {
			continue
		}
		if obj != target && obj.GetVisible() {
			builder.AddObj(obj)
			notifyCount++
		}
		if notifyCount >= MaxNotifierNum {
			net.GetGateConn().SendMessage(uint32(hostId), target.SessionId(), 0, builder.Build())
			builder.Reset()
			notifyCount = 0
		}
	}
	if builder != nil && builder.Length() > 0 {
		net.GetGateConn().SendMessage(uint32(hostId), target.SessionId(), 0, builder.Build())
	}
}

func NewBlockNotifier(stageId int, sceneT *gamedb.SceneConf) *BlockNotifier {

	width := int(math.Ceil(float64(sceneT.ColNum) / float64(BLOCK_SIZE_X)))
	height := int(math.Ceil(float64(sceneT.RowNum) / float64(BLOCK_SIZE_Y)))

	blockNotifier := &BlockNotifier{
		NotifierBase: &NotifierBase{StageId: stageId},
	}
	blockss := make([][]*Block, width)
	for i := range blockss {
		blockss[i] = make([]*Block, height)
		for j := range blockss[i] {
			blockss[i][j] = NewBlock(blockNotifier, i, j)
		}
	}
	blockNotifier.blockss = blockss
	return blockNotifier
}

func (this *BlockNotifier) getBlockPosByPixel(point *Point) (int, int) {
	x := int(math.Floor(float64(point.X()) / float64(BLOCK_SIZE_X)))
	y := int(math.Floor(float64(point.Y()) / float64(BLOCK_SIZE_Y)))
	return x, y
}

func (this *BlockNotifier) getBlockByPixel(point *Point) *Block {
	x := int(math.Floor(float64(point.X()) / float64(BLOCK_SIZE_X)))
	y := int(math.Floor(float64(point.Y()) / float64(BLOCK_SIZE_Y)))
	return this.getBlock(x, y)
}

func (this *BlockNotifier) getBlock(x int, y int) *Block {
	if x < 0 || x >= len(this.blockss) || y < 0 || y >= len(this.blockss[x]) {
		return nil
	}
	return this.blockss[x][y]
}

func (this *BlockNotifier) Add(obj ISceneObj) {
	if obj == nil {
		return
	}
	block := this.getBlockByPixel(obj.Point())
	if block == nil {
		logger.Error("BlockNotifier.Add obj position outbound: (%s)", obj.Point().ToString())
		return
	}
	obj = obj.GetContext()
	block.AddSceneObj(obj)
	if obj.GetVisible() {
		appearMsg := obj.BuildAppearMessage()
		this.NotifyNearby(obj, appearMsg, nil)
	}
	if obj.GetType() == pb.SCENEOBJTYPE_USER {
		this.sendNearbyMsgTo(obj, BatchBuilderTypeEnter, nil)
		//推送玩家进入场景结束
		if obj.GetVisible() && obj.HostId() > 0 && obj.SessionId() > 0 {
			net.GetGateConn().SendMessage(uint32(obj.HostId()), obj.SessionId(), 0, &pb.SceneEnterOverNtf{})
		}
	}
}

// 增加多个对象，并通知相关消息{
func (this *BlockNotifier) Adds(objs []ISceneObj, points []*Point, enterType int) {

	firstObj := objs[0].GetContext()
	var builder = GetBatchBuilder(this.StageId, BatchBuilderTypeEnter)
	builder.SetEnterType(enterType)
	withoutObjs := make(map[int]bool)
	for _, obj := range objs {
		obj = obj.GetContext()
		block := this.getBlockByPixel(obj.Point())
		if block == nil {
			logger.Error("进入场景异常：%，未找到区域块", obj.Point())
			continue
		}
		block.AddSceneObj(obj)
		withoutObjs[obj.GetObjId()] = true
		if obj.GetVisible() {
			builder.AddObj(obj)
		}
	}

	appearMsg := builder.Build()
	this.NotifyNearby(firstObj, appearMsg, nil)

	if enterType != constFight.SCENE_ENTER_FIT && firstObj.GetType() == pb.SCENEOBJTYPE_USER {
		this.sendNearbyMsgTo(firstObj, BatchBuilderTypeEnter, withoutObjs)
		//推送玩家进入场景结束
		if firstObj.GetVisible() && firstObj.HostId() > 0 && firstObj.SessionId() > 0 {
			net.GetGateConn().SendMessage(uint32(firstObj.HostId()), firstObj.SessionId(), 0, &pb.SceneEnterOverNtf{})
		}
	}
}

func (this *BlockNotifier) Update(obj ISceneObj) {
	if obj == nil {
		return
	}
	block := this.getBlockByPixel(obj.Point())
	if block == nil {
		logger.Error("BlockNotifier.Add obj position outbound: (%s)", obj.Point().ToString())
		return
	}
	obj = obj.GetContext()
	block.AddSceneObj(obj)
	if obj.GetVisible() {
		appearMsg := obj.BuildAppearMessage()
		this.NotifyNearby(obj, appearMsg, nil)
	}
}

func (this *BlockNotifier) Relive(obj ISceneObj, oldPoint *Point, reliveType int) {

	block := this.getBlockByPixel(obj.Point())
	if block == nil {
		logger.Error("BlockNotifier.Add obj position outbound: (%s)", obj.Point().ToString())
		return
	}

	isNewBlock := false
	oldBlock := this.getBlockByPixel(oldPoint)
	if oldBlock != nil && (oldBlock.x != block.x || oldBlock.y != block.y) {
		oldBlock.RemoveSceneObj(obj)
		isNewBlock = true
	}
	obj = obj.GetContext()
	block.AddSceneObj(obj)
	if obj.GetVisible() {
		if isNewBlock {
			this.sendNearbyMsgTo(obj, BatchBuilderTypeEnterForTower, nil)
		}
		appearMsg := obj.BuildRelliveMessage()
		if appearMsg != nil {
			appearMsg.(*pb.SceneUserReliveNtf).ReliveType = int32(reliveType)
		}
		this.NotifyNearby(obj, appearMsg, nil)
	}
}

func (this *BlockNotifier) Remove(obj ISceneObj) {
	if obj == nil {
		return
	}
	block := this.getBlockByPixel(obj.Point())
	if block == nil {
		logger.Error("BlockNotifier.Remove obj position outbound: (%s)", obj.Point().ToString())
		return
	}
	block.RemoveSceneObj(obj)
	//if obj.GetVisible() {
	disappearMsg := obj.BuildDisappearMessage()
	this.NotifyNearby(obj, disappearMsg, nil)
	//}
}

func (this *BlockNotifier) Move(obj ISceneObj, oldPoint *Point, moveType int, moveForce bool, sendClient bool) {
	if obj == nil {
		return
	}

	oldBlock := this.getBlockByPixel(oldPoint)
	newBlock := this.getBlockByPixel(obj.Point())
	if newBlock == nil {
		logger.Error("Move obj:%v,x:%d,y:%d newBlock is empty")
		return
	}
	moveMsg := obj.BuildMoveMessage(moveType, moveForce)
	// 没有跨block，直接发送move消息
	if sendClient && oldBlock == newBlock && obj.GetVisible() {
		this.NotifyNearby(obj, moveMsg, nil)
		return
	}

	oldBlock.RemoveSceneObj(obj)
	newBlock.AddSceneObj(obj)

	if !sendClient {
		return
	}

	if moveForce {
		this.moveForce(obj, oldBlock, newBlock, moveType)
	} else {
		this.moveNormal(obj, oldBlock, newBlock, moveType)
	}
}

func (this *BlockNotifier) moveNormal(obj ISceneObj, oldBlock, newBlock *Block, moveType int) {

	moveMsg := obj.BuildMoveMessage(moveType, false)
	appearMsg := obj.BuildAppearMessage()
	disappearMsg := obj.BuildDisappearMessage()
	// minX, maxX, minY, maxY为oldBlock的nearby与newBlock的nearby组合的区域范围
	minX := common.MinIntGet(oldBlock.x-RANGE, newBlock.x-RANGE)
	maxX := common.MaxIntGet(oldBlock.x+RANGE, newBlock.x+RANGE)
	minY := common.MinIntGet(oldBlock.y-RANGE, newBlock.y-RANGE)
	maxY := common.MaxIntGet(oldBlock.y+RANGE, newBlock.y+RANGE)

	var block *Block
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			block = this.getBlock(i, j)
			if block == nil {
				continue
			}
			if common.IntAbs(i-newBlock.x) > RANGE || common.IntAbs(j-newBlock.y) > RANGE {
				//不在新区域周边范围内，移除对象
				if obj.GetVisible() {
					block.Notify(disappearMsg)
				}
				block.sendOthersMsgTo(obj, BatchBuilderTypeLeaveForTower, nil)
			} else if common.IntAbs(i-oldBlock.x) > RANGE || common.IntAbs(j-oldBlock.y) > RANGE {
				//不在旧区域周边范围内，添加对象
				if obj.GetVisible() {
					block.Notify(appearMsg)
				}
				block.sendOthersMsgTo(obj, BatchBuilderTypeEnterForTower, nil)
			} else {
				//新旧区域重合范围内，移动
				block.Notify(moveMsg)
			}
		}
	}
}

func (this *BlockNotifier) moveForce(obj ISceneObj, oldBlock, newBlock *Block, moveType int) {

	moveMsg := obj.BuildMoveMessage(moveType, true)
	appearMsg := obj.BuildAppearMessage()
	disappearMsg := obj.BuildDisappearMessage()
	// minX, maxX, minY, maxY为oldBlock的nearby与newBlock的nearby组合的区域范围
	minX := oldBlock.x - RANGE
	maxX := oldBlock.x + RANGE
	minY := oldBlock.y - RANGE
	maxY := oldBlock.y + RANGE

	var excludeSession map[uint32]bool
	if obj.SessionId() > 0 {
		excludeSession = map[uint32]bool{obj.SessionId(): true}
	}

	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			block := this.getBlock(i, j)
			if block == nil {
				continue
			}
			//不在旧区域周边范围内，添加对象
			if obj.GetVisible() {
				block.NotifyWithExclude(disappearMsg, excludeSession)
			}
			block.sendOthersMsgTo(obj, BatchBuilderTypeLeaveForTower, nil)
		}
	}
	//推送玩家强制移动
	if obj.SessionId() > 0 {
		net.GetGateConn().SendMessage(uint32(obj.HostId()), obj.SessionId(), 0, moveMsg)
	}

	minX = newBlock.x - RANGE
	maxX = newBlock.x + RANGE
	minY = newBlock.y - RANGE
	maxY = newBlock.y + RANGE

	//先通知玩家所在区域
	if obj.GetVisible() {
		newBlock.NotifyWithExclude(appearMsg, excludeSession)
	}
	newBlock.sendOthersMsgTo(obj, BatchBuilderTypeEnterForTower, nil)
	//通知临边区域
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			block := this.getBlock(i, j)
			if block == nil || block == newBlock {
				continue
			}
			//新区域周边范围内，添加对象
			if obj.GetVisible() {
				block.Notify(appearMsg)
			}
			block.sendOthersMsgTo(obj, BatchBuilderTypeEnterForTower, nil)
		}
	}
}

func (this *BlockNotifier) NotifyNearby(obj ISceneObj, msg nw.ProtoMessage, excludeSession map[uint32]bool) {
	if obj == nil {
		return
	}
	x, y := this.getBlockPosByPixel(obj.Point())
	for i := -RANGE; i <= RANGE; i++ {
		for j := -RANGE; j <= RANGE; j++ {
			block := this.getBlock(x+i, y+j)
			if block == nil {
				continue
			}
			if excludeSession != nil {
				block.NotifyWithExclude(msg, excludeSession)
			} else {
				block.Notify(msg)
			}
		}
	}
}

func (this *BlockNotifier) NotifyAll(msg nw.ProtoMessage) {
	for _, blocks := range this.blockss {
		for _, block := range blocks {
			block.Notify(msg)
		}
	}
}

func (this *BlockNotifier) sendNearbyMsgTo(obj ISceneObj, typ BatchBuilderType, withoutobjs map[int]bool) {
	if obj == nil {
		return
	}
	x, y := this.getBlockPosByPixel(obj.Point())

	//优先通知玩家所在区块
	objBlock := this.getBlockByPixel(obj.Point())
	if objBlock != nil {
		objBlock.sendOthersMsgTo(obj, typ, withoutobjs)
	}

	for i := -RANGE; i <= RANGE; i++ {
		for j := -RANGE; j <= RANGE; j++ {
			if i == 0 && j == 0 {
				continue
			}
			block := this.getBlock(x+i, y+j)
			if block == nil {
				continue
			}
			block.sendOthersMsgTo(obj, typ, withoutobjs)
		}
	}
}
