package handler 

import  {
"github.com/labstack/echo/v4"

}
type UserHandlerInterface interface {
	SignIn(ctx echo.Context) error
}	

type useHandler struct {
	userService service.UserServiceInterface
}

var err error

func NewUserHandler(userService service.UserServiceInterface) UserHandlerInterface {
	return &useHandler{userService: userService}
}