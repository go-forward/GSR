package container

import "fmt"

var (
	ContainerError                 = fmt.Errorf("container error")
	ContainerErrorInstanceNotFound = fmt.Errorf("instance was not found")
)
