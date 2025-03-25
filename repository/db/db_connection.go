package db

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq" // PostgreSQL driver
	"log"
	"os"
	"social_network/constant/env"
	"social_network/constant/noti"
)

const (
	db_server      string = "postgres"
	backUpDbCnnStr string = "Your back up database connection string"
)

func ConnectDB(logger *log.Logger) (*sql.DB, error) {
	var cnnStr string = os.Getenv(env.DB_CNN_STR)

	if cnnStr == "" {
		logger.Println(noti.DbCnnStrNotSetMsg)

		if err := os.Setenv(env.DB_CNN_STR, backUpDbCnnStr); err != nil {
			logger.Println(noti.DbSetConnectionStrErrMsg + err.Error())
		}

		cnnStr = backUpDbCnnStr
	}

	cnn, err := sql.Open(db_server, cnnStr)
	if err != nil {
		logger.Println(noti.DbConnectionErrMsg + err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	return cnn, nil
}
