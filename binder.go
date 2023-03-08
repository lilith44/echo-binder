package echo_binder

import (
	"github.com/labstack/echo/v4"
	"github.com/mcuadros/go-defaults"
)

type Binder struct {
	binder *echo.DefaultBinder
}

func New() *Binder {
	return &Binder{
		binder: new(echo.DefaultBinder),
	}
}

func (b *Binder) Bind(i any, c echo.Context) error {
	err := b.binder.Bind(i, c)
	if err != nil {
		return err
	}

	defaults.SetDefaults(i)
	return c.Validate(i)
}
