package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectToDB establece una conexión con la base de datos utilizando GORM.
// Configura la base de datos MySQL con los parámetros definidos.
// Retorna:
//   - Un puntero a la instancia de GORM DB.
//   - Termina la aplicación si no se puede conectar a la base de datos.
func ConnectToDB() *gorm.DB {
	dsn := "root:password1234@tcp(127.0.0.1:3306)/databaseperson?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection established")
	return db
}
