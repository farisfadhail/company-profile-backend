package handler

import (
	"plastindo-back-end/models/entity"
	"plastindo-back-end/models/request"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

func GetAllProductCategoryHandler(ctx *fiber.Ctx) error {
	var productCategories []entity.ProductCategory

	err := db.Debug().Find(&productCategories).Error

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "FAILED TO GET ALL DATAS",
			"error" : err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message" : "SUCCESS GET ALL DATAS",
		"data" : productCategories,
	})
}

func StoreProductCategoryHandler(ctx *fiber.Ctx) error {
	productCategoryRequest := new(request.ProductCategoryRequest)
	err := ctx.BodyParser(productCategoryRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "BAD REQUEST",
			"error" : err.Error(),
		})
	}
	
	err = validate.Struct(productCategoryRequest)
	
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "BAD REQUEST",
			"error" : err.Error(),
		})
	}
	
	productCategory := entity.ProductCategory{
		ParentCategoryId: productCategoryRequest.ParentCategoryId,
		Name: productCategoryRequest.Name,
		Slug: slug.Make(productCategoryRequest.Name),
	}
	
	err = db.Debug().Create(&productCategory).Error
	
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "INTERNAL SERVER ERROR",
			"error" : err.Error(),
		})
	}
	
	return ctx.JSON(fiber.Map{
		"message" : "SUCCESS CREATE DATA",
		"data" : productCategory,
	})
}

func GetBySlugProductCategoryHandler(ctx *fiber.Ctx) error {
	productCategorySlug := ctx.Params("slug")

	var productCategory entity.ProductCategory

	err := db.Debug().Take(&productCategory, "slug = ?", productCategorySlug).Error

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message" : "DATA NOT FOUND",
			"error" : err.Error(),
		})
	}
	
	return ctx.JSON(fiber.Map{
		"message" : "SUCCESS GET DATA",
		"data" : productCategory,
	})
}

func UpdateProductCategoryHandler(ctx *fiber.Ctx) error {
	productCategoryRequest := new(request.ProductCategoryRequest)
	err := ctx.BodyParser(productCategoryRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "BAD REQUEST",
			"error" : err.Error(),
		})
	}

	// FIND DATA
	productCategorySlug := ctx.Params("slug")

	var productCategory entity.ProductCategory

	err = db.Debug().Take(&productCategory, "slug = ?", productCategorySlug).Error

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message" : "DATA NOT FOUND",
			"error" : err.Error(),
		})
	}

	// UPDATE DATA
	if productCategoryRequest.Name != "" {
		productCategoryRequest.Slug = slug.Make(productCategoryRequest.Name)
	}

	err = db.Debug().Model(&productCategory).Updates(productCategoryRequest).Error

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "FAILED TO UPDATE DATA",
			"error" : err.Error(),
		})
	}
	
	return ctx.JSON(fiber.Map{
		"message" : "SUCCESS UPDATE DATA WITH ID : " + strconv.Itoa(productCategory.ID),
		"data" : productCategory,
	})
}

func DeleteProductCategoryHandler(ctx *fiber.Ctx) error {
	// FIND DATA
	productCategorySlug := ctx.Params("slug")

	var productCategory entity.ProductCategory

	err := db.Debug().Take(&productCategory, "slug = ?", productCategorySlug).Error

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message" : "DATA NOT FOUND",
			"error" : err.Error(),
		})
	}

	err = db.Debug().Delete(&productCategory).Error
	
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "FAILED TO DELETE DATA",
			"error" : err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message" : "SUCCESS DELETE DATA",
	})
}