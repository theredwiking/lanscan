package core

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/theredwiking/lanscan/models"
)

func CreateFile(name string) error {
	_, err := os.Stat(name)
	if errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(name)
		defer file.Close()
		if err != nil {
			return errors.New("Failed to create file")
		}
		return nil
	} else {
		return errors.New("File already exist")
	}
}

func WriteFile(name string, data []models.Device) error {
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile(name, jsonData, 0777)
	if err != nil {
		return err
	}

	return nil
}
