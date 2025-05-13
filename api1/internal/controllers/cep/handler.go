package cep

import (
	"github.com/HenriqueOtsuka/cep-weather/api1/internal/handler"
	"github.com/HenriqueOtsuka/cep-weather/api1/internal/models"
	"github.com/gin-gonic/gin"
)

// POST Endpoints

// [POST] /cep
// @Summary Get temperature by zipcode
// @Tags CEP
// @Description Get temperature by zipcode
// @Accept json
// @Produce json
// @Param cep body models.CepPayload true "CEP"
// @Success 200 {object} models.CepResponse "Temperature"
// @Failure 422 {string} string "invalid zipcode"
// @Failure 500 {string} string "internal server error"
// @Router /cep [post]
func HandleGetCepTemperature(c *gin.Context) {
	handler.RequestWithController(c, &models.CepPayload{}, NewControllerGetCepTemperature())
}
