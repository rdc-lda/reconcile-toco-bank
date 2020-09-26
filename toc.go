package main

import (
	"errors"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func getTocTransactions(filename string) (tran *[]transaction, err error) {

	const tocSheetName = "Dados"
	const tocSheetHeadingFirstColumn = "Conta"
	const tocSheetAmountColumn = 12
	const tocSheetDateColumn = 13
	const tocSheetDescriptionColumn = 7

	startLineFound := false
	var rs []transaction

	s, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, errors.New("Error opening TOC Sheet: " + err.Error())
	}

	rows, err := s.Rows(tocSheetName)

	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			return nil, errors.New("Cannot parse rows: " + err.Error())
		}
		if startLineFound {
			t := transaction{}
			t.Amount, err = strconv.ParseFloat(row[tocSheetAmountColumn], 32)
			if err != nil {
				return nil, errors.New("Cannot convert amount: " + err.Error())
			}
			t.Date, err = time.Parse("2006-01-02", row[tocSheetDateColumn])
			if err != nil {
				return nil, errors.New("Cannot convert date: " + err.Error())
			}
			t.Description = row[tocSheetDescriptionColumn]
			rs = append(rs, t)
		} else if len(row) != 0 {
			if row[0] == tocSheetHeadingFirstColumn {
				startLineFound = true
			}
		}
	}

	return &rs, nil
}
