package mpos

import (
	"time"
)

var (
	dateTimeFmt = "2006-01-02 15:04:05"
	dateTimeFmtTicked = `"` + dateTimeFmt + `"`
)

type TransactionTimestamp time.Time

// MarshalJSON override marshalJSON
func (tt *TransactionTimestamp) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(*tt).Format(dateTimeFmtTicked)), nil
}

// UnmarshalJSON override unmarshalJSON
func (tt *TransactionTimestamp) UnmarshalJSON(b []byte) (err error) {
	ts, err := time.Parse(dateTimeFmtTicked, string(b))
	if err != nil {
		return err
	}
	*tt = TransactionTimestamp(ts)
	return nil
}

// String returns it's string representation
func (tt *TransactionTimestamp) String() string {
	return time.Time(*tt).Format(dateTimeFmt)
}
