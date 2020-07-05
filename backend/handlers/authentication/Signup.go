package authentication

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"mrnakumar.com/tamaatar/models"
	"mrnakumar.com/tamaatar/storage"
	"net/http"
	"strings"
	"sync"
)

type SignUpHandler interface {
	SignUp(c echo.Context) error
}

type signUpHandlerImpl struct {
	userDb storage.UserStorage
	lock   sync.Mutex
}

func CreateSignUpHandler(userDb storage.UserStorage) SignUpHandler {
	return signUpHandlerImpl{userDb: userDb, lock: sync.Mutex{}}
}

func (sh signUpHandlerImpl) SignUp(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return echo.ErrBadRequest
	}
	if strings.Contains(user.Name, ";") {
		return c.JSON(http.StatusBadRequest,
			fmt.Sprintf("%v is not allowed in usernmae", storage.UserNameAndPasswordSeparator))
	}
	if strings.Contains(user.Passwd, ";") {
		return c.JSON(http.StatusBadRequest,
			fmt.Sprintf("%v is not allowed in password", storage.UserNameAndPasswordSeparator))
	}
	sh.lock.Lock()
	if sh.userDb.Exists(user.Name, user.Passwd) {
		// user already exists
		return echo.ErrBadRequest
	}
	err := sh.userDb.Save(*user)
	if err != nil {
		return echo.ErrInternalServerError
	}
	sh.lock.Unlock()
	return c.NoContent(http.StatusCreated)
}
