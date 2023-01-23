package golang_database

import (
	"database/sql"
	"time"
)

func getConnection() {

	db, err := sql.Open("mysql", "bad:123qwe@tcp(localhost:3306)/golang_database")

	if err != nil {
		panic(err)
	}

	// maximal idle koneksi
	db.SetMaxIdleConns(10)
	// max open koneksi
	db.SetMaxOpenConns(100)
	// max lifetime koneksi jika 5 menit di close saja
	db.SetConnMaxIdleTime(5 * time.Minute)
	//
	db.SetConnMaxLifetime(60 * time.Minute)
}
