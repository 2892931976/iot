package models

type Product struct {
	P_no                    string `gorm:"primary_key"` // set AnimalId to be primary key
	P_name                  string
	P_key                   string
	P_type                  int //2:网关 1:设备
	P_communication_type_id int // 2:laroWan 1:普通
}

func (Product) TableName() string {
	return "iot_product"
}

func IsGetwayDevice(d_pno string) bool {
	product := Product{}
	return !DB.Where("p_name = ? and p_type = ?", d_pno, "2").First(&product).RecordNotFound()
}

func IsNormalDevice(d_pno string) bool {
	product := Product{}
	return !DB.Where("p_name = ? and p_type = ? and p_communication_type_id = ?", d_pno, "1",
		"1").First(&product).RecordNotFound()
}

func IsLoraWanDevice(d_pno string) bool {
	product := Product{}
	return !DB.Where("p_name = ? and p_type = ? and p_communication_type_id = ?", d_pno, "1",
		"2").First(&product).RecordNotFound()
}

func IsProductNotValid(d_pno string) bool {
	product := Product{}
	return DB.Where("p_name = ?", d_pno).First(&product).RecordNotFound()
}
