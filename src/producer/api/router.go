package api

import (
	"producer/shared"

	"github.com/gofiber/fiber/v2"
)

func ping(c *fiber.Ctx) error {
	res := shared.ContainerItem.Service.Ping()
	return c.SendString(string(res))
}

func sendReceipts(c *fiber.Ctx) error {
	shared.ContainerItem.Service.Kafka()
	return c.JSON("{'ok': true}")
}
