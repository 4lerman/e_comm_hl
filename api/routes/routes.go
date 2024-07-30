package routes

import (
	"net/http"

	"github.com/4lerman/e_com/api/handlers"
	"github.com/gorilla/mux"

	_ "github.com/4lerman/e_com/docs"
    httpSwagger "github.com/swaggo/http-swagger"
)

func Routes(router *mux.Router) {
	router.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods(http.MethodGet)

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router = router.PathPrefix("/api/v1").Subrouter()

	usersRouter := router.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("", handlers.GetUsersHandler).Methods(http.MethodGet)
	usersRouter.HandleFunc("", handlers.CreateUserHandler).Methods(http.MethodPost)
	usersRouter.HandleFunc("/search", handlers.GetUserByQueryHandler).Methods(http.MethodGet)
	usersRouter.HandleFunc("/{id}", handlers.GetUserByIDHandler).Methods(http.MethodGet)
	usersRouter.HandleFunc("/{id}", handlers.UpdateUserHandler).Methods(http.MethodPut)
	usersRouter.HandleFunc("/{id}", handlers.DeleteUserHandler).Methods(http.MethodDelete)


	productsRouter := router.PathPrefix("/products").Subrouter()
	productsRouter.HandleFunc("", handlers.GetProductsHandler).Methods(http.MethodGet)
	productsRouter.HandleFunc("", handlers.CreateProductHandler).Methods(http.MethodPost)
	productsRouter.HandleFunc("/search", handlers.GetProductByQueryHandler).Methods(http.MethodGet)
	productsRouter.HandleFunc("/{id}", handlers.GetProductByIDHandler).Methods(http.MethodGet)
	productsRouter.HandleFunc("/{id}", handlers.UpdateProductHandler).Methods(http.MethodPut)
	productsRouter.HandleFunc("/{id}", handlers.DeleteProductHandler).Methods(http.MethodDelete)

	ordersRouter := router.PathPrefix("/orders").Subrouter()
	ordersRouter.HandleFunc("", handlers.GetOrdersHandler).Methods(http.MethodGet)
	ordersRouter.HandleFunc("", handlers.CreateOrderHandler).Methods(http.MethodPost)
	ordersRouter.HandleFunc("/search", handlers.GetOrdersByQueryHandler).Methods(http.MethodGet)
	ordersRouter.HandleFunc("/{id}", handlers.GetOrderByIDHandler).Methods(http.MethodGet)
	ordersRouter.HandleFunc("/{id}", handlers.UpdateOrderHandler).Methods(http.MethodPut)
	ordersRouter.HandleFunc("/{id}", handlers.DeleteOrderHandler).Methods(http.MethodDelete)

	paymentRouter := router.PathPrefix("/payments").Subrouter()
	paymentRouter.HandleFunc("", handlers.GetPaymentsHandler).Methods(http.MethodGet)
	paymentRouter.HandleFunc("", handlers.CreatePaymentHandler).Methods(http.MethodPost)
	paymentRouter.HandleFunc("/search", handlers.GetPaymentsByQueryHandler).Methods(http.MethodGet)
	paymentRouter.HandleFunc("/{id}", handlers.GetPaymentByIDHandler).Methods(http.MethodGet)
	paymentRouter.HandleFunc("/{id}", handlers.UpdatePaymentHandler).Methods(http.MethodPut)
	paymentRouter.HandleFunc("/{id}", handlers.DeletePaymentHandler).Methods(http.MethodDelete)
}