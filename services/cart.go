package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saipulmuiz/mnc-test-tahap2/helpers"
	"github.com/saipulmuiz/mnc-test-tahap2/models"
	"github.com/saipulmuiz/mnc-test-tahap2/params"
	"github.com/saipulmuiz/mnc-test-tahap2/repositories"
)

type CartService struct {
	cartRepo    repositories.CartRepo
	productRepo repositories.ProductRepo
}

func NewCartService(
	cartRepo repositories.CartRepo,
	productRepo repositories.ProductRepo,
) *CartService {
	return &CartService{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (u *CartService) GetCarts(page, size, userID int) *params.Response {
	var (
		carts *[]models.Cart
		err   error
		count int64
	)

	carts, count, err = u.cartRepo.GetCarts(page, size, userID)
	if err != nil {
		return helpers.HandleErrorService(http.StatusBadRequest, err.Error())
	}

	var cartData []params.GetCartRes
	for _, cart := range *carts {
		cartData = append(cartData, params.GetCartRes{
			CartID:      cart.CartID,
			ProductID:   cart.ProductID,
			ProductName: cart.ProductName,
			Price:       cart.Price,
			Qty:         cart.Qty,
			Total:       cart.Total,
			CreatedAt:   helpers.ParseDateTime(helpers.DATE_FORMAT_YYYY_MM_DD_TIME, cart.CreatedAt),
			UpdatedAt:   helpers.ParseDateTime(helpers.DATE_FORMAT_YYYY_MM_DD_TIME, cart.UpdatedAt),
		})
	}

	pagination := helpers.CalculatePagination(count, page, size, len(*carts))

	result := params.ResponseWithPagination{
		Pagination: pagination,
		Message:    "Success to get cart",
		Data:       cartData,
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: result,
	}
}

func (u *CartService) AddToCart(userID int, req params.AddProductToCart) *params.Response {
	product, err := u.productRepo.FindById(req.ProductID)
	if err != nil {
		return helpers.HandleErrorService(http.StatusNotFound, "Product not found")
	}

	cart := &models.Cart{
		UserID:      userID,
		ProductID:   req.ProductID,
		ProductName: product.ProductName,
		Price:       product.ProductPrice,
		Qty:         req.Qty,
		Total:       product.ProductPrice * float64(req.Qty),
	}

	cart, err = u.cartRepo.AddToCart(cart)
	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	result := params.ResponseSuccess{
		Message: "Add product to cart successfully",
		Data:    cart,
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Payload: result,
	}
}

func (u *CartService) UpdateCart(cartId int, req params.UpdatedCartReq) *params.Response {
	cartDb, _ := u.cartRepo.CheckCartByID(cartId, &models.Cart{})
	if cartDb.CartID == 0 {
		return helpers.HandleErrorService(http.StatusNotFound, "Item cart not found")
	}

	cart := models.Cart{
		Price:     cartDb.Price,
		Qty:       req.Qty,
		Total:     cartDb.Price * float64(req.Qty),
		UpdatedAt: time.Now(),
	}

	cartUpdated, err := u.cartRepo.UpdateCart(cartDb.CartID, &cart)
	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	result := params.ResponseSuccess{
		Message: "Cart successfully updated",
		Data:    cartUpdated,
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: result,
	}
}

func (u *CartService) DeleteItemCart(cartId int) *params.Response {
	checkData, _ := u.cartRepo.CheckCartByID(cartId, &models.Cart{})
	if checkData.CartID == 0 {
		return helpers.HandleErrorService(http.StatusNotFound, "Item cart not found")
	}

	err := u.cartRepo.DeleteCart(cartId)
	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	return &params.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"message": "Success delete item cart",
		},
	}
}
