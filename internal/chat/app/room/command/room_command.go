package command

import (
	"net/http"

	"media.hiway.chat/internal/chat/adapter/persistence"
	"media.hiway.chat/internal/chat/domain/service/room"
)

func AddRome(writer http.ResponseWriter, request *http.Request) {

	db, err := persistence.InitPostgres()
	if err != nil {
		http.Error(writer, "Internal server error. Failed to initialize PostgreSQL", http.StatusInternalServerError)
	}
	defer db.Close()

	repository := persistence.NewRoomRepository(db)
	service := room.NewCreateRoom(repository)

	if err := service.Execute(request.Body); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("room created successfully"))
}
