package api

import (
	"bytes"
	"encoding/json"
	oapimiddleware "github.com/oapi-codegen/echo-middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arke-so/mini-arke/internal/repository"
	"github.com/arke-so/mini-arke/internal/service"
	"github.com/arke-so/mini-arke/internal/testutil"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestHandler(t *testing.T) (*Handler, *echo.Echo, func()) {
	t.Helper()

	db := testutil.SetupTestDB(t)
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	handler := NewHandler(productService)

	e := echo.New()

	// Add OpenAPI validation middleware
	swagger, err := GetSwagger()
	require.NoError(t, err)
	swagger.Servers = nil
	e.Use(oapimiddleware.OapiRequestValidator(swagger))

	cleanup := func() {
		testutil.CleanupTestDB(t, db)
	}

	return handler, e, cleanup
}

func TestListProducts_EmptyList(t *testing.T) {
	handler, e, cleanup := setupTestHandler(t)
	defer cleanup()

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.ListProducts(c, ListProductsParams{})
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response ProductList
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, 0, response.Total)
	assert.Empty(t, response.Products)
}

func TestListProducts_ReturnsCreatedProduct(t *testing.T) {
	handler, e, cleanup := setupTestHandler(t)
	defer cleanup()

	// Create a product first
	createReq := CreateProductRequest{
		Sku:   "TEST-001",
		Name:  "Test Product",
		Price: 29.99,
	}
	body, _ := json.Marshal(createReq)
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.CreateProduct(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Now list products
	req = httptest.NewRequest(http.MethodGet, "/products", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	err = handler.ListProducts(c, ListProductsParams{})
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response ProductList
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, 1, response.Total)
	assert.Len(t, response.Products, 1)
	assert.Equal(t, "TEST-001", response.Products[0].Sku)
	assert.Equal(t, "Test Product", response.Products[0].Name)
	assert.Equal(t, float32(29.99), response.Products[0].Price)
}

func TestCreateProduct_MissingRequiredFields(t *testing.T) {
	handler, e, cleanup := setupTestHandler(t)
	defer cleanup()

	// Register routes
	RegisterHandlers(e, handler)

	// Empty body - missing all required fields
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader([]byte("{}")))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCreateProduct_Success(t *testing.T) {
	handler, e, cleanup := setupTestHandler(t)
	defer cleanup()

	createReq := CreateProductRequest{
		Sku:   "PROD-123",
		Name:  "Awesome Product",
		Price: 99.99,
	}
	stockQty := 50
	createReq.StockQuantity = &stockQty

	body, _ := json.Marshal(createReq)
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.CreateProduct(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var response Product
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.NotEmpty(t, response.Id)
	assert.Equal(t, "PROD-123", response.Sku)
	assert.Equal(t, "Awesome Product", response.Name)
	assert.Equal(t, float32(99.99), response.Price)
	assert.Equal(t, 50, response.StockQuantity)
	assert.NotEmpty(t, response.CreatedAt)
}
