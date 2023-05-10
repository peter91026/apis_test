package purchase

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	purchaseModel "eirc.app/internal/v1/structure/purchase"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *purchaseModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 檢查重複
	purchase, err := r.PurchaseService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, purchase.PurchaseID)
}

func (r *resolver) List(input *purchaseModel.Fields) interface{} {
	input.IsDeleted = util.PointerBool(false)
	output := &purchaseModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, purcahse, err := r.PurchaseService.List(input)
	purchaseByte, err := json.Marshal(purcahse)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(purchaseByte, &output.Purchase)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *purchaseModel.Field) interface{} {
	input.IsDeleted = util.PointerBool(false)
	base, err := r.PurchaseService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	frontCompany := &purchaseModel.Single{}
	companyByte, _ := json.Marshal(base)
	err = json.Unmarshal(companyByte, &frontCompany)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, frontCompany)
}

func (r *resolver) Deleted(input *purchaseModel.Updated) interface{} {
	_, err := r.PurchaseService.GetByID(&purchaseModel.Field{PurchaseID: input.PurchaseID,
		IsDeleted: util.PointerBool(false)})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.PurchaseService.Deleted(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *purchaseModel.Updated) interface{} {
	purchase, err := r.PurchaseService.GetByID(&purchaseModel.Field{PurchaseID: input.PurchaseID,
		IsDeleted: util.PointerBool(false)})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.PurchaseService.Updated(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, purchase.PurchaseID)
}
