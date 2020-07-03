package server

import (
	"errors"
	ip_controller "git.ozon.dev/tvoteva/22_ide/homework/internal/app/ip-controller"
	"github.com/go-chi/chi"
	"net/http"
)

type server struct {
	addr      string
	router    *chi.Mux
	ipService ip_controller.IpService
}

func New(addr string) *server {
	s := &server{
		addr:      addr,
		router:    chi.NewRouter(),
		ipService: ip_controller.New(),
	}
	s.configRouter()
	return s
}

func (s *server) Start() error {
	if err := http.ListenAndServe(s.addr, s.router); err != nil {
		return errors.New("failed to start web server")
	}
	return nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configRouter() {
	router := s.router
	basePath := "/v1"
	router.Route(basePath, func(scope chi.Router) {
		scope.Group(func(public chi.Router) {
			public.Post("/check-access-time", s.HandleCheckAccessTime)
			public.Post("/set-access-time", s.HandleSetAccessTime)
			public.Get("/get-first-access-time", s.HandleGetFirstAccessTime)
		})
	})
}
