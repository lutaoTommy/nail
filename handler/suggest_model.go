package handler

/*建议信息*/
type Suggest struct {
	Id           string `gorm:"primaryKey;column:id;size:20" json:"id"`
    Phone        string `gorm:"column:phone;size:20;" json:"phone"`
    Token        string `gorm:"-" json:"token"`
    UserId       string `gorm:"column:user_id;size:20" json:"user_id"`
    Nickname     string `gorm:"column:nickname;size:100" json:"nickname"`
    Avatar       string `gorm:"column:avatar;size:200" json:"avatar"`
    Time         string `gorm:"column:time;size:20" json:"time"`
    Content      string `gorm:"column:content;size:500;" json:"content"`
}

/*建议信息*/
type SuggestImage struct {
    Id          string   `gorm:"primaryKey;column:id;size:20" json:"id"`
    Type        string   `gorm:"column:type;size:5" json:"type"`
    Size        string   `gorm:"column:size;size:20" json:"size"`
    Token       string   `gorm:"-" json:"token"`
    Width       int      `gorm:"column:width;type:int" json:"width"`
    Height      int      `gorm:"column:height;type:int" json:"height"`
    Time        string   `gorm:"column:time;size:20" json:"time"`
    UserId      string   `gorm:"column:user_id;size:20" json:"user_id"`
    Suggest_id  string   `gorm:"column:suggest_id;size:20" json:"suggest_id"`
}

/*建议信息*/
type SuggestOut struct {
	Id           string `json:"id"`
    Phone        string `json:"phone"`
    Nickname     string `json:"nickname"`
    Avatar       string `json:"avatar"`
    Time         string `json:"time"`
    Content      string `json:"content"`
}

/*建议信息*/
type SuggestDetail struct {
	Id           string   `json:"id"`
    Phone        string   `json:"phone"`
    Nickname     string   `json:"nickname"`
    Avatar       string   `json:"avatar"`
    Time         string   `json:"time"`
    Content      string   `json:"content"`
    Attachs      []Attach `json:"attachs"`
}

/*建议信息*/
type Attach struct {
	Id           string   `json:"id"`
	Url          string   `json:"url"`
}


/*字段检查*/
func (suggest Suggest) checkContent() error {
	var err error
	if suggest.Content == "" {
		err = newError(400, "E_NO_CONTENT")
	} else if len(suggest.Content) > 500 {
		err = newError(400, "E_TOO_LONG")
	}
	return err
}
