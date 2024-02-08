package ui

import (
	"github.com/Johnman67112/dexJoe/internal/controllers"
	"github.com/gin-gonic/gin"

	"github.com/Johnman67112/dexJoe/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/pokemon", controllers.GetPokemon)
	r.GET("/pokemon/:id", controllers.GetOnePokemon)
	r.POST("/pokemon", controllers.RegisterPokemon)
	r.DELETE("/pokemon/:id", controllers.DeletePokemon)
	r.PATCH("/pokemon/:id", controllers.EditPokemon)

	r.GET("/index", controllers.ShowIndex)
	r.NoRoute(controllers.RouteNotFound)

	r.Run()
}
