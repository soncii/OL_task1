package transport

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"login/config"
	"login/model"
	"login/service"
	"net/http"
)

type Handler struct {
	JWTSecret          []byte
	UserService        *service.UserService
	BookService        *service.BookService
	RecordService      *service.RecordService
	TransactionService *service.TransactionService
}

func NewHandler(userService *service.UserService, bookService *service.BookService, recordService *service.RecordService,
	transactionService *service.TransactionService, cfg *config.Config) *Handler {
	return &Handler{UserService: userService, BookService: bookService, RecordService: recordService,
		TransactionService: transactionService, JWTSecret: []byte(cfg.JWTSecret)}
}
func (h *Handler) CreateTransaction(c echo.Context) error {
	req := model.TransactionCreateReq{}
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	res, err := h.TransactionService.Create(req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
func (h *Handler) CreateUser(c echo.Context) error {
	var req model.UserCreateReq
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	fmt.Print(req)
	resp := h.UserService.Create(req)
	return c.JSON(http.StatusOK, resp)
}
func (h *Handler) CreateBook(c echo.Context) error {
	var req model.BookCreateReq
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	fmt.Print(req)
	resp := h.BookService.Create(req)
	return c.JSON(http.StatusOK, resp)
}
func (h *Handler) Validate(c echo.Context) error {
	var req model.UserCreateReq
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	b := h.UserService.Validate(req)
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
	res, err := h.UserService.UpdatePassword(req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
func (h *Handler) GetRecords(c echo.Context) error {
	val := c.Get("email")
	email, err := val.(string)
	if err {
		return fmt.Errorf("could not get email")
	}
	return c.JSON(http.StatusOK, h.UserService.GetUserRecords(email))
}
func (h *Handler) CreateRecord(c echo.Context) error {
	req := model.RecordCreateReq{}
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	d, err := h.RecordService.Create(req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, d)
}
func (h *Handler) GetUsersWithBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, h.UserService.GetUsers())
}
func (h *Handler) GetUsersLastMonth(c echo.Context) error {
	return c.JSON(http.StatusOK, h.UserService.GetUsersLastMonth())
}
