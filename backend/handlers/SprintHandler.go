package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"mrnakumar.com/tamaatar/models"
	"mrnakumar.com/tamaatar/models/request"
	"mrnakumar.com/tamaatar/storage"
	"net/http"
	"sync"
	"time"
)

type SprintHandler interface {
	CreateSprint(c echo.Context) error
}

type sprintHandlerImpl struct {
	lock     sync.Mutex
	sprintDb storage.SprintStorage
}

func GetSprintHandler(sprintDb storage.SprintStorage) SprintHandler {
	return sprintHandlerImpl{sprintDb: sprintDb,
		lock: sync.Mutex{}}
}

func (sh sprintHandlerImpl) CreateSprint(c echo.Context) error {
	createReq := new(request.CreateSprintRequest)
	err := c.Bind(createReq)
	if err != nil {
		return err
	}
	date := time.Now().UTC()
	sprint := models.Sprint{
		Id:       uuid.New().String(),
		UserId:   c.Request().Header.Get("uid"),
		Name:     createReq.Name,
		Duration: createReq.Duration,
		Day:      date.Day(),
		Month:    date.Month().String(),
		Year:     date.Year(),
	}
	sh.lock.Lock()
	err = sh.sprintDb.Create(sprint)
	sh.lock.Unlock()
	if err != nil {
		c.Logger().Errorf("error in creating sprint %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusCreated)
}
