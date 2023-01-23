package golang_database

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDatabase(t *testing.T) {
	fmt.Println("halo")
}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "bad:123qwe@tcp(localhost:3306)/golang_database")

	if err != nil {
		panic(err)
	}
	defer db.Close()

}
