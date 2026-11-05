package handler

import (
	"time"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

/*签到*/
func checkinHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	err = doCheckin(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
	}
}

func doCheckin(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	today := time.Now().Format("2006-01-02")
	var checkin CheckinRecord
	err := db.Where("user_id = ? AND date = ?", userInfo.UserId, today).First(&checkin).Error
	if err == nil {
		return newError(400, "E_ALREADY_CHECKIN")
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	return db.Transaction(func(tx *gorm.DB) error {
		var existing CheckinRecord
		if err := tx.Where("user_id = ? AND date = ?", userInfo.UserId, today).First(&existing).Error; err == nil {
			return newError(400, "E_ALREADY_CHECKIN")
		}
		checkin = CheckinRecord{
			Id:         RandStringBytes(3),
			UserId:     userInfo.UserId,
			Date:       today,
			CreateTime: now,
		}
		if err := tx.Create(&checkin).Error; err != nil {
			return newError(400, "E_ALREADY_CHECKIN")
		}
		record := PointRecord{
			Id:         RandStringBytes(3),
			UserId:     userInfo.UserId,
			Points:     2,
			Type:       "checkin",
			CreateTime: now,
		}
		if err := tx.Create(&record).Error; err != nil {
			return err
		}
		return nil
	})
}

/*签到状态*/
func checkinStatusHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	today := time.Now().Format("2006-01-02")
	var count int64
	db.Model(&CheckinRecord{}).Where("user_id = ? AND date = ?", userInfo.UserId, today).Count(&count)
	ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "data": iris.Map{"checked_in": count > 0}})
}

/*积分汇总*/
func pointSummaryHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	var total int64
	db.Model(&PointRecord{}).Where("user_id = ?", userInfo.UserId).Select("COALESCE(SUM(points), 0)").Scan(&total)
	ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "data": iris.Map{"total_points": total}})
}

/*积分记录列表*/
func pointListHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	params.Page = AtoUI(ctx.URLParam("page"), 0)
	params.Limit = AtoUI(ctx.URLParam("limit"), 10)
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	db = db.Model(&PointRecord{}).Where("user_id = ?", userInfo.UserId)
	err = db.Count(&params.Total).Error
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	data := []PointRecord{}
	if params.Page > 0 {
		db = db.Offset((params.Page - 1) * params.Limit).Limit(params.Limit)
	}
	err = db.Order("create_time desc").Find(&data).Error
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": data})
}
