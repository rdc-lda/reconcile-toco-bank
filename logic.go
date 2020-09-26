package main

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

func reconcile(a []transaction, b []transaction) (r []report, err error) {
	var oneRecord report
	var allRecords []report

	for _, banktr := range a {
		bankDate := banktr.Date
		bankAmount := banktr.Amount

		found := false

		for _, toctr := range b {
			if toctr.Amount == bankAmount && toctr.Date == bankDate {
				found = true
				break
			}
		}

		if found == false {
			oneRecord.Comment = "Not found"
			oneRecord.TransactionDetails.Amount = bankAmount
			oneRecord.TransactionDetails.Date = bankDate
			allRecords = append(allRecords, oneRecord)
		}
	}

	return allRecords, err
}
