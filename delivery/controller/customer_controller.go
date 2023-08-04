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

type CustomerController struct {
	cstmrUsecase usecase.CustomerUsecase
	userUsecase  usecase.UserUsecase
}

func (cstmrCntrl *CustomerController) GetAllCustomer(ctx *gin.Context) {
	arrCustomer, err := cstmrCntrl.cstmrUsecase.GetAllCustomer()
	fmt.Println("----", arrCustomer)
	if err != nil {
		fmt.Printf("customerHandlerImpl.GetAllCustomer() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data customer",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    arrCustomer,
	})
}

func (cstmrCntrl *CustomerController) GetCustomerById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id tidak boleh kosong",
		})
		return
	}

	Cstmr, err := cstmrCntrl.cstmrUsecase.GetCustomerById(id)
	if err != nil {
		fmt.Printf("CustomerHandler.GetCustomerByName() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data Customer",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    Cstmr,
	})
}

func (cstmrCntrl *UserController) GetCustomerByName(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Nama tidak boleh kosong",
		})
		return
	}

	cstmr, err := cstmrCntrl.cstmrUsecase.GetCustomerByname(name)
	if err != nil {
		fmt.Printf("CustomerHandler.GetCustomerByName() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data Customer",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    cstmr,
	})
}

func (cstmrCntrl *CustomerController) DeleteCustomer(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id tidak boleh kosong",
		})
		return
	}

	err := cstmrCntrl.userUsecase.DeleteUser(id)
	if err != nil {
		fmt.Printf("userHandler.userUseCase.DeleteUser(id) : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan dalam menghapus data User",
		})
		return
	}

	err = cstmrCntrl.cstmrUsecase.DeleteCustomer(id)
	if err != nil {
		fmt.Printf("usrHandler.usrUsecase.DeleteUser(id) : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan dalam menghapus data Customer",
		})
		return
	}
}

func (cstmrCntrl *CustomerController) UpdateCustomerStatus(ctx *gin.Context) {
	cstmr := &model.CustomerModel{}
	cstmr.Id, _ = strconv.Atoi(ctx.Param("id"))
	err := ctx.ShouldBindJSON(&cstmr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = cstmrCntrl.cstmrUsecase.UpdateCustomerStatus(cstmr)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("CustomerHandler.UpdateCustomer() 1 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("CustomerHandler.UpdateCustomer() 2 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika menyimpan data Customer",
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (cstmrCntrl *CustomerController) UpdateCustomer(ctx *gin.Context) {
	cstmr := &model.CustomerModel{}
	cstmr.Id, _ = strconv.Atoi(ctx.Param("id"))
	err := ctx.ShouldBindJSON(&cstmr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	if len(cstmr.Name) > 20 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Panjang Nama tidak boleh lebih dari 20 karakter",
		})
		return
	}

	err = cstmrCntrl.cstmrUsecase.UpdateCustomer(cstmr)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("CustomerHandler.UpdateCustomer() 1 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("CustomerHandler.UpdateCustomer() 2 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika menyimpan data Customer",
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (cstmrCntrl *CustomerController) RegisteCustomer(ctx *gin.Context) {
	cstmr := &model.CustomerModel{}
	err := ctx.ShouldBindJSON(&cstmr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	if len(cstmr.Name) > 20 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Panjang Nama tidak boleh lebih dari 20 karakter",
		})
		return
	}

	err = cstmrCntrl.cstmrUsecase.RegisterCustomer(cstmr)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("CustomerHandler.InsertCustomer() 1 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("CustomerHandler.InsertCustomer() 2 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika menyimpan data Customer",
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func NewCustomerController(srv *gin.Engine, cstmrUsecase usecase.CustomerUsecase) *CustomerController {
	cstmrCntrl := &CustomerController{
		cstmrUsecase: cstmrUsecase,
	}
	//FOR ADMIN
	srv.GET("/admin/customer", cstmrCntrl.GetAllCustomer)
	srv.GET("/admin/customer/:id", cstmrCntrl.GetCustomerById)
	srv.GET("/admin/customer/:name", cstmrCntrl.GetCustomerById)
	srv.DELETE("/admin/customer/:id", cstmrCntrl.DeleteCustomer)
	srv.PUT("admin/customer/:id", cstmrCntrl.UpdateCustomer)
	srv.PUT("admin/customer/status:id", cstmrCntrl.UpdateCustomerStatus)
	srv.POST("admin/customer", cstmrCntrl.RegisteCustomer)

	//FOR USER
	srv.GET("/customer/:id", cstmrCntrl.GetCustomerById)
	srv.DELETE("/customer/:id", cstmrCntrl.DeleteCustomer)
	srv.PUT("/customer/:id", cstmrCntrl.UpdateCustomer)
	srv.POST("/customer")
	return cstmrCntrl
}
