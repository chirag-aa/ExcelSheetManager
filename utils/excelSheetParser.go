package utils

import (
	"errors"
	"io"

	"excelsheetmanager.com/models"

	"github.com/xuri/excelize/v2"
)

func isValidationSucessfull(file [][]string) bool {
	if len(file) <= 0 || len(file[0]) < Ecxel_Sheet_Columns {
		return false
	}
	return true
}

func ParseExcelSheet(file io.Reader) ([]models.Employee, error) {
	f, err := excelize.OpenReader(file)

	if err != nil {
		return nil, errors.New(Excel_Sheet_Parsing_Error)
	}
	var employeesData []models.Employee

	sheet := f.GetSheetName(0)
	rows, err := f.GetRows(sheet)

	if err != nil {
		return nil, errors.New(Excel_Sheet_Parsing_Error)
	}

	if !isValidationSucessfull(rows) {
		return nil, errors.New(Validation_Failed)
	}

	for _, row := range rows[1:] {
		if len(row) < 1 {
			continue
		}
		employeesData = append(employeesData, models.Employee{First_name: row[0],
			Last_name:    row[1],
			Company_name: row[2],
			Address:      row[3],
			City:         row[4],
			Country:      row[5],
			Postal:       row[6],
			Phone:        row[7],
			Email:        row[8],
			Web:          row[9],
		})
	}
	return employeesData, nil

}
