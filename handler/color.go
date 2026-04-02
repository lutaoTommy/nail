package handler

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"image"
	"io"
	"nail/config"
	"nail/language"
	"nail/parser"
	"net/http"
	"strconv"
	"strings"
	"time"

	imagerecog20190930 "github.com/alibabacloud-go/imagerecog-20190930/v2/client"
	log "github.com/jeanphorn/log4go"
	"github.com/kataras/iris/v12"
)

/*颜色管理*/
func ColorHandler(color iris.Party) {
	/*查询色系*/
	color.Get("/desc", descColorHandler)
	/*查询颜色*/
	color.Get("/list", listColorHandler)
	/*根据文本解析颜色*/
	color.Get("/parse", parseColorHandler)
	/*根据 ColorDesc 的 X 返回对应 LUT 文件 HTTP 链接*/
	color.Get("/lut", lutUrlHandler)
	/*返回 lut_data 中 lut_id 的最大值*/
	color.Get("/lut/max", lutMaxIdHandler)
	/*用户收藏颜色*/
	color.Post("/favorite", addColorFavoriteHandler)
	color.Post("/favorite/remove", removeColorFavoriteHandler)
	color.Get("/favorite/list", listColorFavoriteHandler)
	/*颜色识别*/
	color.Post("/recognition", colorRecognitionHandler)
}

