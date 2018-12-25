package gsr

import "fmt"

var (
	ContainerError                 = fmt.Errorf("container error")
	ContainerErrorInstanceNotFound = fmt.Errorf("instance was not found")
)

type ContainerInterface interface {
	Get(id interface{}) interface{}
	Has(id interface{}) (bool, error)
}
