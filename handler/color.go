package handler

import (
	"encoding/hex"
	"fmt"
	"nail/parser"
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
