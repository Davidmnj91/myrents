package domain

import (
	"encoding/json"
	"strings"
	"time"
)

const birthDateLayout = "2006-01-02"

type BirthDate time.Time

func (d *BirthDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(birthDateLayout, s)
	if err != nil {
		return err
	}
	*d = BirthDate(t)
	return nil
}

func (d BirthDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d))
}

func NewBirthDate(s string) (BirthDate, error) {
	t, err := time.Parse(birthDateLayout, s)
	if err != nil {
		return BirthDate{}, err
	}
	return BirthDate(t), nil
}

func (d BirthDate) Format() string {
	t := time.Time(d)
	return t.Format(birthDateLayout)
}
