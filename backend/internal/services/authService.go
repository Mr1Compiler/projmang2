package services

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/mustafaameen91/project-managment/backend/internal/auth"
	"github.com/mustafaameen91/project-managment/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid username or password")
)

type AuthService struct {
	userRepo     *repository.UserRepository
	userRoleRepo *repository.UserRoleRepository
	jwtManager   *auth.JWTManager
}

func NewAuthService(userRepo *repository.UserRepository, userRoleRepo *repository.UserRoleRepository, jwtManager *auth.JWTManager) *AuthService {
	return &AuthService{
		userRepo:     userRepo,
		userRoleRepo: userRoleRepo,
		jwtManager:   jwtManager,
	}
}

type LoginResponse struct {
	Token    string `json:"token"`
	UserID   int64  `json:"userId"`
	Username string `json:"username"`
}

func (s *AuthService) Login(ctx context.Context, username, password string) (*LoginResponse, error) {
	creds, err := s.userRepo.GetCredentialsByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(creds.Password), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	token, err := s.jwtManager.GenerateToken(creds.ID, creds.Username)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token:    token,
		UserID:   creds.ID,
		Username: creds.Username,
	}, nil
}

// GetUserPages returns all pages a user has access to with their permissions
func (s *AuthService) GetUserPages(ctx context.Context, userID int64) ([]repository.UserPage, error) {
	return s.userRoleRepo.GetUserPages(ctx, userID)
}
