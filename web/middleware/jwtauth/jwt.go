package jwtauth

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// JWTAuth jwt auth
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  http.StatusText(http.StatusUnauthorized),
			})
			c.Abort()
			return
		}
		if s := strings.Split(token, " "); len(s) == 2 {
			token = s[1]
		}
		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == ErrTokenExpired {
				if token, err = j.RefreshToken(token); err == nil {
					c.Header("Authorization", "Bearer "+token)
					c.JSON(http.StatusOK, gin.H{"error": 0, "message": "refresh token", "token": token})
					c.Abort()
					return
				}
			}
			c.JSON(http.StatusUnauthorized, gin.H{"error": 1, "message": err.Error()})
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

// JWT signingKey
type JWT struct {
	SigningKey []byte
}

var (
	// ErrTokenExpired Token is expired information
	ErrTokenExpired error = errors.New("Token is expired")
	// ErrTokenNotValidYet Token not active yet information
	ErrTokenNotValidYet error = errors.New("Token not active yet")
	// ErrTokenMalformed not even a token"
	ErrTokenMalformed error = errors.New("That's not even a token")
	// ErrTokenInvalid not handle this token
	ErrTokenInvalid error = errors.New("Couldn't handle this token")

	// SignKey sign key
	SignKey = "raytheon"
)

// CustomClaims fields
type CustomClaims struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Tenant string `json:"tenant"`
	jwt.StandardClaims
}

// NewJWT new jwt
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// GetSignKey get sign key
func GetSignKey() string {
	return SignKey
}

// SetSignKey set sign key
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateToken create token information
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken parse token information
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

// RefreshToken refresh token information
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", ErrTokenInvalid
}
