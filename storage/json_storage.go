package storage

import (
	"fmt"
	"os"
)

type JsonStorage struct {
	*Storage
}

func CreateJsonStorage(filename string) *JsonStorage {
	return &JsonStorage{
		&Storage{
			filename: filename,
		},
	}
}

func (s *JsonStorage) Save(data []byte) error {
	err := os.WriteFile(s.GetFilename(), data, 0644)

	if err != nil {
		return fmt.Errorf("%w: %w", ErrSaveStorage, err)
	}

	return nil
}

func (s *JsonStorage) Load() ([]byte, error) {
	data, err := os.ReadFile(s.GetFilename())

	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrLoadStorage, err)
	}

	return data, nil
}
