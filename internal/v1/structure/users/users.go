package User

import (
	model "eirc.app/internal/v1/structure"
)

// User struct is a row record of the companies table in the invoice database
// Table struct is database table struct
type Table struct {
	// 使用者編號
	UserID int `gorm:"primaryKey;uuid_generate_v4();column:User_id;type:UUID;" json:"User_id,omitempty"`
	// 使用者名稱
	UserName string `gorm:"column:UserName;type:VARCHAR;" json:"UserName,omitempty"`
	// 使用者密碼
	PassWord string `gorm:"column:Password;type:VARCHAR;" json:"Password,omitempty"`
	// 使用者電子信箱
	Email string `gorm:"column:Email;type:VARCHAR;" json:"Email,omitempty"`
}

// Base struct is corresponding to table structure file
type Base struct {
	// 使用者編號
	UserID int `json:"User_id,omitempty"`
	// 使用者名稱
	UserName string `json:"UserName,omitempty"`
	// 使用者密碼
	PassWord string `json:"Password,omitempty"`
	// 使用者電子信箱
	Email string `json:"Email,omitempty"`
}

// Single return structure file
type Single struct {
	// 使用者編號
	UserID int `json:"User_id,omitempty"`
	// 使用者名稱
	UserName string `json:"UserName,omitempty"`
	// 使用者密碼
	PassWord string `json:"Password,omitempty"`
	// 使用者電子信箱
	Email string `json:"Email,omitempty"`
}

// Created struct is used to create
type Created struct {
	// 使用者名稱
	UserName string `json:"UserName,omitempty" binding:"required" validate:"required"`
	// 使用者密碼
	PassWord string `json:"PassWord,omitempty" binding:"required" validate:"required"`
	// 使用者電子信箱
	Email string `json:"Email,omitempty" binding:"required" validate:"required"`
}

// Updated struct is used to update
type Updated struct {
	// 使用者編號
	UserID int `json:"User_id,omitempty"  swaggerignore:"true"`
	// 使用者名稱
	UserName string `json:"UserName,omitempty" `
	//使用者密碼
	PassWord string `json:"PassWord,omitempty"  `
	//使用者電子信箱
	Email string `json:"Email,omitempty" `
}

// Field is structure file for search
type Field struct {
	// 使用者編號
	UserID int `json:"User_id,omitempty" swaggerignore:"true"`
	// 使用者名稱
	UserName *string `json:"UserName,omitempty"  swaggerignore:"UserName"`
	// 使用者密碼
	PassWord string `json:"PassWord,omitempty"  swaggerignore:"PassWord"`
	// 使用者電子信箱
	Email string `json:"Email,omitempty"   swaggerignore:"Email"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	Field
	model.InPage
}

// List is multiple return structure files
type List struct {
	Companies []*struct {
		// 使用者編號
		UserID string `json:"User_id,omitempty"`
		// 使用者名稱
		UserName string `json:"UserName,omitempty"`
		// 使用者密碼
		PassWord string `json:"Password,omitempty"`
		// 使用者電子信箱
		Email string `json:"Email,omitempty"`
	} `json:"companies"`
	model.OutPage
}

// TableUserName sets the insert table UserName for this struct type
func (c *Table) TableUserName() string {
	return "companies"
}
