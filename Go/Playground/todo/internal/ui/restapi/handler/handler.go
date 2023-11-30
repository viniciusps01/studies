package handler

import (
	"context"
	"net/http"

	"github.com/viniciusps01/internal/config"
	usecase "github.com/viniciusps01/internal/feature/auth/use_case"
	"github.com/viniciusps01/internal/ui/restapi/middleware"

	"github.com/viniciusps01/pkg/apperrors"
)

var appConfig *config.AppConfig

func SetUp(config *config.AppConfig) {
	appConfig = config
}

func SendAPIError(w http.ResponseWriter, e error) {
	err := apperrors.HttpErrorFrom(e)
	http.Error(w, err.Message, err.Status)
}

func getAuthUser(c context.Context) usecase.AuthUser {
	return *c.Value(middleware.AuthUserKey).(*usecase.AuthUser)
}
