package main

import (
	"net/http"
	"udemy_slack_app/controller"
	"udemy_slack_app/infra"
	"udemy_slack_app/repository"
	"udemy_slack_app/transaction"
	"udemy_slack_app/usecase"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomerValidator struct {
	validator *validator.Validate
}

func (cv *CustomerValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	e.Validator = &CustomerValidator{validator: validator.New()}

	db := infra.Connect()
	transaction := transaction.NewTransaction(db)

	ur := repository.NewUserRepository(db)
	cr := repository.NewChannelRepository(db)
	mr := repository.NewMessageRepository(db)

	uu := usecase.NewUserUsecase(ur)
	cu := usecase.NewChannelUsecase(ur, cr, mr, transaction)
	mu := usecase.NewMessageUsecase(mr)

	uc := controller.NewUserController(uu)
	cc := controller.NewChannelController(cu)
	mc := controller.NewMessageController(mu)

	e.POST("/user", uc.Create)
	e.GET("/user/:user_id", uc.GetByID)
	e.PUT("/user/:user_id", uc.Update)
	e.DELETE("/user/:user_id", uc.Delete)

	e.POST("/channel", cc.Create)
	e.GET("/channel", cc.GetAll)
	e.PUT("/channel/:channel_id", cc.Update)
	e.DELETE("/channel/:channel_id", cc.Delete)

	e.POST("/message", mc.Create)
	e.GET("/message/:channel_id", mc.GetAll)
	e.PUT("/message/:message_id", mc.Update)

	e.Start(":8080")
}
