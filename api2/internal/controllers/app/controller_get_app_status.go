package app

import (
	"net/http"

	"github.com/HenriqueOtsuka/cep-weather/api2/internal/handler"
)

type ControllerGetAppStatus struct {
	handler.TransactionControllerImpl
}

type appstatus struct {
	Status string `json:"status"`
}

func (c *ControllerGetAppStatus) Execute(payload interface{}) (result handler.ResponseController) {
	result = handler.NewJsonResponseController()
	result.SetResult(http.StatusOK, appstatus{Status: "OK"})
	return
}

func NewControllerGetAppStatus() handler.Controller {
	return &ControllerGetAppStatus{}
}
