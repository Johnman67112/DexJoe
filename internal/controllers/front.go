package controllers

import (
	"net/http"

	"github.com/Johnman67112/dexJoe/internal/domain"
	"github.com/Johnman67112/dexJoe/internal/infra"
	"github.com/gin-gonic/gin"
)

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
