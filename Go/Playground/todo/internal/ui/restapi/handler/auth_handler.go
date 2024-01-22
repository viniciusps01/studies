package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	usecase "github.com/viniciusps01/todo/internal/feature/auth/use_case"

	"github.com/viniciusps01/todo/pkg/apperrors"
	"github.com/viniciusps01/todo/pkg/security"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateUserInputDTO

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		SendAPIError(w, apperrors.BadRequestError{Message: "invalid data"})
		return
	}

	u := usecase.NewCreateUserUseCase(appConfig.RepositoryProvider.AuthRepository)

	out, err := u.Exec(input)

	if err != nil {
		SendAPIError(w, err)
		return
	}

	json, err := json.Marshal(out)

	if err != nil {
		SendAPIError(w, err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func ReadUserHandler(w http.ResponseWriter, r *http.Request) {
	var i usecase.ReadUserInputDTO

	authUser := getAuthUser(r.Context())

	i.ID = authUser.ID

	u := usecase.NewReadUserUseCase(appConfig.RepositoryProvider.AuthRepository)

	out, err := u.Exec(i)

	if err != nil {
		SendAPIError(w, err)
		return
	}

	json, err := json.Marshal(out)

	if err != nil {
		SendAPIError(w, apperrors.BadRequestError{
			Message: "Failed to parse response data",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func ReadAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	var i usecase.ReadAllUsersInputDTO

	authUser := getAuthUser(r.Context())

	if !authUser.HasPermission(security.GetAllUsersPermissionCode) {
		SendAPIError(w, apperrors.AuthorizationError{
			Message: "Authorization error",
		})
		return
	}

	if limit, err := strconv.Atoi(r.URL.Query().Get("limit")); err == nil {
		i.Limit = &limit
	}

	if offset, err := strconv.Atoi(r.URL.Query().Get("offset")); err == nil {
		i.Offset = &offset
	}

	u := usecase.NewReadAllUsersUseCase(appConfig.RepositoryProvider.AuthRepository)

	out, err := u.Exec(i)

	if err != nil {
		SendAPIError(w, err)
		return
	}

	json, err := json.Marshal(out)

	if err != nil {
		SendAPIError(w, apperrors.BadRequestError{
			Message: "Failed to parse response data",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func AuthenticateUserHandler(w http.ResponseWriter, r *http.Request) {
	var i usecase.AuthenticateInputDTO

	err := json.NewDecoder(r.Body).Decode(&i)

	if err != nil {
		SendAPIError(w, apperrors.BadRequestError{
			Message: http.StatusText(http.StatusBadRequest),
		})
		return
	}

	u := usecase.NewAuthenticateUseCase(appConfig.RepositoryProvider.AuthRepository)

	token, err := u.Exec(i)

	if err != nil {
		SendAPIError(w, err)
		return
	}

	json, err := json.Marshal(token)

	if err != nil {
		SendAPIError(w, apperrors.InternalServerError{
			Message: "failed to decode output" + err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getAuthUser(r.Context())

	u := usecase.NewDeleteUserUseCase(appConfig.RepositoryProvider.AuthRepository)

	_, err := u.Exec(usecase.DeleteUserInputDTO{ID: user.ID})

	fmt.Println(err)

	if err != nil {
		SendAPIError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
