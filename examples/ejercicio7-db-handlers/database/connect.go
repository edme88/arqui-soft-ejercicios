package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error){
	//Cargar los valores desde las variables de entorno
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dns := fmt.Sprintf("%s:%v@tcp(%s:%v)/%v", 
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_NAME"))
	
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{}) //Se utiliza el pacquete de forma indirecta
	if err != nil {
		return nil, fmt.Errorf("error al conectar a la base de datos: %v", err)
	}
	log.Println("Conexión a la BD exitosa") //Sin el fmt y usando log indica la fecha y hora

	return db, nil
}
