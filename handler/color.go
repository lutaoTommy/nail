package handler

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"image"
	"io"
	"math"
	"nail/nlu"
	"nail/parser"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
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
	/*用户颜色足迹（历史）*/
	color.Post("/history", addColorHistoryHandler)
	color.Post("/history/remove", removeColorHistoryHandler)
	color.Get("/history/list", listColorHistoryHandler)
	color.Get("/history/group", groupColorHistoryHandler)
	/*图片颜色识别*/
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
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	data, err := descColor(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
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
	err = db.Order("no").Find(&data).Error
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
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	colorList, err := listColor(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": colorList})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
	}
}

/*查询颜色*/
func listColor(params *Params) ([]ColorOutWithFavorite, error) {
	baseDB := getMysqlConn()
	var userInfo User
	err := baseDB.Where("token = ?", params.Token).First(&userInfo).Error
	if err != nil {
		return nil, newError(401, "E_NO_TOKEN")
	}
	colorDB := baseDB.Table("colors")
	colorDB = colorDB.Where("count = ?", params.Count)
	if params.Index > 0 {
		colorDB = colorDB.Where("group_id = ?", params.Index)
	}
	if params.Name != "" {
		colorDB = colorDB.Where("name like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	err = colorDB.Count(&params.Total).Error
	if err != nil {
		return nil, err
	}
	colorList := []ColorOut{}
	if err := colorDB.Find(&colorList).Error; err != nil {
		return nil, err
	}
	if len(colorList) == 0 {
		return []ColorOutWithFavorite{}, nil
	}

	colorIds := make([]string, 0, len(colorList))
	for _, c := range colorList {
		if c.Id != "" {
			colorIds = append(colorIds, c.Id)
		}
	}
	favSet := map[string]bool{}
	if len(colorIds) > 0 {
		var favIds []string
		if err := baseDB.Model(&ColorFavorite{}).
			Where("user_id = ? AND color_id IN ?", userInfo.UserId, colorIds).
			Pluck("color_id", &favIds).Error; err != nil {
			return nil, err
		}
		favSet = make(map[string]bool, len(favIds))
		for _, id := range favIds {
			favSet[id] = true
		}
	}

	out := make([]ColorOutWithFavorite, 0, len(colorList))
	for _, c := range colorList {
		out = append(out, ColorOutWithFavorite{
			ColorOut: c,
			Favorite: favSet[c.Id],
		})
	}
	return out, nil
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
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	data, err := parseColor(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
	}
}

/*查询 lut_data 表并以数组形式返回列表（每条 data 为十六进制字符串）；可选入参 start：只返回 lut_id > start 的记录*/
func lutUrlHandler(ctx iris.Context) {
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.JSON(iris.Map{"result_code": 401, "result_msg": Msg(ctx, "E_NO_TOKEN")})
		return
	}
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", token).First(&userInfo).Error; err != nil {
		ctx.JSON(iris.Map{"result_code": 401, "result_msg": Msg(ctx, "E_NO_TOKEN")})
		return
	}
	query := db.Model(&LutData{}).Order("lut_id")
	if startStr := strings.TrimSpace(ctx.URLParam("start")); startStr != "" {
		start, err := strconv.Atoi(startStr)
		if err != nil {
			ctx.JSON(iris.Map{"result_code": 400, "result_msg": Msg(ctx, "E_INVALID_PARAM")})
			return
		}
		query = query.Where("lut_id > ?", start)
	}
	var list []LutData
	if err := query.Find(&list).Error; err != nil {
		ctx.JSON(iris.Map{"result_code": 500, "result_msg": ErrMsg(ctx, err)})
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
		ctx.JSON(iris.Map{"result_code": 401, "result_msg": Msg(ctx, "E_NO_TOKEN")})
		return
	}
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", token).First(&userInfo).Error; err != nil {
		ctx.JSON(iris.Map{"result_code": 401, "result_msg": Msg(ctx, "E_NO_TOKEN")})
		return
	}
	var maxId int
	if err := db.Model(&LutData{}).Select("COALESCE(MAX(lut_id), 0)").Scan(&maxId).Error; err != nil {
		ctx.JSON(iris.Map{"result_code": 500, "result_msg": ErrMsg(ctx, err)})
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
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	added, err := addColorFavorite(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "added": added})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
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
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	err = removeColorFavorite(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
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
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	data, err := listColorFavorite(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
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

/*用户颜色足迹（历史）：添加/置顶（body: {"id": "01"}）*/
func addColorHistoryHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	if err = ctx.ReadJSON(&params); err != nil {
	} else if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if strings.TrimSpace(params.Id) == "" {
		err = newError(400, "E_NO_ID")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	added, err := addColorHistory(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "added": added})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
	}
}

