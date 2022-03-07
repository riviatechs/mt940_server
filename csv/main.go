package csv

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"strconv"

	"github.com/MalukiMuthusi/logger"
	"github.com/riviatechs/mt940_server/cloudstore"
	"github.com/riviatechs/mt940_server/db"
	"github.com/riviatechs/mt940_server/models"
	"go.uber.org/zap"
)

func DownloadCSV(ctx context.Context, input models.DownloadInput) (*string, error) {

	if input.DownLoadType == nil {
		return nil, fmt.Errorf("you must specify  download type")
	}

	if input.Fields == nil {
		return nil, fmt.Errorf("you must specify  download fields")
	}

	return WriteCSVInfo(ctx, input)
}

func WriteCSVInfo(ctx context.Context, input models.DownloadInput) (*string, error) {
	csvHeader := CreateCSVHeader(input.Fields)
	csvRows := GetRecords(ctx, input.Filters, *input.Fields)

	csvInfo := [][]string{}
	csvInfo = append(csvInfo, csvHeader)
	csvInfo = append(csvInfo, csvRows...)

	fileBuff := &bytes.Buffer{}

	writer := csv.NewWriter(fileBuff)
	err := writer.WriteAll(csvInfo)
	if err != nil {
		logger.Info("WriteCSVInfo", zap.Error(err))
		return nil, err
	}

	fileType := input.DownLoadType

	fileName := cloudstore.CreateFileName(*fileType)
	fileURL := cloudstore.CreateFileURL(fileName)

	go cloudstore.UploadFile(fileBuff, fileName)

	return &fileURL, nil

}

func AddToHeader(header []string, field *string, fieldName string) []string {
	if field != nil {
		header := append(header, fieldName)
		return header
	}
	return header
}

func CreateCSVHeader(fields *models.FieldsInput) []string {
	csvHeader := []string{}
	csvHeader = AddToHeader(csvHeader, fields.Date, "Date")
	csvHeader = AddToHeader(csvHeader, fields.TRF, "Transaction REF")
	csvHeader = AddToHeader(csvHeader, fields.AccountName, "Account Name")
	csvHeader = AddToHeader(csvHeader, fields.AccountNumber, "Account Number")
	csvHeader = AddToHeader(csvHeader, fields.Amount, "Amount")
	csvHeader = AddToHeader(csvHeader, fields.Narrative, "Narrative")

	return csvHeader
}

func GetRecords(ctx context.Context, input *models.FilterInput, fields models.FieldsInput) [][]string {
	stmts, err := db.StatementsFiltered(ctx, input)
	if err != nil {
		return nil
	}

	confGroups, err := db.GroupStmtsByDate(stmts)
	if err != nil {
		return nil
	}

	csvInfo := [][]string{}
	for _, confGroup := range confGroups {
		for _, stmt := range confGroup.Confirmations {
			csvRow := CreateCSVRow(*stmt, fields)
			csvInfo = append(csvInfo, csvRow)
		}
	}

	return csvInfo

}

func CreateCSVRow(conf models.Confirmation, fields models.FieldsInput) []string {
	csvRow := []string{}
	if fields.Date != nil {
		d := conf.DateTime
		date := fmt.Sprintf("%d/%d/%d", d.Year(), d.Month(), d.Day())
		csvRow = append(csvRow, date)
	}

	if fields.TRF != nil {
		trf := strconv.FormatUint(uint64(conf.ID), 10)
		csvRow = append(csvRow, trf)
	}

	if fields.AccountName != nil {
		csvRow = append(csvRow, conf.PartyBName)
	}

	if fields.AccountNumber != nil {
		csvRow = append(csvRow, conf.PartyBAccount)
	}

	if fields.Amount != nil {
		amount := strconv.FormatFloat(float64(conf.Amount), 'f', -1, 32)
		csvRow = append(csvRow, amount)
	}

	if fields.Narrative != nil {
		csvRow = append(csvRow, *fields.Narrative)
	}

	return csvRow
}
