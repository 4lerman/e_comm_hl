package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/4lerman/e_com/common/utils"
	orderTypes "github.com/4lerman/e_com/order/types"
	productTypes "github.com/4lerman/e_com/product/types"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store        orderTypes.OrderStore
	productStore productTypes.ProductStore
}

func NewHandler(store orderTypes.OrderStore, productStore productTypes.ProductStore) *Handler {
	return &Handler{
		store:        store,
		productStore: productStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("", h.handleListOrders).Methods(http.MethodGet)
	router.HandleFunc("", h.handleCreateOrder).Methods(http.MethodPost)
	router.HandleFunc("/search", h.handleOrderByStatusOrUser).Methods(http.MethodGet)
	router.HandleFunc("/{id}", h.handleGetOrderById).Methods(http.MethodGet)
	router.HandleFunc("/{id}", h.handleUpdateOrder).Methods(http.MethodPut)
	router.HandleFunc("/{id}", h.handleDeleteOrder).Methods(http.MethodDelete)
	router.HandleFunc("/{id}/order", h.handleCreateOrderItem).Methods(http.MethodPost)
}

func (h *Handler) handleListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.store.ListOrders()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, orders)
}

func (h *Handler) handleCreateOrder(w http.ResponseWriter, r *http.Request) {
	var payload orderTypes.CreateOrderPayload
	if err := utils.ReqParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err := h.store.CreateOrder(orderTypes.Order{
		UserID: payload.UserID,
		Total:  payload.Total,
		Status: payload.Status,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"msg": "Created successfully"})

}

func (h *Handler) handleGetOrderById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("id is not indicated"))
		return
	}

	orderId, _ := strconv.Atoi(id)

	order, err := h.store.GetOrderById(orderId)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("failed to get order by id: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, order)
}

func (h *Handler) handleUpdateOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("id is not indicated"))
		return
	}

	orderId, _ := strconv.Atoi(id)

	var payload orderTypes.UpdateOrderPayload
	if err := utils.ReqParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err := h.store.UpdateOrder(orderId, orderTypes.Order{
		UserID: payload.UserID,
		Total:  payload.Total,
		Status: payload.Status,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"msg": "Updated successfully"})
}

func (h *Handler) handleDeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("id is not indicated"))
		return
	}

	orderId, _ := strconv.Atoi(id)

	if err := h.store.DeleteOrder(orderId); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"msg": "Deleted successfully"})
}

func (h *Handler) handleOrderByStatusOrUser(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	user := r.URL.Query().Get("user")

	var orders []orderTypes.Order
	var err error

	if status != "" {
		orders, err = h.store.GetOrdersByStatus(status)
	} else if user != "" {
		userId, _ := strconv.Atoi(user)
		orders, err = h.store.GetOrdersByUserId(userId)
	} else {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("either status or user query parameter is required"))
		return
	}

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, orders)
}

func (h *Handler) handleCreateOrderItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("id is not indicated"))
		return
	}

	orderId, _ := strconv.Atoi(id)

	var payload orderTypes.CreateOrderItemPayload
	if err := utils.ReqParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	product, err := h.productStore.GetProductByID(payload.ProductID)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if product.Quantity < payload.Quantity {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("product %s is not available in quantity requested", product.Name))
		return
	}

	err = h.productStore.UpdateProduct(payload.ProductID, productTypes.Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity - payload.Quantity,
		Category:    product.Category,
	})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.CreateOrderItem(orderTypes.OrderItem{
		OrderID:   orderId,
		ProductID: payload.ProductID,
		Quantity:  payload.Quantity,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	totalOrder, err := h.store.GetOrderById(orderId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.UpdateOrder(orderId, orderTypes.Order{
		UserID: totalOrder.UserID,
		Status: totalOrder.Status,
		Total:  totalOrder.Total + float64(payload.Quantity)*product.Price,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"msg": "Created successfully"})

}
