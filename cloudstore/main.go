package cloudstore

import (
	"context"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/storage"
	"github.com/riviatechs/mt940_server/util"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

func UploadFile(w io.Reader, fileName string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	wc := client.Bucket(viper.GetString(util.Bucket)).Object(fileName).NewWriter(ctx)
	if _, err = io.Copy(wc, w); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}

	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}

func CreateFileName(fileType string) string {
	id := uuid.NewV4().String()
	return fmt.Sprintf("%s.%s", id, fileType)
}

// create a cloud storage file name that will be publicly accessible
func CreateFileURL(fileName string) string {

	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", viper.GetString(util.Bucket), fileName)
}
