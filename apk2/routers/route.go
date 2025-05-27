package routers

import (
	"duagolang/apk2/controllers"
	"net/http"
)

func init() {
	http.HandleFunc("/apk2", controllers.DashboardController)

	http.Handle("/apk2/static/",
		http.StripPrefix("/apk2/static/",
			http.FileServer(http.Dir("apk2/static"))))
}
