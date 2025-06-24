package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	Filename string
}

func CreateNewStorage[T any](filename string) *Storage[T] {
	return &Storage[T]{Filename: filename}
}

func (s *Storage[T]) save(data T) error {
	filedata, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.Filename, filedata, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	filedata, err := os.ReadFile(s.Filename)

	if err != nil {
		return err
	}
	return json.Unmarshal(filedata, data)
}
