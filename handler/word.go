package handler

import (
	"fmt"
	"strings"
	"github.com/kataras/iris/v12"
)

/*文本管理*/
func WordHandler(word iris.Party) {
	/*敏感词添加*/
	word.Post("/add", addWordHandler)
	/*敏感词列表*/
	word.Get("/list", listWordHandler)
	/*敏感词删除*/
	word.Post("/remove", removeWordHandler)
	/*敏感词过滤*/
	word.Post("/filter", filterWordHandler)
	/*敏感词检查*/
	word.Post("/check", checkWordHandler)
}

/*初始化*/
var trie *Trie
func InitTrie() {
	db := getMysqlConn()
	var words []SensitiveWord
	db = db.Table("sensitive_words")
    err := db.Find(&words).Error
    if err != nil {
    	panic(err)
    }
	/*初始化 Trie 树*/
	trie = NewTrie()
	/*添加敏感词（支持中英文*/
	for _, word := range words {
		trie.Insert(word.Name)
	}
}


/*敏感词添加*/
func addWordHandler(ctx iris.Context) {
	var err error
	var word SensitiveWord
	word.Token = ctx.GetHeader("token")
	if err = ctx.ReadJSON(&word); err != nil {
	} else if err = word.checkName(); err != nil {
	} else if word.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	err = addWord(&word)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*敏感词添加*/
func addWord(word *SensitiveWord) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", word.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	if userInfo.UserId != ADMIN {
		return newError(403, "E_NO_AUTH")
	}
	var exist SensitiveWord
	if err := db.Where("name = ?", word.Name).First(&exist).Error; err == nil {
		return newError(400, "E_WORD_EXIST")
	}
	word.Id = RandStringBytes(2)
	if err := db.Create(word).Error; err != nil {
		return err
	}
	if trie != nil {
		trie.Insert(word.Name)
	}
	return nil
}

/*查询敏感词*/
func listWordHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	params.Name = ctx.URLParam("name")
	params.Page = AtoUI(ctx.URLParam("page"), 1)
    params.Limit = AtoUI(ctx.URLParam("limit"), 10)
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} 
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	data, err := listWord(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*查询敏感词*/
func listWord(params *Params) ([]Word, error) {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return nil, newError(401, "E_NO_TOKEN")
	}
	if userInfo.UserId != ADMIN {
		return nil, newError(403, "E_NO_AUTH")
	}
	db = db.Table("sensitive_words")
	if params.Name != "" {
		db = db.Where("name like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	var err error
	if err = db.Count(&params.Total).Error; err != nil {
		return nil, err
	}
	data := []Word{}
	db = db.Offset((params.Page - 1) * params.Limit).Limit(params.Limit)
	err = db.Find(&data).Error
	return data, err
}

/*敏感词删除*/
func removeWordHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	err = ctx.ReadJSON(&params)
	if err != nil {

	} else if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if params.Id == "" {
		err = newError(400, "E_NO_ID")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	err = removeWord(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*敏感词删除*/
func removeWord(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	if userInfo.UserId != ADMIN {
		return newError(403, "E_NO_AUTH")
	}
	var w SensitiveWord
	if err := db.Where("id = ?", params.Id).First(&w).Error; err != nil {
		return newError(404, "E_NO_WORD")
	}
	if err := db.Delete(&w).Error; err != nil {
		return err
	}
	if trie != nil {
		trie.Remove(w.Name)
	}
	return nil
}

/*敏感词过滤*/
func filterWordHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	if err = ctx.ReadJSON(&params); err != nil {
	} else if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if params.Content == "" {
		err = newError(400, "E_NO_CONTENT")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	err = filterWord(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "data": params.Content})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*敏感词过滤*/
func filterWord(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	if trie != nil {
		params.Content = trie.Filter(params.Content)
	}
	return nil
}

/*敏感词检查*/
func checkWordHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	if err = ctx.ReadJSON(&params); err != nil {
	} else if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if params.Content == "" {
		err = newError(400, "E_NO_CONTENT")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	err = checkWord(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "data": params.Content, "pass": params.Count == 0})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*敏感词检查*/
func checkWord(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	if trie != nil {
		sensitiveWord := trie.Check(params.Content)
		params.Count = len(sensitiveWord)
		params.Content = strings.Join(sensitiveWord, ",")
	} else {
		params.Count = 0
		params.Content = ""
	}
	return nil
}

