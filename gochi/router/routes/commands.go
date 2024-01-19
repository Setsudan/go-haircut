package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendJSONResponse(w http.ResponseWriter, data interface{}) {
	res, err := json.Marshal(data)
	if err != nil {
		SendErrorResponse(w, "Error marshalling JSON", err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
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
