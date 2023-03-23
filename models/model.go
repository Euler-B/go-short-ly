package model

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Goly struct {
	ID       uint64 `json:"id" gorm:"primary_key"`
	Redirect string `json:"redirect" gorm:"not_null"`
	Goly     string `json:"goly" gorm:"unique;not_null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}

func Setup() {
	values := godotenv.Load(".env") // <-- Esta funcion devuelve un error, por lo general asi llama la variable que captura el error, pero le cambie el a values para que no tenga conflictos con la variable que maneja la apertura de la base de datos.
	if values != nil {
		log.Fatal("Error cargando el archivo .env")
	}
	dsn := os.Getenv("DATABASE_URI")

	var err error
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Goly{})
	if err != nil {
		fmt.Println(err)
	}
}
