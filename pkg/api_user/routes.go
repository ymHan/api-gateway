package apiUser

import (
	"api-gateway/pkg/api_user/routes"
	"api-gateway/pkg/common/config"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, c config.Config) {
	svc := InitServiceClient(&c)

	r := app.Group("/api/v1/user")
	r.Post("/", svc.CreateUser)

}

func (svc *UserServiceClient) CreateUser(ctx *fiber.Ctx) error {
	return routes.CreateUser(ctx, svc.UserCommandClient)
}
