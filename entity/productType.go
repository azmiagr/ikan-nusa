package entity

type ProductType struct {
	ProductTypeID int    `json:"product_type_id" gorm:"type:int;primaryKey;autoIncrement"`
	Type          string `json:"type" gorm:"type:varchar(50)"`

	Products []Product `json:"products" gorm:"foreignKey:ProductTypeID"`
}
