package auth

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// SessionMaxAge session cookie max age in seconds. Default: 30 days.
var SessionMaxAge = 30 * 86400 // 30 days

func Middleware(db *sql.DB) echo.MiddlewareFunc {
	// client := entutil.NewClientFromDB(db)
	// cfg := config.MustParseConfig()

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			// w := c.Response()
			// ctx := r.Context()

			if strings.HasPrefix(r.URL.Path, "/static") {
				return next(c)
			}

			sess := getSession(c)
			val, found := sess.Values["sessionId"]
			if found {
				fmt.Println(val)
			}
			return next(c)
		}
	}
}

func LoginUser(c echo.Context, sessionId string) {
	sess := getSession(c)
	sess.Values["sessionId"] = sessionId

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
		sess.Options.MaxAge = SessionMaxAge
		sess.Options.HttpOnly = true
		sess.Options.Secure = true
		sess.Options.SameSite = http.SameSiteNoneMode
		sess.Save(c.Request(), c.Response())
	})
	c.Set(sessionSaveHookAdded, true)

}

const (
	authInfoKey          string = "authInfo"
	sessionSaveHookAdded string = "authSessionSaveHookAdded"
)
