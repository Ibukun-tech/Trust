package model

import "time"

type User struct {
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
