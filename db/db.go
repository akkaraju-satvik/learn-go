package db

import (
	"database/sql"

	"github.com/akkaraju-satvik/learn-go/types"

	_ "github.com/lib/pq"
)

var db, e = sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable port=3006")

func Connect() {
	if e != nil {
		panic(e)
	}
	results, err := db.Query("select * from people;")
	if err != nil {
		panic(err)
	}
	for results.Next() {
		var name string
		var age int
		err = results.Scan(&name, &age)
		if err != nil {
			panic(err)
		}
		println(name, age)
	}
}

func InsertPerson(p types.Person) (string, string, error) {
	var name, age string
	err := db.QueryRow("insert into people values ($1, $2) returning name, age;", p.Name, p.Age).Scan(&name, &age)
	if err != nil {
		return "", "", err
	}

	return name, age, nil
}
