// internal/api/handler.go
package api

import (
	"errors"
	"net/http"

	"github.com/arke-so/mini-arke/internal/repository"
	"github.com/arke-so/mini-arke/internal/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	productService *service.ProductService
}

func NewHandler(productService *service.ProductService) *Handler {
	return &Handler{
		productService: productService,
	}
}

// ListProducts handles GET /products
func (h *Handler) ListProducts(ctx echo.Context, params ListProductsParams) error {
	limit := 20
	offset := 0

	if params.Limit != nil {
		limit = *params.Limit
	}
	if params.Offset != nil {
		offset = *params.Offset
	}

	products, total, err := h.productService.List(ctx.Request().Context(), limit, offset)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, Error{
			Message: "Failed to list products",
			Details: stringPtr(err.Error()),
		})
	}

	return ctx.JSON(http.StatusOK, ProductList{
		Products: toProductListResponse(products),
		Total:    total,
	})
}

// CreateProduct handles POST /products
func (h *Handler) CreateProduct(ctx echo.Context) error {
	var req CreateProductRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, Error{
			Message: "Invalid request body",
			Details: stringPtr(err.Error()),
		})
	}

	input := toCreateProductInput(req)
	product, err := h.productService.Create(ctx.Request().Context(), input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, Error{
			Message: "Failed to create product",
			Details: stringPtr(err.Error()),
		})
	}

	return ctx.JSON(http.StatusCreated, toProductResponse(product))
}

// GetProduct handles GET /products/{productId}
func (h *Handler) GetProduct(ctx echo.Context, productId uuid.UUID) error {
	product, err := h.productService.Get(ctx.Request().Context(), productId)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ctx.JSON(http.StatusNotFound, Error{
				Message: "Product not found",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, Error{
			Message: "Failed to get product",
			Details: stringPtr(err.Error()),
		})
	}

	return ctx.JSON(http.StatusOK, toProductResponse(product))
}

// UpdateProduct handles PUT /products/{productId}
func (h *Handler) UpdateProduct(ctx echo.Context, productId uuid.UUID) error {
	var req UpdateProductRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, Error{
			Message: "Invalid request body",
			Details: stringPtr(err.Error()),
		})
	}

	input := toUpdateProductInput(req)
	product, err := h.productService.Update(ctx.Request().Context(), productId, input)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ctx.JSON(http.StatusNotFound, Error{
				Message: "Product not found",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, Error{
			Message: "Failed to update product",
			Details: stringPtr(err.Error()),
		})
	}

	return ctx.JSON(http.StatusOK, toProductResponse(product))
}

// DeleteProduct handles DELETE /products/{productId}
func (h *Handler) DeleteProduct(ctx echo.Context, productId uuid.UUID) error {
	err := h.productService.Delete(ctx.Request().Context(), productId)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ctx.JSON(http.StatusNotFound, Error{
				Message: "Product not found",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, Error{
			Message: "Failed to delete product",
			Details: stringPtr(err.Error()),
		})
	}

	return ctx.NoContent(http.StatusNoContent)
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}
