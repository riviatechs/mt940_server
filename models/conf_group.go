package models

import "time"

type ConfGroup struct {
	DateTime      time.Time
	Confirmations []*Confirmation
}

type GroupByDateTime []*ConfGroup

func (c GroupByDateTime) Len() int {
	return len(c)
}

func (c GroupByDateTime) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c GroupByDateTime) Less(i, j int) bool {
	return c[i].DateTime.Before(c[j].DateTime)
}
