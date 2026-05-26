package handler

import (
	"regexp"
)

/*通用入参*/
type Params struct {
	Id      string `json:"id"`
	Mac     string `json:"mac"`
	Desc    string `json:"desc"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Content string `json:"content"`
	Page    int    `json:"page"`
	Count   int    `json:"count"`
	Index   int    `json:"index"`
	Total   int64  `json:"total"`
	Limit   int    `json:"limit"`
	Token   string `json:"token"`
	UserId  string `json:"user_id"`
}

/*字段检查*/
func (params Params) checkContent() error {
	if params.Content == "" {
		return newError(400, "E_NO_CONTENT")
	}
	if len(params.Content) > 500 {
		return newError(400, "E_TOO_LONG")
	}
	if sensitiveWord := trie.Check(params.Content); len(sensitiveWord) > 0 {
		return newSensitiveWordError(sensitiveWord)
	}
	return nil
}

/*通用入参*/
type ArrParams struct {
	Id    []string `json:"id"`
	Mac   string   `json:"mac"`
	Name  string   `json:"name"`
	Page  int      `json:"page"`
	Total int64    `json:"total"`
	Limit int      `json:"limit"`
	Token string   `json:"token"`
}

/*字段检查*/
func (params ArrParams) checkMac() error {
	var err error
	rgx := regexp.MustCompile(MacReg)
	if params.Mac == "" {
		err = newError(400, "E_NO_MAC")
	} else if !rgx.MatchString(params.Mac) {
		err = newError(400, "E_INVALID_MAC")
	} else if len(params.Mac) > 20 {
		err = newError(400, "E_TOO_LONG")
	}
	return err
}

/*字段检查*/
func (params ArrParams) checkToken() error {
	var err error
	if params.Token == "" {
		err = newError(400, "E_NO_TOKEN")
	} else if len(params.Token) > 128 {
		err = newError(400, "E_TOO_LONG")
	}
	return err
}

/*字段检查*/
func (params ArrParams) checkId() error {
	var err error
	if len(params.Id) == 0 {
		err = newError(400, "E_NO_ID")
	}
	return err
}
