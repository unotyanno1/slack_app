package controller

import (
	"net/http"
	"udemy_slack_app/usecase"

	"github.com/labstack/echo/v4"
)

type MessageController interface {
	Create(ctx echo.Context) error
	GetAll(ctx echo.Context) error
	Update(ctx echo.Context) error
}

type messageController struct {
	mu usecase.MessageUsecase
}

func NewMessageController(mu usecase.MessageUsecase) MessageController {
	return &messageController{mu}
}

func (mc *messageController) Create(ctx echo.Context) error {
	var req MessageRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	mm := MessageToModel(req)
	mc.mu.Create(ctx.Request().Context(), mm)

	return nil
}

func (mc *messageController) GetAll(ctx echo.Context) error {
	messages, err := mc.mu.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, messages)
}

func (mc *messageController) Update(ctx echo.Context) error {
	messageID := ctx.Param("message_id")

	var req MessageRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	mm := MessageToModel(req)
	mc.mu.Update(ctx.Request().Context(), mm, messageID)

	return nil
}
