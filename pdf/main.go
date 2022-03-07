package pdf

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/MalukiMuthusi/logger"
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

	var b bytes.Buffer
	err = t.Execute(&b, "Maluki Muthusi")
	if err != nil {
		logger.Info("GeneratePDF", zap.Error(err))
		return nil, err
	}

	return HTMLToPDF(&b)
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
