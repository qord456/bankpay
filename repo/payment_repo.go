package repo

import (
	"bank/model"
	"bank/utils"
	"database/sql"
	"fmt"
)

type PaymentRepo interface {
	InsertPayment(pay *model.PaymentModel) error
	GetAllPayment() ([]model.PaymentModel, error)
	GetPaymentByCustomerId(id int) ([]model.PaymentModel, error)
	GetPaymentById(custId int, id int) (*model.PaymentModel, error)
}

type paymentRepoImpl struct {
	db *sql.DB
}

func (payRepo *paymentRepoImpl) InsertPayment(pay *model.PaymentModel) error {
	qry := utils.INSERT_PAYMENT
	_, err := payRepo.db.Exec(qry, pay.CustomerId, pay.Paid, pay.DestinationId, pay.CreatedBy, pay.CreatedAt)
	if err != nil {
		return fmt.Errorf("error on paymentRepoImpl.InsertPayment : %v", err)
	}
	return nil
}

func (payRepo *paymentRepoImpl) GetAllPayment() ([]model.PaymentModel, error) {
	qry := utils.SELECT_ALL_PAYMENT
	rows, err := payRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("getAllPayment() : %w", err)
	}
	defer rows.Close()

	var arrPayment []model.PaymentModel
	for rows.Next() {
		payment := model.PaymentModel{}
		rows.Scan(&payment.Id, &payment.CustomerId, &payment.Paid, &payment.DestinationId, &payment.CreatedBy, payment.CreatedAt)
		arrPayment = append(arrPayment, payment)
	}
	return arrPayment, nil
}

func (payRepo *paymentRepoImpl) GetPaymentByCustomerId(id int) ([]model.PaymentModel, error) {
	qry := utils.SELECT_PAYMENT_BY_CUSTOMER_ID
	rows, err := payRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("getPaymentByCustomerId() : %w", err)
	}
	defer rows.Close()

	var arrPayment []model.PaymentModel
	for rows.Next() {
		payment := model.PaymentModel{}
		rows.Scan(&payment.Id, &payment.CustomerId, &payment.Paid, &payment.DestinationId, &payment.CreatedBy, payment.CreatedAt)
		arrPayment = append(arrPayment, payment)
	}
	return arrPayment, nil
}

func (payRepo *paymentRepoImpl) GetPaymentById(custId int, id int) (*model.PaymentModel, error) {
	qry := utils.SELECT_PAYMENT_BY_ID
	payment := &model.PaymentModel{}
	err := payRepo.db.QueryRow(qry, id).Scan(&payment.Id, &payment.CustomerId, &payment.Paid, &payment.DestinationId, &payment.CreatedBy, payment.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on paymentRepoImpl.GetPaymentById() : %w", err)
	}
	return payment, nil
}

func NewPaymentRepo(db *sql.DB) PaymentRepo {
	return &paymentRepoImpl{
		db: db,
	}
}
