package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lizenshakya/url-shortner-go/helpers"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL            string        `json:"url"`
	CustomShort    string        `json:"short"`
	Expiry         time.Duration `json:"expiry"`
	XRateRemaining int           `json:"rate_limit"`
	XRateLimitRest time.Duration `json:"rate_limit_rest"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	//implement rate limiter

	//check if input is actual url
	if !govalidator.IsUrl(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Not a valid URL"})
	}

	//check for domain error
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "Cannot access URL"})
	}

	//enfore https,SSL
	body.URL = helpers.EnforeHttp(body.URL)
}
