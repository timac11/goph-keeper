package cache

import (
	"context"
	"encoding/json"
	"os"

	"github.com/timac11/goph-keeper/internal/common/logger"
	"go.uber.org/zap"
)

type Cache[T any] struct {
	filePath string
}

func NewCache[T any](path string) *Cache[T] {
	return &Cache[T]{filePath: path}
}

func (c *Cache[T]) Store(ctx context.Context, value T) error {
	log := logger.LoggerFromContext(ctx)

	file, err := os.OpenFile(c.filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		return err
	}

	log.Info("store info to file", zap.String("file", file.Name()))

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

func (c *Cache[T]) Restore(ctx context.Context) (*T, error) {
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
