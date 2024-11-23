package services

import (
	"net/http"
	"time"

	"github.com/saipulmuiz/mnc-test-tahap2/helpers"
	"github.com/saipulmuiz/mnc-test-tahap2/models"
	"github.com/saipulmuiz/mnc-test-tahap2/params"
	"github.com/saipulmuiz/mnc-test-tahap2/repositories"
	"gorm.io/gorm"
)

type ProductService struct {
	productRepo repositories.ProductRepo
}

func NewProductService(repo repositories.ProductRepo) *ProductService {
	return &ProductService{
		productRepo: repo,
	}
}

func (u *ProductService) GetProducts(page, size int) *params.Response {
	var (
		products *[]models.Product
		err      error
		count    int64
	)

	products, count, err = u.productRepo.GetProducts(page, size)
	if err != nil {
		return helpers.HandleErrorService(http.StatusBadRequest, err.Error())
	}

	var productData []params.GetProductRes
	for _, product := range *products {
		productData = append(productData, params.GetProductRes{
			ProductID:    product.ProductID,
			ProductName:  product.ProductName,
			Description:  product.Description,
			ProductPrice: product.ProductPrice,
			Stock:        product.Stock,
			MinStock:     product.MinStock,
			CreatedAt:    helpers.ParseDateTime(helpers.DATE_FORMAT_YYYY_MM_DD_TIME, product.CreatedAt),
			UpdatedAt:    helpers.ParseDateTime(helpers.DATE_FORMAT_YYYY_MM_DD_TIME, product.UpdatedAt),
		})
	}

	pagination := helpers.CalculatePagination(count, page, size, len(*products))

	result := params.ResponseWithPagination{
		Pagination: pagination,
		Message:    "Success to get product",
		Data:       productData,
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: result,
	}
}

func (u *ProductService) GetProductById(productId int) *params.Response {
	product, err := u.productRepo.FindById(productId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return helpers.HandleErrorService(http.StatusNotFound, "Product not found")
		}

		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	resp := params.GetProductRes{
		ProductID:    product.ProductID,
		ProductName:  product.ProductName,
		Description:  product.Description,
		ProductPrice: product.ProductPrice,
		Stock:        product.Stock,
		MinStock:     product.MinStock,
		CreatedAt:    helpers.ParseDateTime(helpers.DATE_FORMAT_YYYY_MM_DD_TIME, product.CreatedAt),
		UpdatedAt:    helpers.ParseDateTime(helpers.DATE_FORMAT_YYYY_MM_DD_TIME, product.UpdatedAt),
	}

	result := params.ResponseSuccess{
		Message: "Success Get Product By Id",
		Data:    resp,
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: result,
	}
}

func (u *ProductService) CreateProduct(request params.CreateProductReq) *params.Response {
	product := models.Product{
		ProductName:  request.ProductName,
		Description:  request.Description,
		ProductPrice: request.ProductPrice,
		Stock:        request.Stock,
		MinStock:     request.MinStock,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	productData, err := u.productRepo.CreateProduct(&product)
	if err != nil {
		return helpers.HandleErrorService(http.StatusBadRequest, err.Error())
	}

	result := params.ResponseSuccess{
		Message: "Create product successfully",
		Data:    productData,
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Payload: result,
	}
}

func (u *ProductService) UpdateProduct(productId int, request params.UpdatedProductReq) *params.Response {
	checkData, _ := u.productRepo.CheckProductByID(productId, &models.Product{})
	if checkData.ProductID == 0 {
		return helpers.HandleErrorService(http.StatusNotFound, "Product not found")
	}

	product := models.Product{
		ProductName:  request.ProductName,
		Description:  request.Description,
		ProductPrice: request.ProductPrice,
		Stock:        request.Stock,
		MinStock:     request.MinStock,
		UpdatedAt:    time.Now(),
	}

	productUpdated, err := u.productRepo.UpdateProduct(nil, checkData.ProductID, &product)

	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	result := params.ResponseSuccess{
		Message: "Product successfully updated",
		Data:    productUpdated,
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: result,
	}
}

func (u *ProductService) DeleteProduct(productId int) *params.Response {
	checkData, _ := u.productRepo.CheckProductByID(productId, &models.Product{})
	if checkData.ProductID == 0 {
		return helpers.HandleErrorService(http.StatusNotFound, "Product not found")
	}

	err := u.productRepo.DeleteProduct(productId)
	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	return &params.Response{
		Status: http.StatusOK,
		Payload: params.ResponseSuccessMessage{
			Message: "Success delete product",
		},
	}
}
