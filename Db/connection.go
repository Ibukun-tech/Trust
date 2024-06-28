package db

import (
	"database/sql"

	model "github.con/Ibukun-tech/trust/Internals/Models"
	"github.con/Ibukun-tech/trust/utils"
)

type DbConnect struct {
	Db *sql.DB
}

// I am must have a connection that must like take in everything so that each associated handlers can make use of it a nd all of it finish
type DB interface {
	CreateUser(*model.User) (*model.User, error)
	ListUsers() ([]*model.User, error)
	CreateAccount(*model.Account) error
}

func NewDbConnection(pd string, c model.ConfigDatabase) (*DbConnect, error) {
	connect := utils.ConnectDb(c)
	db, err := sql.Open(pd, connect)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &DbConnect{
		Db: db,
	}, nil
}

func (d *DbConnect) CreateUser(m *model.User) (*model.User, error) {

	// Write an algorithm that will be able to a create a unique number for all
	//
	query := `insert into users (
		first_name, last_name, email, pass_word, hashed_password,age, active,  created_at                 
 )values ($1, $2, $3, $4, $5, $6)`

	//  Need to write a middleware utils function that will be able to like edit the creating of names and everything
	// I will write a transaction to be able to create the account number at the same time
	_, err := d.Db.Query(query, m.FirstName, m.LastName, m.Email, m.Password, m.HashedPassword, m.Age, m.Active, m.CreatedAt)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (d *DbConnect) ListUsers() ([]model.User, error) {
	var users []model.User
	query := "Select * from Users"
	row, err := d.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var user model.User
		if err := row.Scan(
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Active,
			&user.HashedPassword,
			&user.CreatedAt,
			&user.Password,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := row.Close(); err != nil {
		return nil, err
	}
	if err := row.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (d *DbConnect) CreateAccount(m *model.Account) error {
	// I can create a validator for the creation on the account
	query := `insert into account (
         id, account_number, , balance, created_at                 
	)
	values ($1, $2, $3, $4, $5, $6)
    `

	return nil
}
