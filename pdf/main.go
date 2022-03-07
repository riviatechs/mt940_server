package pdf

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/MalukiMuthusi/logger"
	"github.com/riviatechs/mt940_server/db"
	"github.com/riviatechs/mt940_server/models"
	"github.com/riviatechs/mt940_server/util"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func GeneratePDF(ctx context.Context, input models.DownloadInput) (*string, error) {
	t, err := template.New("statements").Parse(TPL)
	if err != nil {
		logger.Info("GeneratePDF", zap.Error(err))
		return nil, err
	}

	stmts, err := db.StatementsFiltered(ctx, input.Filters)
	if err != nil {
		return nil, err
	}

	fields := CreateFields(*input.Fields)

	confs := ConvertToConfTemplate(stmts)

	templateInput := models.TemplateInput{
		AccountNumber: "01234567891",
		AccountName:   "Maluki Muthusi",
		Ob:            "200000.00",
		Cb:            "20000000.00",
		StartDate:     "04/12/2021",
		EndDate:       "04/02/2022",
		Fields:        fields,
		Conf:          confs,
	}

	var b bytes.Buffer
	err = t.Execute(&b, templateInput)
	if err != nil {
		logger.Info("GeneratePDF", zap.Error(err))
		return nil, err
	}

	return HTMLToPDF(&b)
}

func ConvertToConfTemplate(input []*models.Confirmation) []*models.ConfTemplate {
	var confsTemplate []*models.ConfTemplate

	for _, v := range input {
		conf := CreateConfTemplate(*v)
		confsTemplate = append(confsTemplate, &conf)
	}

	return confsTemplate
}

func CreateConfTemplate(input models.Confirmation) models.ConfTemplate {
	id := strconv.FormatUint(uint64(input.ID), 10)
	amount := strconv.FormatFloat(float64(input.Amount), 'f', -1, 32)
	dateTime := fmt.Sprintf("%d/%d/%d", input.DateTime.Day(), input.DateTime.Month(), input.DateTime.Year())
	var mark string
	if strings.ToLower(input.Mark) == "c" {
		mark = "Credit"
	} else {
		mark = "Debit"
	}
	c := models.ConfTemplate{
		ID:            id,
		Currency:      input.Currency,
		PartyBName:    input.PartyBName,
		PartyBAccount: input.PartyBAccount,
		Amount:        amount,
		DateTime:      dateTime,
		Narrative:     input.Narrative,
		Mark:          mark,
	}

	return c
}

func CreateFields(fields models.FieldsInput) models.Fields {
	newFields := &models.Fields{}

	if fields.TRF != nil {
		newFields.TRF = true
	}
	if fields.Date != nil {
		newFields.Date = true
	}
	if fields.AccountName != nil {
		newFields.AccountName = true
	}
	if fields.AccountNumber != nil {
		newFields.AccountNumber = true
	}
	if fields.Amount != nil {
		newFields.Amount = true
	}
	if fields.Narrative != nil {
		newFields.Narrative = true
	}
	if fields.Mark != nil {
		newFields.Mark = true
	}

	return *newFields
}

func HTMLToPDF(file io.ReadWriter) (*string, error) {
	token, err := util.ExchangeJWT()
	if err != nil {
		return nil, err
	}

	bearer := fmt.Sprintf("Bearer %s", *token)

	header := http.Header{}
	header.Add("Authorization", bearer)
	header.Add("Accept", "application/json, text/plain, */*")
	header.Add("x-api-key", viper.GetString(util.AdobeClientID))
	header.Add("Prefer", "respond-async,wait=0")

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body) // This writer writes to the body buffer

	part, err := writer.CreateFormFile("InputFile0", "InputFile0")
	if err != nil {
		logger.Info("HTMLToPDF", zap.Error(err))
		return nil, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		logger.Info("HTMLToPDF", zap.Error(err))
		return nil, err
	}

	err = writer.WriteField("contentAnalyzerRequests", util.ContentAnalyzerRequest)
	if err != nil {
		logger.Info("HTMLToPDF", zap.Error(err))
		return nil, err
	}

	contentType := writer.FormDataContentType()

	err = writer.Close()
	if err != nil {
		logger.Info("HTMLToPDF", zap.Error(err))
		return nil, err
	}

	r, err := http.NewRequest(http.MethodPost, viper.GetString(util.AdobeHTMLToPDFURL), body)
	if err != nil {
		logger.Info("HTMLToPDF", zap.Error(err))
		return nil, err
	}
	r.Header = header
	r.Header.Set("Content-Type", contentType)

	client := http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		logger.Info("HTMLToPDF", zap.Error(err))
		return nil, err
	}

	key := resp.Header.Get("x-request-id")
	if key == "" {
		return nil, nil
	}

	return &key, nil

}
