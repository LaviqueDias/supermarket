// ============================================
// api/routes/router.go
// ============================================
package routes

import (
	"database/sql"
	"github.com/LaviqueDias/supermarket/internal/domain/repositories"
	"github.com/LaviqueDias/supermarket/internal/handlers"
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

	// Rotas de Produtos
	r.HandleFunc("/product", productHandler.Create).Methods("POST")
	r.HandleFunc("/product", productHandler.GetAll).Methods("GET")
	r.HandleFunc("/product/{id}", productHandler.GetByID).Methods("GET")
	r.HandleFunc("/product/{id}", productHandler.Update).Methods("PUT")
	r.HandleFunc("/product/{id}", productHandler.Delete).Methods("DELETE")

	// Rotas de Clientes
	r.HandleFunc("/client", clientHandler.Register).Methods("POST")
	r.HandleFunc("/client", clientHandler.GetAll).Methods("GET")

	// Rotas de Carrinho
	r.HandleFunc("/cart/add", cartHandler.AddItem).Methods("POST")
	r.HandleFunc("/cart/{client_id}", cartHandler.GetCart).Methods("GET")
	r.HandleFunc("/cart/item/{item_id}", cartHandler.RemoveItem).Methods("DELETE")

	// Rotas de Promoções
	r.HandleFunc("/promotion", promotionHandler.Create).Methods("POST")
	r.HandleFunc("/promotion", promotionHandler.GetAll).Methods("GET")
	r.HandleFunc("/promotion/add-product", promotionHandler.AddProduct).Methods("POST")

	return r
}