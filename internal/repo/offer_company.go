package repo

import (
	"github.com/Safwanseban/voixme-project/internal/types"
	"gorm.io/gorm"
)

type RepoDb struct {
	db *gorm.DB
}

// FindAll implements ProductRepo
func (repo *RepoDb) FindAll() []types.OfferCompany {
	var company []types.OfferCompany
	result := repo.db.Find(&company)

	if result.Error != nil {
		return nil
	}
	return company
}

// Create implements ProductRepo
func (repo *RepoDb) Create(company *types.OfferCompany) (*uint, error) {
	result := repo.db.Create(company)
	if result.Error != nil {

		return nil, result.Error
	}
	return &company.OfferID, nil
}

// FindOne implements ProductRepo
func (*RepoDb) FindOne(*types.OfferCompany) (*types.OfferCompany, error) {
	panic("unimplemented")
}

// FindUsingCountry implements ProductRepo
func (repo *RepoDb) FindUsingCountry(company *types.OfferCompany) ([]types.OfferCompany, error) {
	var offerCompany []types.OfferCompany
	result := repo.db.Where("country", company.Country).Find(&offerCompany)
	if result.Error != nil {

		return nil, result.Error
	}
	if result.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	}
	return offerCompany, nil
}

func NewRepo(db *gorm.DB) ProductRepo {
	return &RepoDb{
		db: db,
	}
}

type ProductRepo interface {
	Create(*types.OfferCompany) (*uint, error)
	FindOne(*types.OfferCompany) (*types.OfferCompany, error)
	FindAll() []types.OfferCompany
	FindUsingCountry(*types.OfferCompany) ([]types.OfferCompany, error)
}
