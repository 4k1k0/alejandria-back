package main

import (
	db "./db"
	rutas "./rutas"
)

func main() {
	// Crear tablas de la base de datos si es que no existen
	db.CrearTablas()
	// Crear el servidor web
	rutas.Init()
}
