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

	// Get data from BCP CSV sheet
	bcpData, err := getBCPTransactions(args.BCPSheet)
	if err != nil {
		log.Fatal(err)
	}

	// Reconcile TOC against BCP
	missingTocTrns, err := reconcileAmounts(*bcpData, *tocData)
	if err != nil {
		log.Fatal(err)
	}

	// Reconcile BCP against TOC
	missingBCPTrns, err := reconcileAmounts(*tocData, *bcpData)
	if err != nil {
		log.Fatal(err)
	}

	// Output
	renderMissingTrnResult(missingTocTrns,
		"These amounts were not found TOC online (compared to BCP)")
	renderMissingTrnResult(missingBCPTrns,
		"These amounts were not found in BCP Millennium (compared to TOC)")
}
