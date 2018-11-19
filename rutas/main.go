package rutas

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Init() {
	r := mux.NewRouter()

	r.HandleFunc("/api/software", ListarSoftware).Methods("GET")
	r.HandleFunc("/api/software/{id}", VerSoftware).Methods("GET")
	r.HandleFunc("/api/software", CrearSoftware).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
