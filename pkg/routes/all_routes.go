package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/cankirma/go-api-for-san-marino/app/handlers"
	"github.com/cankirma/go-api-for-san-marino/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func NotFoundRoute(a *fiber.App) {
	a.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "sorry, endpoint is not found",
			})
		},
	)
}
func AuthorizedRoutes(a *fiber.App) {
	route := a.Group("/api/v1")
	route.Post("/user/sign/out", middleware.JWTProtected(), handlers.UserSignOut)
	route.Post("/token/renew", middleware.JWTProtected(), handlers.RenewTokens)

	route.Post("/category/create",middleware.JWTProtected(),handlers.CreateCategory)
	route.Put("/category/update", middleware.JWTProtected(), handlers.UpdateCategory)
	route.Delete("/category/delete", middleware.JWTProtected(), handlers.DeleteCategory)



	route.Post("/products/create", middleware.JWTProtected(), handlers.CreateProduct)
	route.Put("/products/create", middleware.JWTProtected(), handlers. UpdateProduct)
}
func UnAuthorizedRoutes(a *fiber.App) {
	route := a.Group("/api/v1")
	route.Get("/category/getall", handlers.GetCategories)
	route.Get("/category/get/:id", handlers.GetCategory)

	route.Get(
		"/products/getall",
		handlers.GetProducts,
	)
	route.Get("/products/get/:id", handlers.GetProductById)

	route.Post("/user/signup", handlers.UserSignUp)
	route.Post("/user/signin", handlers.UserSignIn)
}
func SwaggerRoute(a *fiber.App) {
	route := a.Group("/swagger")
	route.Get("*", swagger.Handler)
}
