package handler

/*颜色信息*/
type Color struct {
	Id         string `gorm:"primaryKey;column:id;size:5;" json:"id"`
	R          int    `gorm:"column:r;type:int" json:"r"`
	G          int    `gorm:"column:g;type:int" json:"g"`
	B          int    `gorm:"column:b;type:int" json:"b"`
	X          int    `gorm:"column:x;type:int" json:"x"`
	Y          int    `gorm:"column:y;type:int" json:"y"`
	Desc       string `gorm:"column:desc;size:10" json:"desc"`
	Count      int    `gorm:"column:count;type:int" json:"count"`
	Name       string `gorm:"column:name;size:30" json:"name"`
	Color      string `gorm:"column:color;size:30" json:"color"`
	GroupId    int    `gorm:"column:group_id" json:"group_id"`
	Default    bool   `gorm:"column:default;type:tinyint(1)" json:"default"` /* 是否设备内置颜色：1-62 为 true，63-79 为 false */
	CreateTime string `gorm:"column:create_time;size:20" json:"create_time"`
	UpdateTime string `gorm:"column:update_time;size:20" json:"update_time"`
}

/*颜色信息*/
type ColorOut struct {
	Id      string `json:"id"`
	R       uint8  `json:"r"`
	G       uint8  `json:"g"`
	B       uint8  `json:"b"`
	X       int    `json:"x"`
	Y       int    `json:"y"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Color   string `json:"color"`
	Default bool   `json:"default"` /* 是否设备内置颜色 */
	GroupId int    `json:"group_id"`
}

/*颜色信息*/
type ColorDesc struct {
	Id         int    `gorm:"primaryKey;column:id;" json:"id"`
	Name       string `gorm:"column:name;size:30" json:"name"`
	Color      string `gorm:"column:color;size:30" json:"color"`
	CreateTime string `gorm:"column:create_time;size:20" json:"create_time"`
}

/*根据文本解析颜色：位号 + 颜色 id + 换色位置 1-5*/
type ColorParseItem struct {
	PositionNo string `json:"position_no"` /* 位号，格式 "x,y" 表示色板坐标 */
	Id         string `json:"id"`          /* 颜色 id */
	Slot       int    `json:"slot"`        /* 换色位置 1-5，0 表示未指定 */
	X          int    `json:"x"`
	Y          int    `json:"y"`
	Name       string `json:"name"`
	Color      string `json:"color"`
	Desc       string `json:"desc"`
}

/*LUT 数据表 lut_data：lut_id 对应 lut11.H 的 11，data 为二进制内容*/
type LutData struct {
	LutId int    `gorm:"primaryKey;column:lut_id;type:int" json:"lut_id"`
	Data  []byte `gorm:"column:data;type:mediumblob" json:"data"`
}

func (LutData) TableName() string {
	return "lut_data"
}

/*用户收藏颜色*/
type ColorFavorite struct {
	Id         string `gorm:"primaryKey;column:id;size:20" json:"id"`
	UserId     string `gorm:"column:user_id;size:20;uniqueIndex:uk_user_color" json:"user_id"`
	ColorId    string `gorm:"column:color_id;size:5;uniqueIndex:uk_user_color" json:"color_id"`
	CreateTime string `gorm:"column:create_time;size:20" json:"create_time"`
}

/*图片*/
type ColorRecog struct {
	Id     string `gorm:"primaryKey;column:id;size:20;" json:"id"`
	Url    string `gorm:"column:url;size:100" json:"url"`
	Type   string `gorm:"column:type;size:5" json:"type"`
	Size   string `gorm:"column:size;size:20" json:"size"`
	Token  string `gorm:"-" json:"token"`
	Width  int    `gorm:"column:width;type:int" json:"width"`
	Height int    `gorm:"column:height;type:int" json:"height"`
	Time   string `gorm:"column:time;size:20" json:"time"`
	Color  string `gorm:"column:color;size:20" json:"color"`
	UserId string `gorm:"column:user_id;size:20" json:"user_id"`
}
