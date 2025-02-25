package db

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"social_network/constant/env"
	"social_network/constant/noti"
)

const (
	db_server      string = "postgre"
	backUpDbCnnStr string = "Your back up database connection string"
)

func ConnectDB(table string) (*sql.DB, error) {
	var cnnStr string = os.Getenv(env.DB_CNN_STR)

	if cnnStr == "" {
		log.Println(noti.DbCnnStrNotSetMsg)

		if err := os.Setenv(env.DB_CNN_STR, backUpDbCnnStr); err != nil {
			log.Println(noti.DbSetConnectionStrErrMsg + err.Error())
		}

		cnnStr = backUpDbCnnStr
	}

	cnn, err := sql.Open(db_server, cnnStr)
	if err != nil {
		log.Println(noti.DbConnectionErrMsg + err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	return cnn, nil
}
