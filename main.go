package main

import (
	_ "github.com/cankirma/go-api-for-san-marino/docs"
	"github.com/cankirma/go-api-for-san-marino/pkg/configs"
	"github.com/cankirma/go-api-for-san-marino/pkg/middleware"
	"github.com/cankirma/go-api-for-san-marino/pkg/routes"
	"github.com/cankirma/go-api-for-san-marino/pkg/utils"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	config := configs.FiberConfig()


	app := fiber.New(config)


	middleware.FiberMiddleware(app)

	routes.SwaggerRoute(app)
	routes.UnAuthorizedRoutes(app)
	routes.AuthorizedRoutes(app)
	routes.NotFoundRoute(app)


	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}

}

