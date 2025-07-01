package query

import (
	"encoding/json"
	"net/http"

	"media.hiway.chat/internal/chat/adapter/persistence"
	"media.hiway.chat/internal/chat/adapter/utils"
	"media.hiway.chat/internal/chat/domain/service/room"
)

func GetAllRooms(writer http.ResponseWriter, request *http.Request) {

	db, err := persistence.InitPostgres()
	if err != nil {
		utils.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	repository := persistence.NewRoomRepository(db)
	service := room.NewGetAllRooms(repository)

	rooms, err := service.Execute()
	if err != nil {
		utils.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(*rooms); err != nil {
		utils.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
}

func GetRoomByID(writer http.ResponseWriter, request *http.Request) {

	db, err := persistence.InitPostgres()
	if err != nil {
		utils.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()

	repository := persistence.NewRoomRepository(db)
	service := room.NewGetRoomByID(repository)

	id := request.URL.Query().Get("id")
	if id == "" {
		utils.NewErrorResponse(writer, http.StatusBadRequest, "invalid query parameters")
		return
	}

	rooms, err := service.Execute(id)
	if err != nil {
		utils.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(*rooms); err != nil {
		utils.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
}
