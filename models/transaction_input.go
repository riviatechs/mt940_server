package models

type TransInput struct {
	CustStmtMsgID string  `json:"custStmtMsgID"`
	Mark          string  `json:"mark"`
	DateY         string  `json:"dateY"`
	Currency      string  `json:"currency"`
	Amount        float64 `json:"amount"`
}
