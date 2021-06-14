package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	f := excelize.NewFile()
	f.SetSheetRow("Sheet1", "B6", &[]interface{}{"1", nil, 2})
	err := f.SaveAs("Book3.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func SaveExcelFile(savePath string, records [][]string) (err error) {
	f := excelize.NewFile()
	for i, record := range records {
		f.SetSheetRow("Sheet1", fmt.Sprintf("A%d", i+1), &record)
	}
	err = f.SaveAs(savePath)
	return err
}