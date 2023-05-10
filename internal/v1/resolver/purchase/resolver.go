package purchase

import (
	"eirc.app/internal/v1/service/purchase"
	model "eirc.app/internal/v1/structure/purchase"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
}

type resolver struct {
	PurchaseService purchase.Service
}

func New(db *gorm.DB) Resolver {
	return &resolver{
		PurchaseService: purchase.New(db),
	}
}
