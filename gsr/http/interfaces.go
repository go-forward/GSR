package http

type MiddlewareInterface interface {
	Process(next RequestHandleFunc) RequestHandleFunc
}
