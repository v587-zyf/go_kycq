package scene

import (
	"container/heap"
	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/logger"
)

type OpenList []*Point

func (self OpenList) Len() int           { return len(self) }
func (self OpenList) Less(i, j int) bool { return self[i].fVal < self[j].fVal }
func (self OpenList) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }

func (this *OpenList) Push(x interface{}) {
	*this = append(*this, x.(*Point))
}

func (this *OpenList) Pop() interface{} {
	old := *this
	n := len(old)
	x := old[n-1]
	*this = old[0 : n-1]
	return x
}

type AStarManager struct {
	scene *Scene
}

func NewAStarManager(scene *Scene) *AStarManager {
	sr := &AStarManager{}
	sr.scene = scene
	return sr
}

func (this *AStarManager) FindAStarRoad(startPoint *Point, endPoint *Point) []*Point {

	if startPoint.Equal(endPoint) {
		logger.Info("---------------------------寻路异常：【%v】-【%v】", startPoint.ToString(), endPoint.ToString())
		return nil
	}
	//startTime := time.Now()
	startPoint.newAStarPoint(nil, nil)
	endPoint.newAStarPoint(nil, nil)
	theRoad := this.FindRoad(startPoint, endPoint)
	//roadStr := fmt.Sprintf("astar寻路结果如下,开始：%v,结束：%v", startPoint.ToString(), endPoint.ToString())
	//for _, data := range theRoad {
	//	roadStr += fmt.Sprintf("经过点【%v】", data.ToString())
	//}
	//roadStr += fmt.Sprintf("耗时：%v", time.Now().Sub(startTime).Milliseconds())
	//logger.Debug(roadStr)
	return theRoad
}

func (this *AStarManager) FindRoad(startPoint *Point, endPoint *Point) []*Point {
	openSet := make(map[int32]*Point, 0)
	closeLi := make(map[int32]*Point, 0)
	openSet[int32(gamedb.GetPointIndex(int32(startPoint.X()), int32(startPoint.Y())))] = startPoint
	roadPath := make([]*Point, 0)
	var openList OpenList
	heap.Init(&openList)
	heap.Push(&openList, startPoint)
	for len(openList) > 0 {
		// 将节点从开放列表移到关闭列表当中。
		x := heap.Pop(&openList)
		curPoint := x.(*Point)
		curKey := int32(gamedb.GetPointIndex(int32(curPoint.X()), int32(curPoint.Y())))
		delete(openSet, curKey)
		closeLi[curKey] = curPoint

		points := curPoint.GetNearPoints()
		for _, p := range points {
			if p == nil {
				//logger.Info("附近格子异常", curPoint.ToString())
				continue
			}
			if p.IsBlock() {
				continue
			}

			if endPoint.Equal(p) {
				p.newAStarPoint(curPoint, nil)
				// 找出路径了, 标记路径
				for p.father != nil {
					roadPath = append(roadPath, p)
					p = p.father
				}
				return roadPath
			}

			pKey := gamedb.GetPointIndex(int32(p.X()), int32(p.Y()))
			_, ok := closeLi[int32(pKey)]
			if ok {
				continue
			}

			existAP, ok := openSet[int32(pKey)]
			if !ok {
				p.newAStarPoint(curPoint, endPoint)
				heap.Push(&openList, p)
				openSet[int32(pKey)] = p
			} else {
				oldGVal, oldFather := existAP.gVal, existAP.father
				existAP.father = curPoint
				existAP.calcGVal()
				// 如果新的节点的G值还不如老的节点就恢复老的节点
				if existAP.gVal > oldGVal {
					existAP.father = oldFather
					existAP.gVal = oldGVal
				}
			}
		}
	}

	return roadPath
}

//func (this *AStarManager) GetRoadPoint() (int32, int32) {
//	if len(this.TheRoad) == 0 {
//		return 0, 0
//	}
//	return this.TheRoad[len(this.TheRoad)-1].X, this.TheRoad[len(this.TheRoad)-1].Y
//}
//
//func (this *AStarManager) ClearRoad() {
//	logger.Debug("清楚原有移动路径")
//	this.count = 0
//	this.TheRoad = this.TheRoad[0:0]
//}
