package main

import (
	"fmt"

	"github.com/xlab/termtables"
)

// Render result
func renderMissingTrnResult(r []transaction, m string) {

	table := termtables.CreateTable()
	table.AddHeaders("Ref", "Amount", "Date", "Description")

	for i, reportline := range r {
		table.AddRow(i+1,
			reportline.Amount,
			reportline.Date.Format(layoutISO),
			reportline.Description)
	}

	fmt.Println()
	fmt.Println(m)
	fmt.Println(table.Render())
	fmt.Println()
}
