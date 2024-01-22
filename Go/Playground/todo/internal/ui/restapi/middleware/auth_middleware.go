package middleware

import (
	"context"
	"net/http"
	"strings"

	usecase "github.com/viniciusps01/todo/internal/feature/auth/use_case"

	"github.com/viniciusps01/todo/pkg/apperrors"
)

const AuthUserKey = "auth_user"

func SendAPIError(w http.ResponseWriter, e error) {
	err := apperrors.HttpErrorFrom(e)
	http.Error(w, err.Message, err.Status)
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")

		if !strings.Contains(bearer, "Bearer") {
			SendAPIError(w, apperrors.AuthenticationError{
				Message: "Authentication failure",
			})
			return
		}

		token := strings.Replace(bearer, "Bearer ", "", 1)

		if token == "" {
			SendAPIError(w, apperrors.AuthenticationError{
				Message: "Authentication failure",
			})
			return

		}

		uc := usecase.NewCheckAuthUseCase(appConfig.RepositoryProvider.AuthRepository)

		u, err := uc.Exec(usecase.CheckAuthInputDTO{Token: token})

		if err != nil {
			SendAPIError(w, err)
			return
		}

		ctx := context.WithValue(r.Context(), AuthUserKey, u)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
