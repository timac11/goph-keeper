package cache

import (
	"encoding/json"
	"os"
)

type Cache[T any] struct {
	filePath string
}

func NewCache[T any](path string) *Cache[T] {
	return &Cache[T]{filePath: path}
}

func (c *Cache[T]) Store(value T) error {
	file, err := os.OpenFile(c.filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(&value)

	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cache[T]) Restore() (*T, error) {
	data, err := os.ReadFile(c.filePath)

	if err != nil {
		return nil, err
	}

	var value *T
	err = json.Unmarshal(data, &value)

	if err != nil {
		return nil, err
	}

	return value, nil
}
