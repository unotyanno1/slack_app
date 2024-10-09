package controller

import (
	"net/http"
	"udemy_slack_app/usecase"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	Create(ctx echo.Context) error
	GetByID(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type userController struct {
	uu usecase.UserUsecase
}

func NewUserController(uu usecase.UserUsecase) UserController {
	return &userController{uu}
}

func (uc *userController) Create(ctx echo.Context) error {
	var req UserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	um := UserToModel(req)
	uc.uu.Create(ctx.Request().Context(), um)

	return nil
}

func (uc *userController) GetByID(ctx echo.Context) error {
	userID := ctx.Param("user_id")
	user, err := uc.uu.GetByID(ctx.Request().Context(), userID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, user)
}

func (uc *userController) Update(ctx echo.Context) error {
	userID := ctx.Param("user_id")

	var req UserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	um := UserToModel(req)
	uc.uu.Update(ctx.Request().Context(), um, userID)

	return nil
}

func (uc *userController) Delete(ctx echo.Context) error {
	userID := ctx.Param("user_id")
	uc.uu.Delete(ctx.Request().Context(), userID)

	return nil
}
