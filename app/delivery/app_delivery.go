package delivery

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type appDelivery struct{}

func NewAppDelivery(router fiber.Router) {
	handler := &appDelivery{}

	router.Get("/get-time", handler.GetTime)
}

func (a *appDelivery) GetTime(ctx *fiber.Ctx) error {
	now := time.Now().UTC().Unix()

	return ctx.JSON(now)
}
