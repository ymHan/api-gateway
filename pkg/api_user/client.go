package apiUser

import (
	pb "api-gateway/pkg/api_user/pb"
	"api-gateway/pkg/common/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceClient struct {
	UserQueryClient   pb.UserQueryServiceClient
	UserCommandClient pb.UserCommandServiceClient
}

func InitServiceClient(c *config.Config) *UserServiceClient {
	queryConnection, err := grpc.Dial(c.UserQSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	cmdConnection, err := grpc.Dial(c.UserCSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Cound not connect:", err)
	}

	return &UserServiceClient{
		UserQueryClient:   pb.NewUserQueryServiceClient(queryConnection),
		UserCommandClient: pb.NewUserCommandServiceClient(cmdConnection),
	}
}
