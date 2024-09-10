package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `gorm:"primaryKey;type:varchar(10)"` // student id
	FirstName string         `gorm:"type:varchar(100);not null"`
	LastName  string         `gorm:"type:varchar(100);not null"`
	Password  string         `gorm:"type:varchar(255);not null"` // password's length 255 is used for hashed password
	RoleID    string         `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time      ``
	UpdatedAt time.Time      ``
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Role      Role       `gorm:"foreignKey:RoleID"` // role: SGCU_ADMIN , SGCU_SUPERADMIN , SCCU_ADMIN , SCCU_SUPERADMIN
	Documents []Document `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Role struct {
	ID        string         `gorm:"primaryKey;type:varchar(100)"`
	CreatedAt time.Time      ``
	UpdatedAt time.Time      ``
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Users []User `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type DocumentType struct {
	ID        string         `gorm:"primaryKey;type:varchar(100)"`
	CreatedAt time.Time      ``
	UpdatedAt time.Time      ``
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Documents []Document `gorm:"foreignKey:TypeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Document struct {
	ID        string         `gorm:"primaryKey;type:varchar(100)"`
	Title     string         `gorm:"type:varchar(255);not null"`
	Content   string         `gorm:"type:text;not null"`
	Banner    *string        `gorm:"type:varchar(255)"`
	Cover     *string        `gorm:"type:varchar(255)"`
	UserID    string         `gorm:"type:varchar(10);not null"`
	TypeID    string         `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time      ``
	UpdatedAt time.Time      ``
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Author      User         `gorm:"foreignKey:UserID"`
	Type        DocumentType `gorm:"foreignKey:TypeID"`
	Attachments []Attachment `gorm:"foreignKey:DocumentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type AttachmentType struct {
	ID        string         `gorm:"primaryKey;type:varchar(100)"`
	CreatedAt time.Time      ``
	UpdatedAt time.Time      ``
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Attachments []Attachment `gorm:"foreignKey:TypeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Attachment struct {
	ID          string         `gorm:"primaryKey;type:varchar(100)"`
	DisplayName string         `gorm:"type:varchar(255);not null"`
	DocumentID  string         `gorm:"type:varchar(100);not null"`
	TypeID      string         `gorm:"type:varchar(100);not null"`
	CreatedAt   time.Time      ``
	UpdatedAt   time.Time      ``
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Document Document       `gorm:"foreignKey:DocumentID"`
	Type     AttachmentType `gorm:"foreignKey:TypeID"`
}
