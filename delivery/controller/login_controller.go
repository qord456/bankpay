package controller

import (
	"bank/apperror"
	"bank/model"
	"bank/usecase"
	"bank/utils/authutil"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	lgnUsecase usecase.LoginUsecase
}

func (lgnController LoginController) Login(ctx *gin.Context) {
	loginData := &model.LoginModel{}
	err := ctx.ShouldBindJSON(&loginData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = lgnController.lgnUsecase.GetUserByNameAndPassword(loginData.Username, loginData.Password)
	if err != nil {

		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("loginHandlerImpl.Login() : %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.ErrorMessage,
			})
		} else {
			fmt.Printf("loginHandlerImpl.Login() : %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika login",
			})
		}
		return
	}

	temp, err := authutil.GenerateToken(loginData.Username)
	if err != nil {
		log.Println("Token Invalid")
	}
	ctx.Header("Authorization", "Bearer "+temp)
	ctx.JSON(http.StatusOK, gin.H{
		"token": temp,
	})
}

func (lgnController LoginController) Logout(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	authutil.InvalidateToken(token)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}

func NewLoginController(srv *gin.Engine, lgnUsecase usecase.LoginUsecase) {
	lgnHandler := &LoginController{
		lgnUsecase: lgnUsecase,
	}
	srv.POST("/login", lgnHandler.Login)
	srv.POST("/logout", lgnHandler.Logout)
}
