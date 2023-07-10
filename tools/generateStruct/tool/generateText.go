package tool

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
	"strings"
)

const (
	LANGUAGE_TEMPLATE =
`package gamedb

var(
%s
)

var(
%s
)
`

	TEXT_ERR_TEMP        = "	%s        = initError(%s, \"%s\")\n"
	TEXT_CODE_CONST_TEMP = "	%s = codeTextSign(\"%s\",\"%s\")\n"
)

func GenGameText(readPath, savePath string) {
	file, err := xlsx.OpenFile(readPath + "\\gameText.xlsx")
	if err != nil {
		fmt.Println("ReadExcel|ReadDir is err:%v", err)
		return
	}

	errData := ""
	codeConst := ""
	for _, sheet := range file.Sheets {
		if sheet.Name != "errorText" && sheet.Name != "codeText" {
			continue
		}
		constNameCellIndex, chineseCellIndex := getsettingCellIndex(sheet.Rows[1])
		for k, row := range sheet.Rows {
			if k < lineNumber {
				continue
			}
			if strings.TrimSpace(row.Cells[0].Value) == "" {
				break
			}
			Id := strings.TrimSpace(row.Cells[0].Value)
			constName := strings.ToUpper(strings.TrimSpace(row.Cells[constNameCellIndex].Value))
			if sheet.Name == "errorText" {
				errData += fmt.Sprintf(TEXT_ERR_TEMP, constName, Id, row.Cells[chineseCellIndex])
			} else if sheet.Name == "codeText" {
				codeConst += fmt.Sprintf(TEXT_CODE_CONST_TEMP, constName, constName, row.Cells[chineseCellIndex])
			}
		}
	}

	lanauage := fmt.Sprintf(LANGUAGE_TEMPLATE, errData, codeConst)

	fw, err := os.OpenFile(savePath+"\\language.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("WriteNewFile|OpenFile is err:", err)
		return
	}
	defer fw.Close()
	_, err = fw.Write([]byte(lanauage))
	if err != nil {
		fmt.Println("WriteNewFile|Write is err:", err)
		return
	}
}

func getsettingCellIndex(row *xlsx.Row) (int, int) {

	constNameIndex := -1
	chineseCellIndex := -1
	for k, v := range row.Cells {
		value := strings.TrimSpace(v.Value)
		if value == "constName" {
			constNameIndex = k
		}
		if value == "chinese" {
			chineseCellIndex = k
		}
	}
	return constNameIndex, chineseCellIndex
}
