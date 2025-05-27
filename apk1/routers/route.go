package routers

import (
	"duagolang/apk1/controllers"
	"net/http"
)

func init() {
	http.HandleFunc("/apk1", controllers.DashboardController)

	http.Handle("/apk1/static/",
		http.StripPrefix("/apk1/static/",
			http.FileServer(http.Dir("apk1/static"))))
}
