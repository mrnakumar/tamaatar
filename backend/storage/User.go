package storage

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mrnakumar.com/tamaatar/models"
	"strings"
	"sync"
)

const UserNameAndPasswordSeparator = ";"
const recordSeparator = "\n"
const userDb = "users.db"

type UserStorage interface {
	Save(user models.User) error
	Exists(name string, password string) bool
}

type fileStorage struct {
	lock  sync.Mutex
	users map[string]models.User
}

func GetUserStorage() (UserStorage, error) {
	users, err := readAll()
	if err != nil {
		return nil, err
	}
	return fileStorage{lock: sync.Mutex{}, users: users}, nil
}

// saves a users if does not exist
// or if exists but there is change
func (s fileStorage) Save(user models.User) error {
	s.lock.Lock()
	s.users[user.Name] = models.User{
		Name:   user.Name,
		Passwd: user.Passwd,
	}
	var data bytes.Buffer

	i := 0
	for name, user := range s.users {
		if i > 0 {
			data.WriteString(recordSeparator)
		}
		data.WriteString(fmt.Sprintf("%v%v%v", name, UserNameAndPasswordSeparator, user.Passwd))
		i++
	}
	err := ioutil.WriteFile(userDb, data.Bytes(), 0644)
	s.lock.Unlock()
	return err
}

func readAll() (map[string]models.User, error) {
	data, err := ioutil.ReadFile(userDb)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return make(map[string]models.User), nil
	}
	usersData := strings.Split(string(data), "\n")
	users := make(map[string]models.User)
	for _, e := range usersData {
		uName := strings.Split(e, UserNameAndPasswordSeparator)[0]
		passwd := strings.Split(e, UserNameAndPasswordSeparator)[1]
		users[uName] = models.User{
			Name:   uName,
			Passwd: passwd,
		}
	}
	return users, nil
}

func (s fileStorage) Exists(name string, password string) bool {
	s.lock.Lock()
	v, ok := s.users[name]
	if !ok {
		return false
	}
	s.lock.Unlock()
	return v.Name == name && v.Passwd == password
}
