package handler

import (
	"net/http"

	"github.com/viniciusps01/internal/config"
	"github.com/viniciusps01/pkg/apperrors"
)

var appConfig *config.AppConfig

func SetUpHandlers(config *config.AppConfig) {
	appConfig = config
}

func sendAPIError(w http.ResponseWriter, e error) {
	err := apperrors.HttpErrorFrom(e)
	http.Error(w, err.Message, err.Status)
}
