package purchase

import (
	model "eirc.app/internal/v1/structure/purchase"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	if input.Applicant != nil {
		db.Where("applicant = ?", input.Applicant)
	}

	if input.CompanyName != nil {
		db.Where("company_name like %?%", *input.CompanyName)
	}
	if input.Department != nil {
		db.Where("department = ?", input.Department)
	}

	if input.ProductName != nil {
		db.Where("product_name = ?", input.ProductName)
	}

	if input.PurchaseQuantity != nil {
		db.Where("purchase_quantity = ?", input.PurchaseQuantity)
	}

	if input.Price != nil {
		db.Where("price = ?", input.Price)
	}

	if input.IsDeleted != nil {
		db.Where("is_deleted = ?", input.IsDeleted)
	}

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("created_at desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("purchase_id = ?", input.PurchaseID)
	if input.IsDeleted != nil {
		db.Where("is_deleted = ?", input.IsDeleted)
	}

	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Field) (err error) {
	err = e.db.Model(&model.Table{}).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Save(&input).Error

	return err
}
