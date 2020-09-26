package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"
)

func getBCPTransactions(filename string) (tran *[]transaction, err error) {

	const bcpSheetHeaderRows = 6
	const bcpSheetAmountColumn = 7
	const bcpSheetDateColumn = 4
	const bcpSheetDescriptionColumn = 6

	var oneRecord transaction
	var allRecords []transaction
	var csvBody string

	//
	// We need to open the file and strip some
	// comments from the top first
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("Error opening file: " + err.Error())
	}
	// Strip N lines
	scanner := bufio.NewScanner(f)
	for i := 0; i <= bcpSheetHeaderRows; i++ {
		scanner.Scan()
	}
	// Now fill up the CSV data bucket
	for scanner.Scan() {
		csvBody += scanner.Text() + "\n"
	}

	// Configure the CSV reader
	reader := csv.NewReader(strings.NewReader(csvBody))
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1

	// Load the data using the csv reader
	csvData, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New("Error reading data from BCP sheet: " + err.Error())
	}

	for _, each := range csvData {
		// Hack: replace 1000 sep and swap commas for dots
		amount := strings.Replace(each[bcpSheetAmountColumn], ".", "", -1)
		amount = strings.Replace(amount, ",", ".", -1)
		oneRecord.Amount, err = strconv.ParseFloat(amount, 32)
		if err != nil {
			return nil, errors.New("Cannot convert amount: " + err.Error())
		}
		oneRecord.Date, err = time.Parse("2006/01/02", each[bcpSheetDateColumn])
		if err != nil {
			return nil, errors.New("Cannot convert date: " + err.Error())
		}
		oneRecord.Description = each[bcpSheetDescriptionColumn]
		allRecords = append(allRecords, oneRecord)
	}
	return &allRecords, nil
}
