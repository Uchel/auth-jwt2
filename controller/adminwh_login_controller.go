package controller

import (
	"net/http"

	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/Uchel/auth-jwt2/model"
	"github.com/Uchel/auth-jwt2/usecase"
	"github.com/gin-gonic/gin"
)

type AuthAdminWhLoginController struct {
	AdminLoginUsecase usecase.AdminWhLoginUsecase
}

func (c *AuthAdminWhLoginController) LoginAdmin(ctx *gin.Context) {

	var loginReq model.LoginReq

	if err := ctx.BindJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid input")
		return
	}
	email, password := c.AdminLoginUsecase.FindByEmailAdminWh(loginReq.Email)

	// authenticate user (compare username dan password)
	if loginReq.Email == email && loginReq.Password == password {
		// generate JWT token
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = loginReq.Email
		claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "gagal generate token"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unregistered user"})
	}
}

func NewAdminLoginController(u usecase.AdminWhLoginUsecase) *AuthAdminWhLoginController {
	controller := AuthAdminWhLoginController{

		AdminLoginUsecase: u,
	}

	return &controller
}
