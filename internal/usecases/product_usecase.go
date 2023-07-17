package usecases

import (
	"github.com/Safwanseban/voixme-project/internal/repo"
	"github.com/Safwanseban/voixme-project/internal/types"
)

type ProductUsecase struct {
	Repo repo.ProductRepo
}

type ProductCache map[uint]*types.Product

// CreateProduct implements UsecasesProduct
func (u *ProductUsecase) CreateProduct(product *types.Product) (*uint, error) {
	id, err := u.Repo.Create(product)
	if err != nil {
		return nil, err
	}
	return id, nil
}

// ShowProducts implements UsecasesProduct
func (*ProductUsecase) ShowProducts(*types.Product) ([]types.Product, error) {
	panic("unimplemented")
}

func NewProductUsecase(repo repo.ProductRepo) UsecasesProduct {

	return &ProductUsecase{
		Repo: repo,
	}
}

type UsecasesProduct interface {
	CreateProduct(*types.Product) (*uint, error)
	ShowProducts(*types.Product) ([]types.Product, error)
}
