package routers

import (
	"duagolang/apk3/controllers"
	"net/http"
)

func init() {
	http.HandleFunc("/apk3", controllers.DashboardController)

	http.Handle("/apk3/static/",
		http.StripPrefix("/apk3/static/",
			http.FileServer(http.Dir("apk3/static"))))
}
