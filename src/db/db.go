package db

import (
	"database/sql"
	"fmt"

	easy "github.com/MhunterDev/hunterdev/src/base/encoder"
	logsalot "github.com/MhunterDev/hunterdev/src/base/logs"
	_ "github.com/lib/pq"
)

var connString, connErr = easy.GetConn()

func AddDefault() error {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		logsalot.DbErr(err)
		return err
	}
	defer db.Close()

	var username = "admin"
	var password = "admin"
	var status = "yes"

	var insertUserBase = "INSERT INTO app.users(username,token,status) VALUES(%s)"
	var insertHashBase = "INSERT INTO app.secrets(token,hash) VALUES(%s)"

	token, hash, err := easy.HashAndToken(password)
	if err != nil {
		logsalot.DbErr(err)
	}

	formattedUser := fmt.Sprintf("'%s','%s','%s'", username, token, status)
	formatHash := fmt.Sprintf("'%s','%s'", token, hash)

	finalUser := fmt.Sprintf(insertUserBase, formattedUser)
	finalHash := fmt.Sprintf(insertHashBase, formatHash)

	db.Exec(finalUser)
	db.Exec(finalHash)

	return nil
}
