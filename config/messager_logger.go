package config

import (
	"github.com/DeniesKresna/gobridge/serror"
	"github.com/DeniesKresna/gohelper/utlog"
	"github.com/DeniesKresna/gohelper/utstring"
	"github.com/DeniesKresna/smartm2m/models"
	"github.com/bwmarrin/discordgo"
)

var (
	APPLICATION_ID = utstring.GetEnv(models.EnvDiscordApplicationID)
	PUBLIC_KEY     = utstring.GetEnv(models.EnvDiscordPublicKey)
	BOT_CLIENT_ID  = utstring.GetEnv(models.EnvDiscordBotClientID)
	BOT_URL        = utstring.GetEnv(models.EnvDiscordBotUrl)
	TOKEN          = utstring.GetEnv(models.EnvDiscordToken)
	CHANNEL_ID     = utstring.GetEnv(models.EnvDiscordChannelID)
)

type DiscordConfig struct {
	ChannelID string `yaml:"channel_id"`
	Token     string `yaml:"token"`
}

func (cfg *Config) InitNewDiscordLogger() {
	discordLogger := &DiscordConfig{
		ChannelID: CHANNEL_ID,
		Token:     TOKEN,
	}
	cfg.MessagerLogger = discordLogger
}

func (d *DiscordConfig) SendLogToMessager(errData serror.SError) error {
	joinMessage := "Error on: " + errData.GetComment() + " -> " + errData.GetMessage() + " -> " + errData.GetErrorMessage()

	goBot, err := discordgo.New("Bot " + d.Token)
	if err != nil {
		utlog.Error("Fail create discord bot integration")
		return err
	}

	_, err = goBot.ChannelMessageSend(d.ChannelID, joinMessage)
	if err != nil {
		utlog.Error(err.Error())
		return err
	}
	return nil
}

type IMessagerLogger interface {
	SendLogToMessager(serror.SError) error
}
