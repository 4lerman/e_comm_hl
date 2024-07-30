package handlers

import (
	"net/http"

	configs "github.com/4lerman/e_com/common/config"
	"github.com/4lerman/e_com/common/utils"
	"github.com/gorilla/mux"
)

var productServiceURL = configs.Envs.Products_Url + "/products"

// GetProductsHandler godoc
// @Summary List all products
// @Description Get all products from the product service
// @Tags products
// @Produce  json
// @Success 200 {array} types.Product
// @Failure 500 {object} map[string]string
// @Router /products [get]
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest(http.MethodGet, productServiceURL, nil)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	utils.ResCopy(w, resp.StatusCode, resp)
}

// CreateProductHandler godoc
// @Summary Create a new product
// @Description Create a new product in the product service
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body types.CreateProductPayload true "Product to create"
// @Success 201 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products [post]
func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest(http.MethodPost, productServiceURL, r.Body)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	utils.ResCopy(w, resp.StatusCode, resp)
}

// GetProductByQueryHandler godoc
// @Summary Get products by query
// @Description Get products by name or category from the product service
// @Tags products
// @Produce  json
// @Param name query string false "Product name"
// @Param category query string false "Product category"
// @Success 200 {array} types.Product
// @Failure 500 {object} map[string]string
// @Router /products/search [get]
func GetProductByQueryHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.RawQuery
	productServiceURL = productServiceURL + "/search?" + queryParams

	req, err := http.NewRequest(http.MethodGet, productServiceURL, nil)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	utils.ResCopy(w, resp.StatusCode, resp)
}

// GetProductByIDHandler godoc
// @Summary Get product by ID
// @Description Get a product by ID from the product service
// @Tags products
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} types.Product
// @Failure 500 {object} map[string]string
// @Router /products/{id} [get]
func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	productServiceURL = productServiceURL + "/" + productID

	req, err := http.NewRequest(http.MethodGet, productServiceURL, nil)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	utils.ResCopy(w, resp.StatusCode, resp)
}

// UpdateProductHandler godoc
// @Summary Update a product
// @Description Update a product in the product service
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Param product body types.UpdateProductPayload true "Product to update"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/{id} [put]
func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	productServiceURL = productServiceURL + "/" + productID

	req, err := http.NewRequest(http.MethodPut, productServiceURL, r.Body)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	utils.ResCopy(w, resp.StatusCode, resp)
}

// DeleteProductHandler godoc
// @Summary Delete a product
// @Description Delete a product in the product service
// @Tags products
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/{id} [delete]
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	productServiceURL = productServiceURL + "/" + productID

	req, err := http.NewRequest(http.MethodDelete, productServiceURL, nil)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	utils.ResCopy(w, resp.StatusCode, resp)
}
