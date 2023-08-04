package usecase

import (
	"bank/apperror"
	"bank/model"
	"bank/repo"
	"fmt"
)

type PaymentUsecase interface {
	GetAllPayment() ([]model.PaymentModel, error)
	GetPaymentByCustomerId(id int) ([]model.PaymentModel, error)
	GetPaymentById(custId int, id int) (*model.PaymentModel, error)
	InsertPayment(pay *model.PaymentModel) error
}

type paymentUsecaseImpl struct {
	paymentRepo repo.PaymentRepo
	custRepo    repo.CustomerRepo
}

func (paymentUsecase *paymentUsecaseImpl) GetAllPayment() ([]model.PaymentModel, error) {
	return paymentUsecase.paymentRepo.GetAllPayment()
}

func (paymentUsecase *paymentUsecaseImpl) GetPaymentByCustomerId(id int) ([]model.PaymentModel, error) {
	return paymentUsecase.paymentRepo.GetPaymentByCustomerId(id)
}

func (paymentUsecase *paymentUsecaseImpl) GetPaymentById(custId int, id int) (*model.PaymentModel, error) {
	return paymentUsecase.paymentRepo.GetPaymentById(custId, id)
}

func (paymentUsecase *paymentUsecaseImpl) InsertPayment(pay *model.PaymentModel) error {
	senderId := pay.CustomerId
	recieverId := pay.DestinationId

	senderCust, _ := paymentUsecase.custRepo.GetBlanceById(senderId)
	balanceSender := senderCust.Balance - int(pay.Paid)
	if balanceSender <= 0 {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("Saldo tidak cukup, saldo Anda : %v", senderCust.Balance),
		}
	}
	senderCust.Balance = balanceSender
	paymentUsecase.custRepo.UpdateBalance(senderCust)

	recieverCust, _ := paymentUsecase.custRepo.GetBlanceById(recieverId)
	balanceReciever := recieverCust.Balance + int(pay.Paid)
	recieverCust.Balance = balanceReciever
	paymentUsecase.custRepo.UpdateBalance(recieverCust)

	return paymentUsecase.paymentRepo.InsertPayment(pay)
}

func NewPaymentUsecase(paymentRepo repo.PaymentRepo) PaymentUsecase {
	return &paymentUsecaseImpl{
		paymentRepo: paymentRepo,
	}
}
