package http

import (
	"fmt"
	"net/http"
	"strconv"

	"login/config"
	"login/model"
	"login/service"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	TransactionURL string
	JWTSecret      []byte
	Manager        *service.Manager
}

func NewHandler(manager *service.Manager, cfg *config.Config) *Handler {
	return &Handler{
		Manager:        manager,
		JWTSecret:      []byte(cfg.JWTSecret),
		TransactionURL: cfg.TransactionsURL}
}

// CreateUser godoc
//
// @Summary Create a new User
// @Description  Saves User
// @Tags         User
// @Param input  body model.UserCreateReq true "User details"
// @Accept       json
// @Produce      json
// @Success      200  {object} model.UserCreateResp
// @Failure 500 {string} string "Something went wrong!"
// @Router       /user [post]
func (h *Handler) CreateUser(c echo.Context) error {
	var req model.UserCreateReq
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong!")
	}
	resp, err := h.Manager.UserService.Create(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong!")
	}
	return c.JSON(http.StatusCreated, resp)
}

// CreateBook godoc
// @Summary Create a new book
// @Description Create and save a new book
// @Tags Book
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body model.BookCreateReq true "Book details"
// @Success 201 {object} model.BookCreateResp
// @Failure 500 {string} string "Something went wrong!"
// @Router /book [post]
func (h *Handler) CreateBook(c echo.Context) error {
	var req model.BookCreateReq
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong!")
	}
	fmt.Print(req)
	resp, err := h.Manager.BookService.Create(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong!")
	}
	return c.JSON(http.StatusCreated, resp)
}

// Validate godoc
// @Summary Validate user credentials
// @Description Validate user credentials and generate JWT token
// @Tags User
// @Accept json
// @Produce json
// @Param input body model.UserCreateReq true "User details"
// @Success 200 {string} string "Success"
// @Failure 401 {string} string "Invalid credentials"
// @Failure 500 {string} string "Something went wrong!"
// @Router /login [post]
func (h *Handler) Validate(c echo.Context) error {
	var req model.UserCreateReq
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong!")
	}
	b, err := h.Manager.UserService.Validate(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong!")
	}
	if b {
		token, err := h.GenerateJWT(req.Email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Something went wrong!")
		}
		c.Response().Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return c.JSON(http.StatusOK, "Success")
	}
	return c.JSON(http.StatusForbidden, "Invalid credentials")
}

// ChangePassword godoc
// @Summary Change user password
// @Description Change user password by providing current password and new password
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Param input body model.UserEmailPassReq true "User password details"
// @Success 200 {object} model.UserChangePassResp
// @Failure 500 {object} string "Something went wrong!"
// @Security BearerAuth
// @Router /password [put]
func (h *Handler) ChangePassword(c echo.Context) error {
	req := model.UserEmailPassReq{}
	err := c.Bind(&req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong!")
	}
	val := c.Get("email")
	req.Email, _ = val.(string)
	res, err := h.Manager.UserService.UpdatePassword(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong!")
	}
	return c.JSON(http.StatusOK, res)
}

// GetRecords godoc
//
// @Summary Get user's records
// @Description  Gets all records of a user
// @Tags         User
// @Param Authorization header string true "JWT Token"
// @Accept       json
// @Produce      json
// @Success      200  {array} model.Record
// @Failure      500  {string} error "Something went wrong!"
// @Security BearerAuth
// @Router       /records [get]
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
		return c.JSON(http.StatusInternalServerError, "Something went wrong!")
	}
	return c.JSON(http.StatusOK, records)
}

// CreateRecord godoc
//
// @Summary Borrow book
// @Description  Registers the book borrowing
// @Tags         Record
// @Param Authorization header string true "JWT Token"
// @Param input body model.RecordCreateReq true "Record details"
// @Accept       json
// @Produce      json
// @Success      200  {array} model.Record
// @Failure      400  {string} error ""
// @Failure      500  {string} error "Something went wrong!"
// @Security BearerAuth
// @Router       /borrow [post]
func (h *Handler) CreateRecord(c echo.Context) error {
	req := model.RecordCreateReq{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong!")
	}
	d, err := h.Manager.RecordService.CreateRecord(c.Request().Context(), req, h.TransactionURL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong!")
	}
	err = h.Manager.BookService.UpdateBookRevenue(c.Request().Context(), model.BookRevenue{BookID: d.BookID, Revenue: req.Price})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong!")
	}
	return c.JSON(http.StatusCreated, model.RecordCreateResp{RID: d.ID, CreatedAt: d.CreatedAt})
}

// GetUsersWithBooks godoc
//
// @Summary Get all users with borrowed books
// @Description  Retrieves a list of users along with the books they've borrowed
// @Security      BearerAuth
// @Tags         User
// @Param Authorization header string true "JWT Token"
// @Accept       json
// @Produce      json
// @Success      201  {array} model.User
// @Failure      500  {string} error "Internal Server error"
// @Router       /users [get]
func (h *Handler) GetUsersWithBooks(c echo.Context) error {
	users, err := h.Manager.UserService.GetUsers(c.Request().Context())
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, "Something went wrong!")
	}
	return c.JSON(http.StatusCreated, users)
}

// GetUsersLastMonth godoc
//
// @Summary Get all users who borrowed books last month
// @Description  Retrieves a list of users who registered in the last month
// @Security      BearerAuth
// @Tags         User
// @Param Authorization header string true "JWT Token"
// @Accept       json
// @Produce      json
// @Success      200  {array} model.
// @Failure      500  {string} error "Something went wrong!"
// @Router       /records/month [get]
func (h *Handler) GetUsersLastMonth(c echo.Context) error {
	users, err := h.Manager.UserService.GetUsersLastMonth(c.Request().Context())
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, "Something went wrong!")
	}
	return c.JSON(http.StatusOK, users)
}

// GetBorrowedBooks godoc
//
// @Summary Get all borrowed books with Revenues
// @Description  Retrieves a list of borrowed books
// @Security      BearerAuth
// @Tags         Record
// @Param Authorization header string true "JWT Token"
// @Accept       json
// @Produce      json
// @Success      200  {array} model.Book
// @Failure      500  {string} error "Something went wrong!"
// @Router       /books/borrowed [get]
func (h *Handler) GetBorrowedBooks(c echo.Context) error {
	Books, err := h.Manager.RecordService.GetBorrowedBooks(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong!")
	}
	return c.JSON(http.StatusOK, Books)
}

// GetRecordByID godoc
//
// @Summary Get borrow record with price
// @Description  Retrieves the record of book borrowing with price by ID
// @Security      BearerAuth
// @Tags         Record
// @Param rid  path string true "Record ID"
// @Param Authorization header string true "JWT Token"
// @Accept       json
// @Produce      json
// @Success      200  {array} model.RecordWithTransaction
// @Failure      400  {string} error "Bad Request"
// @Failure      500  {string} error "Internal Server Error"
// @Router       /record/{rid} [get]
func (h *Handler) GetRecordByID(c echo.Context) error {
	str := c.Param("rid")
	rid, err := strconv.Atoi(str)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}
	record, err := h.Manager.RecordService.Get(c.Request().Context(), rid, h.TransactionURL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "something went wrong!")
	}
	return c.JSON(http.StatusOK, record)
}
