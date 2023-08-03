package controller

import (
	"bank/apperror"
	"bank/model"
	"bank/usecase"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func (usrCntrl *UserController) GetAllUser(ctx *gin.Context) {
	arrUser, err := usrCntrl.userUsecase.GetAllUser()
	fmt.Println("----", arrUser)
	if err != nil {
		fmt.Printf("userHandlerImpl.GetAllUser() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    arrUser,
	})
}

func (usrCntrl *UserController) GetUserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	if username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Username tidak boleh kosong",
		})
		return
	}

	usr, err := usrCntrl.userUsecase.GetUserByUsername(username)
	if err != nil {
		fmt.Printf("UserHandler.GetUserByName() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data User",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    usr,
	})
}

func (usrCntrl *UserController) GetUserById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id tidak boleh kosong",
		})
		return
	}

	usr, err := usrCntrl.userUsecase.GetUserById(id)
	if err != nil {
		fmt.Printf("UserHandler.GetUserByName() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data User",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    usr,
	})
}

func (usrCntrl *UserController) RegisterUser(ctx *gin.Context) {
	usr := &model.UserModel{}
	err := ctx.ShouldBindJSON(&usr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	if len(usr.Username) > 20 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Panjang Nama tidak boleh lebih dari 20 karakter",
		})
		return
	}

	err = usrCntrl.userUsecase.RegisterUser(usr)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("UserHandler.InsertUser() 1 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("UserHandler.InsertUser() 2 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika menyimpan data User",
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (usrCntrl *UserController) DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := usrCntrl.userUsecase.DeleteUser(id)
	if err != nil {
		fmt.Printf("userHandler.userUseCase.DeleteUser(id) : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan dalam menghapus data User",
		})
		return
	}
	// err = usrHandler.ctmUsecase.DeleteCustomer(id)
	// if err != nil {
	// 	fmt.Printf("usrHandler.usrUsecase.DeleteUser(id) : %v", err.Error())
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"success":      false,
	// 		"errorMessage": "Terjadi kesalahan dalam menghapus data Customer",
	// 	})
	// 	return
	// }
}

func (usrCntrl *UserController) UpdateUser(ctx *gin.Context) {
	usr := &model.UserModel{}
	usr.Id, _ = strconv.Atoi(ctx.Param("id"))
	err := ctx.ShouldBindJSON(&usr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	if len(usr.Username) > 20 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Panjang Nama tidak boleh lebih dari 20 karakter",
		})
		return
	}

	err = usrCntrl.userUsecase.UpdateUser(usr)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("UserHandler.InsertUser() 1 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("UserHandler.InsertUser() 2 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika menyimpan data User",
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func NewUserController(srv *gin.Engine, userUsecase usecase.UserUsecase) *UserController {
	usrCntrl := &UserController{
		userUsecase: userUsecase,
	}
	srv.GET("/admin/user", usrCntrl.GetAllUser)
	srv.GET("/user/:id", usrCntrl.GetUserById)
	srv.POST("/user", usrCntrl.RegisterUser)
	srv.GET("/user/:name", usrCntrl.GetUserByUsername)
	srv.DELETE("/admin/:id", usrCntrl.DeleteUser)
	srv.PUT("user/:id", usrCntrl.UpdateUser)

	return usrCntrl
}
