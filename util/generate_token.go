package util

import (
	"io/ioutil"
	"time"

	"github.com/MalukiMuthusi/logger"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func GenerateToken() (*string, error) {
	now := time.Now()

	var after30Minutes time.Time
	if now.Minute() <= 45 {
		after30Minutes = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute()+15, now.Second(), now.Nanosecond(), time.UTC)

	} else if now.Minute() > 45 {
		after30Minutes = time.Date(now.Year(), now.Month(), now.Day(), now.Hour()+1, now.Minute(), now.Second(), now.Nanosecond(), time.UTC)
	}

	iss := viper.GetString(AdobeISS)
	sub := viper.GetString(AdobeSUB)
	aud := viper.GetString(AdobeAUD)
	claims := jwt.MapClaims{
		"exp": after30Minutes.Unix(),
		"iss": iss,
		"sub": sub,
		"aud": aud,
		"https://ims-na1.adobelogin.com/s/ent_documentcloud_sdk": true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	var pFile []byte
	if viper.GetBool(DbHostedCloud) {
		p, err := ioutil.ReadFile("/out/private.key")
		if err != nil {
			logger.Info("GenerateToken", zap.Error(err))
			return nil, err
		}
		pFile = p
	} else {
		p, err := ioutil.ReadFile("private.key")
		if err != nil {
			logger.Info("GenerateToken", zap.Error(err))
			return nil, err
		}
		pFile = p
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(pFile)
	if err != nil {
		logger.Info("GenerateToken", zap.Error(err))
		return nil, err
	}

	tokenString, err := token.SignedString(key)
	if err != nil {
		logger.Info("GenerateToken", zap.Error(err))
		return nil, err
	}

	return &tokenString, nil
}
