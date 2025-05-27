package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

func ConnectDB() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(localhost:3306)/go_products_master")
	if err != nil {
		panic("Gagal konek ke database : " + err.Error())
	}

	if err = DB.Ping(); err != nil {
		panic("Gagal ping database : " + err.Error())
	}

	log.Println("Koneksi ke database berhasil")
}
