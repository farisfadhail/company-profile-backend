package handler

import (
	"plastindo-back-end/database"
	"plastindo-back-end/models/entity"
	"plastindo-back-end/models/request"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

var db = database.DatabaseInit()
var validate = validator.New()

func GetAllParentCategoryHandler(ctx *fiber.Ctx) error {

	var parentCategories []entity.ParentCategory

	err := db.Debug().Preload("ProductCategories").Find(&parentCategories).Error

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "CANNOT GET ALL DATAS",
			"error" : err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(parentCategories)
}

func StoreParentCategoryHandler(ctx *fiber.Ctx) error {
	parentCategoryRequest := new(request.ParentCategoryRequest)
	err := ctx.BodyParser(parentCategoryRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "BAD REQUEST",
			"error" : err.Error(),
		})
	}
	
	err = validate.Struct(parentCategoryRequest)
	
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "BAD REQUEST",
			"error" : err.Error(),
		})
	}

	parentCategory := entity.ParentCategory{
		Name: parentCategoryRequest.Name,
		Slug: slug.Make(parentCategoryRequest.Name),
	}

	err = db.Debug().Create(&parentCategory).Error

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "FAILED TO STORE DATA",
			"error" : err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "SUCCESS CREATE PARENT CATEGORY",
		"data" : parentCategory,
	})
}

func GetBySlugParentCategoryHandler(ctx *fiber.Ctx) error {
	parentSlug := ctx.Params("slug")

	var parentCategory entity.ParentCategory

	err := db.Debug().Take(&parentCategory, "slug = ?", parentSlug).Error

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message" : "DATA NOT FOUND",
			"data" : err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "SUCCESS GET DATA",
		"data" : parentCategory,
	})
}

func UpdateParentCategoryHandler(ctx *fiber.Ctx) error {
	parentCategoryRequest := new(request.ParentCategoryRequest)
	err := ctx.BodyParser(parentCategoryRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "BAD REQUEST",
			"error" : err.Error(),
		})
	}

	// FIND DATA
	parentSlug := ctx.Params("slug")

	var parentCategory entity.ParentCategory

	err = db.Debug().Take(&parentCategory, "slug = ?", parentSlug).Error

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message" : "DATA NOT FOUND",
			"error" : err.Error(),
		})
	}

	// UPDATE DATA
	if parentCategoryRequest.Name != "" {
		parentCategoryRequest.Slug = slug.Make(parentCategoryRequest.Name)
	}

	err = db.Debug().Model(&parentCategory).Updates(parentCategoryRequest).Error

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "FAILED TO UPDATE DATA",
			"error" : err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "SUCCESS UPDATE DATA WITH ID : " + strconv.Itoa(parentCategory.ID),
		"data" : parentCategory,
	})
}

func DeleteParentCategoryHandler(ctx *fiber.Ctx) error {
	// FIND DATA
	parentSlug := ctx.Params("slug")

	var parentCategory entity.ParentCategory

	err := db.Debug().Take(&parentCategory, "slug = ?", parentSlug).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message" : "DATA NOT FOUND",
			"error" : err.Error(),
		})
	}

	err = db.Debug().Delete(&parentCategory).Error

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "FAILED TO DELETE DATA",
			"error" : err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "SUCCESS DELETED DATA",
	})
}