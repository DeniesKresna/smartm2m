package config

import (
	"github.com/DeniesKresna/danatest/models"
	"github.com/DeniesKresna/gobridge/sdb"
)

type Config struct {
	Service *models.Service
	DB      *sdb.DBInstance
}
