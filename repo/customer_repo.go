package repo

import (
	"bank/model"
	"bank/utils"
	"database/sql"
	"fmt"
)

type CustomerRepo interface {
	GetAllCustomer() ([]model.CustomerModel, error)
	GetCustomerByName(name string) (*model.CustomerModel, error)
	GetCustomerById(id int) (*model.CustomerModel, error)
	RegisterCustomer(cstmr *model.CustomerModel) error
	UpdateCustomer(cstmr *model.CustomerModel) error
	UpdateCustomerStatus(cstmr *model.CustomerModel) error
	DeleteCustomer(id int) error
	UpdateBalance(pay *model.CustomerModel) error
	GetBlanceById(id int) (*model.CustomerModel, error)
}

type customerRepoImpl struct {
	db *sql.DB
}

func (cstmrRepo *customerRepoImpl) RegisterCustomer(cstmr *model.CustomerModel) error {
	qry := utils.REGISTER_CUSTOMER
	cstmr.Status = "active"
	cstmr.Balance = 0
	_, err := cstmrRepo.db.Exec(qry, cstmr.UserId, cstmr.Nik, cstmr.Name, cstmr.Email, cstmr.Phone, cstmr.Address, cstmr.Birthdate, cstmr.Balance, cstmr.Status)
	if err != nil {
		return fmt.Errorf("error on customerRepoImpl.UpdateCustomer : %v", err)
	}
	return nil
}

func (cstmrRepo *customerRepoImpl) UpdateCustomer(cstmr *model.CustomerModel) error {
	qry := utils.UPDATE_CUSTOMER
	_, err := cstmrRepo.db.Exec(qry, cstmr.Id, cstmr.UserId, cstmr.Nik, cstmr.Name, cstmr.Email, cstmr.Phone, cstmr.Address, cstmr.Birthdate)
	if err != nil {
		return fmt.Errorf("error on customerRepoImpl.UpdateCustomer : %v", err)
	}
	return nil
}

func (cstmrRepo *customerRepoImpl) UpdateCustomerStatus(cstmr *model.CustomerModel) error {
	qry := utils.STATUS_CUSTOMER
	_, err := cstmrRepo.db.Exec(qry, cstmr.Id, cstmr.Status)
	if err != nil {
		return fmt.Errorf("error on customerRepoImpl.UpdateCustomerStatus : %v", err)
	}
	return nil
}

func (cstmrRepo *customerRepoImpl) DeleteCustomer(id int) error {
	qry := utils.DELETE_CUSTOMER
	_, err := cstmrRepo.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("error on customerRepoImpl.DeleteCustomer : %v", &err)
	}
	return nil
}

func (cstmrRepo *customerRepoImpl) GetAllCustomer() ([]model.CustomerModel, error) {
	qry := utils.SELECT_ALL_CUSTOMER
	rows, err := cstmrRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("getAllCustomer() : %w", err)
	}
	defer rows.Close()

	var arrCustomer []model.CustomerModel
	for rows.Next() {
		cstmr := model.CustomerModel{}
		rows.Scan(&cstmr.Id, &cstmr.UserId, &cstmr.Nik, &cstmr.Name, &cstmr.Email, &cstmr.Phone, &cstmr.Address, &cstmr.Birthdate, &cstmr.Balance, &cstmr.Status)
		arrCustomer = append(arrCustomer, cstmr)
	}
	return arrCustomer, nil
}

func (cstmrRepo *customerRepoImpl) GetCustomerByName(name string) (*model.CustomerModel, error) {
	qry := utils.SELECT_CUSTOMER_BY_NAME

	cstmr := &model.CustomerModel{}
	err := cstmrRepo.db.QueryRow(qry, name).Scan(&cstmr.Id, &cstmr.UserId, &cstmr.Nik, &cstmr.Name, &cstmr.Email, &cstmr.Phone, &cstmr.Address, &cstmr.Birthdate, &cstmr.Balance, &cstmr.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on serviceRepoImpl.GetCustomerByName() : %w", err)
	}
	return cstmr, nil
}

func (cstmrRepo *customerRepoImpl) GetCustomerById(id int) (*model.CustomerModel, error) {
	qry := utils.SELECT_CUSTOMER_BY_ID
	cstmr := &model.CustomerModel{}
	err := cstmrRepo.db.QueryRow(qry, id).Scan(&cstmr.Id, &cstmr.UserId, &cstmr.Nik, &cstmr.Name, &cstmr.Email, &cstmr.Phone, &cstmr.Address, &cstmr.Birthdate, &cstmr.Balance, &cstmr.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on serviceRepoImpl.GetCustomerByName() : %w", err)
	}
	return cstmr, nil
}

func (cstmrRepo *customerRepoImpl) UpdateBalance(pay *model.CustomerModel) error {
	qry := utils.UPDATE_BALANCE
	_, err := cstmrRepo.db.Exec(qry, pay.Id, pay.Balance)
	if err != nil {
		return fmt.Errorf("error on paymentRepoImpl.InsertPayment : %v", err)
	}
	return nil
}

func (cstmrRepo *customerRepoImpl) GetBlanceById(id int) (*model.CustomerModel, error) {
	qry := utils.GET_BALANCE
	pay := &model.CustomerModel{}
	err := cstmrRepo.db.QueryRow(qry, id).Scan(&pay.Id, &pay.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on paymentRepoImpl.GetPaymentById() : %w", err)
	}
	return pay, nil
}

func NewCustomerRepo(db *sql.DB) CustomerRepo {
	return &customerRepoImpl{
		db: db,
	}
}
