package util

import "time"

const (
	CusStmtMsgTable              = "cus_stmt_msgs"
	AiTable                      = "ais"
	ObTable                      = "obs"
	SlTable                      = "sls"
	CbTable                      = "cbs"
	CabTable                     = "cabs"
	FwabTable                    = "fwabs"
	StatementsConfirmationsTable = "statements_confirmations"
	StatementsTable              = "statements"
	ConfirmationsTable           = "confirmations"
)

const (
	DbUser           = "DB_USER"
	DbPwd            = "DB_PWD"
	DbName           = "DB_NAME"
	DbPort           = "DB_PORT"
	DbHost           = "DB_HOST"
	DbHostedCloud    = "DB_CLOUD"
	DbConnectionName = "DB_INSTANCE_CONNECTION_NAME"
	Bucket           = "BUCKET"
	DbTimeZone       = "DB_TIMEZONE"
)

const (
	ProjectID = "project_id"
	ProgName  = "mt940"
	Debug     = "debug"
	Wait      = "wait"
	Port      = "port"
)

const (
	TimeFormat = "2006-01-02T15:04:05Z"
	DateFormat = "2006-01-02"
)

func FormatDate(d string) string {
	t, err := time.Parse(time.RFC3339, d)
	if err != nil {
		return ""
	}

	return t.Format(TimeFormat)
}
