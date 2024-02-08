package controllers

import (
	"net/http"

	"github.com/Johnman67112/dexJoe/internal/domain"
	"github.com/Johnman67112/dexJoe/internal/infra"
	"github.com/gin-gonic/gin"

	_ "github.com/swaggo/swag/example/celler/httputil"
	_ "github.com/swaggo/swag/example/celler/model"
)

// RegisterPokemon godoc
// @Summary      Register a new Pokemon
// @Description  With params register a new Pokemon
// @Tags         Pokemon
// @Accept       json
// @Produce      json
// @Param        Pokemon  body  domain.Pokemon  true  "Pokemon Model"
// @Success      200  {object}  domain.Pokemon
// @Failure      400  {object}  error
// @Router       /pokemon [post]
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

// GetPokemon godoc
// @Summary      Show all pokemon
// @Description  Route to show all pokemon
// @Tags         pokemon
// @Accept       json
// @Produce      json
// @Success      200  {object}  []domain.Pokemon
// @Failure      400  {object}  error
// @Router       /pokemon [get]
func GetPokemon(c *gin.Context) {
	var Pokemon []domain.Pokemon
	infra.DB.Find(&Pokemon)
	c.JSON(200, Pokemon)
}

// GetOnePokemon godoc
// @Summary      Show one pokemon
// @Description  Route to show one pokemon with number
// @Tags         pokemon
// @Accept       json
// @Produce      json
// @Success      200  {object}  domain.Pokemon
// @Failure      400  {object}  error
// @Router       /pokemon [get]
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

// DeletePokemon godoc
// @Summary      Deletes a Pokemon
// @Description  With number delete a Pokemon
// @Tags         Pokemon
// @Accept       json
// @Produce      json
// @Success      200  {object}  domain.Pokemon
// @Failure      400  {object}  error
// @Router       /pokemon [delete]
func DeletePokemon(c *gin.Context) {
	var pokemon domain.Pokemon
	number := c.Params.ByName("number")

	infra.DB.Delete(&pokemon, number)

	c.JSON(http.StatusOK, gin.H{"data": "Pokemon deleted sucessfully"})
}

// EditPokemon godoc
// @Summary      Edit a Pokemon
// @Description  With params edit a Pokemon
// @Tags         Pokemon
// @Accept       json
// @Produce      json
// @Param        Pokemon  body  domain.Pokemon  true  "Pokemon Model"
// @Success      200  {object}  domain.Pokemon
// @Failure      400  {object}  error
// @Router       /pokemon [patch]
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
