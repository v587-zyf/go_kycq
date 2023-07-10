package scene

import (
	"cqserver/protobuf/pb"
	"errors"
	"math"
)

var (
	EIGHT_DIR_OFFSET   = []int{-1, -1, 0, -1, 1, -1, 1, 0, 1, 1, 0, 1, -1, 1, -1, 0}
	POINT_ROUND_OFFSET = make([][]int, 20)
)

var (
	ErrNotWalkable = errors.New("not walkable")
	ErrInSceneData = errors.New("in scene data error")
	ErrInvalidMap  = errors.New("invalid scene map")
)

const (
	DIR_NONE         = -1
	DIR_BOTTOM       = pb.SCENEDIR_BOTTOM
	DIR_RIGHT_BOTTOM = pb.SCENEDIR_RIGHT_BOTTOM
	DIR_RIGHT        = pb.SCENEDIR_RIGHT
	DIR_RIGHT_TOP    = pb.SCENEDIR_RIGHT_TOP
	DIR_TOP          = pb.SCENEDIR_TOP
	DIR_LEFT_TOP     = pb.SCENEDIR_LEFT_TOP
	DIR_LEFT         = pb.SCENEDIR_LEFT
	DIR_LEFT_BOTTOM  = pb.SCENEDIR_LEFT_BOTTOM
)

func GetDirOffset(dir int) (int, int) {
	switch dir {
	case DIR_TOP:
		return 0, -1
	case DIR_RIGHT_TOP:
		return 1, -1
	case DIR_RIGHT:
		return 1, 0
	case DIR_RIGHT_BOTTOM:
		return 1, 1
	case DIR_BOTTOM:
		return 0, 1
	case DIR_LEFT_BOTTOM:
		return -1, 1
	case DIR_LEFT:
		return -1, 0
	case DIR_LEFT_TOP:
		return -1, -1
	}
	return 0, 0
}

func GetLeftBottomOffset(dir int, left, back int) (int, int) {

	newPointX, newPointY := 0, 0
	if back > 0 {
		offsetX, offsetY := GetDirOffset((dir + 4) % 8)
		newPointX += offsetX * back
		newPointY += offsetY * back
	}
	if left > 0 {
		offsetX, offsetY := GetDirOffset((dir + 6) % 8)
		newPointX += offsetX * left
		newPointY += offsetY * left
	}
	return newPointX, newPointY
}

func GetDirLeftDir(dir int) int {

	leftDir := ((dir - 2) + 8) % 8
	return leftDir
}

func GetDirRightDir(dir int) int {

	rightDir := (dir + 2) % 8
	return rightDir
}

func GetBackDir(dir int) int {
	backDir := (dir + 4) % 8
	return backDir
}

func Distance(x1, y1, x2, y2 int) int {
	return int(math.Max(math.Abs(float64(x1-x2)), math.Abs(float64(y1-y2))))
}

func DistanceByPoint(point *Point, toPoint *Point) int {
	return Distance(point.X(), point.Y(), toPoint.X(), toPoint.Y())
}

func DistanceByPointMin(point *Point, toPoint *Point) int {
	return int(math.Min(math.Abs(float64(point.X()-toPoint.X())), math.Abs(float64(point.Y()-toPoint.Y()))))
}

func GetFaceDirByPoint(point, targetPoint *Point) int {
	dir := GetDir(int(point.X()), int(point.Y()), int(targetPoint.X()), int(targetPoint.Y()))
	if dir == DIR_NONE {
		dir = DIR_TOP
	}
	return dir
}

////获得朝向目标点的面向
//func GetFaceDir(baseX, baseY, targetX, targetY int) int {
//	if targetX == baseX {
//		if targetY > baseY {
//			return DIR_BOTTOM
//		} else {
//			return DIR_TOP
//		}
//	} else if targetY == baseY {
//		if targetX > baseX {
//			return DIR_RIGHT
//		} else {
//			return DIR_LEFT
//		}
//	}
//	absDeltaX := math.Abs(float64(targetX - baseX))
//	absDeltaY := math.Abs(float64(targetY - baseY))
//	ratio := absDeltaX / absDeltaY
//	baseDir := DIR_BOTTOM
//	AngleTan := []float64{0, 0.414, 1, 2.414}
//	for i := 3; i >= 0; i-- {
//		if ratio > AngleTan[i] {
//			if i == 0 {
//				baseDir = DIR_BOTTOM
//			} else if i == 1 || i == 2 {
//				baseDir = DIR_RIGHT_BOTTOM
//			} else if i == 3 {
//				baseDir = DIR_RIGHT
//			}
//			break
//		}
//	}
//	if targetX > baseX {
//		if targetY > baseY {
//			return baseDir
//		} else {
//			return 4 - baseDir
//		}
//
//	} else {
//		if targetY > baseY {
//			return (8 - baseDir) % 8
//		} else {
//			return baseDir + 4
//		}
//	}
//}

