package repositories

import (
	"github.com/saipulmuiz/mnc-test-tahap2/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CartRepo interface {
	GetCarts(page, size, userID int) (*[]models.Cart, int64, error)
	GetCartByCartIds(userID int, cartIds []int) ([]models.Cart, error)
	CheckCartByID(cartId int, cart *models.Cart) (*models.Cart, error)
	AddToCart(cart *models.Cart) (*models.Cart, error)
	ClearCart(tx *gorm.DB, userID int, cartIds []int) error
	UpdateCart(cartId int, cartUpdate *models.Cart) (*models.Cart, error)
	DeleteCart(cartId int) error
}

type cartRepo struct {
	db         *gorm.DB
	globalRepo GlobalRepo
}

func NewCartRepo(db *gorm.DB, globalRepo GlobalRepo) CartRepo {
	return &cartRepo{db, globalRepo}
}

func (r *cartRepo) GetCarts(page, size, userID int) (*[]models.Cart, int64, error) {
	var (
		carts []models.Cart
		count int64
	)
	err := r.db.
		Order("created_at DESC").
		Preload(clause.Associations).
		Scopes(r.globalRepo.Paginate(page, size)).
		Where("user_id = ?", userID).
		Find(&carts).Error

	if err != nil {
		return nil, count, err
	}

	err = r.db.
		Model(&carts).
		Where("user_id = ?", userID).
		Count(&count).Error

	return &carts, count, err
}

func (r *cartRepo) GetCartByCartIds(userID int, cartIds []int) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Where("user_id = ?", userID).Where("cart_id IN (?)", cartIds).Find(&carts).Error
	return carts, err
}

func (u *cartRepo) CheckCartByID(cartId int, cart *models.Cart) (*models.Cart, error) {
	return cart, u.db.Where("cart_id", cartId).Take(&cart).Error
}

func (r *cartRepo) AddToCart(cart *models.Cart) (*models.Cart, error) {
	existingCart := models.Cart{}
	err := r.db.Where("user_id = ? AND product_id = ?", cart.UserID, cart.ProductID).First(&existingCart).Error
	if err == nil {
		existingCart.Qty += cart.Qty
		existingCart.Total = existingCart.Price * float64(existingCart.Qty)
		return &existingCart, r.db.Save(&existingCart).Error
	}
	return cart, r.db.Create(cart).Error
}

func (r *cartRepo) ClearCart(tx *gorm.DB, userID int, cartIds []int) error {
	query := r.db
	if tx != nil {
		query = tx
	}

	return query.Where("user_id = ?", userID).Where("cart_id IN (?)", cartIds).Delete(&models.Cart{}).Error
}

func (u *cartRepo) UpdateCart(cartId int, cartUpdate *models.Cart) (*models.Cart, error) {
	var cart models.Cart
	result := u.db.Model(&cart).Clauses(clause.Returning{}).Where("cart_id", cartId).Updates(cartUpdate)
	return &cart, result.Error
}

func (u *cartRepo) DeleteCart(cartId int) error {
	var cart models.Cart
	result := u.db.Model(&cart).Where("cart_id", cartId).Delete(cartId)
	return result.Error
}
