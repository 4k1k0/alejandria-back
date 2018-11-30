package rutas

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	db "../db"
	modelo "../db/models"
	"github.com/gorilla/mux"
)

func ListarSoftware(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	base := db.Conectar()
	var programas []modelo.SoftwareItem
	err := base.Model(&programas).Select()
	if err != nil {
		log.Printf("Error")
		db.Desconectar(base)
		return
	}
	db.Desconectar(base)
	json.NewEncoder(w).Encode(&programas)
}

func VerSoftware(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var programa modelo.SoftwareItem

	base := db.Conectar()

	err := base.Model(&programa).Where("id = ?", id).Select()

	if err != nil {
		log.Printf("Error el elemento no existe")
		// Implementar una respuesta json
		db.Desconectar(base)
		return
	}
	db.Desconectar(base)
	json.NewEncoder(w).Encode(&programa)
}

func CrearSoftware(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	var software modelo.SoftwareItem
	// Inicializar estructura de software
	base := db.Conectar()
	software.Name = r.FormValue("name")
	software.Desc = r.FormValue("desc")
	software.Image = r.FormValue("image")
	software.Licence = r.FormValue("licence")
	software.Git = r.FormValue("git")
	software.Website = r.FormValue("website")
	software.OS = r.FormValue("os")
	software.CreatedAt = time.Now()
	software.UpdatedAt = time.Now()
	software.IsActive = false
	// Insertar en la base de datos
	nuevoSoftware := base.Insert(&software)
	if nuevoSoftware != nil {
		log.Printf("Error al insertar %s a la base de datos.", r.FormValue("name"))
		log.Printf("Error: %v", nuevoSoftware)
		db.Desconectar(base)
		return
	}
	log.Printf("%s agregado con éxito.", r.FormValue("name"))
	db.Desconectar(base)
	// Respuesta del servidor
	json.NewEncoder(w).Encode(&software)
}
