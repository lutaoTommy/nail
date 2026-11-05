package handler

/*话题*/
type Topic struct {
	Id         string `gorm:"primaryKey;column:id;size:20" json:"id"`
	Name       string `gorm:"column:name;size:50" json:"name"`
	Status     int    `gorm:"column:status;type:tinyint" json:"status"` // 状态：1 正常
	CreateTime string `gorm:"column:create_time;size:20" json:"create_time"`
}
