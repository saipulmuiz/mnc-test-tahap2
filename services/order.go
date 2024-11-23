package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/saipulmuiz/mnc-test-tahap2/helpers"
	"github.com/saipulmuiz/mnc-test-tahap2/models"
	"github.com/saipulmuiz/mnc-test-tahap2/params"
	"github.com/saipulmuiz/mnc-test-tahap2/repositories"
	"gorm.io/gorm"
)

type OrderService struct {
	productRepo repositories.ProductRepo
	cartRepo    repositories.CartRepo
	orderRepo   repositories.OrderRepo
	db          *gorm.DB
}

func NewOrderService(
	productRepo repositories.ProductRepo,
	cartRepo repositories.CartRepo,
	orderRepo repositories.OrderRepo,
	db *gorm.DB,
) *OrderService {
	return &OrderService{
		productRepo: productRepo,
		cartRepo:    cartRepo,
		orderRepo:   orderRepo,
		db:          db,
	}
}

func (u *OrderService) GetOrders(page, size, userID int) *params.Response {
	var (
		orders *[]models.Order
		err    error
		count  int64
	)

	orders, count, err = u.orderRepo.GetOrders(page, size, userID)
	if err != nil {
		return helpers.HandleErrorService(http.StatusBadRequest, err.Error())
	}

	var orderData []params.GetOrderRes
	for _, order := range *orders {
		orderData = append(orderData, params.GetOrderRes{
			OrderID:    order.ID,
			Status:     order.Status,
			TotalPrice: order.TotalPrice,
			CreatedAt:  helpers.ParseDateTime(helpers.DATE_FORMAT_YYYY_MM_DD_TIME, order.CreatedAt),
			UpdatedAt:  helpers.ParseDateTime(helpers.DATE_FORMAT_YYYY_MM_DD_TIME, order.UpdatedAt),
		})
	}

	pagination := helpers.CalculatePagination(count, page, size, len(*orders))

	result := params.ResponseWithPagination{
		Pagination: pagination,
		Message:    "Success to get order",
		Data:       orderData,
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: result,
	}
}

func (u *OrderService) GetOrderById(orderId int) *params.Response {
	order, err := u.orderRepo.GetOrderByID(orderId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return helpers.HandleErrorService(http.StatusNotFound, "Order not found")
		}

		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	orderDetails, err := u.orderRepo.GetOrderDetailByOrderID(orderId)
	if err != nil {
		return helpers.HandleErrorService(http.StatusBadRequest, err.Error())
	}

	var orderDetailData []params.GetOrderDetailRes
	for _, orderDetail := range *orderDetails {
		orderDetailData = append(orderDetailData, params.GetOrderDetailRes{
			OrderDetailID: orderDetail.ID,
			ProductID:     orderDetail.ProductID,
			ProductName:   orderDetail.ProductName,
			Price:         orderDetail.Price,
			Qty:           orderDetail.Qty,
			Total:         orderDetail.Total,
			CreatedAt:     helpers.ParseDateTime(helpers.DATE_FORMAT_YYYY_MM_DD_TIME, orderDetail.CreatedAt),
			UpdatedAt:     helpers.ParseDateTime(helpers.DATE_FORMAT_YYYY_MM_DD_TIME, orderDetail.UpdatedAt),
		})
	}

	resp := params.GetOrderRes{
		OrderID:     order.ID,
		Status:      order.Status,
		TotalPrice:  order.TotalPrice,
		CreatedAt:   helpers.ParseDateTime(helpers.DATE_FORMAT_YYYY_MM_DD_TIME, order.CreatedAt),
		UpdatedAt:   helpers.ParseDateTime(helpers.DATE_FORMAT_YYYY_MM_DD_TIME, order.UpdatedAt),
		OrderDetail: orderDetailData,
	}

	result := params.ResponseSuccess{
		Message: "Success Get Order By Id",
		Data:    resp,
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: result,
	}
}

func (u *OrderService) CheckoutOrder(userID int, req params.CheckoutOrderReq) *params.Response {
	// Start transaction
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	cartItems, err := u.cartRepo.GetCartByCartIds(userID, req.CartIDs)
	if err != nil {
		tx.Rollback()
		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	if len(cartItems) == 0 {
		return helpers.HandleErrorService(http.StatusNotFound, "Cart still empty")
	}

	var (
		totalPrice   float64
		orderDetails []models.OrderDetail
	)

	for _, cart := range cartItems {
		// Check product stock
		var product *models.Product
		product, err = u.productRepo.FindById(cart.ProductID)
		if err != nil {
			tx.Rollback()
			if err == gorm.ErrRecordNotFound {
				return helpers.HandleErrorService(http.StatusNotFound, "Product not found")
			}

			return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
		}

		// Check the product stock
		if product.Stock < cart.Qty {
			tx.Rollback()
			return helpers.HandleErrorService(http.StatusBadRequest, fmt.Sprintf("Insufficient stock for product with ID = %d and name = '%s'. Please reduce the quantity or choose another product.", cart.ProductID, cart.ProductName))
		}

		// Update product stock
		productUpdate := models.Product{
			Stock:     product.Stock - cart.Qty,
			UpdatedAt: time.Now(),
		}
		_, err = u.productRepo.UpdateProduct(tx, cart.ProductID, &productUpdate)
		if err != nil {
			tx.Rollback()
			if err == gorm.ErrRecordNotFound {
				return helpers.HandleErrorService(http.StatusNotFound, "Product not found")
			}

			return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
		}

		totalPrice += cart.Total

		orderDetails = append(orderDetails, models.OrderDetail{
			ProductID:   cart.ProductID,
			ProductName: cart.ProductName,
			Price:       cart.Price,
			Qty:         cart.Qty,
			Total:       cart.Total,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
	}

	order := &models.Order{
		UserID:     userID,
		TotalPrice: totalPrice,
		Status:     "completed",
	}
	order, err = u.orderRepo.CreateOrder(tx, order, orderDetails)
	if err != nil {
		tx.Rollback()
		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	err = u.cartRepo.ClearCart(tx, userID, req.CartIDs)
	if err != nil {
		tx.Rollback()
		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	// Commit transaction
	err = tx.Commit().Error
	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, "Failed to commit transaction")
	}

	result := params.ResponseSuccess{
		Message: "Checkout order successfully",
		Data:    order,
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Payload: result,
	}
}
