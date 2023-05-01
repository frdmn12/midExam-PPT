package main

import (
	"Mahasiswa"
	"fmt"
	"encoding/json"
	"log"
	"net/http"
)


func main() {
	// Define list seluruh mahasiswa

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":

			if r.URL.Path == "/semuadata" {
				data, err := json.Marshal(Mahasiswa.ListMahasiswa)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(data)
				return
			}
            // jika tidak ada parameter, tampilkan seluruh data mahasiswa
            if len(r.URL.Query()) == 0 {
                data, err := json.Marshal(Mahasiswa.ListMahasiswa)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
                w.Header().Set("Content-Type", "application/json")
                w.Write(data)
                return
            }

        case "POST":
            // jika parameter adalah nama mahasiswa, tampilkan data mahasiswa tersebut
            nama := r.FormValue("nama")
            for _, mhs := range Mahasiswa.ListMahasiswa {
                if mhs.Nama == nama {
                    data, err := json.Marshal(mhs)
                    if err != nil {
                        http.Error(w, err.Error(), http.StatusInternalServerError)
                        return
                    }
                    w.Header().Set("Content-Type", "application/json")
                    w.Write(data)
                    return
                }
            }
            // jika nama mahasiswa tidak ditemukan, tampilkan pesan error
            http.Error(w, "Mahasiswa tidak ditemukan", http.StatusNotFound)
            return
        default:
            http.Error(w, "Method belum dibuat / tidak diizinkan", http.StatusMethodNotAllowed)
            return
        }
    })

    fmt.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
