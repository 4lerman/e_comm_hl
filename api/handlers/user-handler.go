package handlers

import (
    "net/http"
    "github.com/4lerman/e_com/common/config"
    "github.com/4lerman/e_com/common/utils"
    "github.com/gorilla/mux"
)

var userServiceURL = configs.Envs.Users_Url + "/users"

// GetUsersHandler godoc
// @Summary List all users
// @Description Get all users from the user service
// @Tags users
// @Produce  json
// @Success 200 {array} types.User
// @Failure 500 {object} map[string]string
// @Router /users [get]
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
    req, err := http.NewRequest(http.MethodGet, userServiceURL, nil)
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

// CreateUserHandler godoc
// @Summary Create a new user
// @Description Create a new user in the user service
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body types.CreateUserPayload true "User to create"
// @Success 201 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    req, err := http.NewRequest(http.MethodPost, userServiceURL, r.Body)
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

// GetUserByQueryHandler godoc
// @Summary Get users by query
// @Description Get users by name or email from the user service
// @Tags users
// @Produce  json
// @Param name query string false "User name"
// @Param email query string false "User email"
// @Success 200 {array} types.User
// @Failure 500 {object} map[string]string
// @Router /users/search [get]
func GetUserByQueryHandler(w http.ResponseWriter, r *http.Request) {
    queryParams := r.URL.RawQuery
    userServiceURL = userServiceURL + "/search?" + queryParams

    req, err := http.NewRequest(http.MethodGet, userServiceURL, nil)
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

// GetUserByIDHandler godoc
// @Summary Get user by ID
// @Description Get a user by ID from the user service
// @Tags users
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} types.User
// @Failure 500 {object} map[string]string
// @Router /users/{id} [get]
func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userID := vars["id"]
    userServiceURL = userServiceURL + "/" + userID

    req, err := http.NewRequest(http.MethodGet, userServiceURL, nil)
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

// UpdateUserHandler godoc
// @Summary Update a user
// @Description Update a user in the user service
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body types.UpdateUserPayload true "User to update"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [put]
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userID := vars["id"]
    userServiceURL = userServiceURL + "/" + userID

    req, err := http.NewRequest(http.MethodPut, userServiceURL, r.Body)
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

// DeleteUserHandler godoc
// @Summary Delete a user
// @Description Delete a user in the user service
// @Tags users
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [delete]
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userID := vars["id"]
    userServiceURL = userServiceURL + "/" + userID

    req, err := http.NewRequest(http.MethodDelete, userServiceURL, nil)
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
