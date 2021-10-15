package handlers

import (
	"context"
	"encoding/json"

	"github.com/cankirma/go-api-for-san-marino/app/dtos"
	entities "github.com/cankirma/go-api-for-san-marino/app/entities"
	"github.com/cankirma/go-api-for-san-marino/pkg/database"
	"github.com/cankirma/go-api-for-san-marino/pkg/role_repository"
	"github.com/cankirma/go-api-for-san-marino/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

// GetCategories func gets all exists category.
// @Description Get all exists category.
// @Summary gets all exists category
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {array} entities.Category
// @Router /v1/category/getall [get]
func GetCategories(c *fiber.Ctx) error {
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	categories, err := db.GetAllCategories()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":      true,
			"msg":        err.Error(),
			"count":      0,
			"categories": nil,
		})
	}
	connRedis, err := database.RedisConnection()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":      true,
			"msg":        err.Error(),
			"count":      0,
			"categories": nil,
		})
	}
 marshaledCategories ,_:=	json.Marshal(categories)

	errSaveToRedis := connRedis.Set(context.Background(), "categories", marshaledCategories, 0).Err()
	if errSaveToRedis != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   errSaveToRedis.Error(),
		})
	}



	return c.JSON(fiber.Map{
		"error":      false,
		"msg":        nil,
		"count":      len(categories),
		"categories": categories,
	})
}

// GetCategory func gets category by given ID or 404 error.
// @Description Get Category by given ID.
// @Summary get Category by given ID
// @Tags Category
// @Accept json
// @Produce json
// @Param CategoryID path int true "CategoryID"
// @Param CategoryName path string true "CategoryName"
// @Param Description path string  true  "Description"
// @Success 200 {object} entities.Category
// @Router /v1/category/{id} [get]
func GetCategory(c *fiber.Ctx) error {
	request := c.Params("id")
	id, _ := strconv.Atoi(request)

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	category, err := db.GetCategoryById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":    true,
			"msg":      err.Error(),
			"category": nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      "success mustafa tebrik ederim",
		"category": category,
	})
}

// CreateCategory func for creates a new category.
// @Description Create a new category.
// @Summary creates a new category
// @Tags Category
// @Accept json
// @Produce json
// @Param CategoryId body string true "CategoryId"
// @Param CategoryName body string true "CategoryName"
// @Param Description body string true "Description"
// @Param Picture body string true "Picture"
// @Success 200 {object} entities.Category
// @Security ApiKeyAuth
// @Router /v1/category [post]
func CreateCategory(c *fiber.Ctx) error {
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

	credential := claims.Credentials[role_repository.CategoryCreateCredential]

	if !credential {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied, check role_repository of your token",
		})
	}

	category := &dtos.CreateCategoryModel{}

	if err := c.BodyParser(category); err != nil {
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

	validate := utils.NewValidator()

	if err := validate.Struct(category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	if err := db.InsertCategory(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      nil,
		"category": category,
	})
}

// UpdateCategory func for updates Category by given ID.
// @Description Update Category.
// @Summary update Category
// @Tags Category
// @Accept json
// @Produce json
// @Param CategoryId body string true "CategoryId"
// @Param CategoryName body string true "CategoryName"
// @Param Description body string true "Description"
// @Param Picture body string true "Picture"
// @Success 202 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/category [put]
func UpdateCategory(c *fiber.Ctx) error {

	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
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

	credential := claims.Credentials[role_repository.CategoryUpdateCredential]

	if !credential {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied, check role_repository of your token",
		})
	}
	category := &dtos.UpdateCategoryModel{}

	if err := c.BodyParser(category); err != nil {
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

	_, err = db.GetCategoryById(category.CategoryId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "category with this ID not found",
		})
	}

	if err := db.UpdateCategory(category.CategoryId, category); err != nil {

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

// DeleteCategory func for deletes Category by given ID.
// @Description Category  by given ID.
// @Summary Category  by given ID
// @Tags Category
// @Accept json
// @Produce json
// @Param id body string true "category ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/category [delete]
func DeleteCategory(c *fiber.Ctx) error {

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

	// Set credential `category:delete` from JWT data of current category.
	credential := claims.Credentials[role_repository.CategoryDeleteCredential]

	if !credential {

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied, check role_repository of your token",
		})
	}

	// Create new category struct
	foundedCategory := &entities.Category{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(foundedCategory); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	validate := utils.NewValidator()

	// Validate category fields.
	if err := validate.StructPartial(foundedCategory, "CategoryId"); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	*foundedCategory, err = db.CategoryRepository.GetCategoryById(foundedCategory.CategoryId)

	if err := db.DeleteCategory(foundedCategory.CategoryId); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 204 no content.
	return c.SendStatus(fiber.StatusNoContent)
}
