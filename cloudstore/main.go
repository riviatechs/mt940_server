package cloudstore

import (
	"context"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/storage"
	"github.com/riviatechs/mt940_server/util"
	"github.com/spf13/viper"
)

type CloudStore struct{}

func (c *CloudStore) UploadFile(w io.Reader, fileName string) error {
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
