package models

type FilterInput struct {
	Currency    *string
	Tt          *string
	Period      *PeriodInput
	AmountRange *AmountRangeInput
	Search      *string
}

type PeriodInput struct {
	Start *string
	End   *string
	Date  *string
}

type AmountRangeInput struct {
	Lower  *string
	Upper  *string
	Amount *string
}
