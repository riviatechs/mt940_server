package download

import (
	"context"
	"fmt"
	"strings"

	"github.com/riviatechs/mt940_server/csv"
	"github.com/riviatechs/mt940_server/models"
	"github.com/riviatechs/mt940_server/pdf"
)

func HandleDownload(ctx context.Context, input models.DownloadInput) (*string, error) {
	fileType, err := GetFileType(input.DownLoadType)
	if err != nil {
		return nil, err
	}

	if CheckPDF(*fileType) {
		return pdf.GeneratePDF(ctx, input)
	}

	if CheckCSV(*fileType) {
		return csv.DownloadCSV(ctx, input)
	}

	return nil, nil

}

func GetFileType(fileType *string) (*string, error) {
	f := fileType
	if f == nil {
		return nil, fmt.Errorf("accepted file types are pdf or csv")
	}

	if CheckPDF(*f) || CheckCSV(*f) {
		return f, nil
	} else {
		return nil, fmt.Errorf("accepted file types are pdf or csv")
	}
}

func CheckPDF(fileType string) bool {
	if strings.ToLower(fileType) == "pdf" {
		return true
	}

	return false
}

func CheckCSV(fileType string) bool {
	if strings.ToLower(fileType) == "csv" {
		return true
	}

	return false
}
