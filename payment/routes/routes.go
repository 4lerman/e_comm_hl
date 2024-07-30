package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/4lerman/e_com/common/utils"
	"github.com/4lerman/e_com/payment/service"
	"github.com/4lerman/e_com/payment/types"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.PaymentStore
}

func NewHandler(store types.PaymentStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("", h.handleListPayments).Methods(http.MethodGet)
	router.HandleFunc("", h.handleCreatePayment).Methods(http.MethodPost)
	router.HandleFunc("/search", h.handlePaymentByQuery).Methods(http.MethodGet)
	router.HandleFunc("/{id}", h.handleGetPaymentById).Methods(http.MethodGet)
	router.HandleFunc("/{id}", h.handleUpdatePayment).Methods(http.MethodPut)
	router.HandleFunc("/{id}", h.handleDeletePayment).Methods(http.MethodDelete)
}

func (h *Handler) handleListPayments(w http.ResponseWriter, r *http.Request) {
	payments, err := h.store.ListPayments()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, payments)
}

func (h *Handler) handleCreatePayment(w http.ResponseWriter, r *http.Request) {
	var payload types.CreatePaymentPayload
	if err := utils.ReqParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	paymentResponse, err := service.MakePayment()
	if err != nil {
		payload.Status = "failed"
		fmt.Println("Payment failed: ", paymentResponse)
	} else {
		fmt.Println("Payment passed: ", paymentResponse)
		if paymentResponse.Status == "AUTH"{
			payload.Status = "success"
		}
	}

	err = h.store.CreatePayment(types.Payment{
		UserID:  payload.UserID,
		OrderID: payload.OrderID,
		Amount:  payload.Amount,
		Status:  payload.Status,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"msg": "Created successfully"})

}

func (h *Handler) handleGetPaymentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("id is not indicated"))
		return
	}

	paymentId, _ := strconv.Atoi(id)

	payment, err := h.store.GetPaymentById(paymentId)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("failed to get payment by id: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, payment)
}

func (h *Handler) handleUpdatePayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("id is not indicated"))
		return
	}

	paymentId, _ := strconv.Atoi(id)

	var payload types.UpdatePaymentPayload
	if err := utils.ReqParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err := h.store.UpdatePayment(paymentId, types.Payment{
		UserID:  payload.UserID,
		OrderID: payload.OrderID,
		Amount:  payload.Amount,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"msg": "Updated successfully"})
}

func (h *Handler) handleDeletePayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("id is not indicated"))
		return
	}

	paymentId, _ := strconv.Atoi(id)

	if err := h.store.DeletePayment(paymentId); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"msg": "Deleted successfully"})
}

func (h *Handler) handlePaymentByQuery(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	user := r.URL.Query().Get("user")
	order := r.URL.Query().Get("order")

	var payments []types.Payment
	var err error

	userId, _ := strconv.Atoi(user)
	orderId, _ := strconv.Atoi(order)

	if status != "" {
		payments, err = h.store.GetPaymentsByStatus(status)
	} else if user != "" {
		payments, err = h.store.GetPaymentsByUserId(userId)
	} else if order != "" {
		payments, err = h.store.GetPaymentsByOrderId(orderId)
	} else {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("either name or email query parameter is required"))
		return
	}

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, payments)
}
