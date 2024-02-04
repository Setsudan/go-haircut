package routes

import (
	"encoding/json"
	"gohairdresser/structs"
	"io"
	"log"
	"net/http"
)

func SendResponse(w http.ResponseWriter, code int, status, message string, data interface{}, err error) {
	var response structs.APIResponse

	if err != nil {
		response = structs.APIResponse{
			Code:    http.StatusInternalServerError,
			Status:  "Error",
			Message: err.Error(),
			Data:    nil,
		}
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		response = structs.APIResponse{
			Code:    code,
			Status:  status,
			Message: message,
			Data:    data,
		}
		w.WriteHeader(code)
	}

	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(response)
	if jsonErr != nil {
		// Log the error of failing to send the response
		log.Printf("Error sending response: %v", jsonErr)
	}
}

func SendErrorResponse(w http.ResponseWriter, message string, err error, statusCode int) {
	log.Println(err)
	response := struct {
		Code    int         `json:"code"`
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Code:    statusCode,
		Status:  "error",
		Message: message,
		Data:    nil,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func handleJSONDecodingError(w http.ResponseWriter, err error) {
	if err == io.EOF {
		// Handle empty body
		SendErrorResponse(w, "Request body is empty or in wrong format", err, http.StatusBadRequest)
	} else {
		// Handle other JSON decoding errors
		SendErrorResponse(w, "Error decoding JSON", err, http.StatusBadRequest)
	}
}
