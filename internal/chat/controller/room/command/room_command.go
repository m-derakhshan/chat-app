package command

import (
	"encoding/json"
	"net/http"

	"media.hiway.chat/internal/chat/adapter/persistence"
	dto "media.hiway.chat/internal/chat/adapter/rest/dto/room"
	"media.hiway.chat/internal/chat/adapter/utils"
	"media.hiway.chat/internal/chat/domain/service/room"
)

func AddRoom(writer http.ResponseWriter, request *http.Request) {

	db, err := persistence.InitPostgres()
	if err != nil {
		utils.NewErrorResponse(
			writer,
			http.StatusInternalServerError,
			"Internal server error. Failed to connect to the PostgreSQL")
		return
	}
	defer db.Close()

	repository := persistence.NewRoomRepository(db)
	service := room.NewCreateRoom(repository)

	var req dto.CreateRoomRequest

	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		utils.NewErrorResponse(writer, http.StatusBadRequest, "invalid request body. "+err.Error())
		return
	}

	room, err := service.Execute(req.Name)
	if err != nil {
		utils.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}

	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(*room); err != nil {
		utils.NewErrorResponse(writer, http.StatusInternalServerError, "internal server error. "+err.Error())
		return
	}

}

func UpdateRoomName(writer http.ResponseWriter, request *http.Request) {

	db, err := persistence.InitPostgres()
	if err != nil {
		utils.NewErrorResponse(
			writer,
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	defer db.Close()

	roomID := request.URL.Query().Get("id")

	var req dto.UpdateRoomRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		utils.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	repository := persistence.NewRoomRepository(db)
	service := room.NewUpdateRoomName(repository)
	room, err := service.Execute(roomID, req.Name)

	if err != nil {
		utils.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(*room); err != nil {
		utils.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}

}

func DeleteRoom(writer http.ResponseWriter, request *http.Request) {

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

	roomID := request.URL.Query().Get("id")

	repository := persistence.NewRoomRepository(db)
	service := room.NewDeleteRoom(repository)

	room, err := service.Execute(roomID)
	if err != nil {
		utils.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(*room); err != nil {
		utils.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
}
