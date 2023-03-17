package createExcel

import (
	"encoding/json"
	"github.com/tealeg/xlsx"
	"os"
	"path"
	"strconv"
)

type ExportParam struct {
	SheetName string        `json:"sheet_name"`
	Titles    []string      `json:"titles"`
	Columns   []string      `json:"columns"`
	Rows      []interface{} `json:"rows"`
	ExportDir string        `json:"export_dir"`
	Filename  string        `json:"filename"`
}

type SheetParam struct {
	SheetName string        `json:"sheet_name"`
	Titles    []string      `json:"titles"`
	Columns   []string      `json:"columns"`
	Rows      []interface{} `json:"rows"`
	TitleBg   string        `json:"title_bg"`
}

type MultiExportParam struct {
	Sheets    []SheetParam `json:"sheets"`
	ExportDir string       `json:"export_dir"`
	Filename  string       `json:"filename"`
}

func MultiSheetExport(params MultiExportParam) error {
	reportDir := params.ExportDir
	file := xlsx.NewFile()
	// 循环添加sheet
	for _, oneSheet := range params.Sheets {
		// 列名key
		columns := oneSheet.Columns
		// 列名NAME
		titles := oneSheet.Titles
		sheet, err := file.AddSheet(oneSheet.SheetName)
		if err != nil {
			return err
		}
		// 新增行存放标题
		row := sheet.AddRow()
		var cell *xlsx.Cell
		for _, title := range titles {
			cell = row.AddCell()
			cell.Value = title
		}
		// 新增行存放数据内容
		for _, value := range oneSheet.Rows {
			row = sheet.AddRow()
			maps := InterfaceMapToMap(value)
			for _, val := range columns {
				tmpVal := maps[val]
				if tmpVal == nil {
					tmpVal = ""
				}
				cell = row.AddCell()
				cell.Value = Strval(tmpVal)
			}
		}
	}
	filename := params.Filename
	excelPath := reportDir
	os.MkdirAll(excelPath, os.ModePerm)
	fullPath := path.Join(excelPath, filename)
	err := file.Save(fullPath)
	return err
}

func InterfaceMapToMap(u interface{}) map[string]interface{} {
	var value map[string]interface{}
	tmpBytes, _ := json.Marshal(u)
	json.Unmarshal(tmpBytes, &value)
	return value
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}
