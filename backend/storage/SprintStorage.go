package storage

import (
	"encoding/json"
	"io/ioutil"
	"mrnakumar.com/tamaatar/models"
)

const sprintDb = "sprints.db"

type SprintStorage interface {
	Create(sprint models.Sprint) error
	Find(uid string, name string, day int, month string, year int) (models.Sprint, bool)
}

type sprintStorageImpl struct {
	sprints map[string]models.Sprint
}

func GetSprintStorage() SprintStorage {
	sprints, err := loadAll()
	if err != nil {
		panic(err)
	}
	return sprintStorageImpl{sprints: sprints}
}

func (ss sprintStorageImpl) Create(sprint models.Sprint) error {
	ss.sprints[sprint.Id] = sprint
	err := ss.saveAll()
	if err != nil {
		delete(ss.sprints, sprint.Id)
	}
	return err
}

func (ss sprintStorageImpl) Find(uid string, name string, day int, month string, year int) (models.Sprint, bool) {
	for k, v := range ss.sprints {
		if k == name {
			if v.UserId == uid && v.Day == day && v.Month == month && v.Year == year {
				return v, true
			}
		}
	}
	return models.Sprint{}, false
}

func loadAll() (map[string]models.Sprint, error) {
	data, err := ioutil.ReadFile(sprintDb)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return map[string]models.Sprint{}, nil
	}
	var sprints []models.Sprint
	err = json.Unmarshal(data, &sprints)
	if err != nil {
		return nil, err
	}
	result := make(map[string]models.Sprint)
	for _, sprint := range sprints {
		result[sprint.Id] = sprint
	}
	return result, nil
}

func (ss sprintStorageImpl) saveAll() error {
	sprints := make([]models.Sprint, len(ss.sprints))
	i := 0
	for _, sprint := range ss.sprints {
		sprints[i] = sprint
		i++
	}
	bytes, err := json.Marshal(sprints)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(sprintDb, bytes, 0644)
}
