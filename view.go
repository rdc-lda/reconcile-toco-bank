package main

import (
	"fmt"

	"github.com/xlab/termtables"
)

// Render result
func renderResult(r []report, m string) {

	table := termtables.CreateTable()

	table.AddHeaders("Ref", "Amount", "Date", "Description", "Issue")
	for i, reportline := range r {
		table.AddRow(i+1,
			reportline.TransactionDetails.Amount,
			reportline.TransactionDetails.Date.Format(layoutISO),
			reportline.TransactionDetails.Description,
			reportline.Comment)
	}
	fmt.Println()
	fmt.Println(m)
	fmt.Println(table.Render())
	fmt.Println()
}
