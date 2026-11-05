package handler

/*审核日志表*/
type AuditLog struct {
	Id         int64  `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	Object     string `gorm:"column:object;size:500" json:"object"`
	Label      string `gorm:"column:label;size:50" json:"label"`
	Service    string `gorm:"column:service;size:50" json:"service"`
	RiskLevel  string `gorm:"column:risk_level;size:20" json:"risk_level"`
	Bucket     string `gorm:"column:bucket;size:100" json:"bucket"`
	Region     string `gorm:"column:region;size:50" json:"region"`
	RequestId  string `gorm:"column:request_id;size:100" json:"request_id"`
	ScanResult string `gorm:"column:scan_result;type:text" json:"scan_result"`
	CreateTime string `gorm:"column:create_time;size:20" json:"create_time"`
}

func (AuditLog) TableName() string { return "audit_logs" }

/*OSS 内容审核回调 content 字段结构（Code+Data 包装格式）*/
type OssCallbackContent struct {
	Code      int              `json:"Code"`
	Data      *OssCallbackData `json:"Data"`
	RequestId string           `json:"RequestId"`
}

/*回调 Data 字段*/
type OssCallbackData struct {
	OssObjectName string             `json:"OssObjectName"`
	OssBucketName string             `json:"OssBucketName"`
	OssRegionId   string             `json:"OssRegionId"`
	RiskLevel     string             `json:"RiskLevel"`
	Results       []OssServiceResult `json:"Results"`
}

/*单个服务检测结果*/
type OssServiceResult struct {
	Service   string          `json:"Service"`
	RiskLevel string          `json:"RiskLevel"`
	Result    []OssResultItem `json:"Result"`
}

/*检测结果明细*/
type OssResultItem struct {
	Label       string  `json:"Label"`
	RiskLevel   string  `json:"RiskLevel"`
	Description string  `json:"Description"`
	Confidence  float64 `json:"Confidence"`
}
