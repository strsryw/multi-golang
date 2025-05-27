package main

import (
	cfg1 "duagolang/apk1/config"
	_ "duagolang/apk1/routers"
	cfg2 "duagolang/apk2/config" // beri alias cfg2
	_ "duagolang/apk2/routers"
	cfg3 "duagolang/apk3/config" // beri alias cfg3
	_ "duagolang/apk3/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":9090"
	cfg1.ConnectDBBerno()
	cfg2.ConnectDBDpf() // Panggil dengan alias
	cfg3.ConnectDB()    // Panggil dengan alias
	fmt.Println("Server berjalan di http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
