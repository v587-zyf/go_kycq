package ai

import (
	"container/heap"
	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/logger"
	"math"
)

type OpenList []*AStarPoint

func (self OpenList) Len() int           { return len(self) }
func (self OpenList) Less(i, j int) bool { return self[i].fVal < self[j].fVal }
func (self OpenList) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }

func (this *OpenList) Push(x interface{}) {
	*this = append(*this, x.(*AStarPoint))
}

func (this *OpenList) Pop() interface{} {
	old := *this
	n := len(old)
	x := old[n-1]
	*this = old[0 : n-1]
	return x
}

type AStarPoint struct {
	gamedb.Point
	father *AStarPoint
	gVal   int
	hVal   int
	fVal   int
}

func (this *AStarPoint) Equal(p *gamedb.Point) bool {
	return this.X == p.X && this.Y == p.Y
}

func (this *AStarPoint) calcGVal() int {
	if this.father != nil {
		deltaX := math.Abs(float64(this.father.X - this.X))
		deltaY := math.Abs(float64(this.father.Y - this.Y))
		if deltaX == 1 && deltaY == 0 {
			this.gVal = this.father.gVal + 10
		} else if deltaX == 0 && deltaY == 1 {
			this.gVal = this.father.gVal + 10
		} else if deltaX == 1 && deltaY == 1 {
			this.gVal = this.father.gVal + 14
		} else {
			panic("father point is invalid!")
		}
	}
	return this.gVal
}

func (this *AStarPoint) calcHVal(end *AStarPoint) int {
	this.hVal = int(math.Abs(float64(end.X-this.X)) + math.Abs(float64(end.Y-this.Y)))
	return this.hVal
}

func (this *AStarPoint) calcFVal(end *AStarPoint) int {
	this.fVal = this.calcGVal() + this.calcHVal(end)
	return this.fVal
}

type AStarManager struct {
	walkableMap map[int32]*gamedb.Point
	start       *AStarPoint
	end         *AStarPoint
	closeLi     map[int32]*AStarPoint
	openLi      OpenList
	openSet     map[int32]*AStarPoint
	TheRoad     []*AStarPoint
	count		int
}

func NewAStarManager(m map[int32]*gamedb.Point) *AStarManager {
	sr := &AStarManager{}
	sr.walkableMap = m
	sr.TheRoad = make([]*AStarPoint, 0)
	sr.openSet = make(map[int32]*AStarPoint, 0)
	sr.closeLi = make(map[int32]*AStarPoint, 0)
	return sr
}

func (this *AStarManager) newAStarPoint(p *gamedb.Point, father *AStarPoint, end *AStarPoint) (ap *AStarPoint) {
	ap = &AStarPoint{*p, father, 0, 0, 0}
	if end != nil {
		ap.calcFVal(end)
	}
	return ap
}

func (this *AStarManager) getAdjacentPoint(x, y int32) (adjacent []*gamedb.Point) {
	if key := x<<16 | (y - 1); this.walkableMap[key] != nil {
		adjacent = append(adjacent, this.walkableMap[key])
	}
	if key := (x+1)<<16 | (y - 1); this.walkableMap[key] != nil {
		adjacent = append(adjacent, this.walkableMap[key])
	}
	if key := (x+1)<<16 | y; this.walkableMap[key] != nil {
		adjacent = append(adjacent, this.walkableMap[key])
	}
	if key := (x+1)<<16 | (y + 1); this.walkableMap[key] != nil {
		adjacent = append(adjacent, this.walkableMap[key])
	}
	if key := x<<16 | (y + 1); this.walkableMap[key] != nil {
		adjacent = append(adjacent, this.walkableMap[key])
	}
	if key := (x-1)<<16 | (y + 1); this.walkableMap[key] != nil {
		adjacent = append(adjacent, this.walkableMap[key])
	}
	if key := (x-1)<<16 | y; this.walkableMap[key] != nil {
		adjacent = append(adjacent, this.walkableMap[key])
	}
	if key := (x-1)<<16 | (y - 1); this.walkableMap[key] != nil {
		adjacent = append(adjacent, this.walkableMap[key])
	}
	return adjacent
}

