package handler

import (
	"net/http"

	"user-service/internal/adapter/handler/request"
	"user-service/internal/adapter/handler/response"
	"user-service/internal/core/domain/entity"
	"user-service/internal/core/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// UserHandlerInterface mendefinisikan kontrak (interface) untuk semua handler yang berhubungan dengan user.
// Ini membantu kita dalam abstraksi dan mempermudah testing (mocking).
type UserHandlerInterface interface {
	SignIn(c echo.Context) error
	GetUsers(c echo.Context) error
	SignUp(c echo.Context) error
}

// userHandler adalah implementasi dari UserHandlerInterface.
// Struct ini membutuhkan userService untuk menjalankan logika bisnis.
type userHandler struct {
	userService service.UserServiceInterface
}

// SignIn adalah handler untuk menangani request masuk pada endpoint login (/signin).
func (u *userHandler) SignIn(c echo.Context) error {
	// Menyiapkan variabel untuk menampung data request dan format response.
	req := request.SignInRequest{}
	resp := response.DefaultResponse{}
	respSignIn := response.SignInResponse{}

	// c.Bind(&req) otomatis mengambil data dari body request (JSON) dan memasukkannya ke struct req.
	if err := c.Bind(&req); err != nil {
		log.Errorf("[UserHandler] SignIn: %v", err)
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	// c.Validate(req) menjalankan validasi data (misalnya: email harus format email, password tidak boleh kosong).
	if err := c.Validate(req); err != nil {
		log.Errorf("[UserHandler] SignIn: %v", err)
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	// Memanggil layer service untuk mengecek apakah email dan password benar.
	token, user, err := u.userService.SignIn(
		c.Request().Context(),
		req.Email,
		req.Password,
	)

	// Jika terjadi error di service (misal: password salah), kirim response error ke client.
	if err != nil {
		log.Errorf("[UserHandler] SignIn: %v", err)
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	// Jika berhasil, masukkan token dan data user ke dalam struct response.
	respSignIn.AccessToken = token
	if user != nil {
		respSignIn.ID = user.ID
		respSignIn.Name = user.Name
		respSignIn.Email = user.Email
		respSignIn.Role = user.RoleName
		respSignIn.Lat = user.Lat
		respSignIn.Lng = user.Lng
		respSignIn.IsVerified = user.IsVerified
	}
	
	resp.Message = "User signed in successfully"
	resp.Data = respSignIn

	// Kirim response sukses (HTTP 200 OK) dengan data JSON.
	return c.JSON(http.StatusOK, resp)
}

// GetUsers menangani request untuk mengambil semua data user.
func (u *userHandler) GetUsers(c echo.Context) error {
	resp := response.DefaultResponse{}

	users, err := u.userService.FetchAllUsers(c.Request().Context())
	if err != nil {
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Message = "Users fetched successfully"
	resp.Data = users
	return c.JSON(http.StatusOK, resp)
}

// SignUp menangani request untuk mendaftarkan user baru.
func (u *userHandler) SignUp(c echo.Context) error {
	req := request.SignUpRequest{}
	resp := response.DefaultResponse{}

	if err := c.Bind(&req); err != nil {
		log.Errorf("[UserHandler] SignUp Bind: %v", err)
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	if err := c.Validate(req); err != nil {
		log.Errorf("[UserHandler] SignUp Validate: %v", err)
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	// Mapping Request to Entity
	userEnt := &entity.UserEntity{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		RoleName: req.RoleName,
	}

	err := u.userService.SignUp(c.Request().Context(), userEnt)
	if err != nil {
		log.Errorf("[UserHandler] SignUp Service: %v", err)
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Message = "User registered successfully"
	return c.JSON(http.StatusCreated, resp)
}


// NewUserHandler adalah fungsi generator untuk membuat instance handler baru dan mendaftarkan route-nya ke Echo.
func NewUserHandler(e *echo.Echo, userService service.UserServiceInterface) UserHandlerInterface {
	userHandler := &userHandler{
		userService: userService,
	}

	// middleware.Recover() mencegah aplikasi crash total jika terjadi panic/error fatal di handler.
	e.Use(middleware.Recover())

	// Mendaftarkan endpoint
	e.POST("/signin", userHandler.SignIn)
	e.POST("/signup", userHandler.SignUp)
	e.GET("/users", userHandler.GetUsers)

	return userHandler
}
