package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Manan-Rastogi/chezz/helpers"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	db *sql.DB
}

func (d *DB) ConfigureDB(ctx context.Context) {
	user := "root"
	password := "password"
	ip := "localhost:3306"
	dbName := "chezz"

	connString := fmt.Sprintf(`%v:%v@tcp(%v)/%v`, user, password, ip, dbName)

	var err error
	d.db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatalf("Error establishing connection to db: %v", err.Error())
	}

	// Ping DB to check connection
	err = d.db.PingContext(ctx)
	if err != nil {
		log.Fatalf("Error verifing connection to db: %v", err.Error())
	}

	helpers.Logger.Info("DB Connection Successful!")
	ctx = context.WithValue(ctx, "DB", d)
}
