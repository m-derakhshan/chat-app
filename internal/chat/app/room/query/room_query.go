package query

import (
	"net/http"
)

func GetAllRooms(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}
