package setup

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func OpenDatabase() (*sql.DB, error) {
	envConfig, err := LoadEnvConfig()
	if err != nil {
		return nil, fmt.Errorf("could not load environment informations. %v", err)
	}

	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", envConfig.DatabaseHost, envConfig.DatabasePort, envConfig.DatabaseUserName, envConfig.DatabasePassword, envConfig.DatabaseName)
	db, err := sql.Open("postgres", datasource)
	if err != nil {
		return nil, fmt.Errorf("could not load database driver. %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("could not connect to database. %v", err)
	}
	return db, nil
}
