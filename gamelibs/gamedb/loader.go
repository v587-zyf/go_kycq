package gamedb

import (
	"bufio"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sync"

	"time"

	"cqserver/golibs/logger"
	"github.com/tealeg/xlsx"
)

type fileInfo struct {
	fileName   string
	sheetInfos []sheetInfo
}

func (this *fileInfo) GetFileName() string {
	return this.fileName
}

func (this *fileInfo) GetSheetInfos() []sheetInfo {
	return this.sheetInfos
}

type sheetInfo struct {
	SheetName   string
	Initer      func(*GameDb, []interface{}) error
	ObjProptype interface{}
}

func GetFileInfos() []fileInfo {
	return fileInfos
}

func Load(basePath string) error {
	f, err := os.Stat(basePath)
	if err != nil {
		return err
	}
	isCheck := false
	InitGameDb(basePath)
	if f.IsDir() {
		if err := LoadExcel(basePath, "gamedb.dat"); err != nil {
			return err
		}
		isCheck = true
	} else {
		if err := LoadGob(basePath); err != nil {
			return err
		}
		basePath = filepath.Dir(basePath)
	}
	if err := LoatOther(basePath); err != nil {
		return err
	}
	//数据整理
	GetDb().Patch()
	//数据检查
	if isCheck {
		err = GetDb().Check()
		if err != nil {
			return err
		}
	}
	logger.Debug("加载gamedb成功")
	return nil
}

func Reload() {
	path := gameDb.gamedbPath
	if len(path) <= 0 {
		logger.Error("配置路径没找到！！！")
		return
	}
	if err := LoadGob(path); err != nil {
		logger.Error("重加载表配置错误：%v", err)
	}
	//数据整理
	GetDb().Patch()
}

func LoadExcel(basePath string, godname string) error {
	//如果存在dat文件,则先载入.然后对比修改时间,重新加载时间不一致的文件
	godfile := filepath.Join(basePath, godname)
	f, err := os.Stat(godfile)
	if err == nil && !f.IsDir() {
		err = gameDb.loadGob(godfile)
		if err != nil {
			fmt.Println("load gob file error ", err)
			InitGameDb(godfile)
		}
	}

	change, err := gameDb.load(filepath.Join(basePath, "excel"))
	if err != nil {
		return err
	}

	if change {
		gameDb.Ver = time.Now().Format("060102150405")
		gameDb.generateGob(godfile)
	}
	return nil
}

func LoadGob(filename string) error {

	err := gameDb.loadGob(filename)
	if err != nil {
		return err
	}
	return nil
}

func LoatOther(basePath string) error {

	//地图信息加载
	if err := gameDb.LoadAllScenes(basePath); err != nil {
		return err
	}
	////配置模型大小
	//if err := gameDb.LoadModelSizeConf(basePath); err != nil {
	//	return err
	//}
	////屏蔽字库
	if err := LoadSensitivePhrases(filepath.Join(basePath, "filtertext.txt")); err != nil {
		return errors.New(fmt.Sprintf("read sensitive file err: %s", err.Error()))
	}
	logger.Info("敏感词库加载成功")
	return nil
}

func checkUnique(fileName string, keys map[int64]bool, objV reflect.Value) error {
	for _, v := range []string{"Id", "Lvl"} {
		keyFieldV := objV.Elem().FieldByName(v)
		if keyFieldV.IsValid() {
			key := keyFieldV.Int()
			if keys[key] {
				return fmt.Errorf("表　%s 字段 %s, %d,%v 重复了", fileName, v, key, objV)
			}
			keys[key] = true
			break
		}
	}
	return nil

}

func arrayLoader(fieldName string) func(*GameDb, []interface{}) error {
	return func(gameDb *GameDb, objs []interface{}) error {
		fieldV := reflect.ValueOf(gameDb).Elem().FieldByName(fieldName)

		keys := make(map[int64]bool)
		if fieldV.Kind() == reflect.Slice {
			if fieldV.IsNil() || fieldV.Len() > 0 {
				fieldV.Set(reflect.MakeSlice(fieldV.Type(), 0, len(objs)))
			}
			for _, obj := range objs {
				objV := reflect.ValueOf(obj)
				fieldV.Set(reflect.Append(fieldV, objV))
				if err := checkUnique(fieldName, keys, objV); err != nil {
					return err
				}
			}
		} else if fieldV.Kind() == reflect.Array {
			for i, obj := range objs {
				objV := reflect.ValueOf(obj)
				fieldV.Index(i).Set(objV)
				if err := checkUnique(fieldName, keys, objV); err != nil {
					return err
				}
			}
		} else {
			return fmt.Errorf("field %s is not an array", fieldName)
		}
		return nil
	}
}

func mapLoader(fieldName string, keyFieldName string) func(*GameDb, []interface{}) error {
	return func(gameDb *GameDb, objs []interface{}) error {
		fieldV := reflect.ValueOf(gameDb).Elem().FieldByName(fieldName)
		if fieldV.Kind() != reflect.Map {
			return fmt.Errorf("field %s is not a map", fieldName)
		}

		if fieldV.IsNil() || fieldV.Len() > 0 {
			fieldV.Set(reflect.MakeMap(fieldV.Type()))
		}
		for _, obj := range objs {
			objV := reflect.ValueOf(obj)
			keyFieldV := objV.Elem().FieldByName(keyFieldName)
			if !keyFieldV.IsValid() {
				return fmt.Errorf("key field %s wrong filedV:%v, when setting %s", keyFieldName, fieldName)
			}
			if keyFieldV.Kind() != reflect.Int {
				logger.Warn(fmt.Sprintf("key field %s wrong filedV:%v, when setting %s", keyFieldName, fieldName))
				continue
			}
			if fieldV.MapIndex(keyFieldV).IsValid() {
				//return fmt.Errorf(" >%v<. The value of field >%s< in sheet >%s< is duplicate",
				return fmt.Errorf("表 %s 列 %s 值->%v 重复了",
					fieldName, keyFieldName, keyFieldV)
			}
			fieldV.SetMapIndex(keyFieldV, objV)
		}
		return nil
	}
}

