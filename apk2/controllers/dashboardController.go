package controllers

import (
	"net/http"
	"text/template"
)

func DashboardController(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		tmpl, err := template.ParseFiles("apk2/views/layouts/layout.html",
			"apk2/views/dashboard/index.html")
		if err != nil {
			http.Error(w, "Gagal load template", http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"Title":      "Dashboard",
			"ActivePage": "dashboard",
		}

		tmpl.Execute(w, data)
		return

	} else if r.Method == http.MethodPost {

	}

}
