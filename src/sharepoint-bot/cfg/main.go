package cfg

import (
	"github.com/caarlos0/env"
	"errors"
)

type config struct {
	ChromeUrl			string	`env:"CHROME_URL" envDefault:"http://localhost:9222/json"`
	SharepointUrl		string  `env:"SHAREPOINT_URL"`
	BotName				string 	`env:"BOT_NAME" envDefault:"sharepoint-bot"`
	TitleLink			string 	`env:"TITLE_LINK"`
	MainWebhookUrl		string	`env:"WEBHOOK_MAIN_URL"`
	DebugWebhookUrl		string	`env:"WEBHOOK_DEBUG_URL"`
	GoogleCredentials 	string 	`env:"GOOGLE_APPLICATION_CREDENTIALS_JSON"`
	GoogleStorageBucket	string	`env:"GOOGLE_STORAGE_BUCKET"`
	GoogleBucketObject	string 	`env:"GOOGLE_BUCKET_OBJECT"`
}

var c *config

func Load() error {
	c = new(config)
	err := env.Parse(c)
	if err != nil {
		return err
	}
	if c.MainWebhookUrl == `` {
		return errors.New(`env WEBHOOK_MAIN_URL must be specified`)
	}
	if c.DebugWebhookUrl == `` {
		return errors.New(`env WEBHOOK_DEBUG_URL must be specified`)
	}
	return nil
}

func Get() *config {
	return c
}