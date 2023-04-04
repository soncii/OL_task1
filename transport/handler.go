package transport

import (
	"github.com/labstack/echo/v4"
	"login/model"
	"login/service"
	"net/http"
)

type Handler struct {
	UserService *service.UserService
}

func NewHandler(userService *service.UserService) *Handler {
	return &Handler{UserService: userService}
}
func (h Handler) CreateUser(c echo.Context) error {
	var req model.UserCreateReq
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	resp := h.UserService.Create(req)
	return c.JSON(http.StatusOK, resp)
}
