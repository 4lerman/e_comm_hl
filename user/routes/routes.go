package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/4lerman/e_com/common/utils"
	"github.com/4lerman/e_com/user/types"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("", h.handleListUsers).Methods(http.MethodGet)
	router.HandleFunc("", h.handleCreateUser).Methods(http.MethodPost)
	router.HandleFunc("/search", h.handleUserByNameOrEmail).Methods(http.MethodGet)
	router.HandleFunc("/{id}", h.handleGetUserById).Methods(http.MethodGet)
	router.HandleFunc("/{id}", h.handleUpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/{id}", h.handleDeleteUser).Methods(http.MethodDelete)
}

func (h *Handler) handleListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.store.ListUsers()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, users)
}

func (h *Handler) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateUserPayload
	if err := utils.ReqParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err := h.store.CreateUser(types.User{
		FullName: payload.FullName,
		Email:    payload.Email,
		UserRole: payload.UserRole,
		Address:  payload.Address,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"msg": "Created successfully"})

}

func (h *Handler) handleGetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("id is not indicated"))
		return
	}

	userId, _ := strconv.Atoi(id)

	user, err := h.store.GetUserById(userId)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("failed to get user by id: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

func (h *Handler) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("id is not indicated"))
		return
	}

	userId, _ := strconv.Atoi(id)

	var payload types.UpdateUserPayload
	if err := utils.ReqParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err := h.store.UpdateUser(userId, types.User{
		FullName: payload.FullName,
		UserRole: payload.UserRole,
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

	userId, _ := strconv.Atoi(id)

	if err := h.store.DeleteUser(userId); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"msg": "Deleted successfully"})
}

func (h *Handler) handleUserByNameOrEmail(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")

	var users []types.User
	var err error

	if name != "" {
		users, err = h.store.GetUsersByName(name)
	} else if email != "" {
		users, err = h.store.GetUsersByEmail(email)
	} else {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("either name or email query parameter is required"))
		return
	}

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, users)
}
