package manager

import (
	"bank/repo"
	"sync"
)

type RepoManager interface {
	GetUserRepo() repo.UserRepo
}

type repoManagerImpl struct {
	infraManager InfraManager

	usrRepo   repo.UserRepo
	cstmrRepo repo.CustomerRepo
}

var onceLoadUserRepo sync.Once
var onceLoadCustomerRepo sync.Once

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

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManagerImpl{
		infraManager: infraManager,
	}
}
