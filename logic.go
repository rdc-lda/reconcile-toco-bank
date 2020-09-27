package main

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

// Returns those transactions which cannot be matched against "a"
func reconcileAmounts(a []transaction, b []transaction) (r []transaction, err error) {
	var oneRecord transaction
	var allRecords []transaction

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
			oneRecord.Amount = aAmount
			oneRecord.Date = aDate
			oneRecord.Description = aDescr
			allRecords = append(allRecords, oneRecord)
		}
	}

	return allRecords, err
}
