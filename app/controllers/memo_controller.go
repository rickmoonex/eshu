package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// NewMemo handles the uploading of a new memo.
func NewMemo(c *fiber.Ctx) error {
	file, err := c.FormFile("memo")
	if err != nil {
		return err
	}

	dest := fmt.Sprintf("./uploads/%s", file.Filename)
	if err := c.SaveFile(file, dest); err != nil {
		return err
	}

	return nil
}
