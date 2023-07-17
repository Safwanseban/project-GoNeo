package usecases

import (
	"time"

	"github.com/Safwanseban/voixme-project/internal/repo"
	"github.com/Safwanseban/voixme-project/internal/types"
)

type ProductUsecase struct {
	Repo repo.ProductRepo
}

type ProductCache map[types.Product]string

// CreateProduct implements UsecasesProduct
func (u *ProductUsecase) CreateProduct(product *types.Product) (*uint, error) {
	id, err := u.Repo.Create(product)
	if err != nil {
		return nil, err
	}
	return id, nil
}

// ShowProducts implements UsecasesProduct
func (u *ProductUsecase) ShowProducts(product *types.Product) ([]types.Product, error) {
	productCache := u.FetchAndAppend()
	products := make([]types.Product, 0)
	for i, v := range productCache {
		if v == string(product.SpecificCountry) {
			products = append(products, i)
		}

	}
	if len(products) > 0 {

		return products, nil
	}
	products, err := u.Repo.FindUsingCountry(product)
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (u *ProductUsecase) FetchAndAppend() ProductCache {
	ticker := time.NewTicker(10 * time.Second)
	done := make(chan bool)
	productCache := make(ProductCache, 0)
	r := func() ProductCache {

		for {
			select {
			case <-done:
				return nil
			case <-ticker.C:

				products := u.Repo.FindAll()

				for _, product := range products {
					productCache[product] = string(product.SpecificCountry)
				}
				return productCache
			}

		}

	}()
	return r
}
func NewProductUsecase(repo repo.ProductRepo) UsecasesProduct {

	return &ProductUsecase{
		Repo: repo,
	}
}

type UsecasesProduct interface {
	CreateProduct(*types.Product) (*uint, error)
	ShowProducts(*types.Product) ([]types.Product, error)
	FetchAndAppend() ProductCache
}
