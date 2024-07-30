package handlers

import (
	"net/http"

	configs "github.com/4lerman/e_com/common/config"
	"github.com/4lerman/e_com/common/utils"
	"github.com/gorilla/mux"
)

var orderServiceURL = configs.Envs.Orders_Url + "/orders"

// GetOrdersHandler godoc
// @Summary Get all orders
// @Description Get details of all orders
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} types.Order
// @Failure 500 {object} map[string]string
// @Router /orders [get]
func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
    req, err := http.NewRequest(http.MethodGet, orderServiceURL, nil)
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

// CreateOrderHandler godoc
// @Summary Create a new order
// @Description Create a new order
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body types.CreateOrderPayload true "Order payload"
// @Success 201 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders [post]
func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
    req, err := http.NewRequest(http.MethodPost, orderServiceURL, r.Body)
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

// GetOrdersByQueryHandler godoc
// @Summary Get orders by query
// @Description Get orders by status or user
// @Tags orders
// @Accept  json
// @Produce  json
// @Param status query string false "Order status"
// @Param user query int false "User ID"
// @Success 200 {array} types.Order
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/search [get]
func GetOrdersByQueryHandler(w http.ResponseWriter, r *http.Request) {
    queryParams := r.URL.RawQuery
    url := orderServiceURL + "/search?" + queryParams

    req, err := http.NewRequest(http.MethodGet, url, nil)
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

// GetOrderByIDHandler godoc
// @Summary Get order by ID
// @Description Get order details by ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Success 200 {object} types.Order
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /orders/{id} [get]
func GetOrderByIDHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    orderID := vars["id"]

    url := orderServiceURL + "/" + orderID

    req, err := http.NewRequest(http.MethodGet, url, nil)
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

// UpdateOrderHandler godoc
// @Summary Update an order
// @Description Update an existing order
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Param order body types.UpdateOrderPayload true "Order payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/{id} [put]
func UpdateOrderHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    orderID := vars["id"]

    url := orderServiceURL + "/" + orderID

    req, err := http.NewRequest(http.MethodPut, url, r.Body)
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

// DeleteOrderHandler godoc
// @Summary Delete an order
// @Description Delete an order by ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/{id} [delete]
func DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    orderID := vars["id"]

    url := orderServiceURL + "/" + orderID

    req, err := http.NewRequest(http.MethodDelete, url, nil)
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

// CreateOrderItemHandler godoc
// @Summary Create an order item
// @Description Create a new order item for an order
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Param orderItem body types.CreateOrderItemPayload true "Order item payload"
// @Success 201 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/{id}/order [post]
func CreateOrderItemHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    orderID := vars["id"]

    url := orderServiceURL + "/" + orderID + "/order"
    req, err := http.NewRequest(http.MethodPost, url, r.Body)
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
