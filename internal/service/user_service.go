package sevice

import (
	"context"
	"go_api/internal/config"
	"go_api/internal/dto"
	"go_api/internal/repository"

	"github.com/google/uuid"
)

type UserService interface {
	register(ctx context.Context, req *dto.RegisterRequest) (uuid.UUID, int, error)
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
