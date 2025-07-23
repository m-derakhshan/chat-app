package command

import (
	"encoding/json"
	"net/http"

	"media.hiway.chat/internal/chat/adapter/persistence"
	dto "media.hiway.chat/internal/chat/adapter/rest/dto/room"
	"media.hiway.chat/internal/chat/adapter/utils"
	"media.hiway.chat/internal/chat/domain/service/room"
)



// AddRoom godoc
// @Summary      Create a new chat room
// @Description  Creates a new room with the given name
// @Tags         room
// @Accept       json
// @Produce      json
// @Param        room  body      room.CreateRoomRequest  true  "Room creation payload"
// @Success      201   {object}  model.RoomModel
// @Failure      400   {object}  model.ErrorResponse
// @Router       /room [post]
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


// UpdateRoomName godoc
// @Summary      Update a room's name
// @Description  Update the name of a chat room by its ID
// @Tags         room
// @Accept       json
// @Produce      json
// @Param        id    query     string                 true   "Room ID"
// @Param        room  body      room.UpdateRoomRequest  true   "Updated room name"
// @Success      200   {object}  model.RoomModel
// @Failure      400   {object}  model.ErrorResponse
// @Router       /room [patch]
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


// DeleteRoom godoc
// @Summary      Delete a room by ID
// @Description  Deletes a chat room identified by its ID
// @Tags         room
// @Produce      json
// @Param        id   query     string  true  "Room ID to delete"
// @Success      200  {object}  model.RoomModel
// @Failure      500  {object}  model.ErrorResponse
// @Router       /room [delete]
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
