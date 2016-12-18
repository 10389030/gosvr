package tools

import (
	"net/http"
)

// RESTful interface handler
type RESTfulGeter interface {
	Get(http.ResponseWriter, *http.Request)
}

type RESTfulPuter interface {
	Put(http.ResponseWriter, *http.Request)
}

type RESTfulDeleter interface {
	Delete(http.ResponseWriter, *http.Request)
}

type RESTfulPoster interface {
	Post(http.ResponseWriter, *http.Request)
}

// adapt http.ServeHTTP
type RESTfulRouter struct {
	Handler interface{}
}

// impletement http.Handler interface
func (router *RESTfulRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if gh, ok := router.Handler.(RESTfulGeter); ok {
			gh.Get(w, r)
		}
	case "POST":
		if gh, ok := router.Handler.(RESTfulPoster); ok {
			gh.Post(w, r)
		}
	case "PUT":
		if gh, ok := router.Handler.(RESTfulPuter); ok {
			gh.Put(w, r)
		}
	case "DELETE":
		if gh, ok := router.Handler.(RESTfulDeleter); ok {
			gh.Delete(w, r)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

// RESTServeMux is a HTTP request multiplexer that extends http.ServeMux
// Add a ability to route request by http.Method by HandleRESTful.
type RESTServeMux struct {
	*http.ServeMux
}

// HandleRESTful registes handle to route request by http.Method and Path.
// Parameter `handler` should impletement at least one of RESTfulGeter, RESTfulPoster, RESTfulDeleter, RESTfulPuter.
// if handler impletement RESTfulGeter then support http request on pattern with "GET" Method, so as other.
func (mux *RESTServeMux) HandleRESTful(pattern string, handler interface{}) {
	if _, ok := handler.(RESTfulGeter); ok {
		mux.Handle(pattern, &RESTfulRouter{Handler: handler})
	} else if _, ok := handler.(RESTfulPoster); ok {
		mux.Handle(pattern, &RESTfulRouter{Handler: handler})
	} else if _, ok := handler.(RESTfulDeleter); ok {
		mux.Handle(pattern, &RESTfulRouter{Handler: handler})
	} else if _, ok := handler.(RESTfulPuter); ok {
		mux.Handle(pattern, &RESTfulRouter{Handler: handler})
	} else {
		panic("RESTful handler should implement at least one handle of [RESTfulGeter, RESTfulPoster, RESTfulDeleter, RESTfulPuter]")
	}
}

// NewRESTServeMux allocates and returns a new RESTServeMux
func NewRESTServeMux() *RESTServeMux { return &RESTServeMux{http.NewServeMux()} }
