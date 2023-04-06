package config

import (
	"fmt"

	"github.com/DeniesKresna/gobridge/sdb"
	"github.com/DeniesKresna/gohelper/utlog"
	"github.com/DeniesKresna/gohelper/utstring"
	"github.com/DeniesKresna/myqgen2/qgen"
	"github.com/DeniesKresna/smartm2m/models"
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

	// set query generator
	isLogDB := utstring.GetEnv(models.DBIsLogged, "FALSE")
	isLogged := isLogDB == "TRUE"
	q, err := qgen.InitObject(isLogged, models.User{}, models.Item{})
	if err != nil {
		return
	}
	cfg.Q = q

	return
}
