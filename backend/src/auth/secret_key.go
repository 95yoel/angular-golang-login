package auth

import (
	"backend/src/models"
	"encoding/json"
	"os"
)

func GetSecretKey() (string, error) {
	// Read the content of the json
	file, err := os.ReadFile("secret_key.json")
	if err != nil {
		return "", err
	}
	// Unmarsahall the json
	var key models.Secret
	err = json.Unmarshal(file, &key)
	if err != nil {
		return "", err
	}

	// Return the key value
	return key.Key, nil
}
