package db

import (
	"fmt"
	"log"
	"os"

	"github.com/ACordoba15/be-user-maintenance/internal/domain/models"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DSN string = ""
var DB *gorm.DB

func DBConnection() (*gorm.DB, error) {
	var err error

	// Definir los detalles de la conexión
	server := os.Getenv("DB_SERVER")     // Nombre del servidor
	port := os.Getenv("DB_PORT")         // Puerto de SQL Server
	user := os.Getenv("DB_USER")         // Usuario
	password := os.Getenv("DB_PASSWORD") // Contraseña
	database := os.Getenv("DB_NAME")     // Nombre de la base de datos

	// Crear la cadena de conexión
	DSN := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=true&trustServerCertificate=true&tlsmin=1.0", user, password, server, port, database)

	// Conectarse a la base de datos usando GORM y el driver de SQL Server
	DB, err = gorm.Open(sqlserver.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Error conectando a la base de datos: ", err)
		return nil, err
	}

	// Verificar la conexión
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Error obteniendo la conexión SQL: ", err)
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("No se pudo hacer ping a la base de datos: ", err)
		return nil, err
	}

	fmt.Println("Conexión a SQL Server exitosa")
	return DB, nil
}

// MigrateDB se encarga de ejecutar las migraciones de la base de datos
func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{}, &models.Record{})
}
