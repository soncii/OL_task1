package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"transactions/config"
	"transactions/model"
	"transactions/service"
)

type Handler struct {
	TransactionService service.ITransactionService
}

func NewHandler(
	transactionService *service.TransactionService, cfg *config.Config) *Handler {
	return &Handler{TransactionService: transactionService}
}

// CreateTransaction godoc
// @Summary Register transaction
// @Description  Saves the money transaction for borrowing book
// @Tags         Transaction
// @Param input  body model.TransactionCreateReq true "Details of borrowing contains bookID, userID, and the price"
// @Accept       json
// @Produce      json
// @Success      200  {object} model.TransactionCreateResp
// @Router       /transaction [post]
func (h *Handler) CreateTransaction(c echo.Context) error {
	req := model.TransactionCreateReq{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}
	res, err := h.TransactionService.Create(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}
	return c.JSON(http.StatusOK, res)
}

// GetTransaction godoc
//
// @Summary Get transaction by id
// @Description  Returns transaction information by id
// @Tags         Transaction
// @Param tid  path string true "Transaction ID"
// @Accept       json
// @Produce      json
// @Success      200  {object} model.Transaction
// @Failure 	 400 {string} string "Bad Request"
// @Failure 	 500 {string} string "Internal Server Error"
// @Router       /transaction [get]
func (h *Handler) GetTransaction(c echo.Context) error {
	str := c.Param("tid")
	tid, err := strconv.Atoi(str)
	if err != nil || tid < 0 {
		return c.JSON(http.StatusBadRequest, "")
	}
	tr, err := h.TransactionService.Get(c.Request().Context(), uint(tid))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}
	return c.JSON(http.StatusOK, tr)
}
