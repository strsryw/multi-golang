package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DBDpf *sql.DB
)

func ConnectDBDpf() {
	var err error
	DBDpf, err = sql.Open("mysql", "root:@tcp(localhost:3306)/dpfdplnew")
	if err != nil {
		panic("Gagal konek ke database : " + err.Error())
	}

	if err = DBDpf.Ping(); err != nil {
		panic("Gagal ping database : " + err.Error())
	}

	log.Println("Koneksi ke database berhasil")
}
