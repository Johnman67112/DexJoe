package ui

import (
	"github.com/Johnman67112/dexJoe/internal/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/students", controllers.GetPokemon)
	r.GET("/students/:id", controllers.GetOnePokemon)
	r.POST("/students", controllers.RegisterPokemon)
	r.DELETE("/students/:id", controllers.DeletePokemon)
	r.PATCH("/students/:id", controllers.EditPokemon)

	r.Run()
}
