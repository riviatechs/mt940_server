package models

type SlInput struct {
	CustStmtMsgID string  `json:"custStmtMsgID"`
	ValueDate     *string `json:"valueDate"`
	EntryDate     *string `json:"entryDate"`
	Mark          string  `json:"mark"`
	FundsCode     *string `json:"fundsCode"`
	Amount        float32 `json:"amount"`
	Ttic          *string `json:"ttic"`
	RefOwner      *string `json:"refOwner"`
	RefAsi        *string `json:"refAsi"`
	Supp          *string `json:"supp"`
	Iao           *string `json:"iao"`
}
