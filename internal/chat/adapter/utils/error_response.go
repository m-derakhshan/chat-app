package utils

import (
	"encoding/json"
	"net/http"

	"media.hiway.chat/internal/chat/domain/model"
)

func NewErrorResponse(writer http.ResponseWriter, statusCode int, message string) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	errorResponse := model.ErrorResponse{
		Code:    statusCode, 
		Message: message,
	}
	json.NewEncoder(writer).Encode(errorResponse)
}