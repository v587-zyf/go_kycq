package tool

import (
	"cqserver/golibs/common"
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"unicode"
)

/**
 * 将excel中的前四列转化为struct
 * 第一列字段类型		如 int
 * 第二列字段名称		如 显示顺序
 * 第三列字段名		如 id
 * 第四列s,c,all 	s表示服务端使用 c表示客户端使用 all表示都使用
 */
const (
	GAMEDB_TEMP = `
type GameDbBase struct {
	Ver         string
	FileModTime map[string]int64

	//NOTE 关于client的配置：
	//client:对象名,对象类型　，对象名要小写．
	// mapKey 即对应的我们的结构里的key, 要看具体的型中key是什么　，一段是大写的
	InitConf         *InitConf
	%s
}	
`

	GAMEDB_DATA_TEM = "\n    %ss\t\tmap[int]*%s"

	GAMEDB_BASE_FUN_TEMPLE = `
func Get%s( %s int) *%s {
	return gameDb.%s[%s]
}

func Rang%s(f func(conf *%s)bool){
	for _,v := range gameDb.%s{
		if !f(v){
			return
		}
	}
}
`

	LOADER_FILES_TEMP = `
var fileInfos = []fileInfo{

	fileInfo{"game.xlsx", []sheetInfo{
		{"game", (*GameDb).loadGameConf, GameBaseCfg{}},
	}},
	%s
}
`
	FILE_INFO_TEMP = `
	fileInfo{"%s", []sheetInfo{
			%s
	}},`
	SHEET_INFO_TEMP_MAP = `{SheetName: "%s", Initer: mapLoader("%s", "%s"), ObjProptype: %s{}},`
	sheetInfoByArr      = `{SheetName: "%s", Initer: arrayLoader("%s"), ObjProptype: %s{}},`

	GAME_CONF_TEMP = `
type InitConf struct {

%s

}
`
	GAME_FIELD_TEMP = "    %s            %s     `conf:\"%s\"  default:\"%s\"`\n"

	CHECK_ITEM_TEMP = " checker:\"item\""
)

var (
	lineNumber           = 4                                         // 每个工作表需要读取的行数
	structBegin          = "type %s struct {\n"                      // 结构体开始
	structValue          = "    %s %s\t`col:\"%s\" client:\"%s\"%s`" // 结构体的内容
	structValueForServer = "    %s %s\t`col:\"%s\"%s`"               // 服务端使用的结构体内容
	structValueForClient = "    %s %s\t`col:\"%s\"`"                 //客户端使用的结构体内容
	structRemarks        = "	 // %s"                              // 结构体备注
	structValueEnd       = "\n"                                      // 结构体内容结束
	structEnd            = "}\n\n"                                   // 结构体结束
	header               = "package %s\n\r"                          // 文件头
	typeMapping          = map[string]string{
		"number": "float64",
	}
)

type Generate struct {
	savePath   string            // 生成文件的保存路径
	allType    map[string]string // 文件当中的数据类型
	gamedbData string            //所有配置
	objsData   string            // objs表结构生成文件的内容
	loaderData string            //表加载
	gameConf   string            //game表配置
	allFuncs   string            //所有配置基础获取方法
}

func (this *Generate) genFiledType(alltype string) {

	alltypeSlice := common.NewStringSlice(alltype, ",")
	this.allType = make(map[string]string)
	for _, v := range alltypeSlice {
		this.allType[strings.ToLower(strings.TrimSpace(v))] = strings.TrimSpace(v)
	}
	fmt.Println("生成配置表的数据类型为：", this.allType)
}

