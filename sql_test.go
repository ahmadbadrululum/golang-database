package golang_database

import (
	"context"
	"fmt"
	"testing"
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
