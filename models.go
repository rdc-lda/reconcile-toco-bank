package main

import (
	"time"
)

type transaction struct {
	Amount      float64
	Date        time.Time
	Description string
}

type report struct {
	TransactionDetails transaction
	Comment            string
}
