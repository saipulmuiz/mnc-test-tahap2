package repositories

import (
	"github.com/saipulmuiz/mnc-test-tahap2/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderRepo interface {
	GetOrders(page, size, userID int) (*[]models.Order, int64, error)
	GetOrderByID(orderId int) (*models.Order, error)
	GetOrderDetailByOrderID(orderId int) (*[]models.OrderDetail, error)
	CreateOrder(tx *gorm.DB, order *models.Order, orderDetails []models.OrderDetail) (*models.Order, error)
}

type orderRepo struct {
	db         *gorm.DB
	globalRepo GlobalRepo
}

func NewOrderRepo(db *gorm.DB, globalRepo GlobalRepo) OrderRepo {
	return &orderRepo{db, globalRepo}
}

func (r *orderRepo) GetOrders(page, size, userID int) (*[]models.Order, int64, error) {
	var (
		orders []models.Order
		count  int64
	)
	err := r.db.
		Order("created_at DESC").
		Preload(clause.Associations).
		Scopes(r.globalRepo.Paginate(page, size)).
		Where("user_id = ?", userID).
		Find(&orders).Error

	if err != nil {
		return nil, count, err
	}

	err = r.db.
		Model(&orders).
		Where("user_id = ?", userID).
		Count(&count).Error

	return &orders, count, err
}

func (u *orderRepo) GetOrderByID(orderId int) (*models.Order, error) {
	var order models.Order
	err := u.db.Where("id = ?", orderId).First(&order).Error
	return &order, err
}

func (u *orderRepo) GetOrderDetailByOrderID(orderId int) (*[]models.OrderDetail, error) {
	var orderDetail []models.OrderDetail
	err := u.db.Where("order_id = ?", orderId).Find(&orderDetail).Error
	return &orderDetail, err
}

func (u *orderRepo) CreateOrder(tx *gorm.DB, order *models.Order, orderDetails []models.OrderDetail) (*models.Order, error) {
	createOrderAndDetails := func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}

		for _, detail := range orderDetails {
			detail.OrderID = order.ID
			if err := tx.Create(&detail).Error; err != nil {
				return err
			}
		}
		return nil
	}

	if tx != nil {
		if err := createOrderAndDetails(tx); err != nil {
			return nil, err
		}
		return order, nil
	}

	err := u.db.Transaction(func(tx *gorm.DB) error {
		return createOrderAndDetails(tx)
	})

	return order, err
}
