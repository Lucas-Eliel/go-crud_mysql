package controllers

import (
	"crud_mysql/internal/models"
	"crud_mysql/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateHandlers(r *gin.Engine, service services.Service) {
	controller := Controller{service}

	r.GET("/users", controller.GetUsers)
	r.GET("/users/:id", controller.GetUser)
	r.POST("/users", controller.PostUser)
	r.DELETE("/users/:id", controller.DeleteUser)

}

type Controller struct {
	service services.Service
}

func (c Controller) GetUsers(contexto *gin.Context) {
	users := c.service.GetUsuarios()
	contexto.IndentedJSON(http.StatusOK, users)
}

func (c Controller) GetUser(contexto *gin.Context) {
	id := contexto.Param("id")
	user := c.service.GetUsuario(id)
	contexto.IndentedJSON(http.StatusOK, user)
}

func (c Controller) PostUser(contexto *gin.Context) {
	var user models.User

	err := contexto.BindJSON(&user)

	if err != nil {
		panic(err)
	}

	c.service.PostUsuario(user)
	contexto.IndentedJSON(http.StatusCreated, user)
}

func (c Controller) DeleteUser(contexto *gin.Context) {
	id := contexto.Param("id")
	c.service.DeleteUsuario(id)
	contexto.IndentedJSON(http.StatusOK, id)
}
