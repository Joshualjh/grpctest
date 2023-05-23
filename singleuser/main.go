// singleuser/main.go
package main

import (
	"context"

	"github.com/Joshualjh/gotest/data"
	userpb "github.com/Joshualjh/protos/v1/user"
)

const portNumber = "9000"

type userServer struct {
	userpb.UserServer
}

func (s *userServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserReponse, error) {
	userID := req.UserId

	var userMessage *userpb.UserMessage
	for _, u := range data.Users {
		if u.UserId != userID {
			continue
		}
		userMessage = u
		break
	}

	return &userpb.GetUserReponse{
		UserMessage: userMessage,
	}, nil
}
