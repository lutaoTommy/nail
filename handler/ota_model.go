package handler

/*Ota升级*/
type OtaVersion struct {
	Id   string `gorm:"primaryKey;column:id;size:20" json:"id"`
	Time string `gorm:"column:time;size:20" json:"time"`
	Name string `gorm:"column:name;size:50" json:"name"`
	Code int    `gorm:"column:code;type:int" json:"code"`
}

/*表名兼容*/
func (OtaVersion) TableName() string {
	return "ota_versons"
}

/*Ota升级*/
type OtaOut struct {
	Id   string `json:"id"`
	Time string `json:"time"`
	Name string `json:"name"`
	Code int    `json:"code"`
	Url  string `json:"url"`
}
