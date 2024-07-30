package handlers

import (
	"net/http"

	configs "github.com/4lerman/e_com/common/config"
	"github.com/4lerman/e_com/common/utils"
	"github.com/gorilla/mux"
)

var paymentServiceURL = configs.Envs.Payments_Url + "/payments"

// GetPaymentsHandler godoc
// @Summary Get all payments
// @Description Get details of all payments
// @Tags payments
// @Accept  json
// @Produce  json
// @Success 200 {array} types.Payment
// @Failure 500 {object} map[string]string
// @Router /payments [get]
func GetPaymentsHandler(w http.ResponseWriter, r *http.Request) {
    req, err := http.NewRequest(http.MethodGet, paymentServiceURL, nil)
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

// CreatePaymentHandler godoc
// @Summary Create a payment
// @Description Create a new payment
// @Tags payments
// @Accept  json
// @Produce  json
// @Param payment body types.CreatePaymentPayload true "Payment payload"
// @Success 201 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payments [post]
func CreatePaymentHandler(w http.ResponseWriter, r *http.Request) {
    req, err := http.NewRequest(http.MethodPost, paymentServiceURL, r.Body)
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

// GetPaymentsByQueryHandler godoc
// @Summary Get payments by query
// @Description Get payments by status, user, or order
// @Tags payments
// @Accept  json
// @Produce  json
// @Param status query string false "Payment status"
// @Param user query int false "User ID"
// @Param order query int false "Order ID"
// @Success 200 {array} types.Payment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payments/search [get]
func GetPaymentsByQueryHandler(w http.ResponseWriter, r *http.Request) {
    queryParams := r.URL.RawQuery
    url := paymentServiceURL + "/search?" + queryParams

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

// GetPaymentByIDHandler godoc
// @Summary Get payment by ID
// @Description Get payment details by ID
// @Tags payments
// @Accept  json
// @Produce  json
// @Param id path int true "Payment ID"
// @Success 200 {object} types.Payment
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /payments/{id} [get]
func GetPaymentByIDHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    paymentID := vars["id"]

    url := paymentServiceURL + "/" + paymentID

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

// UpdatePaymentHandler godoc
// @Summary Update a payment
// @Description Update an existing payment
// @Tags payments
// @Accept  json
// @Produce  json
// @Param id path int true "Payment ID"
// @Param payment body types.UpdatePaymentPayload true "Payment payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payments/{id} [put]
func UpdatePaymentHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    paymentID := vars["id"]

    url := paymentServiceURL + "/" + paymentID

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

// DeletePaymentHandler godoc
// @Summary Delete a payment
// @Description Delete a payment by ID
// @Tags payments
// @Accept  json
// @Produce  json
// @Param id path int true "Payment ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payments/{id} [delete]
func DeletePaymentHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    paymentID := vars["id"]

    url := paymentServiceURL + "/" + paymentID

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