// 读取excel
func (this *Generate) ReadExcel(readPath, savePath, allType string, generateClient bool, withoutExcel map[string]bool) error {
	if savePath == "" || allType == "" {
		return fmt.Errorf("ReadExcel|savePath or allType is nil")
	}
	this.genFiledType(allType)
	this.savePath = savePath
	files, err := ioutil.ReadDir(readPath)
	if err != nil {
		return fmt.Errorf("ReadExcel|ReadDir is err:%v", err)
	}
	loadFileInfos := ""
	allFileDatas := ""
	allFuncs := ""
	structName := make(map[string]string)
	for _, file := range files {
		fileSuffix := path.Ext(file.Name())
		if fileSuffix != ".xlsx" || file.Name() == "game.xlsx" || hasChineseOrDefault(file.Name()) || JudgeIndex(file.Name(), "~$") || withoutExcel[file.Name()] {
			continue
		}

		wb, err := xlsx.OpenFile(readPath + "\\" + file.Name())
		if err != nil {
			return fmt.Errorf("ReadExcel|xlsx.OpenFile is err :%v", err)
		}
		fileName := strings.TrimSuffix(file.Name(), fileSuffix)
		// 遍历工作表
		sheetInfos := ""
		for _, sheet := range wb.Sheets {
			if hasChineseOrDefault(sheet.Name) {
				continue
			}
			// 判断表格中内容的行数是否小于需要读取的行数
			if sheet.MaxRow < lineNumber {
				return fmt.Errorf("ReadExcel|sheet.MaxRow:%d < lineNumber:%d", sheet.MaxRow, lineNumber)
			}
			sheetStructName := this.getSheetStructName(fileName, sheet.Name)
			if structName[sheetStructName] != "" {
				return fmt.Errorf("Have same sheet name!sheetName:%s, fileName1:%s, fileName2:%s ", sheet.Name, structName[sheet.Name], file.Name())
			}
			structName[sheetStructName] = file.Name()
			sheetData := this.getSheetData(sheet)

			structData, serverUse := this.SplicingData(sheetData, sheetStructName, generateClient)
			if serverUse {
				this.objsData += structData
				sheetDatasName := sheetStructName + "s"
				sheetKey := FirstRuneToUpper(strings.TrimSpace(sheet.Rows[1].Cells[0].Value))
				if len(sheetKey) == 0 {
					return errors.New("主键不能为空")
				}
				sheetInfos += fmt.Sprintf(SHEET_INFO_TEMP_MAP, strings.TrimSpace(sheet.Name), sheetDatasName, sheetKey, sheetStructName)

				allFileDatas += fmt.Sprintf(GAMEDB_DATA_TEM, sheetStructName, sheetStructName)

				allFuncs += fmt.Sprintf(GAMEDB_BASE_FUN_TEMPLE, sheetStructName, sheetKey, sheetStructName, sheetDatasName, sheetKey, sheetDatasName, sheetStructName, sheetDatasName)
			}
		}
		if len(sheetInfos) > 0 {

			loadFileInfos += fmt.Sprintf(FILE_INFO_TEMP, file.Name(), sheetInfos)
		}
	}
	this.loaderData = fmt.Sprintf(LOADER_FILES_TEMP, loadFileInfos)
	this.gamedbData = fmt.Sprintf(GAMEDB_TEMP, allFileDatas)
	this.allFuncs = allFuncs
	if this.objsData == "" {
		return fmt.Errorf("ReadExcel|this.objsData is nil")
	}
	this.gameConf = this.genGameConf(readPath)
	err = this.WriteNewFile()
	if err != nil {
		return err
	}
	return nil
}

func (this *Generate) genGameConf(readPath string) string {
	wb, err := xlsx.OpenFile(readPath + "\\game.xlsx")
	if err != nil {
		fmt.Println(fmt.Sprintf("ReadExcel|xlsx.OpenFile is err :%v", err))
	}
	// 遍历工作表
	var confFiled string
	for _, sheet := range wb.Sheets {
		if sheet.Name != "game" {
			continue
		}
		if hasChineseOrDefault(sheet.Name) {
			continue
		}
		typeCell := -1
		nameCell := -1
		valueCell := -1

		for k, v := range sheet.Rows[1].Cells {
			if v.Value == "clinetType" {
				typeCell = k
			} else if v.Value == "name" {
				nameCell = k
			} else if v.Value == "value" {
				valueCell = k
			}
		}

		for k, v := range sheet.Rows {
			if k < lineNumber {
				continue
			}
			confFiled += fmt.Sprintf(GAME_FIELD_TEMP,
				FirstRuneToUpper(v.Cells[nameCell].Value),
				this.CheckType(v.Cells[typeCell].Value, "game"),
				strings.TrimSpace(v.Cells[nameCell].Value),
				strings.TrimSpace(v.Cells[valueCell].Value))
		}
	}
	return fmt.Sprintf(GAME_CONF_TEMP, confFiled)
}

func (this *Generate) getSheetStructName(fileName, sheetName string) string {
	return FirstRuneToUpper(fileName) + FirstRuneToUpper(sheetName) + "Cfg"
}

type FileObjStruct struct {
	Des         string //字段说明
	Filed       string //字段值
	FileType    string //字段类型
	FileUsetype string //字段应用
}

func (this *Generate) getSheetData(sheet *xlsx.Sheet) []*FileObjStruct {
	sheetData := make([]*FileObjStruct, 0)
	// 遍历列
	for i := 0; i < sheet.MaxCol; i++ {
		// 判断某一列的第二行是否为空
		if sheet.Cell(1, i).Value == "" {
			break
		}
		sheetData = append(sheetData, &FileObjStruct{
			Des:         strings.TrimSpace(sheet.Cell(0, i).Value),
			Filed:       strings.TrimSpace(sheet.Cell(1, i).Value),
			FileType:    strings.TrimSpace(sheet.Cell(2, i).Value),
			FileUsetype: strings.TrimSpace(sheet.Cell(3, i).Value),
		})
	}
	return sheetData
}

