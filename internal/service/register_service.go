package sevice

import (
	"context"
	"errors"
	"go_api/internal/dto"
	"go_api/internal/model"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Register(ctx context.Context, req *dto.RegisterRequest) (uuid.UUID, int, error) {
	userExists, err := s.userRepo.GetUserByEmailOrUsername(ctx, req.Email, req.Username)
	if err != nil {
		return uuid.Nil, http.StatusInternalServerError, err
	}
	if userExists != nil {
		return uuid.Nil, http.StatusBadRequest, errors.New("user already exists")
	}
	passHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, http.StatusInternalServerError, err
	}
	userModel := &model.UserModel{
		Email:           req.Email,
		Username:        req.Username,
		Hashed_password: string(passHash),
	}
}
