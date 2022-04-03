package proto

import (
	"context"
	"food-delivery/internal/repository/proto"
	"food-delivery/pkg/userservice"
)

type UserService struct {
	repo *proto.Repository
}

func NewUserService(repo *proto.Repository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *userservice.User) (*userservice.UserState, error) {
	return s.repo.CreateUser(ctx, user)
}