// 拼装struct
func (this *Generate) SplicingData(data []*FileObjStruct, structObj string, generateClient bool) (string, bool) {

	serverUse := false
	structData := fmt.Sprintf(structBegin, structObj)
	for _, value := range data {
		switch strings.TrimSpace(value.FileUsetype) {
		case "all", "a":
			dataType := this.CheckType(value.FileType, structObj)
			checkTag := this.checkTag(dataType)
			structData += fmt.Sprintf(structValue, FirstRuneToUpper(value.Filed), dataType, value.Filed, value.Filed, checkTag)
			if value.Des != "" {

				structData += fmt.Sprintf(structRemarks, strings.Replace(value.Des, "\n", "", -1))
			}
			structData += fmt.Sprintf(structValueEnd)
			serverUse = true
		case "s":
			dataType := this.CheckType(value.FileType, structObj)
			checkTag := this.checkTag(dataType)
			structData += fmt.Sprintf(structValueForServer, FirstRuneToUpper(value.Filed), dataType, value.Filed, checkTag)
			if value.Des != "" {
				structData += fmt.Sprintf(structRemarks, strings.Replace(value.Des, "\n", "", -1))
			}
			structData += fmt.Sprintf(structValueEnd)
			serverUse = true
		case "c":
			if !generateClient {
				continue
			}
			dataType := this.CheckType(value.FileType, structObj)
			structData += fmt.Sprintf(structValueForClient, FirstRuneToUpper(value.Filed), dataType, value.Filed)
			if value.Des != "" {
				structData += fmt.Sprintf(structRemarks, strings.Replace(value.Des, "\n", "", -1))
			}
			structData += fmt.Sprintf(structValueEnd)
			serverUse = true
		default:
			continue
		}
	}

	structData += structEnd
	return structData, serverUse
}

func (this *Generate) checkTag(dataType string) string {

	checkTag := ""
	if dataType == "ItemInfos" || dataType == "ItemInfo" {
		checkTag = CHECK_ITEM_TEMP
	}
	return checkTag
}

// 拼装好的struct写入新的文件
func (this *Generate) WriteNewFile() error {
	str := strings.Split(this.savePath, "\\")
	if len(str) == 0 {
		return fmt.Errorf("WriteNewFile|len(str) is 0")
	}
	header = fmt.Sprintf(header, str[len(str)-1])
	data := header + "\n" + this.loaderData +
		"\n" + this.gamedbData +
		"\n" + this.allFuncs +
		"\n" + this.objsData +
		"\n" + this.gameConf
	fw, err := os.OpenFile(this.savePath+"\\gamedbBase.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("WriteNewFile|OpenFile is err:%v", err)
	}
	defer fw.Close()
	_, err = fw.Write([]byte(data))
	if err != nil {
		return fmt.Errorf("WriteNewFile|Write is err:%v", err)
	}
	return nil
}

// 检测解析出来的字段类型是否符合要求
func (this *Generate) CheckType(dataType string, source string) string {
	filedType := this.allType[strings.ToLower(strings.TrimSpace(dataType))]
	if filedType != "" {
		return filedType
	}
	if typeMapping[dataType] != "" {
		return typeMapping[dataType]
	}
	panic(fmt.Sprintf("表结构：%v,字段类型错误：%v", source, dataType))
	return ""
}

func (this *Generate) SpecialFile(file *xlsx.File) ([][]string, error) {
	sheetData := make([][]string, 0)
	// 遍历工作表
	for _, sheet := range file.Sheets {
		if hasChineseOrDefault(sheet.Name) {
			continue
		}
		// 遍历行
		for k, row := range sheet.Rows {
			if k < lineNumber {
				continue
			}
			if len(row.Cells) < 5 {
				continue
			}
			if strings.TrimSpace(row.Cells[1].Value) == "" {
				continue
			}
			cellData := make([]string, 0)
			cellData = append(cellData, row.Cells[3].Value, row.Cells[2].Value, row.Cells[1].Value, "all")
			sheetData = append(sheetData, cellData)
		}
	}
	return sheetData, nil
}

// 判断是否存在汉字或者是否为默认的工作表
func hasChineseOrDefault(r string) bool {
	if JudgeIndex(r, "Sheet") {
		return true
	}
	for _, v := range []rune(r) {
		if unicode.Is(unicode.Han, v) {
			return true
		}
	}
	return false
}

// 字符串首字母转换成大写
func FirstRuneToUpper(str string) string {
	data := []byte(str)
	for k, v := range data {
		if k == 0 {
			first := []byte(strings.ToUpper(string(v)))
			newData := data[1:]
			data = append(first, newData...)
			break
		}
	}
	return string(data[:])
}

func JudgeIndex(str, subStr string) bool {
	if strings.Index(str, subStr) != -1 {
		return true
	}
	return false
}
