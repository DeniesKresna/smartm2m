package config

import (
	"github.com/DeniesKresna/gobridge/sdb"
	"github.com/DeniesKresna/myqgen2/qgen"
	"github.com/DeniesKresna/smartm2m/models"
	"github.com/go-playground/validator/v10"
)

type Config struct {
	Service        *models.Service
	DB             *sdb.DBInstance
	Q              *qgen.Obj
	Validator      *validator.Validate
	MessagerLogger IMessagerLogger
}
