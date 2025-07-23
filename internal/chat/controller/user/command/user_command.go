package command

import (
	"encoding/json"
	"net/http"

	"media.hiway.chat/internal/chat/adapter/persistence"
	dto "media.hiway.chat/internal/chat/adapter/rest/dto/user"
	"media.hiway.chat/internal/chat/adapter/utils"
	"media.hiway.chat/internal/chat/domain/service/user"
)

// AddUser godoc
// @Summary      Add a new user
// @Description  Creates a new user with first and last name
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user  body      user.CreateUserRequest  true  "User creation payload"
// @Success      201   {object}  model.UserModel
// @Failure      400   {object}  model.ErrorResponse
// @Router       /user [post]
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

// UpdateUser godoc
// @Summary      Update an existing user
// @Description  Updates first name and last name of a user by ID
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id    query     string                 true  "User ID"
// @Param        user  body      user.UpdateUserRequest true  "User update payload"
// @Success      200   {object}  model.UserModel
// @Failure      400   {object}  model.ErrorResponse
// @Failure      500   {object}  model.ErrorResponse
// @Router       /user [patch]
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

// DeleteUser godoc
// @Summary      Delete a user by ID
// @Description  Deletes a user and returns the deleted user object
// @Tags         user
// @Produce      json
// @Param        id   query     string  true  "User ID"
// @Success      200  {object}  model.UserModel
// @Failure      400  {object}  model.ErrorResponse
// @Failure      500  {object}  model.ErrorResponse
// @Router       /user [delete]
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
