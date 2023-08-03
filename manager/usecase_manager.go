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

	userUsecase usecase.UserUsecase
}

var onceLoadUserUsecase sync.Once

func (um *usecaseManagerImpl) GetUserUsecase() usecase.UserUsecase {
	onceLoadUserUsecase.Do(func() {
		um.userUsecase = usecase.NewUserUsecase(um.repoManagerImpl.GetUserRepo())

	})
	return um.userUsecase
}

func NewUsecaseManager(repoManagerImpl RepoManager) UsecaseManager {
	return &usecaseManagerImpl{
		repoManagerImpl: repoManagerImpl,
	}
}
