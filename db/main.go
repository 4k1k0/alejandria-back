package db

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	model "./models"
	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

type Credentials struct {
	User     string
	Password string
	Addr     string
	Database string
}

func CargarCredenciales() (Credentials, error) {
	file, _ := os.Open("db/credentials.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := Credentials{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return conf, err
	}
	return conf, nil

}

func Conectar() *pg.DB {
	conf, err := CargarCredenciales()
	if err != nil {
		fmt.Printf("Error al cargar configuración: %v", err)
		os.Exit(100)
	}
	options := &pg.Options{
		User:     conf.User,
		Password: conf.Password,
		Addr:     conf.Addr,
		Database: conf.Database,
	}
	db := pg.Connect(options)

	if db == nil {
		log.Printf("Error al conectar a la base de datos")
		os.Exit(100)
	}
	log.Printf("Conexión a la base de datos exitosa")

	return db

}

func CrearTablas() error {
	db := Conectar()
	options := &orm.CreateTableOptions{}
	createErrSoft := db.CreateTable(&model.SoftwareItem{}, options)
	if createErrSoft != nil {
		log.Printf("Error al crear la tabla de software. Error: %v\n", createErrSoft)
	}
	createErrApp := db.CreateTable(&model.AppItem{}, options)
	if createErrApp != nil {
		log.Printf("Error al crear la tabla de apps. Error: %v\n", createErrApp)
	}
	Desconectar(db)
	return nil
}

func Desconectar(db *pg.DB) error {
	err := db.Close()
	if err != nil {
		log.Printf("Error al cerrar la base de datos...\nError: %v", err)
		return err
	}
	log.Printf("Base de datos cerrada\n")
	return nil
}