func addColorHistory(params *Params) (int, error) {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return 0, newError(401, "E_NO_TOKEN")
	}
	colorID := strings.TrimSpace(params.Id)

	// 校验颜色是否存在
	var cnt int64
	if err := db.Table("colors").Where("id = ?", colorID).Count(&cnt).Error; err != nil {
		return 0, err
	}
	if cnt == 0 {
		return 0, nil
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	day := time.Now().Format("2006-01-02")

	// 同一天已存在则更新 update_time 置顶；不同天不覆盖（新增新记录）
	res := db.Model(&ColorHistory{}).
		Where("user_id = ? AND day = ? AND color_id = ?", userInfo.UserId, day, colorID).
		Update("update_time", now)
	if res.Error != nil {
		return 0, res.Error
	}
	if res.RowsAffected > 0 {
		return 0, nil
	}
	row := ColorHistory{
		Id:         RandStringBytes(3),
		UserId:     userInfo.UserId,
		Day:        day,
		ColorId:    colorID,
		CreateTime: now,
		UpdateTime: now,
	}
	if err := db.Create(&row).Error; err != nil {
		return 0, err
	}
	return 1, nil
}

/*用户删除足迹（body: {"id": "01"}）*/
func removeColorHistoryHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	if err = ctx.ReadJSON(&params); err != nil {
	} else if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if strings.TrimSpace(params.Id) == "" {
		err = newError(400, "E_NO_ID")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	err = removeColorHistory(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
	}
}

func removeColorHistory(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	historyID := strings.TrimSpace(params.Id)
	return db.Where("user_id = ? AND id = ?", userInfo.UserId, historyID).Delete(&ColorHistory{}).Error
}

/*用户足迹列表*/
func listColorHistoryHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	params.Page = AtoUI(ctx.URLParam("page"), 0)
	params.Limit = AtoUI(ctx.URLParam("limit"), 0)
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	data, err := listColorHistory(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
	}
}

func listColorHistory(params *Params) ([]ColorHistoryOut, error) {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return nil, newError(401, "E_NO_TOKEN")
	}
	var rows []ColorHistory
	query := db.Model(&ColorHistory{}).Where("user_id = ?", userInfo.UserId).Order("update_time DESC")
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
	if err := query.Find(&rows).Error; err != nil {
		return nil, err
	}
	return buildColorHistoryOut(db, rows)
}

/*用户足迹分组：今天 / 本周（不含今天） / 上周*/
func groupColorHistoryHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	params.Limit = AtoUI(ctx.URLParam("limit"), 0)
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	data, err := groupColorHistory(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
	}
}

