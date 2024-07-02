package db

import (
	"context"
	"database/sql"

	model "github.con/Ibukun-tech/trust/Internals/Models"
	"github.con/Ibukun-tech/trust/utils"
)

// Instead of doing all this I can just create a struct that will be only for the running of
// Tracsaction code then after that *sql.tx then all operations if a code can be doe by it
type DbTransactionFunc func(*DbConnect) error
type DbInterface interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}
type DbConnect struct {
	Db DbInterface
}

func NewDbConnect(db *sql.DB) *DbConnect {
	return &DbConnect{
		Db: db,
	}
}

type StoreWithTx struct {
	*DbConnect
	Db *sql.DB
}

func NewDbConnection(pd string, c model.ConfigDatabase) (*StoreWithTx, error) {
	connect := utils.ConnectDb(c)
	db, err := sql.Open(pd, connect)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &StoreWithTx{
		Db:        db,
		DbConnect: NewDbConnect(db),
	}, nil
}

func (d *DbConnect) AddTx(tx *sql.Tx) *DbConnect {
	return &DbConnect{
		Db: tx,
	}
}

func (s *StoreWithTx) ExecuteTrx(ctx context.Context, st DbTransactionFunc) error {

	tx, err := s.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	dbConn := s.AddTx(tx)
	err = st(dbConn)
	if err != nil {
		return err
	}

	return nil
}

func (d *DbConnect) CreateUser(ctx context.Context, m *model.User) (*model.User, error) {

	// Write an algorithm that will be able to a create a unique number for all
	//
	query := `insert into users (
		first_name, last_name, email, pass_word, hashed_password,age, active,  created_at                 
 )values ($1, $2, $3, $4, $5, $6)`

	//  Need to write a middleware utils function that will be able to like edit the creating of names and everything
	// I will write a transaction to be able to create the account number at the same time
	_, err := d.Db.ExecContext(ctx, query, m.FirstName, m.LastName, m.Email, m.Password, m.HashedPassword, m.Age, m.Active, m.CreatedAt)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (d *DbConnect) ListUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	query := "Select * from Users"
	row, err := d.Db.QueryContext(ctx, query)
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

func (d *DbConnect) CreateAccount(ctx context.Context, m *model.Account) error {
	// I can create a validator for the creation on the account
	query := `insert into account (
         id, user_id, account_number, balance, active, created_at                 
	)
	values ($1, $2, $3, $4, $5, $6)
    `
	_, err := d.Db.ExecContext(ctx, query, m.Id, m.User.Id, m.AccountNumber, m.Balance, m.Active, m.CreatedAt)
	if err != nil {
		return nil
	}
	return nil
}

// Want to add Transaction to the process
