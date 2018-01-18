package models

type Lorawan struct {
	Lda_dno       string `gorm:"primary_key"` // set AnimalId to be primary key
	Lda_dev_addr  string
	Lda_app_key   string
	Lda_app_s_key string
	Lda_nwk_s_key string
}

func (Lorawan) TableName() string {
	return "iot_lorawan_device_attr"
}
