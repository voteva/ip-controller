package api

import (
	"encoding/json"
	"github.com/voteva/ip-controller/pkg/model"
	"net"
	"net/http"
)

func writeErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	result := model.ErrorResponse{
		Error: err.Error(),
	}
	writeResponse(w, statusCode, result)
}

func writeResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	if data != nil {
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(statusCode)
}

func getIP(r *http.Request) string {
	if forwarded := r.Header.Get("X-FORWARDED-FOR"); forwarded != "" {
		return forwarded
	}
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	return host
}
