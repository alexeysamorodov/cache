package cache

import (
	"errors"
)

type Cache interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

var ErrorNotFound = errors.New("item was not found")
