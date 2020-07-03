package server

import (
	"context"
	"encoding/json"
	keys "git.ozon.dev/tvoteva/22_ide/homework/internal/app/constants"
	"git.ozon.dev/tvoteva/22_ide/homework/pkg/model"
	"net/http"
)

func (s *server) HandleSetAccessTime(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), keys.IP, getIP(r))
	s.ipService.HandleSetAccessTime(ctx)
	writeResponse(w, http.StatusOK, nil)
}

func (s *server) HandleGetFirstAccessTime(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), keys.IP, getIP(r))
	response := s.ipService.HandleGetFirstAccessTime(ctx)
	writeResponse(w, http.StatusOK, response)
}

func (s *server) HandleCheckAccessTime(w http.ResponseWriter, r *http.Request) {
	req := &model.IpRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := s.ipService.HandleCheckAccessTime(r.Context(), *req)
	if err != nil {
		writeErrorResponse(w, http.StatusNotFound, err)
		return
	}
	writeResponse(w, http.StatusOK, response)
}
