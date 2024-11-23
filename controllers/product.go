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

type ProductController struct {
	productService services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{
		productService: *service,
	}
}

func (u *ProductController) GetProducts(c *gin.Context) {
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

	result := u.productService.GetProducts(page, limit)

	c.JSON(result.Status, result.Payload)
}

func (u *ProductController) CreateProduct(c *gin.Context) {
	var req params.CreateProductReq
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

	result := u.productService.CreateProduct(req)

	c.JSON(result.Status, result.Payload)
}

func (u *ProductController) GetProductById(c *gin.Context) {
	productId, _ := strconv.Atoi(c.Param("productId"))
	result := u.productService.GetProductById(productId)
	c.JSON(result.Status, result.Payload)
}

func (u *ProductController) UpdateProduct(c *gin.Context) {
	var req params.UpdatedProductReq

	err := c.ShouldBind(&req)

	if err != nil {
		helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
		return
	}

	productId, _ := strconv.Atoi(c.Param("productId"))
	result := u.productService.UpdateProduct(productId, req)
	c.JSON(result.Status, result.Payload)
}

func (u *ProductController) DeleteProduct(c *gin.Context) {
	productId, _ := strconv.Atoi(c.Param("productId"))
	result := u.productService.DeleteProduct(productId)
	c.JSON(result.Status, result.Payload)
}
