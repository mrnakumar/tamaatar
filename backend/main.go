package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"mrnakumar.com/tamaatar/handlers"
	"mrnakumar.com/tamaatar/handlers/authentication"
	"mrnakumar.com/tamaatar/models/tokens"
	"mrnakumar.com/tamaatar/storage"
	"net/http"
)

func main() {
	userStorage, err := storage.GetUserStorage()
	if err != nil {
		log.Fatal(err)
	}

	userVerifier := handlers.CreateUserVerifier(userStorage)
	e := echo.New()
	secrets := accessSecret(e)
	auth := authentication.Authenticator{AccessSecret: []byte(secrets.Access),
		RefreshSecret: []byte(secrets.Refresh),
		UserVerifier:  userVerifier,
	}
	sprintDb := storage.GetSprintStorage()
	sprintHandler := handlers.GetSprintHandler(sprintDb)
	promiseDb := storage.GetVaakurutiStorage()
	promiseHandler := handlers.GetVaakurutiHandler(promiseDb)
	e.Static("/", "frontEnd/")
	e.PUT("/signup", authentication.CreateSignUpHandler(userStorage).SignUp)
	e.POST("/login", auth.Login)
	e.POST("/refresh", auth.Refresh)
	e.GET("/clogin", func(context echo.Context) error {
		return context.NoContent(http.StatusOK)
	}, auth.CheckLogin)
	e.GET("/logout", auth.Logout)
	e.POST("/createSprint", sprintHandler.CreateSprint, auth.CheckLogin)
	e.POST("/createPromise", promiseHandler.CreatePromise, auth.CheckLogin)
	e.GET("/timeBySprintName", sprintHandler.TimeBySprintName, auth.CheckLogin)
	e.GET("/getHisaab", handlers.GetHisaabHandler(promiseDb, sprintDb).GetHisaab, auth.CheckLogin)
	e.Logger.Fatal(e.Start(":8080"))
}

func accessSecret(e *echo.Echo) tokens.TokenSecrets {
	secrets, err := ioutil.ReadFile(".ke_baat_s")
	if err != nil {
		e.Logger.Fatal(err)
	}
	var tokenSecrets tokens.TokenSecrets
	err = json.Unmarshal(secrets, &tokenSecrets)
	if err != nil {
		e.Logger.Fatal(err)
	}
	return tokenSecrets
}
