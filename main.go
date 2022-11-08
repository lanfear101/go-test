package main

import (
	"github.com/gofiber/fiber/v2"
	"myapp/internal/dto"
	"myapp/pkg/image"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/generate_image", generateImage)
	app.Listen(":3000")
}

const (
	// Default setting
	imgColorDefault = "E5E5E5"
	msgColorDefault = "AAAAAA"
	imgWDefault     = 300
	imgHDefault     = 300
	fontSizeDefault = 0

	dpiDefault float64 = 72

	fontfileDefault = "assets/fonts/CeraProMedium.otf"
)

type ImageQueryParam struct {
	Width  int `query:"width"`
	Height int `query:"height"`
}

func generateImage(c *fiber.Ctx) error {
	imageQueryDto := dto.GetImageQueryDto()
	if err := c.QueryParser(imageQueryDto); err != nil {
		return err
	}
	imageQueryDto.FillDefaults()
	buf, _ := image.Do(imageQueryDto)
	c.Set("content-type", "image/png;base64")
	return c.SendString(buf.String())
}
