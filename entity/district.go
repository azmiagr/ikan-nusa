package entity

type District struct {
	DistrictID   int    `json:"district_id" gorm:"type:int;primaryKey;autoIncrement"`
	DistrictName string `json:"district_name" gorm:"type:varchar(100);not null"`
	City         City   `json:"city"`
	CityID       int    `json:"city_id"`

	Addresses []Address `gorm:"foreignKey:DistrictID"`
}
