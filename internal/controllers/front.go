package controllers

import (
	"net/http"

	"github.com/Johnman67112/dexJoe/internal/domain"
	"github.com/Johnman67112/dexJoe/internal/infra"
	"github.com/gin-gonic/gin"
)

// ShowIndex godoc
// @Summary      Show index page
// @Description  Route to show index
// @Tags         front
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Failure      400  {object}  error
// @Router       /index [get]
func ShowIndex(c *gin.Context) {
	var pokemon []domain.Pokemon
	infra.DB.Find(&pokemon)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"pokemon": pokemon,
	})
}

func RouteNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
