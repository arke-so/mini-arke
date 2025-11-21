package service

import (
	"context"
	"time"

	"github.com/arke-so/mini-arke/internal/models"
	"github.com/arke-so/mini-arke/internal/repository"
	"github.com/google/uuid"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

type CreateProductInput struct {
	SKU           string
	Name          string
	Description   *string
	Price         float64
	StockQuantity int
}

type UpdateProductInput struct {
	Name          *string
	Description   *string
	Price         *float64
	StockQuantity *int
}

func (s *ProductService) Create(ctx context.Context, input CreateProductInput) (*models.Product, error) {
	product := &models.Product{
		ID:            uuid.New(),
		SKU:           input.SKU,
		Name:          input.Name,
		Description:   input.Description,
		Price:         input.Price,
		StockQuantity: input.StockQuantity,
		CreatedAt:     time.Now(),
	}

	if err := s.repo.Create(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) Get(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ProductService) List(ctx context.Context, limit, offset int) ([]models.Product, int, error) {
	return s.repo.List(ctx, limit, offset)
}

func (s *ProductService) Update(ctx context.Context, id uuid.UUID, input UpdateProductInput) (*models.Product, error) {
	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if input.Name != nil {
		product.Name = *input.Name
	}
	if input.Description != nil {
		product.Description = input.Description
	}
	if input.Price != nil {
		product.Price = *input.Price
	}
	if input.StockQuantity != nil {
		product.StockQuantity = *input.StockQuantity
	}

	if err := s.repo.Update(ctx, product); err != nil {
		return nil, err
	}

	return s.repo.GetByID(ctx, id)
}

func (s *ProductService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
