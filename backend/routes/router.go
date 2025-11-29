// ============================================
// api/routes/router.go
// ============================================
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
	cartRepo := repositories.NewCartRepository(db)
	promotionRepo := repositories.NewPromotionRepository(db)

	// Inicializar UseCases
	productUC := usecases.NewProductUseCase(productRepo)
	clientUC := usecases.NewClientUseCase(clientRepo)
	cartUC := usecases.NewCartUseCase(cartRepo, productRepo)
	promotionUC := usecases.NewPromotionUseCase(promotionRepo)

	// Inicializar Handlers
	productHandler := handlers.NewProductHandler(productUC)
	clientHandler := handlers.NewClientHandler(clientUC)
	cartHandler := handlers.NewCartHandler(cartUC)
	promotionHandler := handlers.NewPromotionHandler(promotionUC)

	// Configurar Router
	r := mux.NewRouter()

	// Aplicar middlewares globais
	r.Use(middleware.CORSMiddleware)
	r.Use(middleware.LoggerMiddleware)

	// Rotas Públicas (sem autenticação)
	r.HandleFunc("/client/register", clientHandler.Register).Methods("POST")
	r.HandleFunc("/client/login", clientHandler.Login).Methods("POST")
	r.HandleFunc("/product", productHandler.GetAll).Methods("GET")
	r.HandleFunc("/product/{id}", productHandler.GetByID).Methods("GET")

	// Rotas Protegidas (com autenticação JWT)
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(func(next http.Handler) http.Handler {
		return middleware.AuthMiddleware(next)
	})

	// Produtos (protegidos)
	protected.HandleFunc("/product", productHandler.Create).Methods("POST")
	protected.HandleFunc("/product/{id}", productHandler.Update).Methods("PUT")
	protected.HandleFunc("/product/{id}", productHandler.Delete).Methods("DELETE")

	// Clientes (protegidos)
	protected.HandleFunc("/client", clientHandler.GetAll).Methods("GET")

	// Carrinho (protegidos)
	protected.HandleFunc("/cart/add", cartHandler.AddItem).Methods("POST")
	protected.HandleFunc("/cart/{client_id}", cartHandler.GetCart).Methods("GET")
	protected.HandleFunc("/cart/item/{item_id}", cartHandler.RemoveItem).Methods("DELETE")

	// Promoções (protegidas)
	protected.HandleFunc("/promotion", promotionHandler.Create).Methods("POST")
	protected.HandleFunc("/promotion", promotionHandler.GetAll).Methods("GET")
	protected.HandleFunc("/promotion/add-product", promotionHandler.AddProduct).Methods("POST")

	return r
}