package config

import "os"

type Database struct {
	URL string
}

func LoadDatabaseConfiguration() Database {
	return Database{
		URL: os.Getenv("POSTGRESQL_URL"),
	}
}
