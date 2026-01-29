package handler

import (
	"encoding/json"
	"net/http"

	"anon-wa/internal/service"
)

type Request struct {
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

func SendHandler(svc *service.MessageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		defer r.Body.Close()

		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		if err := svc.Send(req.Phone, req.Message); err != nil {
			// log.Println(err)
			http.Error(w, "failed to send message", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	}
}
