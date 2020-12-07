package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/tarokent10/go-sample/sqlboiler/db/entities"
)

const (
	dbuser     string = "user"
	dbpassword string = "user"
	dbhost     string = "localhost"
	dbport     string = "3306"
	database   string = "sampledb"
)

func main() {
	// sqlboiler --wipe mysql

	fmt.Println("main -start")
	defer fmt.Println("main -end")
	// initialization

	sqlinfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpassword, dbhost, dbport, database)
	db, err := sql.Open("mysql", sqlinfo)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	es, err := entities.Users().All(context.Background(), db)
	if err != nil {
		// TODO error handle
		panic(err)
	}

	for _, e := range es {
		fmt.Printf("entity:%#v\n", e)
	}
}
