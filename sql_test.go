package golang_database

import (
	"context"
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
		var id, name, email string
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
		fmt.Println("balance ", balance)
		fmt.Println("rating ", rating)
		fmt.Println("birthDate ", birthDate)
		fmt.Println("married ", married)
		fmt.Println("createdAt ", createdAt)
	}
}
