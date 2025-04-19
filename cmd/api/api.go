package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gabriel-assis7/ecom-api-rest-go/service/product"
	"github.com/gabriel-assis7/ecom-api-rest-go/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Start() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userService := user.NewHandler(userStore)
	userService.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productService := product.NewHandler(productStore)
	productService.RegisterRoutes(subrouter)

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
