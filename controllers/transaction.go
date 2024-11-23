package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/saipulmuiz/mnc-test-tahap2/helpers"
	"github.com/saipulmuiz/mnc-test-tahap2/params"
	"github.com/saipulmuiz/mnc-test-tahap2/services"
)

type TransactionController struct {
	transactionService services.TransactionService
}

func NewTransactionController(service *services.TransactionService) *TransactionController {
	return &TransactionController{
		transactionService: *service,
	}
}

func (u *TransactionController) Topup(c *gin.Context) {
	var req params.Topup
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

	userID := c.GetString("user_id")

	result := u.transactionService.Topup(userID, req)

	c.JSON(result.Status, result.Payload)
}
