package handler

import (
	"net/http"

	"user-service/internal/adapter/handler/request"
	"user-service/internal/adapter/handler/response"
	"user-service/internal/core/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type UserHandlerInterface interface {
	SignIn(c echo.Context) error
}

type userHandler struct {
	userService service.UserServiceInterface
}

func (u *userHandler) SignIn(c echo.Context) error {
	req := request.SignInRequest{}
	resp := response.DefaultResponse{}
	respSignIn := response.SignInResponse{}

	if err := c.Bind(&req); err != nil {
		log.Errorf("[UserHandler] SignIn: %v", err)
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	if err := c.Validate(req); err != nil {
		log.Errorf("[UserHandler] SignIn: %v", err)
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	token, err := u.userService.SignIn(
		c.Request().Context(),
		req.Email,
		req.Password,
	)

	if err != nil {
		log.Errorf("[UserHandler] SignIn: %v", err)
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	respSignIn.AccessToken = token
	resp.Message = "User signed in successfully"
	resp.Data = respSignIn

	return c.JSON(http.StatusOK, resp)
}

func NewUserHandler(e *echo.Echo, userService service.UserServiceInterface) UserHandlerInterface {
	userHandler := &userHandler{
		userService: userService,
	}

	e.Use(middleware.Recover())
	e.POST("/signin", userHandler.SignIn)

	return userHandler
}
