package handler

import (
	"plastindo-back-end/models/entity"
	"plastindo-back-end/models/request"
	"plastindo-back-end/utils"

	"github.com/gofiber/fiber/v2"
)

func GetAllUserHandler(ctx *fiber.Ctx) error {
	var users []entity.User

	err := db.Debug().Find(&users).Error

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "FAILED TO GET ALL DATAS",
			"error" : err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "SUCCESS GET ALL DATAS",
		"data" : users,
	})
}

func GetByIdUserHandler(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User

	err := db.Debug().Take(&user, userId).Error

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message" : "USER NOT FOUND",
			"error" : err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "SUCCESS GET DATA",
		"data" : user,
	})
}

func UpdateUserHandler(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)
	err := ctx.BodyParser(userRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "BAD REQUEST",
			"error" : err.Error(),
		})
	}

	// CHECK EMAIL
	if userRequest.Email != "" {
		err = validate.Struct(userRequest)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message" : "BAD REQUEST",
				"error" : err.Error(),
			})
		}
	}

	// FIND USER
	userId := ctx.Params("id")

	var user entity.User

	err = db.Debug().Take(&user, userId).Error

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "FAILED TO GET DATA",
			"error" : err.Error(),
		})
	}

	// UPDATE USER
	err = db.Debug().Model(&user).Updates(userRequest).Error

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "FAILED TO UPDATE USER DATA",
			"error" : err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "SUCCESS UPDATE USER DATA",
		"data" : user,
	})
}

func DeleteUserHandler(ctx *fiber.Ctx) error {
	var user entity.User

	err := db.Debug().Take(&user).Error

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message" : "USER NOT FOUND",
			"error" : err.Error(),
		})
	}

	err = db.Debug().Delete(&user).Error

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

func UpdatePasswordUserHandler(ctx *fiber.Ctx) error {
	passwordRequest := new(request.PasswordUpdateRequest)
	err := ctx.BodyParser(passwordRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "BAD REQUEST",
			"error" : err.Error(),
		})
	}

	err = validate.Struct(passwordRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "BAD REQUEST",
			"error" : err.Error(),
		})
	}

	if passwordRequest.Password != passwordRequest.ConfirmPassword {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "CONFIRM PASSWORD MUST BE SAME AS PASSWORD",
		})
	}
	
	// FIND USER
	userId := ctx.Params("id")

	var user entity.User

	err = db.Debug().Take(&user, userId).Error

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message" : "USER NOT FOUND",
			"error" : err.Error(),
		})
	}

	// UDPATE PASSWORD
	password, err := utils.HashingPassword(passwordRequest.Password)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "INTERNAL SERVER ERROR",
			"error" : err.Error(),
		})
	}

	err = db.Debug().Model(&user).Update("password", password).Error

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "FAILED TO UPDATE PASSWORD",
			"error" : err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "SUCCESS UPDATE PASSWORD",
	})
}