package main

import (
	"fmt"

	"github.com/xlab/termtables"
)

// Render result
func renderResult(r []report, m string) {

	table := termtables.CreateTable()

	table.AddHeaders("Amount", "Date", "Comment")
	for _, reportline := range r {
		table.AddRow(reportline.TransactionDetails.Amount,
			reportline.TransactionDetails.Date.Format(layoutISO),
			reportline.Comment)
	}
	fmt.Println(m)
	fmt.Println(table.Render())
	fmt.Println()
}
