package main

import (
	"bytes"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var baseFolderPath = "/Users/cccccccccchy/go/src/test/mytest/excel_test/"
var templateFileName = "员工导出模板.xlsx"
var templateFilePath = baseFolderPath + templateFileName
var startRow = 3
var content = [][]string{
	{"1", "1.1", "1.1.1"},
	{"2", "2.2", "2.2.2"},
}
var templateUrl = "https://ccccchy-test.obs.cn-east-3.myhuaweicloud.com/asset/员工导出模板.xlsx"

// maxCharCount 最多26个字符A-Z
const maxCharCount = 26

func main() {
	fmt.Println("start generator excel")

	resp, err := http.Get(templateUrl)
	if err != nil {
		fmt.Println("http.Get(templateUrl)")
		return
	}
	defer func() {
		fmt.Println("start generator done")
		_ = resp.Body.Close()
	}()

	excel, err := excelize.OpenReader(resp.Body)
	if err != nil {
		log.Err(err).Msg("excelize.OpenReader(resp.Body)")
		return
	}
	//excel, err := excelize.OpenFile(templateFilePath)
	//if err != nil {
	//	log.Err(err).Msg("")
	//	return
	//}

	//err = excel.SetCellDefault("Sheet1", "A4", "3")
	//if err != nil {
	//	log.Err(err).Msg("")
	//	return
	//}
	err = excel.SetSheetRow("Sheet1", "A4", &content[0])
	if err != nil {
		log.Err(err).Msg("excel.SetSheetRow(\"Sheet1\", \"A4\", &content[0])")
		return
	}

	var buffer bytes.Buffer
	err = excel.Write(&buffer)
	if err != nil {
		log.Err(err).Msg("excel.Write(&buffer)")
		return
	}

	saveFileName := fmt.Sprintf("%s-%d.xlsx", "员工导出模板", time.Now().Unix())
	saveFilePath := baseFolderPath + saveFileName
	err = ioutil.WriteFile(saveFilePath, buffer.Bytes(), 0644)
	if err != nil {
		log.Err(err).Msg("ioutil.WriteFile")
		return
	}

	//err = excel.SaveAs(saveFilePath)
	//if err != nil {
	//	log.Err(err).Msg("")
	//	return
	//}
	//
	fmt.Println(saveFilePath)

}

// getColumnName 生成列名
// Excel的列名规则是从A-Z往后排;超过Z以后用两个字母表示，比如AA,AB,AC;两个字母不够以后用三个字母表示，比如AAA,AAB,AAC
// 这里做数字到列名的映射：0 -> A, 1 -> B, 2 -> C
// maxColumnRowNameLen 表示名称框的最大长度，假设数据是10行，1000列，则最后一个名称框是J1000(如果有表头，则是J1001),是4位
// 这里根据 maxColumnRowNameLen 生成切片，后面生成名称框的时候可以复用这个切片，而无需扩容
func getColumnName(column, maxColumnRowNameLen int) []byte {
	const A = 'A'
	if column < maxCharCount {
		// 第一次就分配好切片的容量
		slice := make([]byte, 0, maxColumnRowNameLen)
		return append(slice, byte(A+column))
	} else {
		// 递归生成类似AA,AB,AAA,AAB这种形式的列名
		return append(getColumnName(column/maxCharCount-1, maxColumnRowNameLen), byte(A+column%maxCharCount))
	}
}

// getColumnRowName 生成名称框
// Excel的名称框是用A1,A2,B1,B2来表示的，这里需要传入前一步生成的列名切片，然后直接加上行索引来生成名称框，就无需每次分配内存
func getColumnRowName(columnName []byte, rowIndex int) (columnRowName string) {
	l := len(columnName)
	columnName = strconv.AppendInt(columnName, int64(rowIndex), 10)
	columnRowName = string(columnName)
	// 将列名恢复回去
	columnName = columnName[:l]
	return
}

type ExcelVal struct {
	Axis string
	Val  string
}

func convertExcelVal(startRow int, content [][]string) []ExcelVal {
	res := make([]ExcelVal, 0)
	for i, row := range content {
		for j, val := range row {
			res = append(res, f(startRow+i, j, val))
		}
	}
	return res
}

func init() {
	ColName = make(map[int][]byte, 0)
	Lock = sync.RWMutex{}
}

var ColName map[int][]byte
var Lock sync.RWMutex

func f(i, j int, val string) ExcelVal {
	colName := string(ff(j))
	return ExcelVal{
		Axis: colName + fmt.Sprintf("%d", i),
		Val:  val,
	}
}

func ff(i int) []byte {
	if i == 0 {
		Lock.Lock()
		ColName[i] = []byte("A")
		Lock.Unlock()
		return []byte("A")
	}

	if val, ok := ColName[i]; ok {
		return val
	}

	res := make([]byte, 0)
	preVal := ff(i - 1)
	fmt.Println("len(preVal)", len(preVal), string(preVal[len(preVal)-1]), string(preVal[len(preVal)-1]+1))
	t := preVal[len(preVal)-1]
	if t == []byte("Z")[0] {
		preVal[len(preVal)-1] = []byte("A")[0]
		res = append(preVal, []byte("A")...)
	} else {
		preVal[len(preVal)-1] = preVal[len(preVal)-1] + 1
		res = preVal
	}

	Lock.Lock()
	ColName[i] = res
	Lock.Unlock()

	return res
}
