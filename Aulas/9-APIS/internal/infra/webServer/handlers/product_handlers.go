package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Frank-Macedo/PosGoLang/cursoGo/Aulas/9-APIS/internal/dto"
	"github.com/Frank-Macedo/PosGoLang/cursoGo/Aulas/9-APIS/internal/entity"
	entityPKG "github.com/Frank-Macedo/PosGoLang/cursoGo/Aulas/9-APIS/pkg/entity"

	"github.com/Frank-Macedo/PosGoLang/cursoGo/Aulas/9-APIS/internal/infra/database"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {

	return &ProductHandler{
		ProductDB: db,
	}
}

// CreateProduct godoc
// @Summary      Create a new product
// @Description  Create a new product with name and price
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      dto.CreateProductInput  true  "Product data"
// @Success      201  {object}  dto.CreateProductInput
// @Failure      400  {object}  Error
// @Failure      500  {object}  Error
// @Router       /products [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(p)

	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// GetProduct godoc
// @Summary      Get a product by ID
// @Description  Retrieve a product by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "Product ID"
// @Success      200  {object}  entity.Product
// @Failure      400  {object}  Error
// @Failure      404  {object}  Error
// @Router       /products/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "error", http.StatusBadRequest)
		return
	}
	product, err := h.ProductDB.FindByID(id)

	if err != nil {
		http.Error(w, "error", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)

}

// UpdateProduct godoc
// @Summary      Update a product by ID
// @Description  Update a product by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "Product ID"
// @Param        product  body      entity.Product  true  "Product data"
// @Success      200  {object}  entity.Product
// @Failure      400  {object}  Error
// @Failure      404  {object}  Error
// @Router       /products/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "error", http.StatusBadRequest)
		return
	}

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {

		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	product.ID, err = entityPKG.ParseID(id)

	product.Name = product.Name
	product.Price = product.Price

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.ProductDB.FindByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Update(&product)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

// DeleteProduct godoc
// @Summary      Delete a product by ID
// @Description  Delete a product by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "Product ID"
// @Success      200  {object}  string
// @Failure      400  {object}  Error
// @Failure      404  {object}  Error
// @Router       /products/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := h.ProductDB.FindByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

// GetProducts godoc
// @Summary      Get all products
// @Description  Retrieve a list of products with pagination
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        page  query      int  false  "Page number"  default(0)
// @Param        limit  query      int  false  "Number of products per page"
// @Success      200  {array}  entity.Product
// @Failure      400  {object}  Error
// @Router       /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	fmt.Println(page)
	fmt.Println(limit)

	intPage, err := strconv.Atoi(page)

	if err != nil {
		intPage = 0
	}

	intLimit, err := strconv.Atoi(limit)

	if err != nil {
		intLimit = 0
	}

	products, err := h.ProductDB.FindAll(intPage, intLimit, "asc")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)

}
