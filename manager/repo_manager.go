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

	usrRepo repo.UserRepo
}

var onceLoadUserRepo sync.Once

func (rm *repoManagerImpl) GetUserRepo() repo.UserRepo {
	onceLoadUserRepo.Do(func() {
		rm.usrRepo = repo.NewUserRepo(rm.infraManager.GetDB())
	})
	return rm.usrRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManagerImpl{
		infraManager: infraManager,
	}
}