func getangle(startx int, starty int, endx int, endy int, isradian bool) int {

	disX := endx - startx
	disY := endy - starty
	angle := math.Atan2(float64(disY), float64(disX))
	if !isradian {
		angle = angle * 180 / math.Pi
	}
	return int(angle)
}

func getForwardByPoints(fx int, fy int, tox int, toy int) int {
	todir := 0
	angle := getangle(fx, fy, tox, toy, false)
	if angle < 0 {
		angle = 360 + angle
	}

	if angle > 335 || angle < 25 {
		todir = pb.SCENEDIR_RIGHT

	} else if angle > 290 {
		todir = pb.SCENEDIR_RIGHT_TOP
	} else if angle > 245 {
		todir = pb.SCENEDIR_TOP
	} else if angle > 200 {
		todir = pb.SCENEDIR_LEFT_TOP
	} else if angle > 155 {
		todir = pb.SCENEDIR_LEFT
	} else if angle > 110 {
		todir = pb.SCENEDIR_LEFT_BOTTOM
	} else if angle > 65 {
		todir = pb.SCENEDIR_BOTTOM
	} else {
		todir = pb.SCENEDIR_RIGHT_BOTTOM
	}
	return todir

}

func GetDir(fx, fy, tx, ty int) int {

	todir := 0
	angle := getangle(fx, fy, tx, ty, false)
	if angle < 0 {
		angle = 360 + angle
	}

	if angle > 335 || angle < 25 {
		todir = pb.SCENEDIR_RIGHT

	} else if angle > 290 {
		todir = pb.SCENEDIR_RIGHT_TOP
	} else if angle > 245 {
		todir = pb.SCENEDIR_TOP
	} else if angle > 200 {
		todir = pb.SCENEDIR_LEFT_TOP
	} else if angle > 155 {
		todir = pb.SCENEDIR_LEFT
	} else if angle > 110 {
		todir = pb.SCENEDIR_LEFT_BOTTOM
	} else if angle > 65 {
		todir = pb.SCENEDIR_BOTTOM
	} else {
		todir = pb.SCENEDIR_RIGHT_BOTTOM
	}
	return todir
}

func RandomDropPoint(scene *Scene, x, y int) []*Point {

	rangeDis := 1
	for rangeDis < 10 {

		offset := getPointRoundOffset(rangeDis)

		gridList := make([]*Point, 0)
		for i, l := 0, len(offset); i < l; i += 2 {
			offx := x + offset[i]
			offy := y + offset[i+1]

			point := scene.GetPointByXY(offx, offy)
			if point == nil {
				continue
			}

			if !point.IsBlock() && point.GetSceneObjsNum() == 0 {
				gridList = append(gridList, point)
			}
		}

		if len(gridList) > 0 {
			return gridList
		}
		rangeDis++
	}
	return nil
}

func getPointRoundOffset(rangeDis int) []int {

	if rangeDis >= len(POINT_ROUND_OFFSET) { //数组超长，自动进行扩展
		for i := len(POINT_ROUND_OFFSET); i < rangeDis; i++ {
			POINT_ROUND_OFFSET = append(POINT_ROUND_OFFSET, nil)
		}
	}

	ret := POINT_ROUND_OFFSET[rangeDis]
	if ret != nil {
		return ret
	}

	// 0 = 2 * (0 * 2 + 1) * (0 * 2 + 1) = 2; 0 = 1
	// 1 = 2 * (1 * 2 + 1) * (1 * 2 + 1) = 18; 1 = 9
	// 2 = 2 * (2 * 2 + 1) * (2 * 2 + 1) = 50 2 = 25
	// 3 = 2 * (3 * 2 + 1) * (3 * 2 + 1) = 98 3 = 49
	ret = make([]int, 2*(rangeDis*2+1)*(rangeDis*2+1))
	i := 0

	// 1 = 0  到  0
	// 2 = -1 到  1
	// 3 = -2 到   2
	for col := 0 - rangeDis; col <= rangeDis; col++ {
		for row := 0 - rangeDis; row <= rangeDis; row++ {
			ret[i] = col
			i++
			ret[i] = row
			i++
		}
	}

	POINT_ROUND_OFFSET[rangeDis] = ret
	return ret
}

func IsWithinRadius(sourceX, sourceY, targetX, targetY, radius int) bool {
	dx := targetX - sourceX
	dy := targetY - sourceY
	if dx > radius || dy > radius {
		return false
	}
	return true
}

func NextDirPoint(point *Point, dir int) *Point {

	if point == nil {
		//System.out.println("格子为空，找不到");
		return nil
	}
	if dir == DIR_NONE {
		//System.out.println("方向为none，找不到");
		return nil
	}
	nears := point.GetNearPoints()
	ret := nears[dir]
	if ret == nil {
		//System.out.println("八叉树没有数据，找不到");
		return nil
	}

	if ret.IsBlock() {
		//System.out.println("格子为主档，找不到");
		return nil
	}

	return ret
}
