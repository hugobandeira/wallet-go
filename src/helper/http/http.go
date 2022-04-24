package http

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"status"`
	Success    bool        `json:"success"`
	Response   interface{} `json:"response"`
}

func NewResponse(w http.ResponseWriter, response Response) {
	resp, _ := json.Marshal(&response)
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(response.StatusCode)
	w.Write(resp)
}

func ResponseWithSuccess(w http.ResponseWriter, statusCode int, response interface{}) {
	NewResponse(w, Response{
		StatusCode: statusCode,
		Success:    true,
		Response:   response,
	})
}

func ResponseWithError(w http.ResponseWriter, statusCode int, response interface{}) {
	NewResponse(w, Response{
		StatusCode: statusCode,
		Success:    false,
		Response:   response,
	})
}
