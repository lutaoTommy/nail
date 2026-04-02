package handler

import (
	"fmt"
	"time"
	"gorm.io/gorm"
	"github.com/kataras/iris/v12"
)

/*设备管理*/
func DeviceHandler(device iris.Party) {
	/*设备上报*/
	device.Post("/upload", uploadDeviceHandler)
	/*投送记录*/
	device.Post("/post", devicePostHandler)
	/*投送历史*/
	device.Get("/post/history", devicePostHistoryHandler)
	/*投送历史*/
	device.Post("/post/history/remove", removeDevicePostHistoryHandler)
	/*设备别名*/
	device.Post("/rename", deviceRenameHandler)
	/*设备别名*/
	device.Get("/name/list", deviceNameListHandler)
	/*设备删除*/
	device.Post("/remove", deviceRemoveHandler)
	/*设备列表*/
	device.Get("/list", listDeviceHandler)
}

/*设备上报*/
func uploadDeviceHandler(ctx iris.Context) {
	var err error
	var params DeviceUpload
	params.Token = ctx.GetHeader("token")
	err = ctx.ReadJSON(&params)
	if err != nil {
		//nothing
	} else if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	err = uploadDevice(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*设备上报*/
func uploadDevice(params *DeviceUpload) error {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("token = ?", params.Token).First(&userInfo).Error
	if err != nil {
		return newError(401, "E_NO_TOKEN")
	} else if userInfo.UserId == TOURIST {
		return nil
	}
	now := time.Now().Format("2006-01-02 15:04:05")

	/*开启事务*/
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}

	for _, item := range params.Devices {
		if err = item.checkMac(); err != nil {
			tx.Rollback()
			return err
		} else if err = item.checkParams(); err != nil {
			tx.Rollback()
			return err
		}

		/*查找已有设备*/
		var device Device
		err = tx.Where("user_id = ? AND mac = ?", userInfo.UserId, item.Mac).First(&device).Error

		item.Time = now
		item.UserId = userInfo.UserId

		if err == nil {
			/*更新: 保持原有ID*/
			item.Id = device.Id
			if err = tx.Model(&device).Updates(item).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else if err == gorm.ErrRecordNotFound {
			/*新增*/
			item.Id = RandStringBytes(5)
			if err = tx.Create(&item).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}


/*投送记录*/
func devicePostHandler(ctx iris.Context) {
	var err error
	var params ArrParams
	params.Token = ctx.GetHeader("token")
	if err = ctx.ReadJSON(&params); err != nil {
	} else if err = params.checkMac(); err != nil {
	} else if err = params.checkToken(); err != nil {
	} else if err = params.checkId(); err != nil {
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	err = devicePost(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*投送记录*/
func devicePost(params *ArrParams) error {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("token = ?", params.Token).First(&userInfo).Error
	if err != nil {
		return newError(401, "E_NO_TOKEN")
	} else if userInfo.UserId == TOURIST {
		return nil
	}
	var device Device
	err = db.Where("user_id = ?", userInfo.UserId).Where("mac = ?", params.Mac).First(&device).Error
	if err != nil {
		return newError(404, "E_NO_DEVICE")
	}

	/*查询并建立颜色映射*/
	var colors []PostColor
	err = db.Table("colors").Where("id in ?", params.Id).Find(&colors).Error
	if err != nil {
		return err
	}
	colorMap := make(map[string]PostColor)
	for _, c := range colors {
		colorMap[c.Id] = c
	}

	/*入参检查并构建列表*/
	var finalColors []PostColor
	for i, id := range params.Id {
		if color, ok := colorMap[id]; ok {
			newColor := color
			newColor.No = i + 1
			newColor.ColorId = RandStringBytes(3)
			finalColors = append(finalColors, newColor)
		} else {
			return newError(404, "E_INVALID_COLOR")
		}
	}

	var info PostInfo
	info.Colors = finalColors
	info.Mac = device.Mac
	info.Id = RandStringBytes(3)
	info.UserId = userInfo.UserId
	info.DeviceName = device.Name
	info.Time = time.Now().Format("2006-01-02 15:04:05")
	return db.Create(&info).Error
}


/*查询换色历史*/
func devicePostHistoryHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Mac = ctx.URLParam("mac")
	params.Token = ctx.GetHeader("token")
	params.Page = AtoUI(ctx.URLParam("page"), 1)
    params.Limit = AtoUI(ctx.URLParam("limit"), 10)
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} 
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	data, err := devicePostHistory(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*查询换色历史*/
func devicePostHistory(params *Params) ([]PostInfoOut, error) {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("token = ?", params.Token).First(&userInfo).Error
	if err != nil {
		return nil, newError(401, "E_NO_TOKEN")
	}
	/*查询历史*/
	db = db.Table("post_infos")
	db = db.Where("user_id = ?", userInfo.UserId)
	if params.Mac != "" {
		db = db.Where("mac = ?", params.Mac)
	}
	err = db.Count(&params.Total).Error
	if err != nil {
		return nil, err
	}
    data := []PostInfoOut{}
    db = db.Offset((params.Page - 1) * params.Limit).Limit(params.Limit)
    err = db.Preload("Colors", func(db *gorm.DB) *gorm.DB {
    	return db.Order("no ASC")
	}).Order("time desc").Find(&data).Error
    return data, err
}

/*删除历史*/
func removeDevicePostHistoryHandler(ctx iris.Context) {
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
	err = removeDevicePostHistory(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*删除历史*/
func removeDevicePostHistory(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("token = ?", params.Token).First(&userInfo).Error
	if err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	var info PostInfo
	err = db.Where("id = ?", params.Id).First(&info).Error
	if err != nil {
		return newError(404, "E_NO_INFO")
	}
	/*先清理外联数据*/
	db.Where("post_id = ?", info.Id).Delete(&PostColor{})
	return db.Delete(info).Error
}


/*设备别名*/
func deviceRenameHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	err = ctx.ReadJSON(&params)
	if err != nil {
	} else if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if params.Name == "" {
		err = newError(400, "E_NO_NAME")
	} else if len(params.Name) > 50 {
		err = newError(400, "E_TOO_LONG")
	} else if params.Mac == "" {
		err = newError(400, "E_NO_MAC")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	err = deviceRename(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*设备别名*/
func deviceRename(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("token = ?", params.Token).First(&userInfo).Error
	if err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	var device Device
	err = db.Where("user_id = ?", userInfo.UserId).Where("mac = ?", params.Mac).First(&device).Error
	if err != nil {
		return newError(404, "E_NO_DEVICE")
	}
	device.Name = params.Name
	device.Time = time.Now().Format("2006-01-02 15:04:05")
	return db.Updates(device).Error
}

/*设备别名列表*/
func deviceNameListHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Mac = ctx.URLParam("mac")
	params.Token = ctx.GetHeader("token")
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} 
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	data, err := deviceNameList(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*设备别名列表*/
func deviceNameList(params *Params) ([]DeviceName, error) {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("token = ?", params.Token).First(&userInfo).Error
	if err != nil {
		return nil, newError(401, "E_NO_TOKEN")
	}
	db = db.Table("devices")
	db = db.Where("user_id = ?", userInfo.UserId)
	if params.Mac != "" {
		db = db.Where("mac like ?", fmt.Sprintf("%%%s%%", params.Mac))
	}
	db = db.Where("name IS NOT NULL AND name != ?", "")
	err = db.Count(&params.Total).Error
	if err != nil {
		return nil, err
	}
    data := []DeviceName{}
    err = db.Order("time desc").Find(&data).Error
    return data, err
}

/*设备删除*/
func deviceRemoveHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	err = ctx.ReadJSON(&params)
	if err != nil {

	} else if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if params.Mac == "" {
		err = newError(400, "E_NO_MAC")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	err = deviceRemove(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*设备删除*/
func deviceRemove(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("token = ?", params.Token).First(&userInfo).Error
	if err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	var device Device
	err = db.Where("user_id = ?", userInfo.UserId).Where("mac = ?", params.Mac).First(&device).Error
	if err != nil {
		return newError(404, "E_NO_DEVICE")
	}
	return db.Delete(device).Error
}

/*查询设备列表*/
func listDeviceHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	params.Page = AtoUI(ctx.URLParam("page"), 1)
    params.Limit = AtoUI(ctx.URLParam("limit"), 10)
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} 
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	data, err := listDevice(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*查询设备列表*/
func listDevice(params *Params) ([]Device, error) {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("token = ?", params.Token).First(&userInfo).Error
	if err != nil {
		return nil, newError(401, "E_NO_TOKEN")
	}
	db = db.Table("devices")
	db = db.Where("user_id = ?", userInfo.UserId)
	err = db.Count(&params.Total).Error
	if err != nil {
		return nil, err
	}
    data := []Device{}
    db = db.Offset((params.Page - 1) * params.Limit).Limit(params.Limit)
    err = db.Order("time desc").Find(&data).Error
    return data, err
}

