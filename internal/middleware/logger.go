package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CustomLogger() fiber.Handler{
	return func(c *fiber.Ctx) error{
		start := time.Now()
		err := c.Next()
		stop := time.Now()

		log.Printf("[%s] %s %s %d %v", 
			c.IP(),
			c.Method(),
			c.Path(),
			c.Response().StatusCode(),
			stop.Sub(start),
		)

		return err
	}
}