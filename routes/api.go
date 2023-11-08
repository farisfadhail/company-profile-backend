package routes

import (
	"plastindo-back-end/config"
	"plastindo-back-end/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	app.Static("/public", config.ProjectRootPath+"/public/asset")

	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Plastindo Group")
	})

	// Parent Category
	parentCategory := api.Group("/parent-category")
	parentCategory.Get("/", handler.GetAllParentCategoryHandler).Name("parentCategory.index")
	parentCategory.Post("/store", handler.StoreParentCategoryHandler).Name("parentCategory.store")
	parentCategory.Get("/:parentSlug", handler.GetBySlugParentCategoryHandler).Name("parentCategory.show")
	parentCategory.Put("/:parentSlug/update", handler.UpdateParentCategoryHandler).Name("parentCategory.update")
	parentCategory.Delete("/:parentSlug", handler.DeleteParentCategoryHandler).Name("parentCategory.destroy")
}