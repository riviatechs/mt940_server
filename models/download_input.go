package models

type DownloadInput struct {
	DownLoadType *string
	Fields       *FieldsInput
	Filters      *FilterInput
}

type FieldsInput struct {
	TRF           *string
	Amount        *string
	AccountNumber *string
	AccountName   *string
	Date          *string
	Narrative     *string
}
