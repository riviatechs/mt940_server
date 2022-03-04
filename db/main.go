package db

import (
	"fmt"
	"os"

	"github.com/MalukiMuthusi/logger"
	"github.com/riviatechs/mt940_server/models"
	"github.com/riviatechs/mt940_server/util"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Start() {
	dbUser := viper.GetString(util.DbUser)
	dbPwd := viper.GetString(util.DbPwd)
	dbName := viper.GetString(util.DbName)
	dbPort := viper.GetString(util.DbPort)
	dbHost := viper.GetString(util.DbHost)
	dbTimeZone := viper.GetString(util.DbTimeZone)

	var dsn string
	if viper.GetBool(util.DbHostedCloud) {
		instanceConnectionName := viper.GetString(util.DbConnectionName)
		socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
		if !isSet {
			socketDir = "/cloudsql"
		}

		dsn = fmt.Sprintf("host=%s/%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", socketDir, instanceConnectionName, dbUser, dbPwd, dbName, dbPort, dbTimeZone)

	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", dbHost, dbUser, dbPwd, dbName, dbPort, dbTimeZone)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf("StartDB", zap.Error(err))
	}

	err = db.AutoMigrate(&models.CustStmtMsg{}, &models.Ai{}, &models.Ob{}, &models.Sl{}, &models.Sl{}, &models.Cb{}, &models.Cab{}, &models.Fwab{}, &models.Confirmation{}, &models.Statement{}, &models.StatementConfirmation{})
	if err != nil {
		logger.Fatalf("db.AutoMigrate", zap.Error(err))
	}

	Db = db
}

func GroupStmtLinesByDate(sl []*models.Sl) ([]*models.SlGroups, error) {
	sls := make(map[string][]*models.Sl)

	for _, v := range sl {
		if l, ok := sls[v.ValueDate.Format(util.TimeFormat)]; ok {
			s := append(l, v)
			sls[v.ValueDate.Format(util.TimeFormat)] = s
		} else {
			vv := sls[v.ValueDate.Format(util.TimeFormat)]
			vv = append(vv, v)
			sls[v.ValueDate.Format(util.TimeFormat)] = vv
		}
	}

	var slg []*models.SlGroups

	for k, v := range sls {
		slgb := models.SlGroups{ValueDate: k, Sls: v}
		slg = append(slg, &slgb)
	}

	return slg, nil
}
