package controllers

import (
	"net/http"

	"github.com/Johnman67112/dexJoe/internal/domain"
	"github.com/Johnman67112/dexJoe/internal/infra"
	"github.com/gin-gonic/gin"
)

func RegisterPokemon(c *gin.Context) {
	var pokemon domain.Pokemon

	if err := c.ShouldBindJSON(&pokemon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := domain.PokemonValidator(&pokemon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	infra.DB.Create(&pokemon)

	c.JSON(http.StatusOK, pokemon)
}

func GetPokemon(c *gin.Context) {
	var Pokemon []domain.Pokemon
	infra.DB.Find(&Pokemon)
	c.JSON(200, Pokemon)
}

func GetOnePokemon(c *gin.Context) {
	var pokemon domain.Pokemon
	number := c.Params.ByName("number")

	infra.DB.First(&pokemon, number)

	if pokemon.Number == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Pokemon not found"})
		return
	}

	c.JSON(http.StatusOK, pokemon)
}

func DeletePokemon(c *gin.Context) {
	var pokemon domain.Pokemon
	number := c.Params.ByName("number")

	infra.DB.Delete(&pokemon, number)

	c.JSON(http.StatusOK, gin.H{"data": "Pokemon deleted sucessfully"})
}

func EditPokemon(c *gin.Context) {
	var pokemon domain.Pokemon
	number := c.Params.ByName("number")

	infra.DB.First(&pokemon, number)

	if err := c.ShouldBindJSON(&pokemon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := domain.PokemonValidator(&pokemon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	infra.DB.Model(&pokemon).UpdateColumns(pokemon)
	c.JSON(http.StatusOK, pokemon)
}
