package gamedb

import (
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

const CellWidth = 50
const CellHeight = 34

var log = logger.Get("default", true)

var AllScenes = make(map[int]*SceneConf)

type SceneConfJson struct {
	Rc    string         `json:"RC"`    //地图高宽(行列数)
	MapWH string         `json:"mapWH"` //地图宽高（像素） 美术画的，世界为行列宽高
	Rect  map[int]string `json:"rect"`  //区域
	Npc   map[int]int    `json:"npc"`
	Data  string         `json:"data"` //地图点（1不可走，其他可走）
}

type SceneConf struct {
	Id          int
	ColNum      int
	RowNum      int
	Rect        map[int]map[int]bool
	Npc         map[int]int
	WalkableMap map[int32]*Point
}

//LoadAllScenes 加载所有的scene的基础配置
func (this *GameDb) LoadAllScenes(basePath string) error {
	logger.Info("加载地图信息开始......")
	defer func() {
		logger.Info("加载地图信息结束......")
	}()
	basePath = filepath.Join(basePath, "maps.json")
	var err error
	AllScenes, err = LoadScene(basePath)
	if err != nil {
		return err
	}
	return nil
}

type Point struct {
	X int32
	Y int32
}

//Load 加载一个map.json文件，缓存到内存中
func LoadScene(path string) (map[int]*SceneConf, error) {

	file, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Warn("map_util.go load file %s, err:%v", path, err)
		return nil, err
	}

	var m map[int]SceneConfJson
	if err := json.Unmarshal(file, &m); err != nil {
		logger.Warn("map_util.go unmarshal file %s,err:%v", path, err)
		return nil, err
	}
	scenes := make(map[int]*SceneConf)
	for k, v := range m {
		s := &SceneConf{Id: k}
		rc, err := common.IntSliceFromString(v.Rc, "|")
		if err != nil {
			return nil, err
		}
		s.ColNum = rc[1]
		s.RowNum = rc[0]

		s.Rect = make(map[int]map[int]bool)
		for rectIndex, rectArea := range v.Rect {
			s.Rect[rectIndex] = make(map[int]bool, 0)
			monsterRect, errM := common.IntSliceFromString(rectArea, "|")
			if errM != nil {
				return nil, errM
			}

			for _, v := range monsterRect {
				x, y := getRCbbyGridId(v, s.ColNum)
				pointIndex := GetPointIndex(int32(x), int32(y))
				s.Rect[rectIndex][pointIndex] = true
			}
		}

		s.Npc = make(map[int]int)
		for k, v := range v.Npc {
			x, y := getRCbbyGridId(v, s.ColNum)
			s.Npc[k] = GetPointIndex(int32(x), int32(y))
		}

		data, err := common.IntSliceFromString(v.Data, ",")
		if err != nil {
			return nil, err
		}

		s.WalkableMap = make(map[int32]*Point)
		girdId := 0
		for _, v := range data {
			girdType := v % 10
			girdNum := v / 10
			if girdType == 1 {
				girdId += girdNum
				continue
			}
			for i := 0; i < girdNum; i++ {
				x, y := getRCbbyGridId(girdId, s.ColNum)
				s.WalkableMap[int32(GetPointIndex(int32(x), int32(y)))] = &Point{X: int32(x), Y: int32(y)}
				girdId++
			}
		}
		scenes[k] = s
	}

	return scenes, nil
}

/**
 * 获取行列
 * @param gridId
 * @param colNum
 * @param result
 * @returns
 */
func getRCbbyGridId(gridId int, colNum int) (int, int) {

	y := gridId / colNum
	x := gridId % colNum
	return x, y
}

/**
 * 获取客户端格子索引
 * @param colNum
 * @param x，y
 * @returns
 */
func GetRCbGridIdByXY(colNum int, x, y int) int {

	return x + y*colNum
}

//Walkable 某个地图的某个点是否可走
func (this *SceneConf) Walkable(x, y int32) bool {
	return this.WalkableMap[x<<16|y] != nil
}

func (this *SceneConf) GetWalkable() map[int32]*Point {
	return this.WalkableMap
}

func (this *GameDb) GetSceneConf(id int) *SceneConf {
	return AllScenes[id]
}

func GetPointIndex(x, y int32) int {
	return int(x<<16 | y)
}

func GetXYByPointIndex(index int) (int, int) {
	x := index >> 16
	y := index ^ (x << 16)
	return x, y
}
