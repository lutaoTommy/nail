package handler

/*图片*/
type Avatar struct {
    Id          string   `gorm:"primaryKey;column:id;size:20;" json:"id"`
    Url         string   `gorm:"column:url;size:200" json:"url"`
    Type        string   `gorm:"column:type;size:5" json:"type"`
    Size        string   `gorm:"column:size;size:20" json:"size"`
    Token       string   `gorm:"-" json:"token"`
    Width       int      `gorm:"column:width;type:int" json:"width"`
    Height      int      `gorm:"column:height;type:int" json:"height"`
    Time        string   `gorm:"column:time;size:20" json:"time"`
}

/*图片*/
type AvatarOut struct {
    Id          string   `json:"id"`
    Url         string   `json:"url"`
    Time        string   `json:"time"`
}