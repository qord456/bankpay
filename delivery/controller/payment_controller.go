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

type PaymentController struct {
	paymentUsecase usecase.PaymentUsecase
}

func (paymentCntrl *PaymentController) GetAllPayment(ctx *gin.Context) {
	arrPayment, err := paymentCntrl.paymentUsecase.GetAllPayment()
	fmt.Println("----", arrPayment)
	if err != nil {
		fmt.Printf("paymentHandlerImpl.GetAllPayment() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data payment",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    arrPayment,
	})
}

func (paymentCntrl *PaymentController) GetPaymentByCustomerId(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id tidak boleh kosong",
		})
		return
	}

	arrPayment, err := paymentCntrl.paymentUsecase.GetPaymentByCustomerId(id)
	fmt.Println("----", arrPayment)
	if err != nil {
		fmt.Printf("paymentHandlerImpl.GetAllPayment() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data payment",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    arrPayment,
	})
}

func (paymentCntrl *PaymentController) GetPaymentById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	custIdStr := ctx.Param("custid")
	custId, _ := strconv.Atoi(custIdStr)
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id Payment tidak boleh kosong",
		})
		return
	}
	if custIdStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id Customer tidak boleh kosong",
		})
		return
	}

	payment, err := paymentCntrl.paymentUsecase.GetPaymentById(custId, id)
	if err != nil {
		fmt.Printf("PaymentHandler.GetPaymentByName() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data Payment",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    payment,
	})
}

func (paymentCntrl *PaymentController) InsertPayment(ctx *gin.Context) {
	payment := &model.PaymentModel{}
	err := ctx.ShouldBindJSON(&payment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	if payment.Paid <= 10000 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Angka yang ditransfer tidak boleh kosong dan harus lebih dari 10000",
		})
		return
	}

	err = paymentCntrl.paymentUsecase.InsertPayment(payment)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("PaymentHandler.InsertPayment() 1 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("PaymentHandler.InsertPayment() 2 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika menyimpan data Payment",
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func NewPaymentController(srv *gin.Engine, paymentUsecase usecase.PaymentUsecase) *PaymentController {
	paymentCntrl := &PaymentController{
		paymentUsecase: paymentUsecase,
	}
	//FOR ADMIN
	srv.GET("/admin/payment", paymentCntrl.GetAllPayment)
	srv.GET("/admin/payment/user/:custId/:id", paymentCntrl.GetPaymentById)
	srv.GET("/admin/payment/user/:id", paymentCntrl.GetPaymentByCustomerId)
	srv.POST("/admin/payment/user/", paymentCntrl.InsertPayment)
	//FOR USER
	srv.GET("/user/payment/:custid/:id", paymentCntrl.GetPaymentByCustomerId)
	srv.GET("/user/payment/:id", paymentCntrl.GetPaymentByCustomerId)
	srv.POST("/user/payment/", paymentCntrl.InsertPayment)
	return paymentCntrl
}
