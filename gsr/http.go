package gsr

import "net/http"

type RequestHandleFunc func(w http.ResponseWriter, r *http.Request)

type MiddlewareInterface interface {
	Process(next RequestHandleFunc) RequestHandleFunc
}
