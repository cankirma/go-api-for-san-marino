package handlers

import (
	"github.com/cankirma/go-api-for-san-marino/app/entities"
	"github.com/cankirma/go-api-for-san-marino/pkg/database"
	"github.com/cankirma/go-api-for-san-marino/pkg/role_repository"
	"github.com/cankirma/go-api-for-san-marino/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func GetProducts(c *fiber.Ctx) error {
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	products, err := db.GetAllProducts()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":      true,
			"msg":        err.Error(),
			"count":      0,
			"categories": nil,
		})
	}


	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"count": len(products),
		"products": products,
	})
}

func GetProductById(c *fiber.Ctx) error  {
	request := c.Params("id")
	id, _ := strconv.Atoi(request)

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}


	product, err := db.GetProductById(id)
	if err != nil {
		// Return, if category not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":    true,
			"msg":      err.Error(),
			"product": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      "success mustafa tebrik ederim",
		"product": product,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	now := time.Now().Unix()


	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	expires := claims.Expires
	if now > expires {

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	credential := claims.Credentials[role_repository.ProductCreateCredential]

	if !credential {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied",
		})
	}


	product := &entities.Product{}

	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	validate := utils.NewValidator()

	if err := validate.Struct(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}
	if err := db.InsertProduct(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      nil,
		"product": product,
	})
}



func UpdateProduct(c *fiber.Ctx) error {

	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	now := time.Now().Unix()
	expires := claims.Expires
	if now > expires {

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	credential := claims.Credentials[role_repository.ProductUpdateCredential]

	if !credential {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied, check role_repository of your token",
		})
	}
	product := &entities.Product{}

	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}


	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	_, err = db.GetProductById(product.ProductId)
	if err != nil {

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "category with this ID not found",
		})
	}

	if err := db.UpdateProduct(product); err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"error": false,
		"msg":   nil,
	})
}


