package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/saipulmuiz/mnc-test-tahap2/helpers"
	"github.com/saipulmuiz/mnc-test-tahap2/params"
	"github.com/saipulmuiz/mnc-test-tahap2/services"
)

type CartController struct {
	cartService services.CartService
}

func NewCartController(service *services.CartService) *CartController {
	return &CartController{
		cartService: *service,
	}
}

func (u *CartController) GetCarts(c *gin.Context) {
	var (
		page, limit int
		err         error
	)

	pageStr := c.Query("page")
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		page = 1
	}

	limitStr := c.Query("limit")
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		limit = 10
	}

	userID, _ := strconv.Atoi(c.GetString("user_id"))

	result := u.cartService.GetCarts(page, limit, userID)

	c.JSON(result.Status, result.Payload)
}

func (u *CartController) AddToCart(c *gin.Context) {
	var req params.AddProductToCart
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

	userID, _ := strconv.Atoi(c.GetString("user_id"))

	result := u.cartService.AddToCart(userID, req)

	c.JSON(result.Status, result.Payload)
}

func (u *CartController) UpdateCart(c *gin.Context) {
	var req params.UpdatedCartReq

	err := c.ShouldBind(&req)

	if err != nil {
		helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
		return
	}

	cartId, _ := strconv.Atoi(c.Param("cartId"))
	result := u.cartService.UpdateCart(cartId, req)
	c.JSON(result.Status, result.Payload)
}

func (u *CartController) DeleteCart(c *gin.Context) {
	cartId, _ := strconv.Atoi(c.Param("cartId"))
	result := u.cartService.DeleteItemCart(cartId)
	c.JSON(result.Status, result.Payload)
}
