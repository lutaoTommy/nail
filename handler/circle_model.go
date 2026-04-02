package handler

import (
	"fmt"
	"strings"
    "nail/language"
)

/*发布圈子*/
type CirclePost struct {
	Id           string         `gorm:"primaryKey;column:id;size:20" json:"id"`
	Token        string         `gorm:"-" json:"token"`
	UserId       string         `gorm:"column:user_id;size:20" json:"user_id"`
	Content      string         `gorm:"column:content;size:500" json:"content"`
	CreateTime   string         `gorm:"column:create_time;size:20" json:"create_time"`
	UpdateTime   string         `gorm:"column:update_time;size:20" json:"update_time"`
	LikeCount    int            `gorm:"column:like_count;type:int" json:"like_count"`
	CollectCount int            `gorm:"column:collect_count;type:int" json:"collect_count"`
	CommentCount int            `gorm:"column:comment_count;type:int" json:"comment_count"`
	Status       int            `gorm:"column:status;type:tinyint" json:"status"`      // 状态：-1 删除 1 正常
	Likes        []Like         `gorm:"foreignKey:PostId" json:"likes"`
	Images       []PostImage    `gorm:"foreignKey:PostId" json:"images"`
	Comments     []Comment      `gorm:"foreignKey:PostId" json:"comments"`
}

/*发布图片*/
type PostImage struct {
	Id          string   `gorm:"column:id;size:20" json:"id"`
    Type        string   `gorm:"column:type;size:5" json:"type"`
    Size        string   `gorm:"column:size;size:20" json:"size"`
    Token       string   `gorm:"-" json:"token"`
    Width       int      `gorm:"column:width;type:int" json:"width"`
    Height      int      `gorm:"column:height;type:int" json:"height"`
    Time        string   `gorm:"column:time;size:20" json:"time"`
    PostId      string   `gorm:"column:post_id;size:20;index" json:"post_id"`
    Image       string   `gorm:"column:image;size:100" json:"image"`
    Thumbnail   string   `gorm:"column:thumbnail;size:100" json:"thumbnail"`
}

/*字段检查*/
func (circlePost CirclePost) checkContent() error {
	var err error
	if circlePost.Content == "" {
		err = newError(400, "E_NO_CONTENT")
	} else if len(circlePost.Content) > 500 {
		err = newError(400, "E_TOO_LONG")
	} else {
		sensitiveWord := trie.Check(circlePost.Content)
		if len(sensitiveWord) > 0 {
			var httpError HttpError
			httpError.Code = 400
			httpError.Message = language.GetRawMessage("E_SENSITIVEWORD")
			httpError.Message += fmt.Sprintf("(%s)", strings.Join(sensitiveWord, ","))
			return httpError
		}
	}
	return err
}


/*发布圈子*/
type CirclePostOut struct {
	UserId       string         `json:"user_id"`
	Nickname     string         `json:"nickname"`
	Avatar       string         `json:"avatar"`
	Biography    string         `json:"biography"`
	FollowCount  int            `json:"follow_count"`
	FansCount    int            `json:"fans_count"`
	PostCount    int            `json:"post_count"`

	Id           string         `json:"id"`
	Content      string         `json:"content"`
	CreateTime   string         `json:"create_time"`
	LikeCount    int            `json:"like_count"`
	CollectCount int            `json:"collect_count"`
	CommentCount int            `json:"comment_count"`
	Likes        []Like         `gorm:"foreignKey:PostId" json:"likes"`
	Images       []PostImage    `gorm:"foreignKey:PostId" json:"images"`
	Comments     []Comment      `gorm:"foreignKey:PostId" json:"comments"`
}

