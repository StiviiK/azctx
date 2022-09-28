package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/afero"
)

// FileExists checks if a file exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// ReadJson reads a json file, removes utf-8 BOM and unmarshals it
func ReadJson(file afero.File, v any) error {
	// Read the file
	configBytes, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read %s: %s", file.Name(), err.Error())
	}

	// Remove UTF-8 BOM if present
	configBytes = RemoveUTF8BOM(configBytes)

	// Unmarshal the config file
	err = json.Unmarshal(configBytes, v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal %s: %s", file.Name(), err.Error())
	}

	return nil
}

// WriteJson writes a json file
func WriteJson(file afero.File, v any) error {
	// Marshal the config file
	configBytes, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to marshal %s: %s", file.Name(), err.Error())
	}

	// Write the file
	_, err = file.Write(configBytes)
	if err != nil {
		return fmt.Errorf("failed to write %s: %s", file.Name(), err.Error())
	}

	return nil
}
