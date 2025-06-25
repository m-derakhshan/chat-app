package command

import (
	"encoding/json"
	"net/http"

	"media.hiway.chat/internal/chat/adapter/persistence"
	"media.hiway.chat/internal/chat/domain/service/room"
)

type createRoomRequest struct {
	Name string `json:"room_name"`
}

func AddRome(writer http.ResponseWriter, request *http.Request) {
	var req createRoomRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		return
	}

	db, err := persistence.InitPostgres()
	if err != nil {
		http.Error(writer, "Internal server error. Failed to initialize PostgreSQL", http.StatusInternalServerError)
	}
	defer db.Close()

	repository := persistence.NewRoomRepository(db)
	service := room.NewCreateRoom(repository)

	if err := service.Execute(req.Name); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("room created successfully"))
}
