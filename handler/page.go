package handler

import (
	"template-inator/model"
	"template-inator/view/page"

	"github.com/labstack/echo/v4"
)

type PageHandler struct{}

func (h PageHandler) HandleDashShow(c echo.Context) error {
	user := model.User{Email: "josh@me.com"}
	return render(c, page.Dash(user))
} 