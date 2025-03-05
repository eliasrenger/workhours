package utils

import (
	"encoding/json"
	"os"
)

func ReadJSON[T any](filePath string) (T, error) {
	var result T

	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return result, err
	}

	// Unmarshal the JSON
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func AppendJSON[T any, S []T](filePath string, newData S) error {
	oldData, err := ReadJSON[S](filePath)
	if err != nil {
		return err
	}
	rawData := append(oldData, newData...)
	data, err := json.MarshalIndent(rawData, "", "  ")
	return SaveJSON(filePath, data)
}

func SaveJSON(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}
