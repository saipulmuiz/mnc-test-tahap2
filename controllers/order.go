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

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(service *services.OrderService) *OrderController {
	return &OrderController{
		orderService: *service,
	}
}

func (u *OrderController) CheckoutOrder(c *gin.Context) {
	var req params.CheckoutOrderReq
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

	result := u.orderService.CheckoutOrder(userID, req)

	c.JSON(result.Status, result.Payload)
}

func (u *OrderController) GetOrders(c *gin.Context) {
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

	result := u.orderService.GetOrders(page, limit, userID)

	c.JSON(result.Status, result.Payload)
}

func (u *OrderController) GetOrderById(c *gin.Context) {
	orderId, _ := strconv.Atoi(c.Param("orderId"))
	result := u.orderService.GetOrderById(orderId)
	c.JSON(result.Status, result.Payload)
}
