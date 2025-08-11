package model

import (
	"fmt"
	"time"
)

type CustomDate time.Time

func (d CustomDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(d).Format("2006-01-02"))), nil
}

func (d *CustomDate) UnmarshalJSON(b []byte) error {
	str := string(b)
	t, err := time.Parse("\"2006-01-02\"", str)
	if err != nil {
		return err
	}
	*d = CustomDate(t)
	return nil
}

type Transaction struct {
	ID          int        `json:"id_transaction,omitempty"`
	Description string     `json:"description" validate:"required,max=50"`
	Date        CustomDate `json:"date" validate:"required"`
	Value       float64    `json:"value" validate:"required,gt=0"`
}
