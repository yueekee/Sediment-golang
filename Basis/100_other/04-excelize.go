package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
)

func main() {
	f, err := excelize.OpenFile("test_write.xlsx")
	if err != nil {
		log.Println("err:", err)
	}

Label:
	rows := f.GetRows("Sheet1")
	for i := 1; i < len(rows); i++ {
		rows = f.GetRows("Sheet1")
		fmt.Println("年龄：", rows[i][1])
		if rows[i][1] == "" {
			continue
		}
		fmt.Println("----", i)
		testRemoveRow(f, i)
		err = f.Save()
		if err != nil {
			log.Println("save err:", err)
		}
		i--
		goto Label
	}

}

func testRemoveRow(f *excelize.File, row int) {
	f.RemoveRow("Sheet1", row)
	fmt.Println("remove row:", row)
}

func create() {
	xlsx := excelize.NewFile()

	index := xlsx.NewSheet("Sheet1")
	xlsx.SetCellValue("Sheet1", "A1", "姓名")
	xlsx.SetCellValue("Sheet1", "B1", "年龄")
	xlsx.SetCellValue("Sheet1", "A2", "狗子")
	xlsx.SetCellValue("Sheet1", "B2", "18")
	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("test_write.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
