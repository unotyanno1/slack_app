package controller

import (
	"net/http"
	"udemy_slack_app/usecase"

	"github.com/labstack/echo/v4"
)

type ChannelController interface {
	Create(ctx echo.Context) error
	GetAll(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type channelController struct {
	cu usecase.ChannelUsecase
}

func NewChannelController(cu usecase.ChannelUsecase) ChannelController {
	return &channelController{cu}
}

func (cc *channelController) Create(ctx echo.Context) error {
	var req ChannelRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	cm := ChannelToModel(req)
	cc.cu.Create(ctx.Request().Context(), cm)

	return nil
}

func (cc *channelController) GetAll(ctx echo.Context) error {
	channels, err := cc.cu.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, channels)
}

func (cc *channelController) Update(ctx echo.Context) error {
	channelID := ctx.Param("channel_id")

	var req ChannelRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	cm := ChannelToModel(req)
	cc.cu.Update(ctx.Request().Context(), cm, channelID)

	return nil
}

func (cc *channelController) Delete(ctx echo.Context) error {
	channelID := ctx.Param("channel_id")
	cc.cu.Delete(ctx.Request().Context(), channelID)

	return nil
}
