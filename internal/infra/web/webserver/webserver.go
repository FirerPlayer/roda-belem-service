package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router         chi.Router
	GetHandlers    map[string]http.HandlerFunc
	PostHandlers   map[string]http.HandlerFunc
	DeleteHandlers map[string]http.HandlerFunc
	GroupRoutes    []func(r chi.Router)
	WebServerPort  string
}

func NewWebServer(webServerPort string) *WebServer {
	return &WebServer{
		WebServerPort:  webServerPort,
		Router:         chi.NewRouter(),
		GetHandlers:    make(map[string]http.HandlerFunc),
		PostHandlers:   make(map[string]http.HandlerFunc),
		DeleteHandlers: make(map[string]http.HandlerFunc),
		GroupRoutes:    make([]func(r chi.Router), 0),
	}
}

func (s *WebServer) AddGetHandler(path string, handler http.HandlerFunc) {
	s.GetHandlers[path] = handler
}

func (s *WebServer) AddPostHandler(path string, handler http.HandlerFunc) {
	s.PostHandlers[path] = handler
}

func (s *WebServer) AddDeleteHandler(path string, handler http.HandlerFunc) {
	s.DeleteHandlers[path] = handler
}

func (s *WebServer) AddGroupRoute(route func(r chi.Router)) {
	s.GroupRoutes = append(s.GroupRoutes, route)
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, handler := range s.GetHandlers {
		s.Router.Get(path, handler)
	}
	for path, handler := range s.PostHandlers {
		s.Router.Post(path, handler)
	}
	for path, handler := range s.DeleteHandlers {
		s.Router.Delete(path, handler)
	}
	for _, router := range s.GroupRoutes {
		s.Router.Group(router)
	}

	if err := http.ListenAndServe(":"+s.WebServerPort, s.Router); err != nil {
		panic(err.Error())
	}
}
