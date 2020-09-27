package main

import (
	"time"
)

type transaction struct {
	Amount      float64
	Date        time.Time
	Description string
}
