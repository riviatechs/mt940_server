package models

import (
	"fmt"
	"io"
	"strconv"
)

type FilterInput struct {
	Currency    *string
	Tt          *string
	Period      *PeriodInput
	AmountRange *AmountRangeInput
}

type PeriodInput struct {
	Start *string
	End   *string
	Date  *string
}

type AmountRangeInput struct {
	Lower  *float32
	Upper  *float32
	Amount *float32
}

type Float32 float32

func (f *Float32) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("must be a float")
	}

	f32, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return fmt.Errorf("must be a float")
	}

	*f = Float32(f32)
	return nil
}

func (f Float32) MarshalGQL(w io.Writer) {
	ff := float64(f)
	s := strconv.FormatFloat(ff, 'f', -1, 32)
	w.Write([]byte(s))
}
