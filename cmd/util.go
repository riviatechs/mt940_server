package cmd

import (
	"fmt"

	"github.com/MalukiMuthusi/logger"
	"github.com/riviatechs/mt940_server/util"
	"github.com/spf13/viper"
)

func checkMustBeSet() {
	mustBeSet(util.ProjectID)

	// Database
	mustBeSet(util.DbUser)
	mustBeSet(util.DbName)
	mustBeSet(util.DbPwd)
	mustBeSet(util.DbHost)
	mustBeSet(util.DbPort)
	mustBeSet(util.DbConnectionName)
	mustBeSet(util.DbTimeZone)

	// Adobe
	mustBeSet(util.AdobeJWT)
	mustBeSet(util.AdobeClientID)
	mustBeSet(util.AdobeClientSecret)
	mustBeSet(util.AdobeExchangeJWTURL)
}

func mustBeSet(env string) {
	if viper.GetString(env) == "" {
		logger.Fatalf(fmt.Sprintf("%s is not set", env))
	}
}

// bind viper flags to the cobra command
func bind(flag string) {
	viper.BindPFlag(flag, RootCmd.Flags().Lookup(flag))
}
