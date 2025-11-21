package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/arke-so/mini-arke/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(ctx context.Context, product *models.Product) error {
	result := r.db.WithContext(ctx).Create(product)
	if result.Error != nil {
		return fmt.Errorf("failed to create product: %w", result.Error)
	}
	return nil
}

func (r *ProductRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	var product models.Product
	result := r.db.WithContext(ctx).First(&product, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get product: %w", result.Error)
	}
	return &product, nil
}

func (r *ProductRepository) List(ctx context.Context, limit, offset int) ([]models.Product, int, error) {
	var products []models.Product
	var total int64

	if err := r.db.WithContext(ctx).Model(&models.Product{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count products: %w", err)
	}

	result := r.db.WithContext(ctx).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&products)

	if result.Error != nil {
		return nil, 0, fmt.Errorf("failed to list products: %w", result.Error)
	}

	return products, int(total), nil
}

func (r *ProductRepository) Update(ctx context.Context, product *models.Product) error {
	result := r.db.WithContext(ctx).
		Model(&models.Product{}).
		Where("id = ?", product.ID).
		Updates(map[string]interface{}{
			"name":           product.Name,
			"description":    product.Description,
			"price":          product.Price,
			"stock_quantity": product.StockQuantity,
		})

	if result.Error != nil {
		return fmt.Errorf("failed to update product: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *ProductRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&models.Product{}, "id = ?", id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete product: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
