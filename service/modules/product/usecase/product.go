// Di usecase/product.go
package usecase

import (
    "github.com/gin-gonic/gin"
    "github.com/ricnah/workit-be/service/modules/product/repository"
    "github.com/ricnah/workit-be/types/models"
)

type ProductUsecase interface {
    CreateProduct(ctx *gin.Context, product models.Product) (models.Product, error)
    GetProducts(ctx *gin.Context) ([]models.Product, error)
}

type productUsecase struct {
    productRepo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
    return &productUsecase{productRepo: repo}
}

func (u *productUsecase) CreateProduct(ctx *gin.Context, product models.Product) (models.Product, error) {
    err := u.productRepo.CreateProduct(&product)
    if err != nil {
        return models.Product{}, err
    }
    return product, nil
}

func (u *productUsecase) GetProducts(ctx *gin.Context) ([]models.Product, error) {
    return u.productRepo.GetProducts()
}
