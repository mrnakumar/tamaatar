package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"mrnakumar.com/tamaatar/models"
	"mrnakumar.com/tamaatar/models/request"
	"mrnakumar.com/tamaatar/storage"
	"mrnakumar.com/tamaatar/utils"
	"net/http"
	"sync"
	"time"
)

type VaakurutiHandler interface {
	CreatePromise(c echo.Context) error
}

type vaakurutiHandlerImpl struct {
	lock             sync.Mutex
	vaakurutiStorage storage.VaakurutiStorage
}

func GetVaakurutiHandler(vaakurutiStorage storage.VaakurutiStorage) VaakurutiHandler {
	return vaakurutiHandlerImpl{vaakurutiStorage: vaakurutiStorage,
		lock: sync.Mutex{}}
}

func (sh vaakurutiHandlerImpl) CreatePromise(c echo.Context) error {
	createReq := new(request.CreateVakkurutiRequest)
	err := c.Bind(createReq)
	if err != nil {
		return err
	}
	date := time.Now().UTC()
	userId, found := utils.RequestUtils{}.GetHeader(c)
	if !found {
		return c.NoContent(http.StatusBadRequest)
	}
	sprint := models.Vaakuruti{
		Id:       uuid.New().String(),
		UserId:   userId,
		Name:     createReq.Name,
		Duration: createReq.Duration,
		Day:      date.Day(),
		Month:    date.Month().String(),
		Year:     date.Year(),
	}
	sh.lock.Lock()
	err = sh.vaakurutiStorage.Create(sprint)
	sh.lock.Unlock()
	if err != nil {
		c.Logger().Errorf("error in creating promise %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusCreated)
}
