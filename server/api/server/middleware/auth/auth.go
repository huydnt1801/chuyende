package auth

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/huydnt1801/chuyende/internal/driver"
	ss "github.com/huydnt1801/chuyende/internal/session"
	"github.com/huydnt1801/chuyende/internal/user"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type AuthInfo struct {
	User      *user.User
	Driver    *driver.Driver
	SessionID string
}

func Middleware(db *sql.DB) echo.MiddlewareFunc {
	ssSvc := ss.NewSessionService(db)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			ctx := r.Context()

			if strings.HasPrefix(r.URL.Path, "/api/v1/accounts/register") ||
				strings.HasPrefix(r.URL.Path, "/api/v1/accounts/login") ||
				strings.HasPrefix(r.URL.Path, "/api/v1/accounts/phone") {
				LogoutUser(c)
				return next(c)
			}
			sess := getSession(c)
			val, found := sess.Values["sessionId"]
			if found {
				sessID := fmt.Sprintf("%v", val)
				usr, driv, err := ssSvc.GetSession(ctx, sessID)
				if err != nil {
					return err
				}
				LoginUser(c, sessID, usr, driv)
				return next(c)
			}
			return echo.NewHTTPError(http.StatusUnauthorized, "Yêu cầu đăng nhập")
		}
	}
}

func LoginUser(c echo.Context, sessionId string, usr *user.User, driv *driver.Driver) {
	sess := getSession(c)
	sess.Values["sessionId"] = sessionId

	saveSession(c)
	SetAuthInfo(c, usr, driv, sessionId)
}

func LogoutUser(c echo.Context) {
	sess := getSession(c)
	delete(sess.Values, "sessionId")
	sess.Options.MaxAge = -1
	saveSession(c)
}

func getSession(c echo.Context) *sessions.Session {
	sess, _ := session.Get("auth", c)
	return sess
}

func saveSession(c echo.Context) {
	// skip if hook is already added
	if added, ok := c.Get(sessionSaveHookAdded).(bool); ok && added {
		return
	}
	c.Response().Before(func() {
		sess := getSession(c)
		sess.Options.HttpOnly = true
		sess.Options.Secure = true
		sess.Options.SameSite = http.SameSiteNoneMode
		sess.Save(c.Request(), c.Response())
	})
	c.Set(sessionSaveHookAdded, true)

}

func SetAuthInfo(c echo.Context, usr *user.User, driv *driver.Driver, sessionId string) {
	c.Set(authInfoKey, AuthInfo{
		User:      usr,
		Driver:    driv,
		SessionID: sessionId,
	})

}
func GetAuthInfo(c echo.Context) (AuthInfo, bool) {
	val := c.Get(authInfoKey)
	if val == nil {
		return AuthInfo{}, false
	}
	info, ok := val.(AuthInfo)
	if !ok {
		return AuthInfo{}, false
	}
	return info, true
}

const (
	authInfoKey          string = "authInfo"
	sessionSaveHookAdded string = "authSessionSaveHookAdded"
)
