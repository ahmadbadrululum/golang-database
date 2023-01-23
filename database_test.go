package golang_database

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDatabase(t *testing.T) {
	fmt.Println("halo")
}
