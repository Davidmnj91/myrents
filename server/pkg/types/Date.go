package types

import (
	"encoding/json"
	"strings"
	"time"
)

const dateLayout = "2006-01-02"

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(dateLayout, s)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d))
}

func NewDate(s string) (Date, error) {
	t, err := time.Parse(dateLayout, s)
	if err != nil {
		return Date{}, err
	}
	return Date(t), nil
}

func (d Date) Format() string {
	t := time.Time(d)
	return t.Format(dateLayout)
}
