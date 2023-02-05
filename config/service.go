package config

import (
	"github.com/DeniesKresna/danatest/models"
	"github.com/DeniesKresna/gohelper/utstring"
)

func New() *Config {
	var cfg Config

	cfg.Service = &models.Service{
		Name:      utstring.GetEnv(models.AppNameENV),
		Version:   utstring.GetEnv(models.AppVersionENV),
		Host:      utstring.GetEnv(models.AppHostENV),
		Port:      utstring.GetEnv(models.AppPortENV),
		Namespace: utstring.GetEnv(models.AppNamespaceENV),
	}

	return &cfg
}
