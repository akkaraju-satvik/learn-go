package db

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"

	_ "github.com/lib/pq"
	"tudoo.app/cli/types"
)

var db, e = sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable port=3006")

type Person struct {
	Name  sql.NullString
	Age   sql.NullInt64
	Email sql.NullString
}

func (p *Person) IterateColumns() []interface{} {
	px := reflect.ValueOf(p)
	x := reflect.Indirect(px)
	m := make([]interface{}, x.NumField())
	for i := 0; i < x.NumField(); i++ {
		m[i] = px.Elem().Field(i).Addr().Interface()
	}
	return m
}

func Connect() {
	if e != nil {
		panic(e)
	}
	now := time.Now()
	results, err := db.Query("select * from people;")
	if err != nil {
		panic(err)
	}
	fmt.Println("Duration for 'select * from people': ", time.Since(now))
	if err != nil {
		panic(err)
	}
	for results.Next() {
		var person Person
		person.IterateColumns()
		err = results.Scan(person.IterateColumns()...)
		if err != nil {
			panic(err)
		}
		fmt.Println(person.Name.String, person.Age.Int64, person.Email.String)
	}
}

func InsertPerson(p types.Person) (string, string, error) {
	var name, age string
	now := time.Now()
	err := db.QueryRow("insert into people values ($1, $2) returning name, age;", p.Name, p.Age).Scan(&name, &age)
	if err != nil {
		return "", "", err
	}
	fmt.Println("Duration: ", time.Since(now))
	return name, age, nil
}
