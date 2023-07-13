package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"discreto-api-crud-service/models"
	"discreto-api-crud-service/routes"
)

func main() {
	// Configura el router de Gin
	r := gin.Default()

	// Configura la conexión a la base de datos MySQL
	DB, err := gorm.Open(mysql.Open("niafqijcr5qngl30l021:pscale_pw_hIw9O6r2UDDc1kpwTyFMLtscjRN37kZOM6GQuCKQ4zC@tcp(aws.connect.psdb.cloud)/diplomatura?tls=true"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
		os.Exit(1)
	}

	// Realiza la migración de los modelos a la base de datos
	models.MigrateTables(DB)

	// Configura las rutas
	routes.SetupRoutes(r, DB)

	// Inicia el servidor
	r.Run()
}
