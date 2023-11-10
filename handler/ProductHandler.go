package handler

import (
	"log"
	"plastindo-back-end/models/entity"
	"plastindo-back-end/models/request"
	"plastindo-back-end/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

func GetAllProductHandler(ctx *fiber.Ctx) error {
	var products []entity.Product

	err := db.Debug().Find(&products).Error

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "FAILED TO GET ALL DATAS",
			"error" : err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "SUCCESS GET ALL DATAS",
		"data" : products,
	})
}

func StoreProductHandler(ctx *fiber.Ctx) error {
	productRequest := new(request.ProductRequest)
	err := ctx.BodyParser(productRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "BAD REQUEST",
			"error" : err.Error(),
		})
	}
	
	err = validate.Struct(productRequest)
	
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "BAD REQUEST",
			"error" : err.Error(),
		})
	}

	product := entity.Product{
		ProductCategoryId: productRequest.ProductCategoryId,
		Title: productRequest.Title,
		Slug: slug.Make(productRequest.Title),
		Material: productRequest.Material,
		Type: productRequest.Type,
		Static: productRequest.Static,
		Dynamic: productRequest.Dynamic,
		Racking: productRequest.Racking,
		TokopediaLink: productRequest.TokopediaLink,
		ShopeeLink: productRequest.ShopeeLink,
		LazadaLink: productRequest.LazadaLink,
	}

	err = db.Debug().Create(&product).Error

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "FAILED TO CREATE DATA",
			"error" : err.Error(),
		})
	}

	// CREATE FILE
	filenames := ctx.Locals("filenames").([]string)
	if len(filenames) == 0 {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message" : "IMAGE PRODUCT IS REQUIRED!",
		})
	}

	for _, filename := range filenames {
		imageGallery := entity.ImageGallery{
			ProductId: product.ID,
			// Name: strconv.Itoa(product.ID) + filename,
			Name: filename,
		}

		err = db.Debug().Create(&imageGallery).Error

		if err != nil {
			log.Println("SOME DATA FAILED TO CREATE")
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "FAILED TO STORE DATA",
				"error" : err.Error(),
			})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "SUCCESS CREATE DATA",
		"data" : product,
	})
}

func GetBySlugProductHandler(ctx *fiber.Ctx) error {
	productSlug := ctx.Params("slug")

	var product entity.Product

	err := db.Debug().Take(&product, "slug = ?", productSlug).Error

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message" : "DATA NOT FOUND",
			"error" : err.Error(),
		})
	}
	
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "SUCCESS GET DATA",
		"data" : product,
	})
}

func UpdateProductHandler(ctx *fiber.Ctx) error {
	productRequest := new(request.ProductRequest)
	err := ctx.BodyParser(productRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "BAD REQUEST",
			"error" : err.Error(),
		})
	}

	// FIND DATA
	productSlug := ctx.Params("slug")

	var product entity.Product

	err = db.Debug().Take(&product, "slug = ?", productSlug).Error

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message" : "DATA NOT FOUND",
			"error" : err.Error(),
		})
	}

	// DELETE IMAGE
	var imageGalleries []entity.ImageGallery

	err = db.Debug().Find(&imageGalleries, "product_id = ?", product.ID).Error
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message" : "DATA NOT FOUND",
			"error" : err.Error(),
		})
	}

	for _, imageGallery := range imageGalleries {
		// DELETE IMAGE FROM DIRECTORY
		err = utils.HandleRemoveFile(imageGallery.Name)
		if err != nil {
			log.Println("FAILED TO DELETE DATA IN DIRECTORY")
		}

		// DELETE IMAGE FROM DATABASE
		err = db.Debug().Delete(&imageGallery).Error
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "FAILED TO DELETE DATA IMAGE",
				"error" : err.Error(),
			})
		}
	}

	// CREATE FILE
	filenames := ctx.Locals("filenames").([]string)
	if len(filenames) == 0 {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message" : "IMAGE PRODUCT IS REQUIRED!",
		})
	}

	for _, filename := range filenames {
		imageGallery := entity.ImageGallery{
			ProductId: product.ID,
			Name: filename,
		}

		err = db.Debug().Updates(&imageGallery).Error

		if err != nil {
			log.Println("SOME DATA FAILED TO CREATE")
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "FAILED TO STORE DATA",
				"error" : err.Error(),
			})
		}
	}

	// UPDATE DATA
	if productRequest.Title != "" {
		product.Title = productRequest.Title
		product.Slug = slug.Make(productRequest.Title)
	}

	if productRequest.ProductCategoryId != 0 {
        product.ProductCategoryId = productRequest.ProductCategoryId
    }

	product.Material = productRequest.Material
	product.Type = productRequest.Type
	product.Static = productRequest.Static
	product.Dynamic = productRequest.Dynamic
	product.Racking = productRequest.Racking
	product.TokopediaLink = productRequest.TokopediaLink
	product.ShopeeLink = productRequest.ShopeeLink
	product.LazadaLink = productRequest.LazadaLink

	err = db.Debug().Save(&product).Error

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "FAILED TO UPDATE DATA",
			"error" : err.Error(),
		})
	}
	
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "SUCCESS UPDATE DATA WITH ID : " + strconv.Itoa(product.ID),
		"data" : product,
	})
}

func DeleteProductHandler(ctx *fiber.Ctx) error {
	productSlug := ctx.Params("slug")

	var product entity.Product

	err := db.Debug().Take(&product, "slug = ?", productSlug).Error

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message" : "DATA NOT FOUND",
			"error" : err.Error(),
		})
	}

	// DELETE IMAGE
	var imageGalleries []entity.ImageGallery

	err = db.Debug().Find(&imageGalleries, "product_id = ?", product.ID).Error
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message" : "DATA NOT FOUND",
			"error" : err.Error(),
		})
	}

	for _, imageGallery := range imageGalleries {
		// DELETE IMAGE FROM DIRECTORY
		err = utils.HandleRemoveFile(imageGallery.Name)
		if err != nil {
			log.Println("FAILED TO DELETE DATA IN DIRECTORY")
		}

		// DELETE IMAGE FROM DATABASE
		err = db.Debug().Delete(&imageGallery).Error
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "FAILED TO DELETE DATA IMAGE",
				"error" : err.Error(),
			})
		}
	}

	err = db.Debug().Delete(&product).Error

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "FAILED TO DELETE DATA",
			"error" : err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "SUCCESS DELETE DATA",
	})
}