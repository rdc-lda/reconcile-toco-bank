package main

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

func reconcile(a []transaction, b []transaction) (r []report, err error) {
	var oneRecord report
	var allRecords []report

	for _, atr := range a {
		aDate := atr.Date
		aAmount := atr.Amount
		aDescr := atr.Description

		found := false

		for _, btr := range b {
			if btr.Amount == aAmount {
				found = true
				break
			}
		}

		if found == false {
			oneRecord.Comment = "Not found"
			oneRecord.TransactionDetails.Amount = aAmount
			oneRecord.TransactionDetails.Date = aDate
			oneRecord.TransactionDetails.Description = aDescr
			allRecords = append(allRecords, oneRecord)
		}
	}

	return allRecords, err
}
