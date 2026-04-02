package handler

/*评论*/
type Comment struct {
	Id        string          `gorm:"column:id;size:20" json:"id"`
	Time      string          `gorm:"column:time;size:20" json:"time"`
    Content   string          `gorm:"column:content;size:500" json:"content"`
    UserId    string          `gorm:"column:user_id;size:20;index" json:"user_id"`
    PostId    string          `gorm:"column:post_id;size:20;index" json:"post_id"`
    User      UserSimple      `json:"user"`
}

/*表名*/
func (Comment) TableName() string {
    return "comments"
}