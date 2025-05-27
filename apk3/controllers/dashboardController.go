package controllers

import (
	"duagolang/apk3/config"
	"fmt"
	"net/http"
	"text/template"
)

func DashboardController(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		tmpl, err := template.ParseFiles("apk3/views/layouts/layout.html",
			"apk3/views/dashboard/index.html")
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
		action := r.PostFormValue("action")

		if action == "getData" {
			rows, err := config.DB.Query(`
			SELECT * FROM dpfdplnew.distributors;`)
			if err != nil {
				http.Error(w, "Gagal mengambil data absensi", http.StatusInternalServerError)
				return
			}
			defer rows.Close()

			html := ""
			n := 1
			for rows.Next() {
				var id int
				var nama, status, aktif, cons string
				rows.Scan(&id, &nama, &status, &aktif, &cons)
				html += fmt.Sprintf(`
		<tr>
			<td>%v.</td>
			<td id='txtNama%v'>%v</td>
			<td id='txtStatus%v'>%v</td>
			<td id='txtAktif%v'>%v</td>
			<td id='txtCons%v'>%v</td>
			<td><a style='cursor:pointer;' onclick='edit(%v)'>Edit</a> || <a style='cursor:pointer;' onclick='hapus(%v)'>Hapus</a></td>
			
		</tr>`, n, id, nama, id, status, id, aktif, id, cons, id, id)
				n++
			}
			w.Header().Set("Content-Type", "text/html")
			response := html
			w.Write([]byte(response))
		} else if action == "simpan" {
			inpId := r.PostFormValue("inpId")
			inpNama := r.PostFormValue("inpNama")
			inpStatus := r.PostFormValue("inpStatus")
			inpAktif := r.PostFormValue("inpAktif")
			inpCons := r.PostFormValue("inpCons")

			if inpId == "" {
				var totalData int
				query := fmt.Sprintf(`SELECT COUNT(*) FROM dpfdplnew.distributors WHERE nama = '%s'`, inpNama)
				err := config.DB.QueryRow(query).Scan(&totalData)
				if err != nil {
					w.Header().Set("Content-Type", "text/html")
					response := `gagal~!~Gagal mengambil total data`
					w.Write([]byte(response))
					return
				}

				if totalData == 0 {
					result, err := config.DB.Exec(`INSERT INTO dpfdplnew.distributors (nama, status, aktif, cons) VALUES (?, ?, ?, ?)`, inpNama, inpStatus, inpAktif, inpCons)

					if err != nil {

						response := `gagal~!~Gagal menyimpan data`
						w.Write([]byte(response))
						return

					}

					// Periksa jumlah baris yang terpengaruh
					rowsAffected, err := result.RowsAffected()
					if err != nil {
						response := `gagal~!~Gagal mendapat status simpan`
						w.Write([]byte(response))
						return
					}

					if rowsAffected == 0 {
						response := `gagal~!~Tidak ada data yang tersimpan`
						w.Write([]byte(response))
						return
					}

					w.Header().Set("Content-Type", "text/html")
					response := `sukses~!~Data berhasil tersimpan`
					w.Write([]byte(response))
					return
				} else {
					w.Header().Set("Content-Type", "text/html")
					response := `gagal~!~Sudah ada data yang sama`
					w.Write([]byte(response))
					return
				}
			} else {
				//update
				//query cek apakah sudah ada
				query := fmt.Sprintf("SELECT COUNT(*) FROM dpfdplnew.distributors WHERE id != '%v' AND nama = '%s'", inpId, inpNama)

				var totalData int
				err := config.DB.QueryRow(query).Scan(&totalData)
				if err != nil {
					w.Header().Set("Content-Type", "text/html")
					response := `gagal~!~Gagal menghitung total data`
					w.Write([]byte(response))
					return
				}

				if totalData == 0 {

					query := fmt.Sprintf(`UPDATE dpfdplnew.distributors SET nama = '%s', status= '%s', aktif='%s', cons='%s' WHERE id = '%v';`, inpNama, inpStatus, inpAktif, inpCons, inpId)
					result, err := config.DB.Exec(query)
					if err != nil {
						w.Header().Set("Content-Type", "text/html")
						response := `gagal~!~Gagal update data`
						w.Write([]byte(response))
						return
					}

					// Periksa jumlah baris yang terpengaruh
					rowsAffected, err := result.RowsAffected()
					if err != nil {
						w.Header().Set("Content-Type", "text/html")
						response := `gagal~!~Gagal mendapatkan baris yang terupdate`
						w.Write([]byte(response))
						return
					}

					if rowsAffected == 0 {
						w.Header().Set("Content-Type", "text/html")
						response := `gagal~!~Tidak ada data yang diperbarui`
						w.Write([]byte(response))
						return
					} else {
						w.Header().Set("Content-Type", "text/html")
						response := `sukses~!~Data berhasil diperbarui`
						w.Write([]byte(response))
					}

				} else {
					response := `gagal~!~Data sudah ada`
					w.Write([]byte(response))
					return
				}

			}
		} else if action == "hapus" {
			id := r.PostFormValue("id")

			query := fmt.Sprintf(`DELETE FROM dpfdplnew.distributors WHERE id = '%v';`, id)

			result, err := config.DB.Exec(query)
			if err != nil {
				w.Header().Set("Content-Type", "text/html")
				response := `gagal~!~Gagal menghapus data`
				w.Write([]byte(response))
				return
			}

			rowsAffected, err := result.RowsAffected()
			if err != nil {
				w.Header().Set("Content-Type", "text/html")
				response := `gagal~!~Gagal mendapatkan baris yang dihapus`
				w.Write([]byte(response))
				return
			}

			if rowsAffected == 0 {
				w.Header().Set("Content-Type", "text/html")
				response := `gagal~!~Tidak ada data yang dihapus`
				w.Write([]byte(response))
				return
			} else {
				w.Header().Set("Content-Type", "text/html")
				response := `sukses~!~Data berhasil dihapus`
				w.Write([]byte(response))
				return
			}
		}

		return
	}

}
