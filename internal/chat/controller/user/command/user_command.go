package command

import (
	"encoding/json"
	"net/http"

	"media.hiway.chat/internal/chat/adapter/persistence"
	dto "media.hiway.chat/internal/chat/adapter/rest/dto/user"
	"media.hiway.chat/internal/chat/adapter/utils"
	"media.hiway.chat/internal/chat/domain/service/user"
)

func AddUser(writer http.ResponseWriter, request *http.Request) {
	db, err := persistence.InitPostgres()
	if err != nil {
		utils.NewErrorResponse(
			writer,
			http.StatusInternalServerError,
			"Internal server error. Failed to connect to the PostgreSQL")
		return
	}
	defer db.Close()

	repository := persistence.NewUserRepository(db)
	service := user.NewCreateUser(repository)

	var req dto.CreateUserRequest

	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		utils.NewErrorResponse(writer, http.StatusBadRequest, "invalid request body. "+err.Error())
		return
	}

	user, err := service.Execute(req.FirstName, req.LastName)
	if err != nil {
		utils.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}

	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(*user); err != nil {
		utils.NewErrorResponse(writer, http.StatusInternalServerError, "internal server error. "+err.Error())
		return
	}
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	db, err := persistence.InitPostgres()
	if err != nil {
		utils.NewErrorResponse(
			writer,
			http.StatusInternalServerError,
			"Internal server error. Failed to connect to the PostgreSQL")
		return
	}
	defer db.Close()

	repository := persistence.NewUserRepository(db)
	service := user.NewUpdateUser(repository)

	var req dto.UpdateUserRequest
	userID := request.URL.Query().Get("id")

	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		utils.NewErrorResponse(writer, http.StatusBadRequest, "invalid request body. "+err.Error())
		return
	}

	user, err := service.Execute(userID, req.FirstName, req.LastName)
	if err != nil {
		utils.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(*user); err != nil {
		utils.NewErrorResponse(writer, http.StatusInternalServerError, "internal server error. "+err.Error())
		return
	}
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	db, err := persistence.InitPostgres()
	if err != nil {
		utils.NewErrorResponse(
			writer,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	defer db.Close()

	userID := request.URL.Query().Get("id")

	repository := persistence.NewUserRepository(db)
	service := user.NewDeleteUser(repository)

	user, err := service.Execute(userID)
	if err != nil {
		utils.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(*user); err != nil {
		utils.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
}
