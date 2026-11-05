package handler

import (
	"regexp"
	"time"
)

/*防伪码状态*/
const (
	AntiFakeStatusInactive = 0 // 未激活
	AntiFakeStatusActive   = 1 // 已激活
)

/*防伪码*/
type AntiFakeCode struct {
	Id           uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Code         string     `gorm:"column:code;size:24;uniqueIndex" json:"code"`
	ProductModel string     `gorm:"column:product_model;size:2" json:"product_model"`
	Name         string     `gorm:"column:name;size:100" json:"name"`
	Batch        string     `gorm:"column:batch;size:6" json:"batch"`
	Number       string     `gorm:"column:number;size:2" json:"number"`
	RandomPart   string     `gorm:"column:random_part;size:14" json:"random_part"`
	Status       int        `gorm:"column:status;type:tinyint;default:0" json:"status"`
	VerifyCount  int        `gorm:"column:verify_count;type:int;default:0" json:"verify_count"`
	VerifiedBy   string     `gorm:"column:verified_by;size:20" json:"verified_by"`
	VerifiedAt   *time.Time `gorm:"column:verified_at" json:"verified_at"`
	Province     string     `gorm:"column:province;size:50" json:"province"`
	City         string     `gorm:"column:city;size:50" json:"city"`
	District     string     `gorm:"column:district;size:50" json:"district"`
	Address      string     `gorm:"column:address;size:200" json:"address"`
	Latitude     float64    `gorm:"column:latitude" json:"latitude"`
	Longitude    float64    `gorm:"column:longitude" json:"longitude"`
	ActivatedBy  string     `gorm:"column:activated_by;size:20" json:"activated_by"`
	ActivatedAt  *time.Time `gorm:"column:activated_at" json:"activated_at"`
	CreatedAt    time.Time  `gorm:"column:created_at" json:"created_at"`
}

/*生成防伪码请求*/
type AntiFakeGenerateReq struct {
	Token        string `json:"token"`
	ProductModel string `json:"product_model"`
	Name         string `json:"name"`
	Batch        string `json:"batch"`
	Number       string `json:"number"`
}

/*激活防伪码请求*/
type AntiFakeActivateReq struct {
	Token string `json:"token"`
	Code  string `json:"anti_fake_code"`
}

/*验证防伪码请求*/
type AntiFakeVerifyReq struct {
	Code      string  `json:"anti_fake_code"`
	Province  string  `json:"province"`
	City      string  `json:"city"`
	District  string  `json:"district"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

/*防伪码响应*/
type AntiFakeResp struct {
	Code         string `json:"anti_fake_code"`
	ProductModel string `json:"product_model"`
	Batch        string `json:"batch"`
	Number       string `json:"number"`
	Status       int    `json:"status"`
	ActivatedAt  string `json:"activated_at,omitempty"`
}

/*字段检查-生成请求*/
func (req AntiFakeGenerateReq) checkParams() error {
	if len(req.ProductModel) != 2 {
		return newError(400, "E_INVALID_PRODUCT_MODEL")
	}
	if len(req.Batch) != 6 {
		return newError(400, "E_INVALID_BATCH")
	}
	if len(req.Number) != 2 {
		return newError(400, "E_INVALID_NUMBER")
	}
	if req.Name == "" {
		return newError(400, "E_NO_NAME")
	}
	if len(req.Name) > 100 {
		return newError(400, "E_TOO_LONG")
	}
	rgx2 := regexp.MustCompile(`^[A-Za-z0-9]{2}$`)
	rgx6 := regexp.MustCompile(`^[A-Za-z0-9]{6}$`)
	if !rgx2.MatchString(req.ProductModel) {
		return newError(400, "E_INVALID_PRODUCT_MODEL")
	}
	if !rgx6.MatchString(req.Batch) {
		return newError(400, "E_INVALID_BATCH")
	}
	if !rgx2.MatchString(req.Number) {
		return newError(400, "E_INVALID_NUMBER")
	}
	return nil
}

/*字段检查-激活请求*/
func (req AntiFakeActivateReq) checkParams() error {
	if len(req.Code) != 24 {
		return newError(400, "E_INVALID_ANTI_FAKE_CODE")
	}
	rgx := regexp.MustCompile(`^[A-Za-z0-9]{24}$`)
	if !rgx.MatchString(req.Code) {
		return newError(400, "E_INVALID_ANTI_FAKE_CODE")
	}
	return nil
}

/*字段检查-验证请求*/
func (req AntiFakeVerifyReq) checkParams() error {
	if len(req.Code) != 24 {
		return newError(400, "E_INVALID_ANTI_FAKE_CODE")
	}
	rgx := regexp.MustCompile(`^[A-Za-z0-9]{24}$`)
	if !rgx.MatchString(req.Code) {
		return newError(400, "E_INVALID_ANTI_FAKE_CODE")
	}
	return nil
}
