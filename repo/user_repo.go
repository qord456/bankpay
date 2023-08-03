package repo

import (
	"bank/model"
	"bank/utils"
	"database/sql"
	"fmt"
	"time"
)

type UserRepo interface {
	GetAllUser() ([]model.UserModel, error)
	GetUserByUsername(name string) (*model.UserModel, error)
	GetUserById(id int) (*model.UserModel, error)
	RegisterUser(usr *model.UserModel) error
	UpdateUser(usr *model.UserModel) error
	DeleteUser(id int) error
}

type userRepoImpl struct {
	db *sql.DB
}

func (usrRepo *userRepoImpl) UpdateUser(usr *model.UserModel) error {
	qry := utils.UPDATE_USER
	usr.UpdatedAt = time.Now()
	_, err := usrRepo.db.Exec(qry, usr.Id, usr.Username, usr.Password, usr.UpdatedAt)
	if err != nil {
		return fmt.Errorf("error on userRepoImpl.UpdateUser : %v", err)
	}
	return nil
}

func (usrRepo *userRepoImpl) DeleteUser(id int) error {
	qry := utils.DELETE_USER
	_, err := usrRepo.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("error on userRepoImpl.DeleteUser : %v", &err)
	}
	return nil
}

func (usrRepo *userRepoImpl) GetAllUser() ([]model.UserModel, error) {
	qry := utils.SELECT_ALL_USER
	rows, err := usrRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("getAllUser() : %w", err)
	}
	defer rows.Close()

	var arrUser []model.UserModel
	for rows.Next() {
		usr := model.UserModel{}
		rows.Scan(&usr.Id, &usr.Username, &usr.Password, &usr.CreatedAt, &usr.UpdatedAt)
		arrUser = append(arrUser, usr)
	}
	return arrUser, nil
}

func (usrRepo *userRepoImpl) RegisterUser(usr *model.UserModel) error {
	qry := utils.INSERT_USER

	usr.CreatedAt = time.Now()
	usr.UpdatedAt = time.Now()

	_, err := usrRepo.db.Exec(qry, usr.Username, usr.Password, usr.CreatedAt, usr.UpdatedAt)
	if err != nil {
		return fmt.Errorf("error on userRepoImpl.RegisterUser() : %w", err)
	}
	return nil
}

func (usrRepo *userRepoImpl) GetUserByUsername(name string) (*model.UserModel, error) {
	qry := utils.SELECT_USER_BY_NAME

	usr := &model.UserModel{}
	err := usrRepo.db.QueryRow(qry, name).Scan(&usr.Id, &usr.Username, &usr.Password, &usr.CreatedAt, &usr.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on serviceRepoImpl.GetUserByName() : %w", err)
	}
	return usr, nil
}

func (usrRepo *userRepoImpl) GetUserById(id int) (*model.UserModel, error) {
	qry := utils.SELECT_USER_BY_ID
	usr := &model.UserModel{}
	err := usrRepo.db.QueryRow(qry, id).Scan(&usr.Id, &usr.Username, &usr.Password, &usr.CreatedAt, &usr.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on serviceRepoImpl.GetUserByName() : %w", err)
	}
	return usr, nil
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepoImpl{
		db: db,
	}
}
