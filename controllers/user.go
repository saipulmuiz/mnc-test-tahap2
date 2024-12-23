package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/saipulmuiz/mnc-test-tahap2/helpers"
	"github.com/saipulmuiz/mnc-test-tahap2/params"
	"github.com/saipulmuiz/mnc-test-tahap2/services"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		userService: *service,
	}
}

func (u *UserController) RegisterUser(c *gin.Context) {
	var req params.RegisterUser
	validate := validator.New()

	err := c.ShouldBind(&req)
	if err != nil {
		helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
		return
	}

	err = validate.Struct(req)
	if err != nil {
		validationMessage := helpers.BuildAndGetValidationMessage(err)

		helpers.HandleErrorController(c, http.StatusBadRequest, validationMessage)
		return
	}

	result := u.userService.RegisterUser(req)

	c.JSON(result.Status, result.Payload)
}

func (u *UserController) Login(c *gin.Context) {
	var req params.UserLogin

	err := c.ShouldBind(&req)

	if err != nil {
		helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
		return
	}

	result := u.userService.Login(req)

	c.JSON(result.Status, result.Payload)
}

func (u *UserController) UpdateProfile(c *gin.Context) {
	var req params.UpdateProfile

	err := c.ShouldBind(&req)

	if err != nil {
		helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
		return
	}

	userId := c.GetString("user_id")
	result := u.userService.UpdateProfile(userId, req)
	c.JSON(result.Status, result.Payload)
}

func (u *UserController) RefreshToken(c *gin.Context) {
	var req params.RefreshToken

	err := c.ShouldBind(&req)

	if err != nil {
		helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
		return
	}

	result := u.userService.RefreshToken(req)

	c.JSON(result.Status, result.Payload)
}
