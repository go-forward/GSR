package container

type ContainerInterface interface {
	Get(id interface{}) interface{}
	Has(id interface{}) (bool, error)
}
