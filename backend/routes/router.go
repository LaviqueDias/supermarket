package routes

import (
	"database/sql"
	"net/http"
	"github.com/LaviqueDias/supermarket/internal/domain/repositories"
	"github.com/LaviqueDias/supermarket/internal/handlers"
	"github.com/LaviqueDias/supermarket/internal/middleware"
	"github.com/LaviqueDias/supermarket/internal/usecases"

	"github.com/gorilla/mux"
)

func SetupRouter(db *sql.DB) *mux.Router {
	// Inicializar Repositories
	productRepo := repositories.NewProductRepository(db)
	clientRepo := repositories.NewClientRepository(db)
	employeeRepo := repositories.NewEmployeeRepository(db)
	cartRepo := repositories.NewCartRepository(db)
	promotionRepo := repositories.NewPromotionRepository(db)

	// Inicializar UseCases
	productUC := usecases.NewProductUseCase(productRepo)
	clientUC := usecases.NewClientUseCase(clientRepo)
	employeeUC := usecases.NewEmployeeUseCase(employeeRepo)
	cartUC := usecases.NewCartUseCase(cartRepo, productRepo)
	promotionUC := usecases.NewPromotionUseCase(promotionRepo)

	// Inicializar Handlers
	productHandler := handlers.NewProductHandler(productUC)
	clientHandler := handlers.NewClientHandler(clientUC)
	employeeHandler := handlers.NewEmployeeHandler(employeeUC)
	cartHandler := handlers.NewCartHandler(cartUC)
	promotionHandler := handlers.NewPromotionHandler(promotionUC)

	r := mux.NewRouter()

	// Aplicar middlewares globais
	r.Use(middleware.CORSMiddleware)
	r.Use(middleware.LoggerMiddleware)

	// ROTAS PÚBLICAS (sem autenticação)
	
	// Clientes
	r.HandleFunc("/client/register", clientHandler.Register).Methods("POST")
	r.HandleFunc("/client/login", clientHandler.Login).Methods("POST")
	
	// Funcionários
	r.HandleFunc("/employee/register", employeeHandler.Register).Methods("POST")
	r.HandleFunc("/employee/login", employeeHandler.Login).Methods("POST")
	
	// Produtos (visualização pública)
	r.HandleFunc("/product", productHandler.GetAll).Methods("GET")
	r.HandleFunc("/product/{id}", productHandler.GetByID).Methods("GET")

	// ROTAS PROTEGIDAS - APENAS FUNCIONÁRIOS
	
	employeeRoutes := r.PathPrefix("/").Subrouter()
	employeeRoutes.Use(func(next http.Handler) http.Handler {
		return middleware.AuthMiddleware(middleware.EmployeeOnlyMiddleware(next))
	})

	// Produtos (CRUD completo - apenas employees)
	employeeRoutes.HandleFunc("/product", productHandler.Create).Methods("POST")
	employeeRoutes.HandleFunc("/product/{id}", productHandler.Update).Methods("PUT")
	employeeRoutes.HandleFunc("/product/{id}", productHandler.Delete).Methods("DELETE")

	// Promoções (CRUD completo - apenas employees)
	employeeRoutes.HandleFunc("/promotion", promotionHandler.Create).Methods("POST")
	employeeRoutes.HandleFunc("/promotion", promotionHandler.GetAll).Methods("GET")
	employeeRoutes.HandleFunc("/promotion/add-product", promotionHandler.AddProduct).Methods("POST")

	// Gerenciar clientes e funcionários
	employeeRoutes.HandleFunc("/client", clientHandler.GetAll).Methods("GET")
	employeeRoutes.HandleFunc("/employee", employeeHandler.GetAll).Methods("GET")

	// ROTAS PROTEGIDAS - APENAS CLIENTES
	
	clientRoutes := r.PathPrefix("/").Subrouter()
	clientRoutes.Use(func(next http.Handler) http.Handler {
		return middleware.AuthMiddleware(middleware.ClientOnlyMiddleware(next))
	})

	// Carrinho (apenas clientes)
	clientRoutes.HandleFunc("/cart/add", cartHandler.AddItem).Methods("POST")
	clientRoutes.HandleFunc("/cart/{client_id}", cartHandler.GetCart).Methods("GET")
	clientRoutes.HandleFunc("/cart/item/{item_id}", cartHandler.RemoveItem).Methods("DELETE")

	return r
}