package auth

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strings"
)

func HeaderAuthorization() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if len(strings.Split(auth, " ")) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"msg": "Unauthorized",
			})
		}
		accessToken := strings.Split(auth, " ")[1]

		claims, err := VerifyToken(accessToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"msg": "Unauthorized",
			})
		}

		userID, ok := claims["userID"]
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"msg": "Unauthorized",
			})
		}

		email, ok := claims["email"]
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"msg": "Unauthorized",
			})
		}

		c.Locals("userID", userID)
		c.Locals("email", email)
		c.Locals("access_token", accessToken)
		log.Println("userID: ", c.Locals("userID"))

		return c.Next()
	}
}