func (this *AStarManager) pointAsKey(x, y int32) int32 {
	return x<<16 | y
}

func (this *AStarManager) MoveByBeforeRoad() bool {
	if this.count > 10 {
		this.count = 0
		return false
	}
	logger.Debug("使用原有路径继续走")
	saveDataLen := len(this.TheRoad)
	if saveDataLen <= 1 {
		logger.Debug("旧路径已全部走完,需要重新寻路")
		this.count = 0
		return false
	} else {
		this.TheRoad = append(this.TheRoad[:(saveDataLen-1)])
		return true
	}
}

func (this *AStarManager) FindAStarRoad(startX, startY, endX, endY int32) bool {
	logger.Debug("startX:%d, startY:%d, endX:%d, endY:%d",startX, startY, endX, endY)
	this.start = this.newAStarPoint(&gamedb.Point{startX, startY}, nil, nil)
	this.end = this.newAStarPoint(&gamedb.Point{endX, endY}, nil, nil)
	saveData := make([]*AStarPoint, 0)
	saveData = append(saveData, this.TheRoad...)
	this.TheRoad = make([]*AStarPoint, 0)
	this.openSet = make(map[int32]*AStarPoint, 0)
	this.closeLi = make(map[int32]*AStarPoint, 0)

	heap.Init(&this.openLi)
	heap.Push(&this.openLi, this.start) // 首先把起点加入开放列表
	this.openSet[this.pointAsKey(this.start.X, this.start.Y)] = this.start
	r := this.FindRoad()
	logger.Debug("astar寻路结果如下")
	for i, data := range this.TheRoad {
		logger.Debug("[%d]x:%d, y:%d", i, data.X, data.Y)
	}
	saveDataLen := len(saveData)
	if !r {
		if len(saveData) <= 1 {
			logger.Debug("目标点不可走,旧路径已全部走完")
		} else {
			logger.Debug("目标点不可走,使用原有路径继续走!")
			this.TheRoad = append(saveData[:(saveDataLen-1)])
		}
	}
	return r
}

func (this *AStarManager) FindRoad() bool {
	for len(this.openLi) > 0 {
		// 将节点从开放列表移到关闭列表当中。
		x := heap.Pop(&this.openLi)
		curPoint := x.(*AStarPoint)
		curKey := this.pointAsKey(curPoint.X, curPoint.Y)
		delete(this.openSet, curKey)
		this.closeLi[curKey] = curPoint

		points := this.getAdjacentPoint(curPoint.X, curPoint.Y)
		for _, p := range points {
			theAP := this.newAStarPoint(p, curPoint, this.end)
			if this.end.Equal(p) {
				// 找出路径了, 标记路径
				for theAP.father != nil {
					this.TheRoad = append(this.TheRoad, theAP)
					theAP = theAP.father
				}
				return true
			}

			pKey := this.pointAsKey(p.X, p.Y)
			_, ok := this.closeLi[pKey]
			if ok {
				continue
			}
			if this.walkableMap[pKey] == nil {
				continue
			}

			existAP, ok := this.openSet[pKey]
			if !ok {
				heap.Push(&this.openLi, theAP)
				this.openSet[this.pointAsKey(theAP.X, theAP.Y)] = theAP
			} else {
				oldGVal, oldFather := existAP.gVal, existAP.father
				existAP.father = curPoint
				existAP.calcGVal()
				// 如果新的节点的G值还不如老的节点就恢复老的节点
				if existAP.gVal > oldGVal {
					// restore father
					existAP.father = oldFather
					existAP.gVal = oldGVal
				}
			}

		}
	}

	return false
}

func (this *AStarManager) GetRoadPoint() (int32, int32) {
	if len(this.TheRoad) == 0 {
		return 0, 0
	}
	return this.TheRoad[len(this.TheRoad)-1].X, this.TheRoad[len(this.TheRoad)-1].Y
}

func (this *AStarManager) ClearRoad() {
	logger.Debug("清楚原有移动路径")
	this.count = 0
	this.TheRoad = this.TheRoad[0:0]
}