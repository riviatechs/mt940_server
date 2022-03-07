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
	oauthToken, err := GenerateToken()
	logger.Info("ExchangeJWT", zap.String("TOKEN", *oauthToken))
	if err != nil {
		return nil, err
	}

	clientID := viper.GetString(AdobeClientID)
	clientSecret := viper.GetString(AdobeClientSecret)

	formData := url.Values{
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"jwt_token":     {*oauthToken},
	}

	resp, err := http.PostForm(viper.GetString(AdobeExchangeJWTURL), formData)
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
