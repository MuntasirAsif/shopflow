package response

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type APIResponse struct {
	Status     bool   `json:"status"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
}

func (r *APIResponse) JSON(ctx context.Context, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Println("Error encoding response:", err)
	}
}

func (r *APIResponse) Error(ctx context.Context, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Println("Error encoding response:", err)
	}
}

func NewAPIResponse() *APIResponse {
	return &APIResponse{}
}

func Response(status bool, statusCode int, message string, data any) *APIResponse {
	return &APIResponse{
		Status:     status,
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}
