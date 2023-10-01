package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// @Title		WireGuard Configuration Manager API
// @Version		0.0.1
// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	err := godotenv.Load()
	handleErr(err)

	gin.SetMode(gin.ReleaseMode)
	initAPI()
}
