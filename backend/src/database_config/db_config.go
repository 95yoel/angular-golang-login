package dbconfig

import (
	"backend/src/models"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

/*
Returns the connection string by loading configuration from db_config.json.
*/
func GetConnectionString() string {
	//Obtain the content in config.json
	config, err := loadConfig("db_config.json")
	if err != nil {
		log.Fatal("Error loading configuration", err)
	}
	connectionString := createConnectionString(config)
	return connectionString
}

/*
	Read the filename , parses it content to JSON , and returns a Config struct.
*/

func loadConfig(filename string) (models.Config, error) {
	var config models.Config

	// Obtain the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		return config, err
	}

	// Create the complete route
	filePath := filepath.Join(currentDir, filename)

	// Read file from the complete route
	file, err := os.ReadFile(filePath)
	if err != nil {
		return config, err
	}

	// Unmarshal the JSON content into the Config struct
	err = json.Unmarshal(file, &config)
	if err != nil {
		return config, err
	}

	//return the config
	return config, nil
}

/*
Create the connectionString in the format :
connectionString := "user=user password=password dbname=dbname  sslmode=disable"
*/
func createConnectionString(config models.Config) string {
	return "user=" + config.User + " password=" + config.Password + " dbname=" + config.DBname + " sslmode=disable"
}
