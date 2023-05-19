package main

import (
	"github.com/gin-gonic/gin"
	"crud_mysql/pkg/mysql"
	"crud_mysql/internal/controllers"
	"crud_mysql/internal/repositories"
	"crud_mysql/internal/services"
)

func main() {

	r := gin.Default()
	
	db := mysql.CreateClient()

	repository := repositories.NewRepositoryImpl(db)
	service := services.NewServiceImpl(repository)

	controllers.CreateHandlers(r, service)

	r.Run("localhost:8090")


}
