package handler

/*积分记录表*/
type PointRecord struct {
	Id         string `gorm:"primaryKey;column:id;size:20" json:"id"`
	UserId     string `gorm:"column:user_id;size:20;index" json:"user_id"`
	Points     int    `gorm:"column:points;type:int" json:"points"`
	Type       string `gorm:"column:type;size:20" json:"type"`
	CreateTime string `gorm:"column:create_time;size:20" json:"create_time"`
}

/*签到记录表*/
type CheckinRecord struct {
	Id         string `gorm:"primaryKey;column:id;size:20" json:"id"`
	UserId     string `gorm:"column:user_id;size:20;uniqueIndex:idx_user_date" json:"user_id"`
	Date       string `gorm:"column:date;size:10;uniqueIndex:idx_user_date" json:"date"`
	CreateTime string `gorm:"column:create_time;size:20" json:"create_time"`
}
