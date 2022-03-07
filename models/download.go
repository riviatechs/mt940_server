package models

type TemplateInput struct {
	AccountNumber string
	AccountName   string
	Ob            string
	Cb            string
	StartDate     string
	EndDate       string
	Conf          []*ConfTemplate
	Fields        Fields
}

type Fields struct {
	TRF           bool
	Amount        bool
	AccountNumber bool
	AccountName   bool
	Date          bool
	Narrative     bool
	Mark          bool
}

type ConfTemplate struct {
	ID            string
	Currency      string
	PartyBName    string
	PartyBAccount string
	Amount        string
	DateTime      string
	Narrative     string
	PartyAName    string
	PartyAAccount string
	Mark          string
}
