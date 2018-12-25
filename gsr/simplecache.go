package gsr

import (
	"fmt"
	"time"
)

var (
	SimpleCacheError                = fmt.Errorf("simplecache error")
	SimpleCacheErrorInvalidArgument = fmt.Errorf("simplecache invalid argument")
)

type CacheInterface interface {
	Get(key, def interface{}) (interface{}, error)
	Set(key, val interface{}, ttl time.Duration) (bool, error)
	Clear() (bool, error)
	GetMultiple(keys, defs []interface{}) ([]interface{}, error)
	SetMultiple(vals map[interface{}]interface{}, ttl time.Duration) (bool, error)
	DeleteMultiple(keys []interface{}) (bool, error)
	Has(key interface{}) (bool, error)
}
