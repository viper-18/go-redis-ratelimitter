package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/viper-18/go-redis-ratelimitter/app"
	r "github.com/viper-18/go-redis-ratelimitter/redis"
)

func main() {
	godotenv.Load()
	r.InitRedisClient(
		os.Getenv("REDIS_ADDR"),
		os.Getenv("REDIS_PASSWORD"),
		0,
	)
	defer r.Client.Close()

	rateLimiter := app.NewRateLimitter(r.Client, 3, time.Minute)

	fiberApp := fiber.New()

	fiberApp.Use(logger.New())

	fiberApp.Use(app.RateLimitterMiddleware(rateLimiter))

	fiberApp.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Welcome!"})
	})

	fiberApp.Get("/endpoint", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "SUCCESS"})
	})

	// Start the server
	log.Fatal(fiberApp.Listen(":8000"))

}