/*查询色系*/
func descColorHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Name = ctx.URLParam("name")
	params.Token = ctx.GetHeader("token")
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	data, err := descColor(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*查询色系*/
func descColor(params *Params) ([]ColorDesc, error) {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("token = ?", params.Token).First(&userInfo).Error
	if err != nil {
		return nil, newError(401, "E_NO_TOKEN")
	}
	db = db.Table("color_descs")
	if params.Name != "" {
		db = db.Where("name like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	err = db.Count(&params.Total).Error
	if err != nil {
		return nil, err
	}
	data := []ColorDesc{}
	err = db.Find(&data).Error
	return data, err
}

/*查询颜色*/
func listColorHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Name = ctx.URLParam("name")
	params.Token = ctx.GetHeader("token")
	params.Index = AtoUI(ctx.URLParam("index"), 0)
	params.Count = AtoUI(ctx.URLParam("count"), 61)
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	colorList, err := listColor(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": colorList})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*查询颜色*/
func listColor(params *Params) ([]ColorOut, error) {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("token = ?", params.Token).First(&userInfo).Error
	if err != nil {
		return nil, newError(401, "E_NO_TOKEN")
	}
	db = db.Table("colors")
	db = db.Where("count = ?", params.Count)
	if params.Index > 0 {
		db = db.Where("group_id = ?", params.Index)
	}
	if params.Name != "" {
		db = db.Where("name like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	err = db.Count(&params.Total).Error
	if err != nil {
		return nil, err
	}
	colorList := []ColorOut{}
	err = db.Find(&colorList).Error
	return colorList, err
}

/*根据用户输入文本解析：返回位号和颜色 id*/
func parseColorHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Name = strings.TrimSpace(ctx.URLParam("text"))
	params.Token = ctx.GetHeader("token")
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	}
	if params.Name == "" {
		err = newError(400, "E_NO_CONTENT")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	data, err := parseColor(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*查询 lut_data 表并以数组形式返回列表（每条 data 为十六进制字符串）；可选入参 start：只返回 lut_id > start 的记录*/
func lutUrlHandler(ctx iris.Context) {
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.JSON(iris.Map{"result_code": 401, "result_msg": language.GetRawMessage("E_NO_TOKEN")})
		return
	}
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", token).First(&userInfo).Error; err != nil {
		ctx.JSON(iris.Map{"result_code": 401, "result_msg": language.GetRawMessage("E_NO_TOKEN")})
		return
	}
	query := db.Model(&LutData{}).Order("lut_id")
	if startStr := strings.TrimSpace(ctx.URLParam("start")); startStr != "" {
		start, err := strconv.Atoi(startStr)
		if err != nil {
			ctx.JSON(iris.Map{"result_code": 400, "result_msg": language.GetRawMessage("E_INVALID_PARAM")})
			return
		}
		query = query.Where("lut_id > ?", start)
	}
	var list []LutData
	if err := query.Find(&list).Error; err != nil {
		ctx.JSON(iris.Map{"result_code": 500, "result_msg": err.Error()})
		return
	}
	out := make([]iris.Map, 0, len(list))
	for _, row := range list {
		out = append(out, iris.Map{"lut_id": row.LutId, "data": hex.EncodeToString(row.Data)})
	}
	ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": len(list), "data": out})
}

/*返回 lut_data 表中 lut_id 的最大值*/
func lutMaxIdHandler(ctx iris.Context) {
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.JSON(iris.Map{"result_code": 401, "result_msg": language.GetRawMessage("E_NO_TOKEN")})
		return
	}
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", token).First(&userInfo).Error; err != nil {
		ctx.JSON(iris.Map{"result_code": 401, "result_msg": language.GetRawMessage("E_NO_TOKEN")})
		return
	}
	var maxId int
	if err := db.Model(&LutData{}).Select("COALESCE(MAX(lut_id), 0)").Scan(&maxId).Error; err != nil {
		ctx.JSON(iris.Map{"result_code": 500, "result_msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "max_id": maxId})
}

/*用户收藏颜色（支持批量，body: {"id": ["01", "02", "03"]}）*/
func addColorFavoriteHandler(ctx iris.Context) {
	var err error
	var params ArrParams
	params.Token = ctx.GetHeader("token")
	if err = ctx.ReadJSON(&params); err != nil {
	} else if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if len(params.Id) == 0 {
		err = newError(400, "E_NO_ID")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	added, err := addColorFavorite(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "added": added})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

func addColorFavorite(params *ArrParams) (int, error) {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return 0, newError(401, "E_NO_TOKEN")
	}
	// 去重、去空，得到待处理的 color_id 列表
	uniqueIds := make([]string, 0, len(params.Id))
	seen := make(map[string]bool)
	for _, id := range params.Id {
		id = strings.TrimSpace(id)
		if id == "" || seen[id] {
			continue
		}
		seen[id] = true
		uniqueIds = append(uniqueIds, id)
	}
	if len(uniqueIds) == 0 {
		return 0, nil
	}
	// 批量查询存在的颜色 id
	var validColorIds []string
	if err := db.Table("colors").Where("id IN ?", uniqueIds).Pluck("id", &validColorIds).Error; err != nil {
		return 0, err
	}
	if len(validColorIds) == 0 {
		return 0, nil
	}
	// 批量查询已收藏的 (user_id, color_id)
	var existing []string
	if err := db.Model(&ColorFavorite{}).Where("user_id = ? AND color_id IN ?", userInfo.UserId, validColorIds).Pluck("color_id", &existing).Error; err != nil {
		return 0, err
	}
	existingSet := make(map[string]bool, len(existing))
	for _, id := range existing {
		existingSet[id] = true
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	var toAdd []ColorFavorite
	for _, colorId := range validColorIds {
		if existingSet[colorId] {
			continue
		}
		toAdd = append(toAdd, ColorFavorite{
			Id:         RandStringBytes(3),
			UserId:     userInfo.UserId,
			ColorId:    colorId,
			CreateTime: now,
		})
	}
	if len(toAdd) == 0 {
		return 0, nil
	}
	if err := db.Create(&toAdd).Error; err != nil {
		return 0, err
	}
	return len(toAdd), nil
}

/*用户取消收藏颜色*/
func removeColorFavoriteHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	if err = ctx.ReadJSON(&params); err != nil {
	} else if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if params.Id == "" {
		err = newError(400, "E_NO_ID")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	err = removeColorFavorite(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

func removeColorFavorite(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	return db.Where("user_id = ? AND color_id = ?", userInfo.UserId, params.Id).Delete(&ColorFavorite{}).Error
}

/*用户收藏颜色列表*/
func listColorFavoriteHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	params.Page = AtoUI(ctx.URLParam("page"), 0)
	params.Limit = AtoUI(ctx.URLParam("limit"), 0)
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	data, err := listColorFavorite(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

func listColorFavorite(params *Params) ([]ColorOut, error) {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return nil, newError(401, "E_NO_TOKEN")
	}
	var favs []ColorFavorite
	query := db.Model(&ColorFavorite{}).Where("user_id = ?", userInfo.UserId).Order("create_time DESC")
	if err := query.Count(&params.Total).Error; err != nil {
		return nil, err
	}
	if params.Limit > 0 {
		page := params.Page
		if page < 1 {
			page = 1
		}
		query = query.Offset((page - 1) * params.Limit).Limit(params.Limit)
	}
	if err := query.Find(&favs).Error; err != nil {
		return nil, err
	}
	if len(favs) == 0 {
		return []ColorOut{}, nil
	}
	colorIds := make([]string, 0, len(favs))
	for _, f := range favs {
		colorIds = append(colorIds, f.ColorId)
	}
	var list []ColorOut
	if err := db.Table("colors").Where("id IN ?", colorIds).Find(&list).Error; err != nil {
		return nil, err
	}
	// 按收藏时间顺序返回（favs 顺序），用 map 一次遍历 O(n)
	colorByID := make(map[string]ColorOut, len(list))
	for _, c := range list {
		colorByID[c.Id] = c
	}
	result := make([]ColorOut, 0, len(favs))
	for _, f := range favs {
		if c, ok := colorByID[f.ColorId]; ok {
			result = append(result, c)
		}
	}
	return result, nil
}

/*根据文本解析颜色：使用 parser 包最长匹配，返回位号和颜色 id*/
func parseColor(params *Params) ([]ColorParseItem, error) {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("token = ?", params.Token).First(&userInfo).Error
	if err != nil {
		return nil, newError(401, "E_NO_TOKEN")
	}
	var list []ColorOut
	err = db.Table("colors").Where("count = ?", 61).Find(&list).Error
	if err != nil {
		return nil, err
	}
	entries := make([]parser.ColorEntry, 0, len(list))
	for _, c := range list {
		entries = append(entries, parser.ColorEntry{
			Id:    c.Id,
			X:     c.X,
			Y:     c.Y,
			Name:  c.Name,
			Color: c.Color,
			Desc:  c.Desc,
		})
	}
	matches := parser.Parse(params.Name, entries)
	params.Total = int64(len(matches))
	result := make([]ColorParseItem, 0, len(matches))
	for _, m := range matches {
		result = append(result, ColorParseItem{
			PositionNo: m.PositionNo,
			Id:         m.Id,
			Slot:       m.Slot,
			X:          m.X,
			Y:          m.Y,
			Name:       m.Name,
			Color:      m.Color,
			Desc:       m.Desc,
		})
	}
	return result, nil
}

/*图片颜色识别*/
func colorRecognitionHandler(ctx iris.Context) {
	var cType string
	var fileBytes []byte
	var picture ColorRecog
	picture.Token = ctx.GetHeader("token")
	returnData := iris.Map{"result_code": 500}
	file, info, err := ctx.FormFile("image")
	if err != nil {
		err = newError(400, "E_INVALID_PICTURE")
	} else if c, _, e := image.DecodeConfig(file); e != nil {
		err = e
	} else {
		/*再次读取*/
		file.Seek(0, 0)
		defer file.Close()
		fileBytes, err = io.ReadAll(file)
		if err != nil {
			/*do nothing*/
		} else if picture.Token == "" {
			err = newError(401, "E_NO_TOKEN")
		} else if cType = http.DetectContentType(fileBytes); cType != JPEG && cType != PNG {
			err = newError(400, "E_INVALID_PICTURE")
		}
		picture.Width = c.Width
		picture.Height = c.Height
	}
	if err != nil {
		returnData["result_code"] = getErrCode(err)
		returnData["result_msg"] = err.Error()
	} else {
		picture.Size = strconv.FormatInt(info.Size, 10)
		picture.Type = strings.Split(cType, "/")[1]
		data, colors, err := colorRecognition(fileBytes, &picture)
		if err == nil {
			returnData["result_code"] = 200
			returnData["result_msg"] = "success"
			returnData["data"] = data.ColorTemplateList
			returnData["recommend"] = colors
		} else {
			returnData["result_code"] = getErrCode(err)
			returnData["result_msg"] = err.Error()
		}
	}
	ctx.JSON(returnData)
}

/*图片颜色识别*/
func colorRecognition(imgData []byte, picture *ColorRecog) (data *imagerecog20190930.RecognizeImageColorResponseBodyData, colors []ColorOut, err error) {
	db := getMysqlConn()
	var userInfo User
	err = db.Where("token = ?", picture.Token).First(&userInfo).Error
	if err != nil {
		return data, colors, newError(401, "E_NO_TOKEN")
	}
	/*解析图片*/
	var img image.Image
	img, err = parseImageData(bytes.NewBuffer(imgData))
	if err != nil {
		return data, colors, err
	}
	/*保存图片*/
	picture.Id = RandStringBytes(3)
	name := picture.Id + ".png"
	path := "public/picture/" + name
	err = savePngImage(img, path)
	if err != nil {
		return data, colors, err
	}
	/*保存oss*/
	err = putOssObject("ai/"+name, path)
	if err != nil {
		return data, colors, err
	}
	picture.Url = fmt.Sprintf("%s/ai/%s", config.GetOssDomain(), name)
	data, err = imageRecog(picture.Url)
	if err != nil {
		log.LOGGER("app").Debug("颜色识别失败: %s - %v", picture.Id, err)
		return data, colors, err
	}
	log.LOGGER("app").Debug("颜色识别成功: %s - %v", picture.Id, data)
	var colorList []ColorOut
	err = db.Table("colors").Find(&colorList).Error
	if err != nil {
		return data, colors, err
	}
	var ids []string
	colors = []ColorOut{}
	colors, err = colorRecommend(colorList, data)
	if err != nil {
		return data, colors, err
	}
	log.LOGGER("app").Debug("颜色推荐: %s - %v", picture.Id, colors)
	picture.Time = time.Now().Format("2006-01-02 15:04:05")
	for _, item := range colors {
		ids = append(ids, item.Id)
	}
	picture.UserId = userInfo.UserId
	picture.Color = strings.Join(ids, ",")
	err = db.Create(picture).Error
	return data, colors, err
}
