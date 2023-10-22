package repository

import (
	"fmt"
	"github.com/XandaoTP/golandcodepix/domain/model"
	"gorm.io/gorm"
)

type PixKeyRepositoryDb struct {
	Db *gorm.DB
}

func (r PixKeyRepositoryDb) AddBank(bank *model.Bank) error {
	err := r.Db.Create(bank).Error
	if err != nil {
		return err
	}
	return nil
}

func (r PixKeyRepositoryDb) AddAccount(account *model.Account) error {
	err := r.Db.Create(account).Error
	if err != nil {
		return err
	}
	return nil
}

func (r PixKeyRepositoryDb) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) {
	err := r.Db.Create(pixKey).Error
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}

func (r PixKeyRepositoryDb) FindKeyById(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	r.Db.Preload("Account.Bank").First(&pixkey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}
	return &pixKey, nil
}

func (r PixKeyRepositoryDb) FindKeyById(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	r.Db.Preload("Account.Bank").First(&pixkey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}
	return &pixKey, nil
}
