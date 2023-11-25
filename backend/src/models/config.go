package models

type Config struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	DBname   string `json:"dbname"`
	Password string `json:"password"`
}
