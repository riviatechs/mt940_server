package models

type CustStmtMsgInput struct {
	Trn  string       `json:"trn"`
	Rr   *string      `json:"rr"`
	Ai   AiInput      `json:"ai"`
	Sn   string       `json:"sn"`
	Ob   TransInput   `json:"ob"`
	Sl   []SlInput    `json:"sl"`
	Cb   TransInput   `json:"cb"`
	Cab  *TransInput  `json:"cab"`
	Fwab []TransInput `json:"fwab"`
	Iao  *string      `json:"iao"`
}