func groupColorHistory(params *Params) (iris.Map, error) {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return nil, newError(401, "E_NO_TOKEN")
	}

	loc := time.Local
	now := time.Now().In(loc)

	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)

	wd := int(now.Weekday())
	if wd == 0 { // Sunday
		wd = 7
	}
	startOfThisWeek := startOfDay.AddDate(0, 0, -(wd - 1)) // 周一 00:00
	startOfLastWeek := startOfThisWeek.AddDate(0, 0, -7)

	// 使用字符串比较避免在循环中解析时间
	todayStr := startOfDay.Format("2006-01-02 15:04:05")
	thisWeekStr := startOfThisWeek.Format("2006-01-02 15:04:05")
	lastWeekStr := startOfLastWeek.Format("2006-01-02 15:04:05")

	// 只查上周周一以来的数据
	var rows []ColorHistory
	if err := db.Model(&ColorHistory{}).
		Where("user_id = ? AND update_time >= ?", userInfo.UserId, lastWeekStr).
		Order("update_time DESC").
		Find(&rows).Error; err != nil {
		return nil, err
	}

	var todayRows, thisWeekRows, lastWeekRows []ColorHistory
	seenToday := make(map[string]bool)
	seenThisWeek := make(map[string]bool)
	seenLastWeek := make(map[string]bool)

	for _, r := range rows {
		if r.UpdateTime >= todayStr {
			if !seenToday[r.ColorId] {
				todayRows = append(todayRows, r)
				seenToday[r.ColorId] = true
			}
		} else if r.UpdateTime >= thisWeekStr {
			if !seenThisWeek[r.ColorId] {
				thisWeekRows = append(thisWeekRows, r)
				seenThisWeek[r.ColorId] = true
			}
		} else if r.UpdateTime >= lastWeekStr {
			if !seenLastWeek[r.ColorId] {
				lastWeekRows = append(lastWeekRows, r)
				seenLastWeek[r.ColorId] = true
			}
		}
	}

	// 每组应用 limit
	if params.Limit > 0 {
		if len(todayRows) > params.Limit {
			todayRows = todayRows[:params.Limit]
		}
		if len(thisWeekRows) > params.Limit {
			thisWeekRows = thisWeekRows[:params.Limit]
		}
		if len(lastWeekRows) > params.Limit {
			lastWeekRows = lastWeekRows[:params.Limit]
		}
	}

	// 统一获取颜色详情，避免多次查询
	allRows := append(append(todayRows, thisWeekRows...), lastWeekRows...)
	colorHistoryOut, err := buildColorHistoryOut(db, allRows)
	if err != nil {
		return nil, err
	}

	// 建立索引
	historyMap := make(map[string]ColorHistoryOut, len(colorHistoryOut))
	for _, item := range colorHistoryOut {
		historyMap[item.HistoryId] = item
	}

	// 构建结果
	return iris.Map{
		"today":     mapHistory(todayRows, historyMap),
		"this_week": mapHistory(thisWeekRows, historyMap),
		"last_week": mapHistory(lastWeekRows, historyMap),
	}, nil
}

// 辅助函数：根据 history rows 和已获取的详情 map 构建输出
func mapHistory(rows []ColorHistory, historyMap map[string]ColorHistoryOut) []ColorHistoryOut {
	res := make([]ColorHistoryOut, 0, len(rows))
	for _, r := range rows {
		if item, ok := historyMap[r.Id]; ok {
			res = append(res, item)
		}
	}
	return res
}

// 统一辅助函数：将 ColorHistory rows 转成 []ColorHistoryOut，并保持 rows 顺序
func buildColorHistoryOut(db *gorm.DB, rows []ColorHistory) ([]ColorHistoryOut, error) {
	if len(rows) == 0 {
		return []ColorHistoryOut{}, nil
	}
	// 收集去重后的颜色 ID
	colorIdSet := make(map[string]bool)
	colorIds := make([]string, 0, len(rows))
	for _, r := range rows {
		if !colorIdSet[r.ColorId] {
			colorIdSet[r.ColorId] = true
			colorIds = append(colorIds, r.ColorId)
		}
	}
	var list []ColorOut
	if err := db.Table("colors").Where("id IN ?", colorIds).Find(&list).Error; err != nil {
		return nil, err
	}
	colorByID := make(map[string]ColorOut, len(list))
	for _, c := range list {
		colorByID[c.Id] = c
	}
	result := make([]ColorHistoryOut, 0, len(rows))
	for _, r := range rows {
		if c, ok := colorByID[r.ColorId]; ok {
			result = append(result, ColorHistoryOut{
				HistoryId:  r.Id,
				UpdateTime: r.UpdateTime,
				ColorOut:   c,
			})
		}
	}
	return result, nil
}

