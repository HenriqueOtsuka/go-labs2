package app

import (
	"github.com/HenriqueOtsuka/cep-weather/api1/internal/handler"
	"github.com/gin-gonic/gin"
)

// GET Endpoints

// [GET] /appstatus
// @Summary Get app status
// @Tags App
// @Description Get app status
// @Accept json
// @Produce json
// @Success 200 {object} string "App Status"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /appstatus [get]
func HandleGetAppStatus(c *gin.Context) {
	handler.RequestWithController(c, nil, NewControllerGetAppStatus())
}
