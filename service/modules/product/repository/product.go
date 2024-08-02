package repository

import (
	"github.com/ricnah/workit-be/types/models"
	"gorm.io/gorm"
)

// ProductRepository interface untuk operasi produk
type ProductRepository interface {
    CreateProduct(product *models.Product) error
    GetProducts() ([]models.Product, error)
}

// productRepository implementasi dari ProductRepository
type productRepository struct {
	db *gorm.DB
}

// NewProductRepository membuat instance baru dari productRepository
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

// CreateProduct menambahkan produk baru ke database
func (r *productRepository) CreateProduct(product *models.Product) error {
	return r.db.Create(product).Error
}

// GetProducts mengambil semua produk dari database
func (r *productRepository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}
