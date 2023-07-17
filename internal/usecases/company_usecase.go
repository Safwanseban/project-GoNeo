package usecases

import (
	"time"

	"github.com/Safwanseban/voixme-project/internal/repo"
	"github.com/Safwanseban/voixme-project/internal/types"
)

type ProductUsecase struct {
	Repo repo.ProductRepo
}

var CompanyCache = make(map[types.OfferCompany]string, 0)

// CreateProduct implements UsecasesProduct
func (u *ProductUsecase) CreateProduct(company *types.OfferCompany) (*uint, error) {
	id, err := u.Repo.Create(company)
	if err != nil {
		return nil, err
	}
	return id, nil
}

// ShowProducts implements UsecasesProduct
func (u *ProductUsecase) ShowOfferCompany(company *types.OfferCompany) ([]types.OfferCompany, error) {

	offerCompany := make([]types.OfferCompany, 0)
	for i, v := range CompanyCache {
		if v == string(company.Country) {
			offerCompany = append(offerCompany, i)
		}

	}
	if len(offerCompany) > 0 {
		return offerCompany, nil
	}
	companies, err := u.Repo.FindUsingCountry(company)
	if err != nil {
		return nil, err
	}
	return companies, nil
}
func (u *ProductUsecase) FetchAndAppend() {
	ticker := time.NewTicker(10 * time.Second)
	done := make(chan bool)

	go func() {

		for {
			select {
			case <-done:
				return
			case <-ticker.C:

				company := u.Repo.FindAll()

				for _, company := range company {
					CompanyCache[company] = company.Country
				}

			}

		}

	}()

}
func NewCompanyUsecase(repo repo.ProductRepo) UsecasesCompany {

	return &ProductUsecase{
		Repo: repo,
	}
}

type UsecasesCompany interface {
	CreateProduct(*types.OfferCompany) (*uint, error)
	ShowOfferCompany(*types.OfferCompany) ([]types.OfferCompany, error)
	FetchAndAppend()
}
