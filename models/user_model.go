package models

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID         string  `json:"id,omitempty" bson:"_id"`
	FirstName  string  `json:"first_name,omitempty" bson:"first_name"`
	FamilyName string  `json:"family_name,omitempty" bson:"family_name"`
	Email      string  `json:"email,omitempty" bson:"email"`
	Password   string  `json:"password,omitempty" bson:"password"`
	Savings    float64 `json:"savings,omitempty" bson:"savings"`
	Expenses   float64 `json:"expenses,omitempty" bson:"expenses"`
	Income     float64 `json:"income,omitempty" bson:"income"`
	IsVerified bool    `json:"is_verified,omitempty" bson:"is_verified"`
	Wallet	 float64  `json:"wallet,omitempty" bson:"wallet"`
	Target	 float64 `json:"target,omitempty" bson:"target"`
}

func (u *User) SetInitialAttributes(hashedPassword string) {
	u.Savings = 0
	u.Expenses = 0
	u.Income = 0
	u.Wallet = 0
	u.Target = 0
	u.IsVerified = false
	u.ID = uuid.New().String()
	u.Password = hashedPassword
}
func (u *User) IncreaseIncome(amount float64) error {
	if amount < 0 {
		return errors.New("amount cannot be negative")
	}
	u.Income += amount
	u.Wallet += amount
	return nil
}

func (u *User) DecreaseIncome(amount float64) error {
	if amount < 0 {
		return errors.New("amount cannot be negative")
	}
	if amount > u.Income {
		return errors.New("insufficient income")
	}
	u.Income -= amount
	u.Wallet -= amount
	return nil
}

func (u *User) GetIncome() float64 {
	return u.Income
}

func (u *User) IncreaseExpense(amount float64) error {
	if amount < 0 {
		return errors.New("amount cannot be negative")
	}
	u.Expenses += amount
	u.Wallet -= amount
	return nil
}

func (u *User) DecreaseExpense(amount float64) error {
	if amount < 0 {
		return errors.New("amount cannot be negative")
	}
	if amount > u.Expenses {
		return errors.New("insufficient expenses")
	}
	u.Expenses -= amount
	u.Wallet += amount
	return nil
}

func (u *User) GetExpenses() float64 {
	return u.Expenses
}

func (u *User) DepositSavings(amount float64) error {
	if amount < 0 {
		return errors.New("amount cannot be negative")
	}
	u.Savings += amount
	return nil
}

func (u *User) GetSavings() float64 {
	return u.Savings
}

func (u *User) SetWallet(amount float64) error {
	if amount < 0 {
		return errors.New("amount cannot be negative")
	}
	u.Wallet = amount
	return nil
}

func (u *User) GetWallet() float64 {
	return u.Wallet
}

func (u *User) SetTarget(amount float64) error {
	if amount < 0 {
		return errors.New("amount cannot be negative")
	}
	u.Target = amount
	return nil
}

func (u *User) GetTarget() float64 {
	return u.Target
}
