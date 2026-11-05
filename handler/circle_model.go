package handler

import "unicode/utf8"

/*发布圈子*/
type CirclePost struct {
	Id           string      `gorm:"primaryKey;column:id;size:20" json:"id"`
	Token        string      `gorm:"-" json:"token"`
	UserId       string      `gorm:"column:user_id;size:20" json:"user_id"`
	Content      string      `gorm:"column:content;size:500" json:"content"`
	Topic        string      `gorm:"column:topic;size:50" json:"topic"`
	CreateTime   string      `gorm:"column:create_time;size:20" json:"create_time"`
	UpdateTime   string      `gorm:"column:update_time;size:20" json:"update_time"`
	LikeCount    int         `gorm:"column:like_count;type:int" json:"like_count"`
	CollectCount int         `gorm:"column:collect_count;type:int" json:"collect_count"`
	CommentCount int         `gorm:"column:comment_count;type:int" json:"comment_count"`
	Status       int         `gorm:"column:status;type:tinyint" json:"status"`                               // 状态：-1 删除 1 正常
	ReviewStatus int         `gorm:"column:review_status;type:tinyint;default:0;index" json:"review_status"` // 审核状态：0 未审核 1 已审核
	ReviewTime   string      `gorm:"column:review_time;size:20" json:"review_time"`                          // 审核时间
	Reviewer     string      `gorm:"column:reviewer;size:20" json:"reviewer"`                                // 审核人 user_id
	Likes        []Like      `gorm:"foreignKey:PostId" json:"likes"`
	Images       []PostImage `gorm:"foreignKey:PostId" json:"images"`
	Comments     []Comment   `gorm:"foreignKey:PostId" json:"comments"`
}

/*审核动态*/
type CircleReview struct {
	Id           string `json:"id"`
	ReviewStatus int    `json:"review_status"` // 0 未审核 1 已审核
	Token        string `gorm:"-" json:"-"`
}

/*发布图片/视频*/
type PostImage struct {
	Id        string  `gorm:"column:id;size:20" json:"id"`
	Type      string  `gorm:"column:type;size:10" json:"type"`
	MediaType string  `gorm:"column:media_type;size:10" json:"media_type"` // image / video
	Size      string  `gorm:"column:size;size:20" json:"size"`
	Token     string  `gorm:"-" json:"token"`
	Width     int     `gorm:"column:width;type:int" json:"width"`
	Height    int     `gorm:"column:height;type:int" json:"height"`
	Duration  float64 `gorm:"column:duration;type:decimal(8,2)" json:"duration"` // 视频时长（秒）
	Time      string  `gorm:"column:time;size:20" json:"time"`
	PostId    string  `gorm:"column:post_id;size:20;index" json:"post_id"`
	Image     string  `gorm:"column:image;size:100" json:"image"`
	Thumbnail string  `gorm:"column:thumbnail;size:100" json:"thumbnail"`
}

/*字段检查*/
func (circlePost CirclePost) checkContent() error {
	var err error
	if circlePost.Content == "" {
		err = newError(400, "E_NO_CONTENT")
	} else if utf8.RuneCountInString(circlePost.Content) > 500 {
		err = newError(400, "E_TOO_LONG")
	} else if err = checkSensitiveContent(circlePost.Content); err != nil {
	} else {
		err = checkSensitiveContent(circlePost.Topic)
	}
	return err
}

/*发布圈子*/
type CirclePostOut struct {
	UserId      string `json:"user_id"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Biography   string `json:"biography"`
	FollowCount int    `json:"follow_count"`
	FansCount   int    `json:"fans_count"`
	PostCount   int    `json:"post_count"`
	IsFollowed  bool   `json:"is_followed"`

	Id           string      `json:"id"`
	Content      string      `json:"content"`
	Topic        string      `json:"topic"`
	CreateTime   string      `json:"create_time"`
	LikeCount    int         `json:"like_count"`
	CollectCount int         `json:"collect_count"`
	CommentCount int         `json:"comment_count"`
	ReviewStatus int         `json:"review_status"`
	ReviewTime   string      `json:"review_time"`
	Reviewer     string      `json:"reviewer"`
	Likes        []Like      `gorm:"foreignKey:PostId" json:"likes"`
	Images       []PostImage `gorm:"foreignKey:PostId" json:"images"`
	Comments     []Comment   `gorm:"foreignKey:PostId" json:"comments"`
}
