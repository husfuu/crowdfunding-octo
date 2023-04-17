package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/husfuu/crowdfunding-octo/app/models"
)

func GetCampaigns(c *fiber.Ctx) error {

	campains := 1

	return c.JSON(fiber.Map{
		"error":     false,
		"msg":       nil,
		"campaigns": campains,
	})
}

func GetCampaign(c *fiber.Ctx) error {
	campains := 1

	return c.JSON(fiber.Map{
		"error":     false,
		"msg":       nil,
		"campaigns": campains,
	})
}

func CreateCampaign(c *fiber.Ctx) error {
	// now := time.Now().Unix()

	// create a new campaign struct
	campaign := &models.Campaign{}

	if err := c.BodyParser(campaign); err != nil {
		// return 400 and error message
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	campains := 1

	return c.JSON(fiber.Map{
		"error":     false,
		"msg":       nil,
		"campaigns": campains,
	})
}

func UpdateCampaign(c *fiber.Ctx) error {
	campaign := &models.Campaign{}

	if err := c.BodyParser(campaign); err != nil {
		// return 400 and error message
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":     false,
		"msg":       nil,
		"campaigns": campaign,
	})
}
