package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phk13/poc-go-ambassador/src/controllers"
	"github.com/phk13/poc-go-ambassador/src/middlewares"
)

func Setup(app *fiber.App) {
	api := app.Group("api")

	// Admin routes
	admin := api.Group("admin")
	admin.Post("register", controllers.Register)
	admin.Post("login", controllers.Login)

	//// Authenticated admin routes
	adminAuthenticated := admin.Use(middlewares.IsAuthenticated)
	adminAuthenticated.Get("user", controllers.User)
	adminAuthenticated.Post("logout", controllers.Logout)
	adminAuthenticated.Put("users/info", controllers.UpdateInfo)
	adminAuthenticated.Put("users/password", controllers.UpdatePassword)

	adminAuthenticated.Get("ambassadors", controllers.Ambassadors)

	adminAuthenticated.Get("products", controllers.Products)
	adminAuthenticated.Post("products", controllers.CreateProducts)
	adminAuthenticated.Get("products/:id", controllers.GetProduct)
	adminAuthenticated.Put("products/:id", controllers.UpdateProduct)
	adminAuthenticated.Delete("products/:id", controllers.DeleteProduct)

	adminAuthenticated.Get("users/:id/links", controllers.Link)

	adminAuthenticated.Get("orders", controllers.Orders)

	// Ambassador routes
	ambassador := api.Group("ambassador")
	ambassador.Post("register", controllers.Register)
	ambassador.Post("login", controllers.Login)
	ambassador.Get("products/frontend", controllers.ProductsFrontend)
	ambassador.Get("products/backend", controllers.ProductsBackend)

	//// Authenticated ambassador routes
	ambassadorAuthenticated := ambassador.Use(middlewares.IsAuthenticated)
	ambassadorAuthenticated.Get("user", controllers.User)
	ambassadorAuthenticated.Post("logout", controllers.Logout)
	ambassadorAuthenticated.Put("users/info", controllers.UpdateInfo)
	ambassadorAuthenticated.Put("users/password", controllers.UpdatePassword)

	ambassadorAuthenticated.Post("links", controllers.CreateLink)
	ambassadorAuthenticated.Get("stats", controllers.Stats)
	ambassadorAuthenticated.Get("rankings", controllers.Rankings)

	// Checkout routes
	checkout := api.Group("checkout")
	checkout.Get("links/:code", controllers.GetLink)
	checkout.Post("orders", controllers.CreateOrder)
	checkout.Post("orders/confirm", controllers.CompleteOrder)
}
