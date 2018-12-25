package http

import "net/http"

type RequestHandleFunc func(w http.ResponseWriter, r *http.Request)
