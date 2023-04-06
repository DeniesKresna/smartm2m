package models

// app init
const (
	AppNameENV      = "APP_NAME"
	AppVersionENV   = "APP_VERSION"
	AppHostENV      = "APP_HOST"
	AppPortENV      = "APP_PORT"
	AppNamespaceENV = "APP_NAMESPACE"
	AppENV          = "APP_ENV"
	AppApiSecret    = "API_SECRET"
)

// mysql conn
const (
	DBDriverENV   = "DB_DRIVER"
	DBHostENV     = "DB_HOST"
	DBPortENV     = "DB_PORT"
	DBAppPortENV  = "DB_APP_PORT"
	DBNameENV     = "DB_NAME"
	DSNENV        = "DB_DSN"
	DBUserENV     = "DB_USER"
	DBPasswordENV = "DB_PASSWORD"
	DBIsLogged    = "DB_IS_LOGGED"
)

// discord logger
const (
	EnvDiscordApplicationID = "DISCORD_APPLICATION_ID"
	EnvDiscordPublicKey     = "DISCORD_PUBLIC_KEY"
	EnvDiscordBotClientID   = "DISCORD_BOT_CLIENT_ID"
	EnvDiscordBotUrl        = "DISCORD_BOT_URL"
	EnvDiscordToken         = "DISCORD_TOKEN"
	EnvDiscordChannelID     = "DISCORD_CHANNEL_ID"
)

// aws s3
const (
	AWSS3Region     = "AWS_S3_REGION"
	AWSS3PrivateKey = "AWS_S3_PRIVATE_KEY"
	AWSS3PublicKey  = "AWS_S3_PUBLIC_KEY"
	AWSS3Bucket     = "AWS_S3_BUCKET"
)

// email
const (
	EmailServerEnv   = "EMAIL_SERVER"
	EmailPortEnv     = "EMAIL_SERVER_PORT"
	EmailAddressEnv  = "EMAIL_USER"
	EmailPasswordEnv = "EMAIL_PASSWORD"
)
