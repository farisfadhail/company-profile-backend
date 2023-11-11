package handler

import (
	"log"
	"plastindo-back-end/models/entity"
	"plastindo-back-end/models/request"
	"plastindo-back-end/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func SignUpHandler(ctx *fiber.Ctx) error {
	userRequest := new(request.UserCreateRequest)
	err := ctx.BodyParser(userRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "BAD REQUEST",
			"error":   err.Error(),
		})
	}

	err = validate.Struct(userRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "BAD REQUEST",
			"error":   err.Error(),
		})
	}

	if userRequest.Password != userRequest.ConfirmPassword {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "CONFIRM PASSWORD MUST BE SAME AS PASSWORD",
		})
	}

	password, err := utils.HashingPassword(userRequest.Password)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "INTERNAL SERVER ERROR",
			"error":   err.Error(),
		})
	}

	user := entity.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: password,
	}

	err = db.Debug().Create(&user).Error

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "FAILED TO SIGN UP USER",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "SUCCESS SIGN UP USER",
		"data":    user,
	})
}

func SignInHandler(ctx *fiber.Ctx) error {
	signInRequest := new(request.SignInRequest)
	err := ctx.BodyParser(signInRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "BAD REQUEST",
			"error" : err.Error(),
		})
	}

	err = validate.Struct(signInRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "BAD REQUEST",
			"error" : err.Error(),
		})
	}

	// CHECK AVAIL USER
	var user entity.User

	err = db.Debug().Take(&user, "email = ?", signInRequest.Email).Error

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message" : "WRONG CREDENTIAL",
		})
	}

	// CHECK VALIDATION PASSWORD
	isValid := utils.CheckHashPassword(signInRequest.Password, user.Password)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message" : "WRONG CREDENTIAL",
		})
	}

	// GENERATE JWT
	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(5 * time.Minute).Unix()

	token, err := utils.GenerateJWT(&claims)
	if err != nil {
		log.Println(err.Error())
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message" : "WRONG CREDENTIAL",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token" : token,
	})
}
