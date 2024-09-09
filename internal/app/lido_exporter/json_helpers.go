package lido_exporter

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func readJSON(filename string) (map[string]string, error) {
	var data map[string]string

	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return data, fmt.Errorf("File does not exist")
		}
		return data, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return data, fmt.Errorf("Error parsing JSON file: %v", err)
	}

	return data, nil
}

// Función para escribir en un archivo JSON
func writeJSON(filename string, data map[string]string) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("Error serializing JSON file: %v", err)
	}

	err = os.WriteFile(filename, bytes, 0644)
	if err != nil {
		return fmt.Errorf("Error writing file: %v", err)
	}

	return nil
}
