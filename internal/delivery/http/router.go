package http

import "github.com/gin-gonic/gin"

func NewRouter(registroH *RegistroHandler) *gin.Engine {
	r := gin.Default()

	registros := r.Group("/registros")
	{
		registros.POST("/create", registroH.Create)
		registros.GET("/find-all", registroH.FindAll)
	}

	return r
}