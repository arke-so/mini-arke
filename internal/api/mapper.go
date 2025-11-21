package api

import (
	"github.com/arke-so/mini-arke/internal/models"
	"github.com/arke-so/mini-arke/internal/service"
)

// toProductResponse converts a domain Product to API Product response
func toProductResponse(p *models.Product) Product {
	return Product{
		Id:            p.ID,
		Sku:           p.SKU,
		Name:          p.Name,
		Description:   p.Description,
		Price:         float32(p.Price),
		StockQuantity: p.StockQuantity,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

// toProductListResponse converts a slice of domain Products to API Product responses
func toProductListResponse(products []models.Product) []Product {
	result := make([]Product, len(products))
	for i, p := range products {
		result[i] = toProductResponse(&p)
	}
	return result
}

// toCreateProductInput converts API CreateProductRequest to service CreateProductInput
func toCreateProductInput(req CreateProductRequest) service.CreateProductInput {
	input := service.CreateProductInput{
		SKU:           req.Sku,
		Name:          req.Name,
		Description:   req.Description,
		Price:         float64(req.Price),
		StockQuantity: 0,
	}

	if req.StockQuantity != nil {
		input.StockQuantity = *req.StockQuantity
	}

	return input
}

// toUpdateProductInput converts API UpdateProductRequest to service UpdateProductInput
func toUpdateProductInput(req UpdateProductRequest) service.UpdateProductInput {
	input := service.UpdateProductInput{}

	if req.Name != nil {
		input.Name = req.Name
	}

	if req.Description != nil {
		input.Description = req.Description
	}

	if req.Price != nil {
		price := float64(*req.Price)
		input.Price = &price
	}

	if req.StockQuantity != nil {
		input.StockQuantity = req.StockQuantity
	}

	return input
}
