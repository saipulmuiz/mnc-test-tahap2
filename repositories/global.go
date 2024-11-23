package repositories

import (
	"gorm.io/gorm"
)

type GlobalRepo interface {
	Paginate(page, size int) func(db *gorm.DB) *gorm.DB
}

type globalRepo struct{}

func NewGlobalRepo() *globalRepo {
	return &globalRepo{}
}

func (g *globalRepo) Paginate(page, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}