func (this *GameDb) generateGob(fileName string) error {

	now := time.Now()
	defer func() {
		fmt.Println("generateGob use time", time.Since(now).Seconds())
	}()

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	enc := gob.NewEncoder(w)
	enc.Encode(this)
	return w.Flush()
}

func (this *GameDb) loadGob(fileName string) error {

	now := time.Now()
	defer func() {
		fmt.Println("loadGod use time", time.Since(now).Seconds())
	}()

	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	dec := gob.NewDecoder(r)
	return dec.Decode(&this)
}

func (this *GameDb) load(baseDir string) (bool, error) {
	now := time.Now()
	defer func() {
		fmt.Println("load all use time", time.Since(now).Seconds())
	}()

	var wg sync.WaitGroup
	var loadErr error
	var num int
	for i := range fileInfos {

		filename := filepath.Join(baseDir, fileInfos[i].fileName)
		finfo, err := os.Stat(filename)
		if err != nil {
			loadErr = errors.New("filename not found, " + err.Error())
			break
		}
		fmodtime := finfo.ModTime().UnixNano()
		if this.FileModTime[fileInfos[i].fileName] == fmodtime {
			logger.Info("文件 %v 未修改!", fileInfos[i].fileName)
			continue
		}
		this.FileModTime[fileInfos[i].fileName] = fmodtime
		num++

		wg.Add(1)
		go func(index int, filename string) {

			err := this.loadFile(filename, fileInfos[index].sheetInfos)
			if err != nil {
				fmt.Println("加载:", fileInfos[index].fileName, "error:", err)
				loadErr = err
			}
			logger.Info("加载完成: %v", fileInfos[index].fileName)
			wg.Done()
		}(i, filename)
	}

	wg.Wait()
	if loadErr != nil {
		return false, loadErr
	}
	logger.Info("共加载: %v 个文件", num)
	if num == 0 {
		return false, nil
	}
	fmt.Println("All gamedb rem cal is ok")
	return true, loadErr
}

func (this *GameDb) loadFile(filename string, sheetInfos []sheetInfo) error {

	xlsFile, err := xlsx.OpenFile(filename)
	if err != nil {
		return err
	}
	for _, sheetInfo := range sheetInfos {
		sheet, ok := xlsFile.Sheet[sheetInfo.SheetName]
		if !ok {
			return fmt.Errorf("no %s sheet found", sheetInfo.SheetName)
		}
		objProptype := reflect.New(reflect.TypeOf(sheetInfo.ObjProptype)).Interface()
		objs, err := ReadXlsxSheet(sheet, objProptype, 2, 1, nil) // read from 3rd line,2nd row
		if err != nil {
			return err
		}
		err = sheetInfo.Initer(this, objs)
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *GameDb) loadGameConf(objs []interface{}) error {
	gameConfs := make(map[string]*GameBaseCfg)
	for _, obj := range objs {
		game := obj.(*GameBaseCfg)
		if _, ok := gameConfs[game.Name]; ok {
			return errors.New(fmt.Sprintf("gameconf key:%d namd:%s 重复了", game.Id, game.Name))
		}
		gameConfs[game.Name] = game
	}
	return DecodeConfValues(GetConf(), gameConfs)
}

var ModelSizeconf map[string]map[string][]int

func (this *GameDb) LoadModelSizeConf(basePath string) error {
	filePath := filepath.Join(basePath, "jsons/modelsize.json")

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.Warn("jsonLoad.go load file %s, err:%v", filePath, err)
		return err
	}
	err = json.Unmarshal(file, &ModelSizeconf)
	if err != nil {
		logger.Warn("jsonLoad.go unmarshal file %s,err:%v", filePath, err)
		return err
	}

	return nil
}

func (this *GameDb) GetMonsterModelSizeConf() map[string][]int {
	return ModelSizeconf["monster"]
}

func (this *GameDb) GetRoleModelSizeConf() map[string][]int {
	return ModelSizeconf["role"]
}

func (this *GameDb) GetHeroModelSizeConf() map[string][]int {
	return ModelSizeconf["hero"]
}

func (this *GameDb) GetNpcModelSizeConf() map[string][]int {
	return ModelSizeconf["npc"]
}

func (this *GameDb) InitMaskwordConf() {
	//this.Maskwords = make(map[string]int, 1000)
	//this.MaskWordsHeader = make(map[rune][]string, 256)
	//for _, conf := range this.MaskwordConfs {
	//	this.Maskwords[conf.Name] = conf.Id
	//	r := []rune(conf.Name)
	//	if len(r) == 0 {
	//		continue
	//	}
	//	if _, ok := this.MaskWordsHeader[r[0]]; ok == false {
	//		this.MaskWordsHeader[r[0]] = make([]string, 0)
	//	}
	//	//屏蔽字的第一个字符为key,相同字符开头的屏蔽字列表为value
	//	this.MaskWordsHeader[r[0]] = append(this.MaskWordsHeader[r[0]], conf.Name)
	//}
}
