package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"login/config"
	"login/model"
	"login/service"
	"net/http"
)

type Handler struct {
	JWTSecret []byte
	Manager   *service.Manager
}

func NewHandler(manager *service.Manager, cfg *config.Config) *Handler {
	return &Handler{Manager: manager, JWTSecret: []byte(cfg.JWTSecret)}
}

func (h *Handler) CreateUser(c echo.Context) error {
	var req model.UserCreateReq
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	fmt.Print(req)
	resp, err := h.Manager.UserService.Create(c.Request().Context(), req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *Handler) CreateBook(c echo.Context) error {
	var req model.BookCreateReq
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	fmt.Print(req)
	resp, err := h.Manager.BookService.Create(c.Request().Context(), req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *Handler) Validate(c echo.Context) error {
	var req model.UserCreateReq
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	b, err := h.Manager.UserService.Validate(c.Request().Context(), req)
	if err != nil {
		return err
	}
	if b {
		token, err := h.GenerateJWT(req.Email)
		if err != nil {
			// Handle error
			return err
		}
		c.Response().Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return c.JSON(http.StatusOK, "Success")
	}
	return c.JSON(http.StatusForbidden, "Invalid credentials")
}

func (h *Handler) ChangePassword(c echo.Context) error {
	req := model.UserEmailPassReq{}
	err := c.Bind(&req)

	if err != nil {
		return err
	}
	val := c.Get("email")
	req.Email, _ = val.(string)
	res, err := h.Manager.UserService.UpdatePassword(c.Request().Context(), req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
func (h *Handler) GetRecords(c echo.Context) error {
	val := c.Get("email")
	email, err := val.(string)
	if !err {
		fmt.Println("could not get email")
		return fmt.Errorf("could not get email")
	}
	records, err1 := h.Manager.UserService.GetUserRecords(c.Request().Context(), email)
	if err1 != nil {
		fmt.Println(err)
		return err1
	}
	return c.JSON(http.StatusOK, records)
}
func (h *Handler) CreateRecord(c echo.Context) error {
	req := model.RecordCreateReq{}
	err := c.Bind(&req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	d, err := h.Manager.RecordService.Create(c.Request().Context(), req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, d)
}
func (h *Handler) GetUsersWithBooks(c echo.Context) error {

	users, err := h.Manager.UserService.GetUsers(c.Request().Context())
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, users)
}
func (h *Handler) GetUsersLastMonth(c echo.Context) error {
	users, err := h.Manager.UserService.GetUsersLastMonth(c.Request().Context())
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, users)
}
