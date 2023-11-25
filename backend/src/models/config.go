package models

type Config struct {
	User     string `json:"user"`
	DBname   string `json:"dbname"`
	Password string `json:"password"`
}

type ConfigDev struct {
	User     string `json:"user"`
	DBname   string `json:"dbname"`
	Password string `json:"password"`
}

type ConfigProd struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	DBname   string `json:"dbname"`
	Password string `json:"password"`
}
