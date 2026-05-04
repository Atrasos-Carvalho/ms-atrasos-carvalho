package http

import (
	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func NewRouter(registroH *RegistroHandler) *gin.Engine {
	r := gin.Default()
	
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))

	registros := r.Group("/registros")
	{
		registros.POST("/create", registroH.Create)
		registros.GET("/find-all", registroH.FindAll)
		registros.GET("/find-count", registroH.FindCount)
	}


	return r
}