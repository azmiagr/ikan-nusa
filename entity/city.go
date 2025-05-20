package entity

type City struct {
	CityID     int      `json:"city_id" gorm:"type:int;primaryKey;autoIncrement"`
	CityName   string   `json:"city_name" gorm:"type:varchar(100);not null"`
	Province   Province `json:"province"`
	ProvinceID int      `json:"province_id"`

	Districts []District `json:"district" gorm:"foreignKey:CityID"`
}
