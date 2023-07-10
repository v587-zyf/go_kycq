package scene

import (
	"cqserver/protobuf/pb"
	"fmt"
	"math"
)

type Point struct {
	id        int               //格子id
	x         int               //列（X）
	y         int               //行（Y）
	near      []*Point          //附近格子
	sceneObjs map[int]ISceneObj //格子中对象
	isBlock   bool              //是否阻塞的格子
	isSafe    bool              //是否安全

	//寻路使用
	father *Point
	gVal   int
	hVal   int
	fVal   int
}

func (this *Point) IsSafe() bool {
	return this.isSafe
}

func (this *Point) SetIsSafe(isSafe bool) {
	this.isSafe = isSafe
}

func (this *Point) Y() int {
	return this.y
}

func (this *Point) X() int {
	return this.x
}

func NewGird(id int, x, y int, isBlock bool) *Point {
	g := &Point{}
	g.id = id
	g.x = x
	g.y = y
	g.near = make([]*Point, 0)
	g.sceneObjs = make(map[int]ISceneObj)
	g.isBlock = isBlock
	return g
}

func (this *Point) GetSceneObjsNum() int {
	num := 0
	for _,v := range this.sceneObjs{
		if v.GetVisible(){
			num ++
		}
	}
	return num
}

func (this *Point) setNears(near []*Point) {
	this.near = near
}

func (this *Point) GetOneObject() ISceneObj {

	for _, v := range this.sceneObjs {
		return v
	}
	return nil
}

func (this *Point) GetAllObject() map[int]ISceneObj {
	return this.sceneObjs
}

func (this *Point) addObject(obj ISceneObj) {
	this.sceneObjs[obj.GetObjId()] = obj
}

func (this *Point) removeObject(obj ISceneObj) {

	delete(this.sceneObjs, obj.GetObjId())
}

func (this *Point) GetNewNearPointByDir(dir int) *Point {
	if dir == DIR_NONE || dir > DIR_LEFT_TOP {
		return nil
	}
	if this.near[dir] != nil {
		return this.near[dir]
	}
	return nil
}

func (this *Point) GetNearPoints() []*Point {
	return this.near
}

func (this *Point) SetIsBlock(isBlock bool) {
	this.isBlock = isBlock
}

func (this *Point) IsBlock() bool {
	return this.isBlock
}

func (this *Point) CanStand(obj ISceneObj) bool {
	if this.IsBlock() {
		return false
	}

	for _, v := range this.sceneObjs {

		if v.GetObjId() == obj.GetObjId(){
			continue
		}

		if v.IsSceneObj() {
			continue
		} else {
			if v.GetVisible() {

				if obj.GetType() == pb.SCENEOBJTYPE_MONSTER || obj.IsSceneObj() {
					return false
				} else {
					if v.GetType() == pb.SCENEOBJTYPE_MONSTER {
						return false
					}
				}
			}
		}
	}
	return true
}

func (this *Point) ToString() string {

	return fmt.Sprintf("x:%v，y:%v", this.x, this.y)
}

func (this *Point) ToPbPoint() *pb.Point {
	return &pb.Point{
		X: int32(this.x),
		Y: int32(this.y),
	}
}

func (this *Point) newAStarPoint(father *Point, end *Point) {
	this.father = father
	if end != nil {
		this.calcFVal(end)
	}
}

func (this *Point) Equal(p *Point) bool {
	return this.X() == p.X() && this.Y() == p.Y()
}

func (this *Point) calcGVal() int {
	if this.father != nil {
		deltaX := math.Abs(float64(this.father.X() - this.X()))
		deltaY := math.Abs(float64(this.father.Y() - this.Y()))
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

func (this *Point) calcHVal(end *Point) int {
	this.hVal = int(math.Abs(float64(end.X()-this.X())) + math.Abs(float64(end.Y()-this.Y())))
	return this.hVal
}

func (this *Point) calcFVal(end *Point) int {
	this.fVal = this.calcGVal() + this.calcHVal(end)
	return this.fVal
}
