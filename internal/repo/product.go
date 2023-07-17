package repo

import (
	"github.com/Safwanseban/voixme-project/internal/types"
	"gorm.io/gorm"
)

type RepoDb struct {
	db *gorm.DB
}

// Create implements ProductRepo
func (repo *RepoDb) Create(product *types.Product) (*uint, error) {
	// db.Create(product)
	result := repo.db.Create(product)
	if result.Error != nil {

		return nil, result.Error
	}
	return &product.ID, nil
}

// FindOne implements ProductRepo
func (*RepoDb) FindOne(*types.Product) (*types.Product, error) {
	panic("unimplemented")
}

// FindUsingCountry implements ProductRepo
func (repo *RepoDb) FindUsingCountry(product *types.Product) ([]types.Product, error) {
	var products []types.Product
	result := repo.db.Where("specific_country", product.SpecificCountry).Find(&products)
	if result.Error != nil {

		return nil, result.Error
	}
	return products, nil
}

func NewRepo(db *gorm.DB) ProductRepo {
	return &RepoDb{
		db: db,
	}
}

type ProductRepo interface {
	Create(*types.Product) (*uint, error)
	FindOne(*types.Product) (*types.Product, error)
	FindUsingCountry(*types.Product) ([]types.Product, error)
}
