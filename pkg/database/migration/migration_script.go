package main

import (
	"fmt"

	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/entities"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/config"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/database"
	"github.com/isd-sgcu/sucu-backend-2024/utils/constant"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	cfg := config.GetConfig()
	db := database.NewGormDatabase(cfg)

	// migrate schema
	if err := db.AutoMigrate(entities.Role{}); err != nil {
		panic("Error while migrating roles table: " + err.Error())
	}
	if err := db.AutoMigrate(entities.DocumentType{}); err != nil {
		panic("Error while migrating document_types table: " + err.Error())
	}
	if err := db.AutoMigrate(entities.AttachmentType{}); err != nil {
		panic("Error while migrating attachment_types table: " + err.Error())
	}
	if err := db.AutoMigrate(entities.User{}); err != nil {
		panic("Error while migrating users table: " + err.Error())
	}
	if err := db.AutoMigrate(entities.Document{}); err != nil {
		panic("Error while migrating documents table: " + err.Error())
	}
	if err := db.AutoMigrate(entities.Attachment{}); err != nil {
		panic("Error while migrating attachments table: " + err.Error())
	}

	// init data
	var roles []entities.Role = []entities.Role{
		{ID: constant.SGCU_SUPERADMIN},
		{ID: constant.SGCU_ADMIN},
		{ID: constant.SCCU_SUPERADMIN},
		{ID: constant.SCCU_ADMIN},
	}

	var documentTypes []entities.DocumentType = []entities.DocumentType{
		{ID: constant.ANNOUNCEMENT},
		{ID: constant.BUDGET},
		{ID: constant.STATISTIC},
	}

	var attachmentTypes []entities.AttachmentType = []entities.AttachmentType{
		{ID: constant.DOCS},
		{ID: constant.IMAGE},
	}

	var user entities.User = entities.User{
		ID:        "6633221100",
		FirstName: "Methee",
		LastName:  "Hephong",
		Password: func() string {
			hashed, err := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
			if err != nil {
				panic("Error while generating hash password: " + err.Error())
			}
			return string(hashed)
		}(),
		RoleID: constant.SGCU_SUPERADMIN,
	}

	// migrate init data
	if err := db.Table("roles").Create(&roles).Error; err != nil {
		panic("Error while migrating roles data: " + err.Error())
	}
	if err := db.Table("document_types").Create(&documentTypes).Error; err != nil {
		panic("Error while migrating document_types data: " + err.Error())
	}
	if err := db.Table("attachment_types").Create(&attachmentTypes).Error; err != nil {
		panic("Error while migrating attachment_types data: " + err.Error())
	}
	if err := db.Table("users").Create(&user).Error; err != nil {
		panic("Error while migrating users data: " + err.Error())
	}

	fmt.Println("migration successful")
}
