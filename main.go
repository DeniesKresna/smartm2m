package main

import (
	"github.com/DeniesKresna/gohelper/utlog"
	"github.com/DeniesKresna/smartm2m/app"
	"github.com/DeniesKresna/smartm2m/config"
	"github.com/joho/godotenv"
)

type Stock struct {
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Availability int     `json:"availability"`
	IsActive     bool    `json:"is_active"`
}

var inventory = map[string]Stock{}

func main() {
	err := godotenv.Load()
	if err != nil {
		utlog.Error("Error loading .env file")
	}

	conf := config.New()

	err = conf.InitDB()
	if err != nil {
		utlog.Errorf("error while connecting DB. %+v", err)
		return
	}

	conf.InitNewDiscordLogger()

	// start Http server
	app := app.InitApp(conf)
	err = app.GateOpen()
	if err != nil {
		utlog.Errorf("error while open gate. %+v", err)
		return
	}
}
