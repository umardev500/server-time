package injector

import (
	"server-time/app/delivery"

	"github.com/gofiber/fiber/v2"
)

func NewAppInjector(router fiber.Router) {
	delivery.NewAppDelivery(router)
}
