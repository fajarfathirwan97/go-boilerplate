package config

import "os"

type Env struct {
	AppPort      string
	AccessKey    string
	AccessSecret string
	DBUri        string
}

func GetEnv() Env {
	return Env{
		AppPort:      os.Getenv("APP_PORT"),
		AccessKey:    os.Getenv("ACCESS_KEY"),
		AccessSecret: os.Getenv("ACCESS_SECRET"),
		DBUri:        os.Getenv("DB_URI"),
	}
}
