package handler

/*收藏*/
type Collect struct {
	Id     string     `gorm:"column:id;size:20" json:"id"`
	Time   string     `gorm:"column:time;size:20" json:"time"`
	UserId string     `gorm:"column:user_id;size:20;index" json:"user_id"`
	PostId string     `gorm:"column:post_id;size:20;index" json:"post_id"`
	User   UserSimple `json:"user"`
}

/*表名*/
func (Collect) TableName() string {
	return "collects"
}
