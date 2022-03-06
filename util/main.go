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
	AdobeExchangeJWTURL = "ADOBE_JWT_EXCHANGE_URL"
	AdobeClientID       = "ADOBE_CLIENT_ID"
	AdobeClientSecret   = "ADOBE_CLIENT_SECRET"
	AdobeJWT            = "ADOBE_JWT"
	AdobeHTMLToPDFURL   = "ADOBE_HTML_TO_PDF_URL"
)

const (
	TimeFormat = "2006-01-02T15:04:05Z"
	DateFormat = "2006-01-02"
)

type TokenResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   uint   `json:"expires_in"`
}

func FormatDate(d string) string {
	t, err := time.Parse(time.RFC3339, d)
	if err != nil {
		return ""
	}

	return t.Format(TimeFormat)
}

const (
	ContentAnalyzerRequest = `{
		"cpf:inputs": {
			"documentIn": {
				"cpf:location": "InputFile0",
				"dc:format": "text/html"
			}
		},
		"cpf:engine": {
			"repo:assetId": "urn:aaid:cpf:Service-e2ee120a2b06427cb449592f5db967e7"
		},
		"cpf:outputs": {
			"documentOut": {
				"cpf:location": "multipartLabelOut",
				"dc:format": "application/pdf"
			}
		}
	}`
)
