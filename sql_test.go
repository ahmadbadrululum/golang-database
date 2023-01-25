package golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// prepare statement
	script := `INSERT INTO customer (id, name) VALUES ("bad1", "rules1")`

	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("masuk")
}

func TestGetSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// prepare statement
	script := `select * from customer`

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string

		// fmt.Println(rows.Scan())

		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("id " + id)
		fmt.Println("name " + name)
	}
}

func TestSqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// prepare statement
	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate, createdAt time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("id ", id)
		fmt.Println("name ", name)
		if email.Valid {
			fmt.Println("email ", email.String)
		}

		fmt.Println("balance ", balance)
		fmt.Println("rating ", rating)
		fmt.Println("birthDate ", birthDate)
		fmt.Println("married ", married)
		fmt.Println("createdAt ", createdAt)
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	script := "SELECT username FROM user WHERE username = '" + username +
		"' AND password = '" + password + "' LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}
