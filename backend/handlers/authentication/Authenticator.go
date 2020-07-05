package authentication

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"mrnakumar.com/tamaatar/handlers"
	"mrnakumar.com/tamaatar/models"
	"net/http"
	"time"
)

const expire_access_after = 15           // 15 minutes
const expire_refresh_after = 7 * 24 * 60 // 7 days

type Authenticator struct {
	AccessSecret  []byte
	RefreshSecret []byte
	UserVerifier  handlers.UserVerifier
}

func (a Authenticator) CheckLogin(hf echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("access")
		if err != nil {
			return echo.ErrUnauthorized
		}
		access := cookie.Value
		token, err := jwt.Parse(access, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return a.AccessSecret, nil
		})

		if err != nil {
			return echo.ErrUnauthorized
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.ErrUnauthorized
		}
		uid, ok := claims["id"]
		if !ok {
			// id should be present
			return echo.ErrUnauthorized
		}
		// TODO: may we also verify xsrf token
		setUserName(uid.(string), c)
		return hf(c)
	}
}

func (a Authenticator) Login(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(echo.ErrBadRequest.Code, "invalid json")
	}
	if !a.UserVerifier.CheckCredentials(user.Name, user.Passwd) {
		return c.JSON(echo.ErrUnauthorized.Code, "invalid user name or password")
	}
	err := a.setTokens(c, user.Name)
	if err != nil {
		return echo.ErrInternalServerError
	}
	setUserName(user.Name, c)
	return c.NoContent(http.StatusOK)
}

func (a Authenticator) Refresh(c echo.Context) error {
	rt, err := c.Cookie("refresh")
	if err != nil {
		return echo.ErrUnauthorized
	}
	refreshToken, err := jwt.Parse(rt.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.RefreshSecret, nil
	})
	if err != nil {
		return echo.ErrUnauthorized
	}
	rtClaims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return echo.ErrUnauthorized
	}
	uname, ok := rtClaims["id"]
	if !ok {
		return echo.ErrUnauthorized
	}
	err = a.setTokens(c, uname.(string))
	if err != nil {
		return echo.ErrInternalServerError
	}
	setUserName(uname.(string), c)
	return nil
}

func (e Authenticator) Logout(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
func (a Authenticator) setTokens(c echo.Context, uname string) error {
	at, err := a.createAccessToken(uname)
	if err != nil {
		return echo.ErrUnauthorized
	}
	rt, err := a.createRefreshToken(uname)
	if err != nil {
		return echo.ErrUnauthorized
	}
	setCookie(c, "access", at)
	setCookie(c, "refresh", rt)
	return nil
}

func (a Authenticator) createAccessToken(uname string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = uname
	claims["exp"] = time.Now().Add(time.Minute * expire_access_after).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	access, err := at.SignedString(a.AccessSecret)
	return access, err
}

func (a Authenticator) createRefreshToken(uname string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = uname
	claims["exp"] = time.Now().Add(time.Minute * expire_refresh_after).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	access, err := at.SignedString(a.RefreshSecret)
	return access, err
}

func setCookie(c echo.Context, name string, value string) {
	c.SetCookie(&http.Cookie{
		Name:     name,
		MaxAge:   15 * 60, // same as jwt expiration time
		Secure:   false,   // TODO: change to true in prod
		Value:    value,
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
	})
}

func setUserName(userName string, c echo.Context) {
	c.Request().Header.Set("uid", userName)
}
