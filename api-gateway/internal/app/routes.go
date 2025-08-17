package app

import (
	"net/http"

	"github.com/alibekkenny/simple-marketplace/api-gateway/internal/middleware"
	"github.com/alibekkenny/simple-marketplace/api-gateway/internal/transport/handler"
	"github.com/justinas/alice"
)

func (a *App) routes(user *handler.UserHandler,
	category *handler.CategoryHandler,
	product *handler.ProductHandler,
	productOffer *handler.ProductOfferHandler,
	cart *handler.CartHandler,
	order *handler.OrderHandler) http.Handler {

	standartChain := alice.New(
		middleware.LoggingMiddleware,
		middleware.RecoveryMiddleware,
	)
	authChain := standartChain.Append(middleware.AuthMiddleware(a.cfg.JWTSecret))

	adminRole := authChain.Append(middleware.RoleMiddleware("admin"))
	supplierRole := authChain.Append(middleware.RoleMiddleware("supplier"))
	// buyerRole := authChain.Append(middleware.RoleMiddleware("buyer"))

	mux := http.NewServeMux()

	mux.Handle("POST /user/register", standartChain.Then(http.HandlerFunc(user.Register)))
	mux.Handle("POST /user/login", standartChain.Then(http.HandlerFunc(user.Login)))

	mux.Handle("POST /category", adminRole.Then(http.HandlerFunc(category.CreateCategory)))
	mux.Handle("PUT /category/{id}", adminRole.Then(http.HandlerFunc(category.UpdateCategory)))
	mux.Handle("DELETE /category/{id}", adminRole.Then(http.HandlerFunc(category.DeleteCategory)))
	mux.Handle("GET /category", standartChain.Then(http.HandlerFunc(category.ListCategories)))
	mux.Handle("GET /category/{category_id}/products", standartChain.Then(http.HandlerFunc(product.ListProductsByCategory)))

	mux.Handle("POST /product", supplierRole.Then(http.HandlerFunc(product.CreateProduct)))
	mux.Handle("PUT /product/{id}", supplierRole.Then(http.HandlerFunc(product.UpdateProduct)))
	mux.Handle("DELETE /product/{id}", adminRole.Then(http.HandlerFunc(product.DeleteProduct)))
	mux.Handle("GET /product/{id}", standartChain.Then(http.HandlerFunc(product.GetProductByID)))

	mux.Handle("POST /product/{product_id}/offer", supplierRole.Then(http.HandlerFunc(productOffer.CreateProductOffer)))
	mux.Handle("PUT /product/{product_id}/offer", supplierRole.Then(http.HandlerFunc(productOffer.UpdateProductOffer)))
	mux.Handle("DELETE /product/{product_id}/offer/{offer_id}", supplierRole.Then(http.HandlerFunc(productOffer.DeleteProductOffer)))
	mux.Handle("GET /product/{product_id}/offer", standartChain.Then(http.HandlerFunc(productOffer.ListProductOffersByProductID)))
	mux.Handle("GET /product/{product_id}/offer/{offer_id}", standartChain.Then(http.HandlerFunc(productOffer.GetProductOfferByID)))

	mux.Handle("GET /cart", authChain.Then(http.HandlerFunc(cart.GetCart)))
	mux.Handle("POST /cart/item", authChain.Then(http.HandlerFunc(cart.AddToCart)))
	mux.Handle("PUT /cart/item/{offer_id}", authChain.Then(http.HandlerFunc(cart.UpdateCartItem)))
	mux.Handle("DELETE /cart/item/{offer_id}", authChain.Then(http.HandlerFunc(cart.RemoveCartItem)))
	mux.Handle("DELETE /cart/item", authChain.Then(http.HandlerFunc(cart.ClearCart)))
	mux.Handle("POST /cart/checkout", authChain.Then(http.HandlerFunc(order.Checkout)))

	mux.Handle("GET /order/{id}", authChain.Then(http.HandlerFunc(order.GetOrderByID)))
	mux.Handle("GET /order", authChain.Then(http.HandlerFunc(order.ListOrdersByUserID)))

	mux.Handle("GET /supplier/{supplier_id}/product", standartChain.Then(http.HandlerFunc(productOffer.ListProductOffersBySupplierID)))

	return mux
}
