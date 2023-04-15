package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"login/config"
	"login/model"
	"login/service"
	"login/storage"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

//E2E Test where the user registers, logins,
//and accesses protected endpoint with jwt token
func TestHandler_RegisterAndLogin(t *testing.T) {
	//Initializing Server
	s, err := initialization()
	assert.NoError(t, err)

	//Initializing Create Request
	createReqBody := `{"Name":"Test_User","Email":"test_12@email.com","Password":"123"}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(createReqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recCreate := httptest.NewRecorder()
	c := s.App.NewContext(req, recCreate)
	c.SetPath("/api/v1/user")

	//Serving Create request
	assert.NoError(t, s.handler.CreateUser(c))
	assert.Equal(t, http.StatusCreated, recCreate.Code)
	respBody := model.UserCreateResp{}
	assert.NoError(t, json.Unmarshal(recCreate.Body.Bytes(), &respBody))
	uid := respBody.UID
	assert.NotEmpty(t, uid)
	assert.True(t, compareDates(respBody.CreatedAt, time.Now()))

	//Initializing Login Request
	loginReqBody := `{"Email":"test_12@email.com","Password":"123"}`
	loginReq := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(loginReqBody))
	loginReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recLogin := httptest.NewRecorder()
	cLogin := s.App.NewContext(loginReq, recLogin)
	cLogin.SetPath("/api/v1/login")

	//Serving Login Request
	assert.NoError(t, s.handler.Validate(cLogin))
	assert.Equal(t, http.StatusOK, recLogin.Code)
	//Receiving jwt-token
	token := recLogin.Header().Get("Authorization")
	assert.NotEmpty(t, token)

	//Initializing and Serving Get Request on protected endpoint
	usersGetReq := httptest.NewRequest(http.MethodGet, "/", nil)
	loginReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	loginReq.Header.Set("Authorization", token)
	recUsers := httptest.NewRecorder()
	cGet := s.App.NewContext(usersGetReq, recUsers)
	cGet.SetPath("/api/v1/users")
	assert.NoError(t, s.handler.GetUsersWithBooks(cGet))
	assert.Equal(t, http.StatusOK, recLogin.Code)

	//Cleaning db
	_ = s.handler.Manager.UserService.Delete(context.Background(), model.UserDeleteReq{UID: uid, Hard: true})

}

//E2E test Where User Adds a book
func TestHandler_AddBook(t *testing.T) {
	//Initializing Server
	s, err := initialization()
	assert.NoError(t, err)

	//Getting JWT Token by registring test user
	token, uid, err := getToken(s)
	assert.NoError(t, err)

	//Creating Requests and Response objects
	reqBody := `{"Title": "test_book", "Author": "test_author"}`
	AddBookReq := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(reqBody))
	AddBookReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	AddBookReq.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()
	c := s.App.NewContext(AddBookReq, rec)
	c.SetPath("/api/v1/book")

	//Sending Request
	assert.NoError(t, s.handler.CreateBook(c))
	assert.Equal(t, http.StatusCreated, rec.Code)
	respBody := model.BookCreateResp{}
	err = json.Unmarshal(rec.Body.Bytes(), &respBody)

	//Checking if the book was added in db
	b, err := s.handler.Manager.BookService.GetByID(context.Background(), model.BookGetByIDReq{BID: respBody.BID})
	assert.NoError(t, err)
	assert.Equal(t, b.Title, "test_book")
	assert.Equal(t, b.Author, "test_author")

	//Cleaning db
	_ = s.handler.Manager.UserService.Delete(context.Background(), model.UserDeleteReq{UID: uid, Hard: true})
	_, _ = s.handler.Manager.BookService.DeleteByID(context.Background(), model.BookGetByIDReq{BID: b.ID})
}
func compareDates(resp time.Time, now time.Time) bool {
	year, month, day := resp.Date()
	y1, m1, d1 := now.Date()
	return year == y1 && month == m1 && day == d1
}
func initialization() (*Server, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	st, err := storage.NewStorage(cfg)
	if err != nil {
		return nil, err
	}
	m := service.NewManager(st, cfg)
	h := NewHandler(m, cfg)
	s := NewServer(cfg, h)
	s.App = echo.New()
	s.SetupRoutes()
	go func() {
		if err := s.App.Start(s.cfg.Port); err != http.ErrServerClosed {
			fmt.Printf("%v", err)
		}
	}()
	return s, nil
}
func getToken(s *Server) (string, uint, error) {
	createReqBody := `{"Name":"Test_User","Email":"test_23@email.com","Password":"123"}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(createReqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recCreate := httptest.NewRecorder()
	c := s.App.NewContext(req, recCreate)
	c.SetPath("/api/v1/user")
	err := s.handler.CreateUser(c)
	if err != nil {
		return "", 0, err
	}
	respBody := model.UserCreateResp{}
	err = json.Unmarshal(recCreate.Body.Bytes(), &respBody)
	if err != nil {
		return "", 0, err
	}
	loginReqBody := `{"Email":"test_23@email.com","Password":"123"}`
	loginReq := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(loginReqBody))
	loginReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recLogin := httptest.NewRecorder()
	cLogin := s.App.NewContext(loginReq, recLogin)
	cLogin.SetPath("/api/v1/login")
	err = s.handler.Validate(cLogin)
	if err != nil {
		return "", 0, err
	}
	token := recLogin.Header().Get("Authorization")
	return token, respBody.UID, nil

}
