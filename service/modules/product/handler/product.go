package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/ricnah/workit-be/service/modules/product/usecase"
    "github.com/ricnah/workit-be/types/models"
    "net/http"
)

type ProductHandler struct {
    usecase usecase.ProductUsecase // Sesuaikan dengan interface yang benar
}

func NewProductHandler(usecase usecase.ProductUsecase) *ProductHandler {
    return &ProductHandler{usecase: usecase}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.BindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    createdProduct, err := h.usecase.CreateProduct(c, product)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, createdProduct)
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
    products, err := h.usecase.GetProducts(c)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, products)
}
