package rutas

import (
	"encoding/json"
	"fmt"
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
	var software modelo.SoftwareItem
	_ = json.NewDecoder(r.Body).Decode(&software)
	fmt.Println(software)
	// Insertar en la base
	base := db.Conectar()
	software.CreatedAt = time.Now()
	software.UpdatedAt = time.Now()
	nuevoSoftware := base.Insert(&software)
	if nuevoSoftware != nil {
		log.Printf("Error al insertar nuevo software a la base de datos. Error: %v", nuevoSoftware)
		db.Desconectar(base)
		return
	}
	log.Printf("Nuevo software agregado con éxito...")
	db.Desconectar(base)
	// Respuesta del servidor
	json.NewEncoder(w).Encode(&software)
}
