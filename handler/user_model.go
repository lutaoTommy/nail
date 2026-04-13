package handler

import (
	"gorm.io/gorm"
	"regexp"
)

/*用户信息*/
type User struct {
	Cert         string `gorm:"column:cert;size:20;" json:"cert"`
	Phone        string `gorm:"column:phone;size:20;" json:"phone"`
	Token        string `gorm:"column:token;size:128" json:"token"`
	Passwd       string `gorm:"column:passwd;size:100" json:"passwd"` //bcrypt hash(新)
	OpenID       string `gorm:"column:openid;size:30" json:"openid"`
	UserId       string `gorm:"primaryKey;column:user_id;size:20" json:"user_id"`
	Nickname     string `gorm:"column:nickname;size:100" json:"nickname"`
	Email        string `gorm:"column:email;size:100" json:"email"`
	Avatar       string `gorm:"column:avatar;size:200" json:"avatar"`
	Language     string `gorm:"column:language;size:5" json:"language"`
	Biography    string `gorm:"column:biography;size:200" json:"biography"`
	CertTime     string `gorm:"column:cert_time;size:20" json:"cert_time"`
	LoginTime    string `gorm:"column:login_time;size:20" json:"login_time"`
	RegisterTime string `gorm:"column:register_time;size:20" json:"register_time"`
	Status       int    `gorm:"column:status;type:tinyint" json:"status"`
	FollowCount  int    `gorm:"column:follow_count;type:int" json:"follow_count"` // 关注数
	FansCount    int    `gorm:"column:fans_count;type:int" json:"fans_count"`     // 粉丝数
	PostCount    int    `gorm:"column:post_count;type:int" json:"post_count"`     // 发布动态数
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

/*表名*/
func (User) TableName() string {
	return "users"
}

/*用户信息*/
type UserSimple struct {
	UserId      string `gorm:"primaryKey;column:user_id;size:20" json:"user_id"`
	Nickname    string `gorm:"column:nickname;size:100" json:"nickname"`
	Avatar      string `gorm:"column:avatar;size:200" json:"avatar"`
	Biography   string `gorm:"column:biography;size:200" json:"biography"`
	FollowCount int    `gorm:"column:follow_count;type:int" json:"follow_count"` // 关注数
	FansCount   int    `gorm:"column:fans_count;type:int" json:"fans_count"`     // 粉丝数
	PostCount   int    `gorm:"column:post_count;type:int" json:"post_count"`     // 发布动态数
}

/*表名*/
func (UserSimple) TableName() string {
	return "users"
}

/*用户信息*/
type UserOut struct {
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	UserId       string `json:"user_id"`
	Nickname     string `json:"nickname"`
	Avatar       string `json:"avatar"`
	Biography    string `json:"biography"`
	LoginTime    string `json:"login_time"`
	RegisterTime string `json:"register_time"`
	FollowCount  int    `json:"followCount"`
	FansCount    int    `json:"fansCount"`
	PostCount    int    `json:"postCount"`
}

/*字段检查*/
func (user User) checkPhone() error {
	var err error
	rgx := regexp.MustCompile(PhoneReg)
	if user.Phone == "" {
		err = newError(400, "E_NO_PHONE")
	} else if !rgx.MatchString(user.Phone) {
		err = newError(400, "E_INVALID_PHONE")
	} else if len(user.Phone) > 20 {
		err = newError(400, "E_TOO_LONG")
	}
	return err
}

/*字段检查*/
func (user User) checkMail() error {
	var err error
	rgx := regexp.MustCompile(MailReg)
	if user.Email == "" {
		err = newError(400, "E_NO_EMAIL")
	} else if !rgx.MatchString(user.Email) {
		err = newError(400, "E_INVALID_EMAIL")
	} else if len(user.Email) > 100 {
		err = newError(400, "E_TOO_LONG")
	}
	return err
}

/*字段检查*/
func (user User) checkPasswd() error {
	var err error
	if user.Passwd == "" {
		err = newError(400, "E_NO_PWD")
		// bcrypt 最大输入 72 bytes（超出会截断）
	} else if len(user.Passwd) > 72 {
		err = newError(400, "E_TOO_LONG")
	}
	return err
}

/*字段检查*/
func (user User) checkCert() error {
	var err error
	if user.Cert == "" {
		err = newError(400, "E_NO_CERT")
	} else if len(user.Cert) > 6 {
		err = newError(400, "E_TOO_LONG")
	}
	return err
}

/*字段检查*/
func (user User) checkToken() error {
	var err error
	if user.Token == "" {
		err = newError(400, "E_NO_TOKEN")
	} else if len(user.Token) > 128 {
		err = newError(400, "E_TOO_LONG")
	}
	return err
}

/*销毁账号请求*/
type UserDestory struct {
	Token string `gorm:"-" json:"token"`
}

/*字段检查*/
func (user UserDestory) checkToken() error {
	var err error
	if user.Token == "" {
		err = newError(400, "E_NO_TOKEN")
	} else if len(user.Token) > 128 {
		err = newError(400, "E_TOO_LONG")
	}
	return err
}

/*修改密码请求*/
type ChangePasswdReq struct {
	OldPasswd string `json:"old_passwd"`
	NewPasswd string `json:"new_passwd"`
}

