package manager

import (
	"bank/repo"
	"sync"
)

type RepoManager interface {
	GetUserRepo() repo.UserRepo
	GetCustomerRepo() repo.CustomerRepo
	GetPaymentRepo() repo.PaymentRepo
}

type repoManagerImpl struct {
	infraManager InfraManager

	usrRepo     repo.UserRepo
	cstmrRepo   repo.CustomerRepo
	paymentRepo repo.PaymentRepo
}

var onceLoadUserRepo sync.Once
var onceLoadCustomerRepo sync.Once
var onceLoadPaymentRepo sync.Once

func (rm *repoManagerImpl) GetUserRepo() repo.UserRepo {
	onceLoadUserRepo.Do(func() {
		rm.usrRepo = repo.NewUserRepo(rm.infraManager.GetDB())
	})
	return rm.usrRepo
}

func (rm *repoManagerImpl) GetCustomerRepo() repo.CustomerRepo {
	onceLoadCustomerRepo.Do(func() {
		rm.cstmrRepo = repo.NewCustomerRepo(rm.infraManager.GetDB())
	})
	return rm.cstmrRepo
}
func (rm *repoManagerImpl) GetPaymentRepo() repo.PaymentRepo {
	onceLoadPaymentRepo.Do(func() {
		rm.paymentRepo = repo.NewPaymentRepo(rm.infraManager.GetDB())
	})
	return rm.paymentRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManagerImpl{
		infraManager: infraManager,
	}
}
