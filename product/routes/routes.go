package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/4lerman/e_com/common/utils"
	"github.com/4lerman/e_com/product/types"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{
		store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("", h.handleGetProducts).Methods(http.MethodGet)
	router.HandleFunc("", h.handleCreateProduct).Methods(http.MethodPost)
	router.HandleFunc("/search", h.handleProductByNameOrCategory).Methods(http.MethodGet)
	router.HandleFunc("/{id}", h.handleGetProductById).Methods(http.MethodGet)
	router.HandleFunc("/{id}", h.handleUpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/{id}", h.handleDeleteUser).Methods(http.MethodDelete)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateProductPayload
	if err := utils.ReqParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.CreateProduct(types.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		Quantity:    payload.Quantity,
		Category:    payload.Category,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetProducts()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ps)
}

func (h *Handler) handleGetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("id is not indicated"))
		return
	}

	productId, _ := strconv.Atoi(id)

	product, err := h.store.GetProductByID(productId)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("failed to get product by id: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, product)
}

func (h *Handler) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("id is not indicated"))
		return
	}

	product, _ := strconv.Atoi(id)

	var payload types.UpdateProductPayload
	if err := utils.ReqParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err := h.store.UpdateProduct(product, types.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		Quantity:    payload.Quantity,
		Category:    payload.Category,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"msg": "Updated successfully"})
}

func (h *Handler) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("id is not indicated"))
		return
	}

	productId, _ := strconv.Atoi(id)

	if err := h.store.DeleteProduct(productId); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"msg": "Deleted successfully"})
}

func (h *Handler) handleProductByNameOrCategory(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("category")

	var products []types.Product
	var err error

	if name != "" {
		products, err = h.store.GetProductsByName(name)
	} else if email != "" {
		products, err = h.store.GetProductsByCategory(email)
	} else {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("either name or category query parameter is required"))
		return
	}

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, products)
}
