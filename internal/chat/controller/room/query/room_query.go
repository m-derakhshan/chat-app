package query

import (
	"encoding/json"
	"net/http"

	"media.hiway.chat/internal/chat/adapter/persistence"
	"media.hiway.chat/internal/chat/domain/service/room"
)

func GetAllRooms(writer http.ResponseWriter, request *http.Request) {

	db, err := persistence.InitPostgres()
	if err != nil {
		http.Error(writer, "Internal server error. Failed to connect to the PostgreSQL", http.StatusInternalServerError)
	}
	defer db.Close()

	repository := persistence.NewRoomRepository(db)
	service := room.NewGetAllRooms(repository)

	rooms, err := service.Execute()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(*rooms); err != nil {
		http.Error(writer, "Failed to encode rooms to JSON", http.StatusInternalServerError)
		return
	}
}

func GetRoomByID(writer http.ResponseWriter, request *http.Request) {

	db, err := persistence.InitPostgres()
	if err != nil {
		http.Error(writer, "Internal server error. Failed to connect to the PostgreSQL", http.StatusInternalServerError)
	}
	defer db.Close()

	repository := persistence.NewRoomRepository(db)
	service := room.NewGetRoomByID(repository)

	rooms, err := service.Execute(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(*rooms); err != nil {
		http.Error(writer, "Failed to encode rooms to JSON", http.StatusInternalServerError)
		return
	}
}
