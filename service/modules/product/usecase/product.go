package usecase

import (
    "github.com/gin-gonic/gin"
    "github.com/ricnah/workit-be/service/extensions/terror"
    "github.com/ricnah/workit-be/service/modules/product/repository"
    "github.com/ricnah/workit-be/types/models"
)

type productUsecase struct {
    productRepo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
    return &productUsecase{productRepo: repo}
}

func (u *productUsecase) CreateProduct(ctx *gin.Context, product models.Product) (models.Product, terror.ErrInterface) {
    err := u.productRepo.CreateProduct(&product)
    if err != nil {
        return models.Product{}, err 
    }
    return product, nil
}

func (u *productUsecase) GetProducts(ctx *gin.Context) ([]models.Product, terror.ErrInterface) {
    products, err := u.productRepo.GetProducts()
    if err != nil {
        return nil, err 
    }
    return products, nil
}

func (u *productUsecase) GetProductByID(ctx *gin.Context, id int64) (models.Product, terror.ErrInterface) {
    product, err := u.productRepo.GetProductByID(id)
    if err != nil {
        return models.Product{}, err 
    }
    return product, nil
}

func (u *productUsecase) UpdateProduct(ctx *gin.Context, product models.Product) (models.Product, terror.ErrInterface) {
    err := u.productRepo.UpdateProduct(&product)
    if err != nil {
        return models.Product{}, err 
    }
    return product, nil
}

func (u *productUsecase) DeleteProduct(ctx *gin.Context, id int64) (terror.ErrInterface) {
    err := u.productRepo.DeleteProduct(id)
    if err != nil {
        return err 
    }
    return nil
}
