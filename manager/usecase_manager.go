package manager

import (
	"bank/usecase"
	"sync"
)

type UsecaseManager interface {
	GetUserUsecase() usecase.UserUsecase
}

type usecaseManagerImpl struct {
	repoManagerImpl RepoManager

	userUsecase    usecase.UserUsecase
	cstmrUsecase   usecase.CustomerUsecase
	paymentUsecase usecase.PaymentUsecase
}

var onceLoadUserUsecase sync.Once
var onceLoadCustomerUsecase sync.Once
var onceLoadPaymentUsecase sync.Once

func (um *usecaseManagerImpl) GetUserUsecase() usecase.UserUsecase {
	onceLoadUserUsecase.Do(func() {
		um.userUsecase = usecase.NewUserUsecase(um.repoManagerImpl.GetUserRepo())
	})
	return um.userUsecase
}

func (um *usecaseManagerImpl) GetCustomerUsecase() usecase.CustomerUsecase {
	onceLoadCustomerUsecase.Do(func() {
		um.cstmrUsecase = usecase.NewCustomerUsecase(um.repoManagerImpl.GetCustomerRepo())
	})
	return um.cstmrUsecase
}

func (um *usecaseManagerImpl) GetPaymentUsecase() usecase.PaymentUsecase {
	onceLoadPaymentUsecase.Do(func() {
		um.paymentUsecase = usecase.NewPaymentUsecase(um.repoManagerImpl.GetPaymentRepo())
	})
	return um.paymentUsecase
}

func NewUsecaseManager(repoManagerImpl RepoManager) UsecaseManager {
	return &usecaseManagerImpl{
		repoManagerImpl: repoManagerImpl,
	}
}
