package entity

type Province struct {
	ProvinceID   int    `json:"province_id" gorm:"type:int;primaryKey;autoIncrement"`
	ProvinceName string `json:"province_name" gorm:"type:varchar(30);not null"`

	Cities []City `json:"city" gorm:"foreignKey:ProvinceID"`
}
