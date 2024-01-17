package middlewares

import (
	"os"

	"github.com/AlvinMrema/kighala-api/pkg/utils/jwt_utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// tokenValue := c.GetReqHeaders()["Authorization"][0]
		tokenValue := c.Get("X-Auth-Token") // c.Get("Authorization")
		claims := &utils.Claims{}

		tkn, err := jwt.ParseWithClaims(tokenValue, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
		if  !tkn.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   "Unauthorised Request",
			})
		}
		return c.Next()
	}
}
