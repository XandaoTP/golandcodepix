package repository

import "gorm.io/gorm"

type PixKeyRepositoryDb struct {
	Db *gorm.DB
}

func (r PixKeyRepositoryDb) AddBank(bank *model) {

}
