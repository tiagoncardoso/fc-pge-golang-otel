package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	Handler http.HandlerFunc
	Method  string
	Path    string
}

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]Handler
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]Handler),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, method string, handler http.HandlerFunc) {
	s.Handlers[path+"-"+method] = Handler{
		handler,
		method,
		path,
	}
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)

	for _, handler := range s.Handlers {
		switch handler.Method {
		case "GET":
			s.Router.Get(handler.Path, handler.Handler)
		case "POST":
			s.Router.Post(handler.Path, handler.Handler)
		case "PUT":
			s.Router.Put(handler.Path, handler.Handler)
		case "DELETE":
			s.Router.Delete(handler.Path, handler.Handler)
		default:
			s.Router.Head(handler.Path, handler.Handler)
		}
	}
	err := http.ListenAndServe(":"+s.WebServerPort, s.Router)
	if err != nil {
		panic(err)
	}
}
