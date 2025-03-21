package main

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   uint32
	name string
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Neuspešno povezivanje sa bazom:", err)
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatal("Neuspešna migracija baze:", err)
	}

	log.Println("Tabela uspešno kreirana!")
}
