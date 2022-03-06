package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/MalukiMuthusi/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func ExchangeJWT() (*string, error) {
	u := viper.GetString(AdobeExchangeJWTURL)

	clientID := viper.GetString(AdobeClientID)
	clientSecret := viper.GetString(AdobeClientSecret)
	jwt := viper.GetString(AdobeJWT)

	formData := url.Values{
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"jwt_token":     {jwt},
	}

	resp, err := http.PostForm(u, formData)
	if err != nil {
		logger.Info("ExchangeJWT", zap.Error(err))
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Info("ExchangeJWT", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	var tokenResponse TokenResponse
	err = json.Unmarshal(b, &tokenResponse)
	if err != nil {
		logger.Info("ExchangeJWT", zap.Error(err))
		return nil, err
	}

	accessToken := tokenResponse.AccessToken
	return &accessToken, nil
}
