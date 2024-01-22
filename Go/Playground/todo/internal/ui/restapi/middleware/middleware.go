package middleware

import "github.com/viniciusps01/todo/internal/config"

var appConfig *config.AppConfig

func SetUp(config *config.AppConfig) {
	appConfig = config
}
