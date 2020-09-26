package main

import (
	"log"

	"github.com/alexflint/go-arg"
)

// cliArgs defines the configuration that the CLI
// expects. By using a struct we can very easily
// aggregate them into an object and check what are
// the expected types.
type cliArgs struct {
	TOCSheet string `arg:"-t,help:full path to TOC Online Excel file"`
	BCPSheet string `arg:"-b,help:full path to Millennium BCP Excel file"`
}

var (
	// args is a reference to an instantiation of
	// the configuration that the CLI expects but
	// with some values set.
	args = &cliArgs{
		TOCSheet: "./testdata/extratocontas_514106654_25_09_2020.xlsx",
		BCPSheet: "./testdata/EXTMV120825856552.xls",
	}
)

func main() {
	// Parse to provided arguments
	arg.MustParse(args)

	// Get data from TOC sheet
	tocData, err := getTocTransactions(args.TOCSheet)
	if err != nil {
		log.Fatal(err)
	}

	bcpData, err := getBCPTransactions(args.BCPSheet)
	if err != nil {
		log.Fatal(err)
	}

	resultA, err := reconcile(*bcpData, *tocData)
	if err != nil {
		log.Fatal(err)
	}
	renderResult(resultA, "Reconcile TOC online (compared to BCP)")

	resultB, err := reconcile(*tocData, *bcpData)
	if err != nil {
		log.Fatal(err)
	}
	renderResult(resultB, "Reconcile Millennium BCP (compared to TOC Online)")
}
