package storage

import (
	"encoding/json"
	"io/ioutil"
	"mrnakumar.com/tamaatar/models"
)

const promiseDb = "promise.db"

type VaakurutiStorage interface {
	Create(sprint models.Vaakuruti) error
}

type vaakurutiStorageImpl struct {
	promises map[string]models.Vaakuruti
}

func GetVaakurutiStorage() VaakurutiStorage {
	promises, err := loadAllPromises()
	if err != nil {
		panic(err)
	}
	return vaakurutiStorageImpl{promises: promises}
}

func (ss vaakurutiStorageImpl) Create(vaakuruti models.Vaakuruti) error {
	ss.promises[vaakuruti.Id] = vaakuruti
	err := ss.saveAll()
	if err != nil {
		delete(ss.promises, vaakuruti.Id)
	}
	return err
}

func loadAllPromises() (map[string]models.Vaakuruti, error) {
	data, err := ioutil.ReadFile(promiseDb)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return map[string]models.Vaakuruti{}, nil
	}
	var promises []models.Vaakuruti
	err = json.Unmarshal(data, &promises)
	if err != nil {
		return nil, err
	}
	result := make(map[string]models.Vaakuruti)
	for _, promise := range promises {
		result[promise.Id] = promise
	}
	return result, nil
}

func (ss vaakurutiStorageImpl) saveAll() error {
	promises := make([]models.Vaakuruti, len(ss.promises))
	i := 0
	for _, sprint := range ss.promises {
		promises[i] = sprint
		i++
	}
	bytes, err := json.Marshal(promises)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(promiseDb, bytes, 0644)
}