/*初始化语音指令中的颜色关键词*/
func InitColorTriggers() {
	db := getMysqlConn()
	var colors []Color
	if err := db.Table("colors").Select("name, color").Where("count = ?", 61).Find(&colors).Error; err != nil {
		return
	}
	names := make([]string, 0, len(colors)*2)
	for _, c := range colors {
		names = append(names, c.Name)
		if c.Color != "" {
			names = append(names, c.Color)
		}
	}
	nlu.AddColorTriggers(names)
}

/*图片颜色识别：上传图片，返回图中占比最高的几种颜色 count：返回前 N 种颜色，默认 5，范围 1~20*/
func colorRecognitionHandler(ctx iris.Context) {
	token := ctx.GetHeader("token")
	if token == "" {
		err := newError(401, "E_NO_TOKEN")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", token).First(&userInfo).Error; err != nil {
		err = newError(401, "E_NO_TOKEN")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	count := AtoUI(ctx.URLParam("count"), 5)
	if count < 1 {
		count = 1
	} else if count > 20 {
		count = 20
	}
	file, _, err := ctx.FormFile("image")
	if err != nil {
		err = newError(400, "E_INVALID_PICTURE")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	defer file.Close()
	fileBytes, err := io.ReadAll(file)
	if err != nil || len(fileBytes) == 0 {
		err = newError(400, "E_INVALID_PICTURE")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	if cType := http.DetectContentType(fileBytes); cType != JPEG && cType != PNG {
		err = newError(400, "E_INVALID_PICTURE")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	img, err := parseImageData(bytes.NewBuffer(fileBytes))
	if err != nil {
		err = newError(400, "E_INVALID_PICTURE")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	/*加载 colors 表全部颜色，用于识别结果匹配最近色*/
	var palette []Color
	if err := db.Table("colors").Select("id, name, r, g, b, x, y, `desc`").Find(&palette).Error; err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	data, err := analyzeImageColors(img, count, palette)
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "data": data})
}

/*在 palette 中找 RGB 欧氏距离最近的颜色，palette 为空返回 nil*/
func matchNearestPaletteColor(r, g, b uint8, palette []Color) *Color {
	var best *Color
	var bestDist2 int64
	for i := range palette {
		c := &palette[i]
		dr := int64(r) - int64(c.R)
		dg := int64(g) - int64(c.G)
		db := int64(b) - int64(c.B)
		d2 := dr*dr + dg*dg + db*db
		if best == nil || d2 < bestDist2 {
			best = c
			bestDist2 = d2
		}
	}
	return best
}

/*
分析图片颜色：缩放到最长边 200px 采样，按每通道 4bit 量化分桶合并相近色；
匹配 colors 表（palette）最近色，同一库颜色的条目合并去重，返回占比最高的前 count 种颜色
*/
func analyzeImageColors(img image.Image, count int, palette []Color) ([]ColorRecognitionOut, error) {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	if w == 0 || h == 0 {
		return nil, newError(400, "E_INVALID_PICTURE")
	}
	/*大图缩放采样，加速统计*/
	if maxSide := max(w, h); maxSide > 200 {
		scale := maxSide / 200
		if scale < 1 {
			scale = 1
		}
		resized, err := resizeImg(img, uint(w/scale), uint(h/scale))
		if err != nil {
			return nil, err
		}
		img = resized
		bounds = img.Bounds()
	}
	/*量化分桶：key 为每通道高 4bit，桶内累加原始 RGB 与像素数*/
	type colorBucket struct {
		r, g, b uint64
		count   int
	}
	buckets := make(map[uint16]*colorBucket)
	total := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r16, g16, b16, a16 := img.At(x, y).RGBA()
			if a16 == 0 {
				continue
			}
			r, g, b := uint8(r16>>8), uint8(g16>>8), uint8(b16>>8)
			key := uint16(r>>4)<<8 | uint16(g>>4)<<4 | uint16(b>>4)
			if bucket, ok := buckets[key]; ok {
				bucket.r += uint64(r)
				bucket.g += uint64(g)
				bucket.b += uint64(b)
				bucket.count++
			} else {
				buckets[key] = &colorBucket{r: uint64(r), g: uint64(g), b: uint64(b), count: 1}
			}
			total++
		}
	}
	if total == 0 {
		return nil, newError(400, "E_INVALID_PICTURE")
	}
	/*按像素数降序排序*/
	list := make([]*colorBucket, 0, len(buckets))
	for _, bucket := range buckets {
		list = append(list, bucket)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].count > list[j].count
	})
	/*合并相近色：平均 RGB 欧氏距离不超过 32 视为同一种颜色（消除量化桶边界割裂）*/
	const similarColorDistance2 = 32 * 32
	merged := make([]*colorBucket, 0, len(list))
	for _, bucket := range list {
		r := bucket.r / uint64(bucket.count)
		g := bucket.g / uint64(bucket.count)
		b := bucket.b / uint64(bucket.count)
		isMerged := false
		for _, m := range merged {
			mr := m.r / uint64(m.count)
			mg := m.g / uint64(m.count)
			mb := m.b / uint64(m.count)
			dr, dg, db := int64(r)-int64(mr), int64(g)-int64(mg), int64(b)-int64(mb)
			if dr*dr+dg*dg+db*db <= similarColorDistance2 {
				m.r += bucket.r
				m.g += bucket.g
				m.b += bucket.b
				m.count += bucket.count
				isMerged = true
				break
			}
		}
		if !isMerged {
			merged = append(merged, bucket)
		}
	}
	/*合并后桶大小会变化，重新按像素数降序排序，保证占比高的在前*/
	sort.Slice(merged, func(i, j int) bool {
		return merged[i].count > merged[j].count
	})
	/*匹配 colors 表最近色：同一库颜色的条目合并去重（占比相加），保留首条（占比最高）的图片 RGB*/
	type matchedBucket struct {
		r, g, b uint8
		count   int
		match   *ColorRecognitionMatch
	}
	results := make([]*matchedBucket, 0, len(merged))
	for _, bucket := range merged {
		r := uint8(bucket.r / uint64(bucket.count))
		g := uint8(bucket.g / uint64(bucket.count))
		b := uint8(bucket.b / uint64(bucket.count))
		m := matchNearestPaletteColor(r, g, b, palette)
		isMerged := false
		if m != nil {
			for _, res := range results {
				if res.match != nil && res.match.Id == m.Id {
					res.count += bucket.count
					isMerged = true
					break
				}
			}
		}
		if !isMerged {
			var match *ColorRecognitionMatch
			if m != nil {
				match = &ColorRecognitionMatch{
					Id:   m.Id,
					Name: m.Name,
					X:    m.X,
					Y:    m.Y,
					R:    uint8(m.R),
					G:    uint8(m.G),
					B:    uint8(m.B),
					Desc: m.Desc,
				}
			}
			results = append(results, &matchedBucket{r: r, g: g, b: b, count: bucket.count, match: match})
		}
	}
	/*匹配合并后占比变化，再按像素数降序排一次并截断前 count 个*/
	sort.Slice(results, func(i, j int) bool {
		return results[i].count > results[j].count
	})
	if len(results) > count {
		results = results[:count]
	}
	/*占比保留 4 位小数*/
	data := make([]ColorRecognitionOut, 0, len(results))
	for _, res := range results {
		data = append(data, ColorRecognitionOut{
			R:          res.r,
			G:          res.g,
			B:          res.b,
			Color:      fmt.Sprintf("#%02X%02X%02X", res.r, res.g, res.b),
			Percentage: math.Round(float64(res.count)/float64(total)*10000) / 10000,
			Match:      res.match,
		})
	}
	return data, nil
}
