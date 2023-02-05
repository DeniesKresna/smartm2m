package config

import (
	"fmt"

	"github.com/DeniesKresna/danatest/models"
	"github.com/DeniesKresna/gobridge/sdb"
	"github.com/DeniesKresna/gohelper/utlog"
	"github.com/DeniesKresna/gohelper/utstring"
)

func (cfg *Config) InitDB() (err error) {
	var (
		dbDriver   = utstring.GetEnv(models.DBDriverENV, "mysql")
		dbUser     = utstring.GetEnv(models.DBUserENV)
		dbPassword = utstring.GetEnv(models.DBPasswordENV)
		dbHost     = utstring.GetEnv(models.DBHostENV)
		dbPort     = utstring.GetEnv(models.DBPortENV)
		dbName     = utstring.GetEnv(models.DBNameENV)
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	dBInstance, err := sdb.InitDB(dbDriver, dsn)
	if err != nil {
		utlog.Error(err)
		return
	}

	cfg.DB = dBInstance
	return
}
