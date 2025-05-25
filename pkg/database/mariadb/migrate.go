package mariadb

import (
	"ikan-nusa/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Store{},
		&entity.Product{},
		&entity.OtpCode{},
		&entity.Cart{},
		&entity.Province{},
		&entity.City{},
		&entity.District{},
		&entity.Address{},
		&entity.CartItems{},
	)
	if err != nil {
		return err
	}

	return nil
}
