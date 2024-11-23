package repositories

import (
	"github.com/saipulmuiz/mnc-test-tahap2/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepo interface {
	GetProducts(page, size int) (*[]models.Product, int64, error)
	FindById(productId int) (*models.Product, error)
	CreateProduct(product *models.Product) (*models.Product, error)
	CheckProductByID(productId int, product *models.Product) (*models.Product, error)
	UpdateProduct(tx *gorm.DB, productId int, productUpdate *models.Product) (*models.Product, error)
	DeleteProduct(productId int) error
}

type productRepo struct {
	db         *gorm.DB
	globalRepo GlobalRepo
}

func NewProductRepo(db *gorm.DB, globalRepo GlobalRepo) ProductRepo {
	return &productRepo{db, globalRepo}
}

func (u *productRepo) GetProducts(page, size int) (*[]models.Product, int64, error) {
	var (
		products []models.Product
		count    int64
	)
	err := u.db.
		Order("created_at DESC").
		Preload(clause.Associations).
		Scopes(u.globalRepo.Paginate(page, size)).
		Find(&products).Error

	if err != nil {
		return nil, count, err
	}

	err = u.db.
		Model(&products).
		Count(&count).Error

	return &products, count, err
}

func (u *productRepo) FindById(productId int) (*models.Product, error) {
	var product models.Product
	err := u.db.Where("product_id = ?", productId).First(&product).Error
	return &product, err
}

func (u *productRepo) CreateProduct(product *models.Product) (*models.Product, error) {
	return product, u.db.Create(&product).Error
}

func (u *productRepo) CheckProductByID(productId int, product *models.Product) (*models.Product, error) {
	return product, u.db.Where("product_id", productId).Take(&product).Error
}

func (u *productRepo) UpdateProduct(tx *gorm.DB, productId int, productUpdate *models.Product) (*models.Product, error) {
	var (
		product models.Product
		result  *gorm.DB
	)
	if tx != nil {
		result = tx.Model(&product).Clauses(clause.Returning{}).Where("product_id", productId).Updates(productUpdate)
	} else {
		result = u.db.Model(&product).Clauses(clause.Returning{}).Where("product_id", productId).Updates(productUpdate)
	}

	return &product, result.Error
}

func (u *productRepo) DeleteProduct(productId int) error {
	var product models.Product
	result := u.db.Model(&product).Where("product_id", productId).Delete(productId)
	return result.Error
}
