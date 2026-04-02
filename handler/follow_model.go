package handler

/*关注关系表*/
type Follow struct {
	Id           string           `gorm:"primaryKey;column:id;size:20" json:"id"`
	FollowerID   string           `gorm:"column:follower_id;size:20" json:"follower_id"`   // 关注者ID
	FollowingID  string           `gorm:"column:following_id;size:20" json:"following_id"` // 被关注者ID
	Status       int              `gorm:"column:status;type:tinyint" json:"status"`        // 状态：-1 取消关注 1 已关注
	FollowTime   string           `gorm:"column:follow_time;size:20" json:"follow_time"`
	CancelTime   string           `gorm:"column:cancel_time;size:20" json:"cancel_time"`
}

/*关注关系表*/
type FollowOut struct {
	Id             string    `json:"id"`
	FollowingID    string    `json:"following_id"` // 被关注者ID
	Nickname       string    `json:"nickname"`
	Avatar         string    `json:"avatar"`
	Biography      string    `json:"biography"`
	FollowTime      string   `json:"follow_time"`
	FollowCount    int       `json:"followCount"`
	FansCount      int       `json:"fansCount"`
	PostCount      int       `json:"postCount"`
}