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
)

func (s *userService) RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest, userId uuid.UUID) (string, string, int, error) {
	userExists, err := s.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return "", "", http.StatusBadRequest, err
	}
	if userExists == nil {
		return "", "", http.StatusNotFound, errors.New("user not found")
	}
	refreshTokenExists, err := s.userRepo.GetRefreshToken(ctx, userId, time.Now())
	if err != nil {
		return "", "", http.StatusBadRequest, err
	}
	if refreshTokenExists == nil {
		return "", "", http.StatusUnauthorized, errors.New("refresh token was expired")
	}
	if req.RefreshToken != refreshTokenExists.RefreshToken {
		return "", "", http.StatusUnauthorized, errors.New("refresh token not found")
	}
	token, err := jwt.CreateToken(userId, userExists.Username, s.cfg.SecretJWT)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}
	_, err = s.userRepo.DeleteRefreshTokenByUserId(ctx, userId)
	if err != nil {
		return "", "", http.StatusBadRequest, err
	}
	refreshToken, err := refreshtoken.GenerateRefreshToken()
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}
	s.userRepo.CreateRefreshToken(ctx, &model.RefreshTokenModel{
		ID:           uuid.New(),
		UserID:       userId,
		RefreshToken: refreshToken,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		ExpiredAt:    time.Now().Add(7 * 24 * time.Hour),
	})

	return token, refreshToken, http.StatusOK, nil
}
