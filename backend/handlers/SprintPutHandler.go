package handlers

import (
	"github.com/labstack/echo/v4"
	"mrnakumar.com/tamaatar/models"
	"mrnakumar.com/tamaatar/storage"
	"net/http"
)

func SprintPutHandler(c echo.Context) error {
	sprint := new(models.Sprint)
	c.Bind(sprint)
	 c.String(http.StatusOK,"o")
	return storage.SaveSprint(sprint)
}