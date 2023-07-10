package base

import (
	"cqserver/fightserver/internal/scene"
	"cqserver/protobuf/pb"
	"fmt"
)

var (
	skillRangTypeOffsetFunc map[int]func(rangeDis int) map[int][]int
)

func init() {
	skillRangTypeOffsetFunc = map[int]func(rangeDis int) map[int][]int{
		pb.SKILLRANGETYPE_LINE:        LineOffset,
		pb.SKILLRANGETYPE_SQUARE:      SquareOffset,
		pb.SKILLRANGETYPE_SQUARE_HALF: SquareHalfOffset,
	}
}

/**
 *  @Description: 线行偏移
 *  @param rangeDis	释放范围
 *  @return map[int][]int
 */
func LineOffset(rangeDis int) map[int][]int {

	offset := make(map[int][]int)
	for _, v := range pb.SCENEDIR_ARRAY {
		offset[v] = make([]int, 0)
		offset[v] = append(offset[v], 0, 0)
		offsetX, offsetY := scene.GetDirOffset(v)
		startX := offsetX
		startY := offsetY
		for j := 0; j < rangeDis; j++ {
			offset[v] = append(offset[v], startX+offsetX*j, startY+offsetY*j)
		}
	}
	return offset
}

func SquareHalfOffset(rangeDis int) map[int][]int {

	offSet := make(map[int][]int)
	offsetMap := make(map[int]map[string]bool)
	var calc func(dir, x, y, dis int)
	calc = func(dir, x, y, dis int) {
		for i := 0; i < 5; i++ {
			offsetX, offsetY := scene.GetDirOffset((i + dir + 6) % 8)
			if dis > 1 {
				calc(dir, x+offsetX, y+offsetY, dis-1)
			}
			key := fmt.Sprintf("%d_%d",x+offsetX, y+offsetY)
			if _,ok := offsetMap[dir][key];!ok{
				offsetMap[dir][key] = true
				offSet[dir] = append(offSet[dir], x+offsetX, y+offsetY)
			}
		}
	}
	for _, v := range pb.SCENEDIR_ARRAY {

		//offSet[v] = append(offSet[v], 0, 0)
		offsetMap[v] = make(map[string]bool)
		calc(v, 0, 0, rangeDis)
	}
	return offSet
}

func SquareOffset(rangeDis int) map[int][]int {

	offset := make(map[int][]int)
	startX := - rangeDis
	endX := rangeDis
	startY := - rangeDis
	endY := rangeDis
	points := make([]int, 0)
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			//if x == 0 && y == 0 {
			//	continue
			//}
			points = append(points, x, y)
		}
	}
	for _, v := range pb.SCENEDIR_ARRAY {

		offset[v] = make([]int, 0)
		offset[v] = points
	}
	return offset
}
