package jwtauth

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type CustomClaimsTest struct {
	*CustomClaims
	siginKey string
	wanted   string
}

type ExpiredClaimsTest struct {
	CustomClaims
	siginKey string
}

var claims = []CustomClaimsTest{
	{
		CustomClaims: &CustomClaims{
			ID:    "1",
			Name:  "awh521",
			Email: "1044levellevel176017@qq.com",
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
				Issuer:    "test",
			},
		},
		siginKey: "test",
		wanted:   "",
	},
}
var expiredClaims = []ExpiredClaimsTest{
	{
		CustomClaims: CustomClaims{
			ID:    "1",
			Name:  "awh521",
			Email: "1044176017@qq.com",
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: 1500,
				Issuer:    "test",
			},
		},
		siginKey: "test",
	},
}

var jt = &JWT{
	SigningKey: []byte("test"),
}

var foreverClaims = CustomClaims{
	ID:    "1000",
	Name:  "defaul",
	Email: "default@qq.com",
	StandardClaims: jwt.StandardClaims{
		ExpiresAt: 0,
		Issuer:    "default",
	},
}

func TestCreateForeverTokens(t *testing.T) {
	token, err := jt.CreateToken(foreverClaims)
	assert.NoError(t, err)
	claims, err := jt.ParseToken(token)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), claims.StandardClaims.ExpiresAt)
}

func TestJWTCreateToken(t *testing.T) {
	for _, c := range claims {
		j := &JWT{SigningKey: []byte(c.siginKey)}
		token, err := j.CreateToken(*c.CustomClaims)
		assert.NoError(t, err)
		assert.IsType(t, "string", token)
	}
}

func TestJWTParseToken(t *testing.T) {
	for _, c := range claims {
		j := &JWT{SigningKey: []byte(c.siginKey)}
		var err error
		c.wanted, err = j.CreateToken(*c.CustomClaims)
		if err != nil {
			t.Error(err)
		}
		result, err := j.ParseToken(c.wanted)
		assert.NoError(t, err)
		assert.Equal(t, c.CustomClaims.ID, result.ID)
		assert.Equal(t, c.CustomClaims.Email, result.Email)
		assert.Equal(t, c.CustomClaims.Name, result.Name)
		assert.Equal(t, c.CustomClaims.StandardClaims.ExpiresAt, result.StandardClaims.ExpiresAt)
		assert.Equal(t, c.CustomClaims.StandardClaims.Issuer, result.StandardClaims.Issuer)
	}
}

func TestRefreshToken(t *testing.T) {
	for _, c := range expiredClaims {
		j := &JWT{SigningKey: []byte(c.siginKey)}
		token, err := j.CreateToken(c.CustomClaims)
		assert.NoError(t, err)
		claims, err := j.ParseToken(token)
		assert.EqualError(t, err, ErrTokenExpired.Error())
		assert.Nil(t, claims)
		token, err = j.RefreshToken(token)
		assert.NoError(t, err)
		assert.IsType(t, "string", token)
	}
}
