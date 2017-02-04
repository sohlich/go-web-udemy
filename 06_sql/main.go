package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

type RowMap struct {
	Title  string
	Author string
}

func main() {
	log.Printf("Application %s starting.", "SQL Training")

	db, err = sql.Open("mysql", "user:user@tcp(localhost:3306)/attendence?charset=utf8")
	if err != nil {
		log.Println("Cannot connect to DB: ", err.Error())
		return
	}
	defer db.Close()

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)

	rows, err := db.Query("SELECT title,author from book")
	if err != nil {
		log.Println("Cannot query the DB: ", err.Error())
		return
	}

	// All rows must be read
	for rows.Next() {
		auth := &RowMap{}
		err := rows.Scan(&auth.Title, &auth.Author)
		if err != nil {
			log.Println("Cannot read row: ", err.Error())
			continue
		}
		log.Println(auth)
	}
}

// The way how to resolve
// the row into a map
func rowToMap(rows *sql.Rows) map[string]string {
	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for i, _ := range columns {
		valuePtrs[i] = &values[i]
	}
	rows.Scan(valuePtrs...)
	rowMap := map[string]string{}

	for i, col := range columns {
		var v interface{}
		val := values[i]
		b, ok := val.([]byte)
		if ok {
			v = string(b)
		} else {
			v = val
		}
		rowMap[col] = fmt.Sprintf("%s", v)
	}
	return rowMap
}
