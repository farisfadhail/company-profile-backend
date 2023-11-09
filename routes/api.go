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
	parentCategory.Get("/:slug", handler.GetBySlugParentCategoryHandler).Name("parentCategory.show")
	parentCategory.Put("/:slug/update", handler.UpdateParentCategoryHandler).Name("parentCategory.update")
	parentCategory.Delete("/:slug", handler.DeleteParentCategoryHandler).Name("parentCategory.destroy")

	// Product Category
	productCategory := api.Group("/product-category")
	productCategory.Get("/", handler.GetAllProductCategoryHandler).Name("productCategory.index")
	productCategory.Post("/store", handler.StoreProductCategoryHandler).Name("productCategory.store")
	productCategory.Get("/:slug", handler.GetBySlugProductCategoryHandler).Name("productCategory.show")
	productCategory.Put("/:slug/update", handler.UpdateProductCategoryHandler).Name("productCategory.update")
	productCategory.Delete("/:slug", handler.DeleteProductCategoryHandler).Name("productCategory.destroy")

	product := api.Group("/product")
	product.Get("/", handler.GetAllProductHandler).Name("product.index")
	product.Post("/store", handler.StoreProductHandler).Name("product.store")
	product.Get("/:slug", handler.GetBySlugProductHandler).Name("product.show")
	product.Put("/:slug/update", handler.UpdateProductHandler).Name("product.update")
	product.Delete("/:slug", handler.DeleteProductHandler).Name("product.destroy")
}