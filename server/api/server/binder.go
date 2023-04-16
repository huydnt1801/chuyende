package server

import "github.com/labstack/echo/v4"

type CustomBinder struct{}

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
	// You may use default binder
	defaultBinder := new(echo.DefaultBinder)
	if err = defaultBinder.Bind(i, c); err != nil {
		return
	}

	// Define your custom implementation here

	// Using custom validator to validate
	if err = c.Validate(i); err != nil {
		return
	}
	return
}
