package query

import (
	"encoding/json"
	"net/http"

	"media.hiway.chat/internal/chat/adapter/persistence"
	"media.hiway.chat/internal/chat/adapter/utils"
	"media.hiway.chat/internal/chat/domain/service/user"
)

// GetUserByID godoc
// @Summary      Get a user by ID
// @Description  Returns a single user by their UUID
// @Tags         user
// @Produce      json
// @Param        id   query     string     true  "User ID (UUID)"
// @Success      200  {object}  model.UserModel
// @Failure      400  {object}  model.ErrorResponse
// @Router       /user [get]
func GetUserByID(writer http.ResponseWriter, request *http.Request) {

	db, err := persistence.InitPostgres()
	if err != nil {
		utils.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	repository := persistence.NewUserRepository(db)
	service := user.NewGetUserByID(repository)

	id := request.URL.Query().Get("id")
	if id == "" {
		utils.NewErrorResponse(writer, http.StatusBadRequest, "invalid query parameters")
		return
	}

	user, err := service.Execute(id)
	if err != nil {
		utils.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(*user); err != nil {
		utils.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}

}
