package handlers

import (
	"github.com/labstack/echo/v4"
	"mrnakumar.com/tamaatar/models/response"
	"mrnakumar.com/tamaatar/storage"
	"mrnakumar.com/tamaatar/utils"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
)

type HisaabHandler interface {
	GetHisaab(c echo.Context) error
}

type hisaabHandlerImpl struct {
	PromiseDb storage.VaakurutiStorage
	SprintDb  storage.SprintStorage
	lock      sync.Mutex
}

func GetHisaabHandler(promiseDb storage.VaakurutiStorage, sprintDb storage.SprintStorage) HisaabHandler {
	return hisaabHandlerImpl{
		PromiseDb: promiseDb,
		SprintDb:  sprintDb,
		lock:      sync.Mutex{},
	}
}

func (hh hisaabHandlerImpl) GetHisaab(c echo.Context) error {
	userId, found := utils.RequestUtils{}.GetUserNameHeader(c)
	if !found {
		return c.NoContent(http.StatusBadRequest)
	}
	now := time.Now()
	hh.lock.Lock()
	promises := hh.PromiseDb.GetAll(userId, now.Day(), now.Month().String(), now.Year())
	sprints := hh.SprintDb.List(userId, now.Day(), now.Month().String(), now.Year())
	hh.lock.Unlock()
	byName := make(map[string]response.HisaabResponse)
	for _, p := range promises {
		byName[p.Name] = response.HisaabResponse{
			Name:             p.Name,
			PromisedDuration: p.Duration,
		}
	}
	for _, s := range sprints {
		if _, ok := byName[s.Name]; !ok {
			byName[s.Name] = response.HisaabResponse{
				Name:             s.Name,
				FinishedDuration: s.Duration,
			}
		} else {
			byName[s.Name] = response.HisaabResponse{
				Name:             s.Name,
				PromisedDuration: byName[s.Name].PromisedDuration,
				FinishedDuration: byName[s.Name].FinishedDuration + s.Duration,
			}
		}
	}
	var result []response.HisaabResponse
	for _, v := range byName {
		result = append(result, v)
	}
	sort.Slice(result, func(i, j int) bool {
		a := result[i]
		b := result[j]
		return strings.Compare(a.Name, b.Name) <= 0
	})
	return c.JSON(http.StatusOK, result)
}
