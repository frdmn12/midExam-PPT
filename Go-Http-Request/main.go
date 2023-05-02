package main

import (
	"Mahasiswa"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Membuat requs http
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 1. method:Get  parameter:/semuadata
		// 2. method:Post  parameter:/nama

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
		case "POST":
			// parameter:/nama
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
			http.Error(w, "Mahasiswa tidak ditemukan", http.StatusNotFound)
		default:
			http.Error(w, "Method belum dibuat / tidak diizinkan", http.StatusMethodNotAllowed)
		}

	})
	fmt.Println("Port berjalan http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
