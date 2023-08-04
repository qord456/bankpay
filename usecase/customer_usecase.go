package usecase

import (
	"bank/apperror"
	"bank/model"
	"bank/repo"
	"fmt"
)

type CustomerUsecase interface {
	GetAllCustomer() ([]model.CustomerModel, error)
	GetCustomerById(id int) (*model.CustomerModel, error)
	GetCustomerByname(name string) (*model.CustomerModel, error)
	RegisterCustomer(cstmr *model.CustomerModel) error
	DeleteCustomer(id int) error
	UpdateCustomer(cstmrArr *model.CustomerModel) error
	UpdateCustomerStatus(cstmrArr *model.CustomerModel) error
}

type customerUsecaseImpl struct {
	cstmrRepo repo.CustomerRepo
}

func (cstmrUsecase *customerUsecaseImpl) DeleteCustomer(id int) error {
	cstmr, err := cstmrUsecase.cstmrRepo.GetCustomerById(id)
	if err != nil {
		return fmt.Errorf("cstmrUsecase.cstmrRepo.GetCustomerByName() : %w", err)
	}
	if cstmr == nil {
		return apperror.AppError{
			ErrorCode:    400,
			ErrorMessage: fmt.Sprintf("data customer dengan id : %d tidak ada", id),
		}
	}

	return nil
}

func (cstmrUsecase *customerUsecaseImpl) UpdateCustomer(cstmrArr *model.CustomerModel) error {
	return cstmrUsecase.cstmrRepo.UpdateCustomer(cstmrArr)
}

func (cstmrUsecase *customerUsecaseImpl) UpdateCustomerStatus(cstmrArr *model.CustomerModel) error {
	return cstmrUsecase.cstmrRepo.UpdateCustomerStatus(cstmrArr)
}

func (cstmrUsecase *customerUsecaseImpl) GetCustomerById(id int) (*model.CustomerModel, error) {
	cstmr, err := cstmrUsecase.cstmrRepo.GetCustomerById(id)
	if err != nil {
		return nil, fmt.Errorf("cstmrUsecase.cstmrRepo.GetCustomerByName() : %w", err)
	}
	if cstmr == nil {
		return nil, apperror.AppError{
			ErrorCode:    400,
			ErrorMessage: fmt.Sprintf("data customer dengan id : %d tidak ada", id),
		}
	}

	return cstmr, nil
}

func (cstmrUsecase *customerUsecaseImpl) GetCustomerByname(name string) (*model.CustomerModel, error) {
	cstmr, err := cstmrUsecase.cstmrRepo.GetCustomerByName(name)
	if err != nil {
		return nil, fmt.Errorf("cstmrUsecase.cstmrRepo.GetCustomerByName() : %w", err)
	}
	if cstmr == nil {
		return nil, apperror.AppError{
			ErrorCode:    400,
			ErrorMessage: fmt.Sprintf("data customer dengan nama : %s tidak ada", name),
		}
	}

	return cstmr, nil
}

func (cstmrUsecase *customerUsecaseImpl) RegisterCustomer(cstmr *model.CustomerModel) error {
	cstmrAvaliable, err := cstmrUsecase.cstmrRepo.GetCustomerByName(cstmr.Name)
	if err != nil {
		return fmt.Errorf("cstmrUsecase.cstmrRepo.GetCustomerByName() : %w", err)
	}
	if cstmrAvaliable != nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("data customer dengan nama : %s sudah ada", cstmr.Name),
		}
	}
	return cstmrUsecase.cstmrRepo.RegisterCustomer(cstmr)
}

func (cstmrUsecase *customerUsecaseImpl) GetAllCustomer() ([]model.CustomerModel, error) {
	return cstmrUsecase.cstmrRepo.GetAllCustomer()
}

func NewCustomerUsecase(cstmrRepo repo.CustomerRepo) CustomerUsecase {
	return &customerUsecaseImpl{
		cstmrRepo: cstmrRepo,
	}
}
