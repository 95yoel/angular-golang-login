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

	// Obtain the current enviroment
	env := os.Getenv("ENV")
	connectionString := ""

	if env == "production" { // if the enviroment is prod , use the dbconfig_.prod file
		config, err := loadConfigProd("db_config.prod.json")
		if err != nil {
			log.Fatal("Error loading configuration", err)
		}
		connectionString = createConnectionStringProd(config)
	} else { // if the enviroment isn't prod , use the dbconfig_.dev file
		config, err := loadConfigDev("db_config.dev.json")
		if err != nil {
			log.Fatal("Error loading configuration", err)
		}
		connectionString = createConnectionStringDev(config)
	}

	return connectionString
}

/*
Create the connectionString in the format :
connectionString := "user=user password=password dbname=dbname  sslmode=disable"
*/

func createConnectionStringDev(config models.ConfigDev) string {
	return "user=" + config.User + " password=" + config.Password + " dbname=" + config.DBname + " sslmode=disable"
}

/*
Create the connectionString in the format :
connectionString := "user=user password=password dbname=dbname host=host  sslmode=disable"
*/

func createConnectionStringProd(config models.ConfigProd) string {
	return "user=" + config.User + " password=" + config.Password + " dbname=" + config.DBname + " host=" + config.Host + " sslmode=disable"
}

/*
Load the configuration from a JSON file and unmarshal it into the specified structure.
*/
func loadConfig(filename string, config interface{}) error {
	// Obtain the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	// Create the complete route
	filePath := filepath.Join(currentDir, filename)

	// Read file from the complete route
	file, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Unmarshal the JSON content into the specified structure
	return json.Unmarshal(file, config)
}

/*
Load the dev configuration and return it as a ConfigDev struct.
*/
func loadConfigDev(filename string) (models.ConfigDev, error) {
	var config models.ConfigDev
	err := loadConfig(filename, &config)
	return config, err
}

/*
Load the prod configuration and return it as a ConfigProd struct.
*/
func loadConfigProd(filename string) (models.ConfigProd, error) {
	var config models.ConfigProd
	err := loadConfig(filename, &config)
	return config, err
}
