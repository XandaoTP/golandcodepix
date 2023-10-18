package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	AccountID string   `json:"account_id" valid:"notnull"`
	Account   *Account `valid:"-"`
	Status    string   `json:"status" valid:"notnull"`
}

func (pixKey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixKey)
	if err != nil {
		return err
	}
	return nil
}

func NewPix(kind string, account *Account, key string) (*PixKey, error) {
	pixKey := PixKey{
		Kind:    kind,
		Key:     key,
		Account: account,
		Status:  "active",
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()
	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}
	return &pixKey, nil
}
