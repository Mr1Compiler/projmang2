package services

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafa91ameen/prjalgo/backend/internal/auth"
	"github.com/mustafa91ameen/prjalgo/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials  = errors.New("invalid username or password")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
	ErrExpiredRefreshToken = errors.New("refresh token has expired")
	ErrRevokedRefreshToken = errors.New("refresh token has been revoked")
	ErrUserInactive        = errors.New("user account is inactive")
)

type AuthService struct {
	db               *pgxpool.Pool
	userRepo         *repository.UserRepository
	userRoleRepo     *repository.UserRoleRepository
	refreshTokenRepo *repository.RefreshTokenRepository
	jwtManager       *auth.JWTManager
	refreshExpiry    time.Duration
}

func NewAuthService(
	db *pgxpool.Pool,
	userRepo *repository.UserRepository,
	userRoleRepo *repository.UserRoleRepository,
	refreshTokenRepo *repository.RefreshTokenRepository,
	jwtManager *auth.JWTManager,
	refreshExpiry time.Duration,
) *AuthService {
	return &AuthService{
		db:               db,
		userRepo:         userRepo,
		userRoleRepo:     userRoleRepo,
		refreshTokenRepo: refreshTokenRepo,
		jwtManager:       jwtManager,
		refreshExpiry:    refreshExpiry,
	}
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	UserID       int64  `json:"userId"`
	Username     string `json:"username"`
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

	// Check if user is active
	if creds.Status != nil && *creds.Status == "inactive" {
		return nil, ErrUserInactive
	}

	// Update last login time - non-critical, log error but don't fail login
	if err := s.userRepo.UpdateLastLogin(ctx, creds.ID); err != nil {
		log.Printf("WARN: failed to update last login for user %d: %v", creds.ID, err)
	}

	accessToken, err := s.jwtManager.GenerateToken(creds.ID, creds.Username)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.generateRefreshToken()
	if err != nil {
		return nil, err
	}

	tokenHash := s.hashToken(refreshToken)
	expiresAt := time.Now().UTC().Add(s.refreshExpiry)

	_, err = s.refreshTokenRepo.Create(ctx, creds.ID, tokenHash, expiresAt)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       creds.ID,
		Username:     creds.Username,
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*LoginResponse, error) {
	tokenHash := s.hashToken(refreshToken)

	storedToken, err := s.refreshTokenRepo.GetByTokenHash(ctx, tokenHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrInvalidRefreshToken
		}
		return nil, err
	}

	if storedToken.Revoked {
		return nil, ErrRevokedRefreshToken
	}

	now := time.Now().UTC()
	log.Printf("DEBUG RefreshToken: now=%v, expiresAt=%v, isExpired=%v", now, storedToken.ExpiresAt, now.After(storedToken.ExpiresAt))
	if now.After(storedToken.ExpiresAt) {
		log.Printf("DEBUG RefreshToken: Token EXPIRED - rejecting")
		return nil, ErrExpiredRefreshToken
	}

	user, err := s.userRepo.GetByID(ctx, storedToken.UserID)
	if err != nil {
		return nil, err
	}

	accessToken, err := s.jwtManager.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := s.generateRefreshToken()
	if err != nil {
		return nil, err
	}

	newTokenHash := s.hashToken(newRefreshToken)
	newExpiresAt := time.Now().UTC().Add(s.refreshExpiry)

	// Start transaction for atomic token update
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// Update token hash within transaction
	err = s.refreshTokenRepo.UpdateTokenHashWithTx(ctx, tx, storedToken.ID, newTokenHash, newExpiresAt)
	if err != nil {
		return nil, err
	}

	// Commit transaction
	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		UserID:       user.ID,
		Username:     user.Username,
	}, nil
}

// LogoutResult contains information about the logged out user
type LogoutResult struct {
	UserID int64 `json:"userId"`
}

func (s *AuthService) Logout(ctx context.Context, refreshToken string) (*LogoutResult, error) {
	tokenHash := s.hashToken(refreshToken)

	storedToken, err := s.refreshTokenRepo.GetByTokenHash(ctx, tokenHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	err = s.refreshTokenRepo.Revoke(ctx, storedToken.ID)
	if err != nil {
		return nil, err
	}

	return &LogoutResult{UserID: storedToken.UserID}, nil
}

func (s *AuthService) generateRefreshToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (s *AuthService) hashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

// GetUserPages returns all pages a user has access to with their permissions
func (s *AuthService) GetUserPages(ctx context.Context, userID int64) ([]repository.UserPage, error) {
	return s.userRoleRepo.GetUserPages(ctx, userID)
}
