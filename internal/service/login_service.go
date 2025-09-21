package service

import (
	"context"
	"errors"
	"go_api/internal/dto"
	"go_api/internal/model"
	"go_api/pkg/jwt"
	refreshtoken "go_api/pkg/refreshToken"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Login(ctx context.Context, req *dto.LoginRequest) (string, string, int, error) {
	userExists, err := s.userRepo.GetUserByEmailOrUsername(ctx, req.Email, "")
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if userExists == nil {
		return "", "", http.StatusNotFound, errors.New("wrong email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userExists.Hashed_password), []byte(req.Password))
	if err != nil {
		return "", "", http.StatusNotFound, errors.New("wrong email or password")
	}

	token, err := jwt.CreateToken(userExists.ID, userExists.Username, s.cfg.SecretJWT)
	if err != nil {
		return "", "", http.StatusBadRequest, err
	}
	refreshTokenExists, err := s.userRepo.GetRefreshToken(ctx, userExists.ID, time.Now())
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}
	if refreshTokenExists != nil {
		return token, refreshTokenExists.RefreshToken, http.StatusOK, nil
	}
	refreshToken, err := refreshtoken.GenerateRefreshToken()
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}
	err = s.userRepo.CreateRefreshToken(ctx, &model.RefreshTokenModel{
		ID:           uuid.New(),
		UserID:       userExists.ID,
		RefreshToken: refreshToken,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		ExpiredAt:    time.Now().Add(7 * 24 * time.Hour),
	})
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}
	return token, refreshToken, http.StatusOK, nil
}
