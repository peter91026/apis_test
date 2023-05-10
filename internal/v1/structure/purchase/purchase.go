package purchase

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

// Table struct is database table struct
type Table struct {
	// 編號
	PurchaseID string `gorm:"primaryKey;column:purchase_id;uuid_generate_v4()type:UUID;" json:"purchase_id,omitempty"`
	// 申請人
	Applicant string `gorm:"column:applicant;type:VARCHAR;" json:"applicant,omitempty"`
	// 公司名稱
	CompanyName string `gorm:"column:company_name;type:VARCHAR;" json:"company_name,omitempty"`
	// 部門
	Department string `gorm:"column:department;type:VARCHAR;" json:"department,omitempty"`
	// 品名
	ProductName string `gorm:"column:product_name;type:VARCHAR;" json:"product_name,omitempty"`
	// 價格
	Price int64 `gorm:"column:price;type:INT4;" json:"price,omitempty"`
	// 數量
	PurchaseQuantity int64 `gorm:"column:purchase_quantity;type:INT4" json:"purchase_quantity,omitempty"`
	// 創建者
	CreatedBy string `gorm:"column:created_by;type:VARCHAR;" json:"created_by,omitempty"`
	// 創建時間
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at"`
	// 更新時間
	UpdatedAt *time.Time `gorm:"column:updated_at;type:TIMESTAMP;" json:"updated_at,omitempty"`
	// 更新者
	UpdatedBy *string `gorm:"column:updated_by;type:UUID;" json:"updated_by,omitempty"`
	// 是否刪除
	IsDeleted bool `gorm:"column:is_deleted;type:bool;false" json:"is_deleted,omitempty"`
}

// Base struct is corresponding to table structure file
type Base struct {
	// 編號
	PurchaseID string `json:"purchase_id,omitempty"`
	// 申請人
	Applicant string `json:"applicant,omitempty"`
	// 公司名稱
	CompanyName string `json:"company_name,omitempty"`
	// 部門
	Department string `json:"department,omitempty"`
	// 品名
	ProductName string `json:"product_name,omitempty"`
	// 價格
	Price int64 `json:"price,omitempty"`
	// 數量
	PurchaseQuantity int64 `json:"purchase_quantity,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 創建時間
	CreatedAt time.Time `json:"created_at"`
	// 更新時間
	UpdatedAt *time.Time `gorm:"column:updated_at;type:TIMESTAMP;" json:"updated_at,omitempty"`
	// 更新者
	UpdatedBy *string `gorm:"column:updated_by;type:UUID;" json:"updated_by,omitempty"`
	// 是否刪除
	IsDeleted bool `json:"is_deleted,omitempty"`
}

// Single return structure file
type Single struct {
	// 編號
	PurchaseID string `json:"purchase_id,omitempty"`
	// 申請人
	Applicant string `json:"applicant,omitempty"`
	// 公司名稱
	CompanyName string `json:"company_name,omitempty"`
	// 部門
	Department string `json:"department,omitempty"`
	// 品名
	ProductName string `json:"product_name,omitempty"`
	// 價格
	Price int64 `json:"price,omitempty"`
	// 數量
	PurchaseQuantity int64 `json:"purchase_quantity,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 創建時間
	CreatedAt time.Time `json:"created_at"`
	// 更新時間
	UpdatedAt *time.Time ` json:"updated_at,omitempty"`
	// 更新者
	UpdatedBy *string ` json:"updated_by,omitempty"`
}

// Created struct is used to create
type Created struct {
	// 申請人
	Applicant string `json:"applicant,omitempty" binding:"required" validate:"required"`
	// 公司名稱
	CompanyName string `json:"company_name,omitempty" binding:"required" validate:"required"`
	// 部門
	Department string `json:"department,omitempty" binding:"required" validate:"required"`
	// 品名
	ProductName string `json:"product_name,omitempty" binding:"required" validate:"required"`
	// 價格
	Price int64 `json:"price,omitempty" binding:"required" validate:"required"`
	// 數量
	PurchaseQuantity int64 `json:"purchase_quantity,omitempty" binding:"required" validate:"required"`
	// 創建者
	CreatedBy string `json:"created_by" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 編號
	PurchaseID string `json:"purchase_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 申請人
	Applicant *string `json:"applicant,omitempty" form:"applicant"`
	// 公司名稱
	CompanyName *string `json:"company_name,omitempty" form:"company_name"`
	// 部門
	Department *string `json:"department,omitempty" form:"department"`
	// 品名
	ProductName *string `json:"product_name,omitempty" form:"product_name"`
	// 價格
	Price *int64 `json:"price,omitempty" form:"price"`
	// 數量
	PurchaseQuantity *int64 `json:"purchase_quantity,omitempty" form:"purchase_quantity"`
	// 是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty" swaggerignore:"true"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	Field
	model.InPage
}

// List is multiple return structure files
type List struct {
	Purchase []*struct {
		// 編號
		PurchaseID string `json:"purchase_id,omitempty"`
		// 申請人
		Applicant string `json:"applicant,omitempty"`
		// 公司名稱
		CompanyName string `json:"company_name,omitempty"`
		// 部門
		Department string `json:"department,omitempty"`
		// 品名
		ProductName string `json:"product_name,omitempty"`
		// 價格
		Price int64 `json:"price,omitempty"`
		// 數量
		PurchaseQuantity int64 `json:"purchase_quantity,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
	} `json:"purchase"`
	model.OutPage
}

// Updated struct is used to update
type Updated struct {
	// 編號
	PurchaseID string `json:"purchase_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 公司名稱
	CompanyName *string `json:"company_name,omitempty"`
	// 申請人
	Applicant *string `json:"applicant,omitempty"`
	// 部門
	Department *string `json:"department,omitempty"`
	// 品名
	ProductName *string `json:"product_name,omitempty"`
	// 價格
	Price *int64 `json:"price,omitempty"`
	// 數量
	PurchaseQuantity *int64 `json:"purchase_quantity,omitempty"`
	// 更新者
	UpdatedBy *string `gorm:"column:updated_by;type:UUID;" json:"updated_by,omitempty"`
	// 是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty" swaggerignore:"true"`
}

// TableName sets the insert table name for this struct type
func (a *Table) TableName() string {
	return "purchase"
}
