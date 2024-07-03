package model

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrNoFirstName error = errors.New("your first name is required")
	ErrNoLastName        = errors.New("your last name is required")
	ErrNoEmail           = errors.New("your email is required")
	ErrUnderAge          = errors.New(" you must be older than 10 years")
)

type User struct {
	Id             int64     `json:"-"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Password       string    `json:"pass_word"`
	HashedPassword string    `json:"hashed_password"`
	Active         bool      `json:"active"`
	Age            int       `json:"age"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type ConfigDatabase struct {
	DbUser   string
	DbPass   string
	DbRemote string
	DbPort   string
	DbName   string
}

func (u *User) ValidateUser() error {
	if u.FirstName == "" {
		return ErrNoFirstName
	}
	if u.LastName == "" {
		return ErrNoLastName
	}
	if u.Email == "" {
		return ErrNoEmail
	}
	if u.Age < 10 {
		return ErrUnderAge
	}
	u.Email = strings.ToLower(u.Email)
	u.FirstName = strings.ToUpper(string(u.FirstName[0])) + strings.ToLower(u.FirstName[1:])
	u.LastName = strings.ToUpper(string(u.LastName[0])) + strings.ToLower(u.LastName[1:])
	u.Active = true
	u.Create(true)
	return nil
}

func (u *User) Create(t ...bool) {
	if len(t) > 0 || t[0] {
		u.CreatedAt = time.Now().UTC()
	}
	u.UpdatedAt = time.Now().UTC()
}

type Account struct {
	Id            int       `json:"-"`
	AccountNumber string    `json:"account_number"`
	Active        bool      `json:"active"`
	Balance       float64   `json:"balance"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	User
}

type Transaction struct {
}
