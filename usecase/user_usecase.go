package usecase

import (
	"bank/apperror"
	"bank/model"
	"bank/repo"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	DeleteUser(id int) error
	UpdateUser(usrArr *model.UserModel) error
	GetUserById(id int) (*model.UserModel, error)
	GetUserByUsername(name string) (*model.UserModel, error)
	RegisterUser(usr *model.UserModel) error
	GetAllUser() ([]model.UserModel, error)
}

type userUsecaseImpl struct {
	usrRepo repo.UserRepo
}

func (usrUsecase *userUsecaseImpl) DeleteUser(id int) error {
	usr, err := usrUsecase.usrRepo.GetUserById(id)
	if err != nil {
		return fmt.Errorf("usrUsecase.usrRepo.GetUserByName() : %w", err)
	}
	if usr == nil {
		return apperror.AppError{
			ErrorCode:    400,
			ErrorMessage: fmt.Sprintf("data user dengan id : %d tidak ada", id),
		}
	}

	return nil
}

func (usrUsecase *userUsecaseImpl) UpdateUser(usrArr *model.UserModel) error {
	return usrUsecase.usrRepo.UpdateUser(usrArr)
}

func (usrUsecase *userUsecaseImpl) GetUserById(id int) (*model.UserModel, error) {
	usr, err := usrUsecase.usrRepo.GetUserById(id)
	if err != nil {
		return nil, fmt.Errorf("usrUsecase.usrRepo.GetUserByName() : %w", err)
	}
	if usr == nil {
		return nil, apperror.AppError{
			ErrorCode:    400,
			ErrorMessage: fmt.Sprintf("data user dengan id : %d tidak ada", id),
		}
	}

	return usr, nil
}

func (usrUsecase *userUsecaseImpl) GetUserByUsername(name string) (*model.UserModel, error) {
	usr, err := usrUsecase.usrRepo.GetUserByUsername(name)
	if err != nil {
		return nil, fmt.Errorf("usrUsecase.usrRepo.GetUserByName() : %w", err)
	}
	if usr == nil {
		return nil, apperror.AppError{
			ErrorCode:    400,
			ErrorMessage: fmt.Sprintf("data user dengan nama : %s tidak ada", name),
		}
	}

	return usr, nil
}

func (usrUsecase *userUsecaseImpl) RegisterUser(usr *model.UserModel) error {
	usrAvaliable, err := usrUsecase.usrRepo.GetUserByUsername(usr.Username)
	if err != nil {
		return fmt.Errorf("usrUsecase.usrRepo.GetUserByName() : %w", err)
	}
	if usrAvaliable != nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("data user dengan nama : %s sudah ada", usr.Username),
		}
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost) //hashing
	if err != nil {
		return nil
	}
	usr.Password = string(hashedPassword)

	return usrUsecase.usrRepo.RegisterUser(usr)
}

func (usrUsecase *userUsecaseImpl) GetAllUser() ([]model.UserModel, error) {
	return usrUsecase.usrRepo.GetAllUser()
}

func NewUserUsecase(usrRepo repo.UserRepo) UserUsecase {
	return &userUsecaseImpl{
		usrRepo: usrRepo,
	}
}
