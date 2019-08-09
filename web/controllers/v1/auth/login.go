package auth

import (
	"net/http"
	"raytheon/datamodels"
	"raytheon/utils"
	"raytheon/web/middleware/jwtauth"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginResult is a representation of a user token information
type LoginResult struct {
	Token    string `json:"token"`
	UserName string `json:"username"`
	Avatar   string `json:"avatar"`
}

//GenerateToken 生成token
func GenerateToken(c *gin.Context, user datamodels.User, tenant string) {
	j := &jwtauth.JWT{
		SigningKey: []byte("raytheon"),
	}
	claims := jwtauth.CustomClaims{
		Name:   user.UserName,
		Role:   user.Role,
		Tenant: tenant,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,  //签名生效时间
			ExpiresAt: time.Now().Unix() + 86400, //签名过期时间 24h
			Issuer:    "raytheon",                //签名发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(utils.APIResponseMeta(c, http.StatusForbidden, nil))
		return
	}

	c.JSON(utils.APIResponseMeta(c, http.StatusOK, LoginResult{
		Token:    token,
		UserName: user.UserName,
		Avatar:   user.Avatar,
	}))
}

// Login user login
func Login(c *gin.Context) {
	var loginStruct datamodels.User
	err := c.Bind(&loginStruct)
	if err != nil {
		c.JSON(utils.APIResponseMeta(c, http.StatusForbidden, nil))
		return
	}

	var user datamodels.User
	utils.DBConn.Where("username = ?", loginStruct.UserName).First(&user)
	if user.ID == 0 {
		user.UserName = loginStruct.UserName
		user.Password = loginStruct.Password
		user.Role = utils.WebViewPerm
		utils.DBConn.Create(&user)
		if user.ID == 0 {
			c.JSON(utils.APIResponseMeta(c, http.StatusForbidden, nil))
			return
		}
	}
	var isLogin bool
	if utils.APIConfig.LDAPConfig.Enable {
		isLogin, err = utils.UserAuthentication(loginStruct.UserName, loginStruct.Password)
		if err != nil {
			c.JSON(utils.APIResponseMeta(c, http.StatusForbidden, nil))
			return
		}
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginStruct.Password))
		if err != nil {
			c.JSON(utils.APIResponseMeta(c, http.StatusForbidden, nil))
			return
		}
		isLogin = true
	}
	if !isLogin {
		c.JSON(utils.APIResponseMeta(c, http.StatusForbidden, nil))
		return
	}
	GenerateToken(c, user, "user")
}
