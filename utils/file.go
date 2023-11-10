package utils

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

const defaultPath = "./public/assets/"

func HandleMultipleFile(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()

	if err != nil {
		log.Println("Error Read Multiple Form Request. Error :", err.Error())
	}

	files := form.File["image_galleries"]
	var filenames []string
	for idx, file := range files {
		if file != nil {
			err := CheckContentType(file, "image/jpg", "image/png", "image/jpeg")
			if err != nil {
				return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"message" : err.Error(),
				})
			}

			extFile := filepath.Ext(file.Filename)
			filename := fmt.Sprintf("%d-%d%s", idx, time.Now().UnixMilli(), extFile)

			err = ctx.SaveFile(file, defaultPath + filename)
			if err != nil {
				log.Println("Failed to store image. Error :", err.Error())
			}

			filenames = append(filenames, filename)
		} else {
			log.Println("Nothing to upload")
		}
	}

	ctx.Locals("filenames", filenames)

	return ctx.Next()
}

func HandleRemoveFile(filename string, path ...string) error {
	if len(path) > 0 {
		err := os.Remove(path[0] + filename)

		if err != nil {
			log.Println("FAILED TO REMOVE FILE")
			return err
		} 
	} else {
		err := os.Remove(defaultPath + filename)

		if err != nil {
			log.Println("FAILED TO REMOVE FILE")
			return err
		} 
	}
	
	return nil
}

func CheckContentType(file *multipart.FileHeader, contentTypes ...string) error {
	if len(contentTypes) > 0 {
		for _, contentType := range contentTypes {
			typeFile := file.Header.Get("content-type")
			if typeFile == contentType {
				return nil
			}
		}

		return errors.New("only allowed png/jpg/jpeg file")
	} else {
		return errors.New("file not found to checking")
	}
}