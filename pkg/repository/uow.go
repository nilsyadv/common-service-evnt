package repository

import (
	"gorm.io/gorm"

	"github.com/Nilesh-Coherent/common-service-evnt/pkg/db"
)

type UOW struct {
	DB        *gorm.DB
	readonly  bool
	committed bool
}

func NewUnitOfWork(db *db.DB, flag bool) *UOW {
	if flag {
		return &UOW{
			DB:       db.Getdb(),
			readonly: flag,
		}
	}
	return &UOW{
		DB:       db.Getdb().Begin(),
		readonly: flag,
	}
}

func (uow *UOW) Commit() {
	if !uow.readonly {
		uow.DB.Commit()
	}
	uow.committed = true
}

func (uow *UOW) RollBack() {
	if !uow.committed && !uow.readonly {
		uow.DB.Rollback()
	}
}
