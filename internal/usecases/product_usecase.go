package usecases

import (
	"fmt"
	"time"

	"github.com/Safwanseban/voixme-project/internal/repo"
	"github.com/Safwanseban/voixme-project/internal/types"
)

type ProductUsecase struct {
	Repo repo.ProductRepo
}

var ProductCache = make(map[types.Product]string, 0)

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
	fmt.Println(ProductCache)
	products := make([]types.Product, 0)
	for i, v := range ProductCache {
		if v == string(product.SpecificCountry) {
			products = append(products, i)
		}

	}
	if len(products) > 0 {
		fmt.Println("this accessed")
		return products, nil
	}
	products, err := u.Repo.FindUsingCountry(product)
	if err != nil {
		return nil, err
	}
	return products, nil
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

				products := u.Repo.FindAll()

				for _, product := range products {
					ProductCache[product] = string(product.SpecificCountry)
				}

			}

		}

	}()

}
func NewProductUsecase(repo repo.ProductRepo) UsecasesProduct {

	return &ProductUsecase{
		Repo: repo,
	}
}

type UsecasesProduct interface {
	CreateProduct(*types.Product) (*uint, error)
	ShowProducts(*types.Product) ([]types.Product, error)
	FetchAndAppend()
}
