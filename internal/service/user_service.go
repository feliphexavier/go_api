package service

import (
	"context"
	"go_api/internal/config"
	"go_api/internal/dto"
	"go_api/internal/repository"

	"github.com/google/uuid"
)

type UserService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (uuid.UUID, int, error)
	Login(ctx context.Context, req *dto.LoginRequest) (string, string, int, error)
	RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest, userId uuid.UUID) (string, string, int, error)
}
type userService struct {
	cfg      *config.Config
	userRepo repository.UserRepository
}

func NewService(cfg *config.Config, userRepo repository.UserRepository) UserService {
	return &userService{
		cfg:      cfg,
		userRepo: userRepo,
	}
}
