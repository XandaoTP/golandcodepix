package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Account struct {
	Base      `valid:"required"`
	OwnerName string    `gorm:"column:owner_name;type:varchar(255);notnull" valid:"notnull"`
	Bank      *Bank     `valid:"-"`
	BankID    string    `gorm:"column:bank_id;type:uuid;notnull" valid:"-"`
	Number    string    `json:"number" gorm:"" valid:"-"`
	PixKey    []*PixKey `gorm:"ForeingKey:AccountID" valid:"_"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}
	return nil
}

func NewAccount(bank *Bank, number string, ownerName string) (*Account, error) {
	account := Account{
		OwnerName: ownerName,
		Bank:      bank,
		Number:    number,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()
	err := account.isValid()
	if err != nil {
		return nil, err
	}
	return &account, nil
}
