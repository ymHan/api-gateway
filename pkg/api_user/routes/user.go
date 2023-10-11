package routes

import (
	pb "api-gateway/pkg/api_user/pb"
	"context"
	"github.com/gofiber/fiber/v2"
)

type CreateUserRequestBody struct {
	Email    string `json:"email"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type UpdateUserRequestBody struct {
	Email    string `json:"email"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type UpdateUserPasswordRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(ctx *fiber.Ctx, c pb.UserCommandServiceClient) error {
	body := CreateUserRequestBody{}

	if err := ctx.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := c.CreateUser(context.Background(), &pb.CreateUserRequest{
		Email:    body.Email,
		Username: body.UserName,
		Password: body.Password,
	})

	if err != nil {
		return err
	}

	return ctx.Status(int(res.Status)).JSON(&res)
}

func UpdateUser(ctx *fiber.Ctx, c pb.UserCommandServiceClient) error {
	body := UpdateUserRequestBody{}

	if err := ctx.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := c.UpdateUser(context.Background(), &pb.UpdateUserRequest{
		Email:    body.Email,
		Username: body.UserName,
		Password: body.Password,
	})

	if err != nil {
		return err
	}

	return ctx.Status(int(res.Status)).JSON(&res)
}
