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
	script := `INSERT INTO customer (id, name) VALUES ("bad", "rules")`

	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("masuk")
}
