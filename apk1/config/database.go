package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DBBerno *sql.DB
)

func ConnectDBBerno() {
	var err error
	DBBerno, err = sql.Open("mysql", "root:@tcp(localhost:3306)/bernodb")
	if err != nil {
		panic("Gagal konek ke database : " + err.Error())
	}

	if err = DBBerno.Ping(); err != nil {
		panic("Gagal ping database : " + err.Error())
	}

	log.Println("Koneksi ke database berhasil")
}
