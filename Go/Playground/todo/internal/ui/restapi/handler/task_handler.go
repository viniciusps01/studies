package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	usecase "github.com/viniciusps01/internal/feature/task/use_case"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateTaskInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		status := http.StatusBadRequest
		http.Error(w, http.StatusText(status), status)
		return
	}

	authUser := getAuthUser(r.Context())
	input.UserID = authUser.ID

	u := usecase.NewCreateTask(appConfig.RepositoryProvider.TaskRepository)

	out, err := u.Exec(input)

	if err != nil {
		SendAPIError(w, err)
		return
	}

	json, err := json.Marshal(out)

	if err != nil {
		SendAPIError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(json)
}

func ReadTaskHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.ReadTaskInputDTO
	var err error

	input.ID, err = strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		status := http.StatusBadRequest
		http.Error(w, http.StatusText(status), status)
		return
	}

	authUser := getAuthUser(r.Context())
	input.UserID = authUser.ID

	out, err := usecase.NewReadTaskUseCase(appConfig.RepositoryProvider.TaskRepository).Exec(input)

	if err != nil {
		SendAPIError(w, err)
		return
	}

	json, err := json.Marshal(out)

	if err != nil {
		SendAPIError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.UpdateTaskInputDTO
	var err error

	err = json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		status := http.StatusBadRequest
		http.Error(w, http.StatusText(status), status)
		return
	}

	input.ID, err = strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		status := http.StatusBadRequest
		http.Error(w, http.StatusText(status), status)
		return
	}

	authUser := getAuthUser(r.Context())
	input.UserID = authUser.ID

	u := usecase.NewUpdateTaskUseCase(appConfig.RepositoryProvider.TaskRepository)

	out, err := u.Exec(input)

	if err != nil {
		SendAPIError(w, err)
		return
	}

	json, err := json.Marshal(out)

	if err != nil {
		status := http.StatusBadRequest
		http.Error(w, "failed to parse output data", status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.DeleteTaskInputDTO
	var err error

	input.ID, err = strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		status := http.StatusBadRequest
		http.Error(w, http.StatusText(status), status)
		return
	}

	authUser := getAuthUser(r.Context())
	input.UserID = authUser.ID

	u := usecase.NewDeleteTaskUseCase(appConfig.RepositoryProvider.TaskRepository)

	err = u.Exec(input)

	if err != nil {
		SendAPIError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func ReadAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	input := usecase.ReadAllTaskInputDTO{}

	authUser := getAuthUser(r.Context())
	input.UserID = authUser.ID

	if limit, err := strconv.Atoi((r.URL.Query().Get("limit"))); err == nil {
		input.Limit = &limit
	}

	if offset, err := strconv.Atoi((r.URL.Query().Get("offset"))); err == nil {
		input.Offset = &offset
	}

	u := usecase.NewReadAllTaskUseCase(appConfig.RepositoryProvider.TaskRepository)

	out, err := u.Exec(input)

	if err != nil {
		SendAPIError(w, err)
	}

	json, err := json.Marshal(out)

	if err != nil {
		status := http.StatusInternalServerError
		http.Error(w, "failed to parse output data", status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
