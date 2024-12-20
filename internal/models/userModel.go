package models

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type User struct {
	ID             string         `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Email          string         `json:"email" gorm:"uniqueIndex;not null" validate:"required,email"`
	Username       string         `json:"username" gorm:"unique"`
	PasswordHash   string         `json:"-"` //never return in JSON
	Name           string         `json:"name"`
	ProfilePicture string         `json:"profile_picture"`
	IsActive       bool           `json:"is_active" gorm:"default:true"`
	EmailVerified  bool           `json:"email_verified" gorm:"default:false"`
	Role           string         `json:"role" gorm:"default:'user'" validate:"required,oneof=admin user"`
	LastLogin      time.Time      `json:"last_login"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Validate performs validation on User fields
func (u *User) Validate() error {
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}
	if u.Username == "" {
		return fmt.Errorf("username is required")
	}
	if u.Role != "user" && u.Role != "admin" {
		return fmt.Errorf("invalid role")
	}
	return nil
}

// OAuthAccount represents a user's OAuth connection
type OAuthAccount struct {
	ID           string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	User         User      `json:"user" gorm:"foreignKey:UserID"`
	Provider     string    `json:"provider"` // e.g., "google", "github"
	ProviderID   string    `json:"provider_id" gorm:"uniqueIndex"`
	AccessToken  string    `json:"-"` // Encrypted in DB
	TokenType    string    `json:"-"`
	ExpiresAt    time.Time `json:"-"`
	RefreshToken string    `json:"-"` // Encrypted in DB
}

// JWTClaims extends standard JWT claims with custom user information
type JWTClaims struct {
	jwt.RegisteredClaims
	UserID   string   `json:"user_id"`
	Email    string   `json:"email"`
	Username string   `json:"username"`
	Role     string   `json:"role"`
	Scopes   []string `json:"scopes"`
}

// OAuthToken extends the oauth2 token with additional metadata
type OAuthToken struct {
	oauth2.Token
	DeviceID   string    `json:"device_id,omitempty"`
	LastUsedAt time.Time `json:"last_used_at"`
	IsPrimary  bool      `json:"is_primary"`
}

// RefreshToken represents a refresh token for maintaining session
type RefreshToken struct {
	ID          string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID      string    `json:"user_id"`
	User        User      `json:"user" gorm:"foreignKey:UserID"`
	Token       string    `json:"token" gorm:"unique"`
	ExpiresAt   time.Time `json:"expires_at"`
	IsRevoked   bool      `json:"is_revoked" gorm:"default:false"`
	RotatedFrom string    `json:"rotated_from,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

// for security purposes
type LoginAttempt struct {
	ID         string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID     string    `json:"user_id"`
	User       User      `json:"user" gorm:"foreignKey:UserID"`
	IPAddress  string    `json:"ip_address"`
	Successful bool      `json:"successful"`
	Timestamp  time.Time `json:"timestamp"`
}

// for OAuth providers
type OAuthConfig struct {
	ClientID     string          `json:"client_id"`
	ClientSecret string          `json:"-"` // Never expose
	RedirectURL  string          `json:"redirect_url"`
	Scopes       []string        `json:"scopes"`
	Endpoint     oauth2.Endpoint `json:"-"`
}

// just in case we need it in the future
// type UserSettings struct {
// 	ID                   string `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
// 	UserID               string `json:"user_id"`
// 	User                 User   `json:"user" gorm:"foreignKey:UserID"`
// 	NotificationsEnabled bool   `json:"notifications_enabled" gorm:"default:true"`
// 	Theme                string `json:"theme" gorm:"default:'light'"`
// 	Language             string `json:"language" gorm:"default:'en'"`
// }

// AuditLog would implemetn in future
// type AuditLog struct {
// 	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
// 	UserID    string    `json:"user_id"`
// 	User      User      `json:"user" gorm:"foreignKey:UserID"`
// 	Action    string    `json:"action"`
// 	IPAddress string    `json:"ip_address"`
// 	Timestamp time.Time `json:"timestamp"`
// }
