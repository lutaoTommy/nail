package handler

/*敏感词*/
type SensitiveWord struct {
    Id       string `gorm:"primaryKey;column:id;size:20" json:"id"`
    Name     string `gorm:"column:name;size:100" json:"name"`
    Token    string `gorm:"-" json:"token"`
}

/*字段检查*/
func (word SensitiveWord) checkName() error {
    var err error
    if word.Name == "" {
        err = newError(400, "E_NO_NAME")
    } else if len(word.Name) > 100 {
        err = newError(400, "E_TOO_LONG")
    }
    return err
}

/*敏感词*/
type Word struct {
    Id       string `json:"id"`
    Name     string `json:"name"`
}