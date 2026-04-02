package handler

import (
	"time"
	"gorm.io/gorm"
	"github.com/kataras/iris/v12"
)


/*关注账号*/
func followUserHandler(ctx iris.Context) {
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
	err = followUser(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*关注账号*/
func followUser(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	var following User
	if err := db.Where("user_id = ?", params.Id).First(&following).Error; err != nil {
		return newError(404, "E_INVALID_USER")
	}
	if userInfo.UserId == following.UserId {
		return newError(400, "E_NO_SELF_FOLLOW")
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	return db.Transaction(func(tx *gorm.DB) error {
		var follow Follow
		err := tx.Where("follower_id = ? AND following_id = ?", userInfo.UserId, following.UserId).First(&follow).Error
		if err == nil {
			if follow.Status == 1 {
				return newError(400, "E_FOLLOWED")
			}
			follow.Status = 1
			follow.FollowTime = now
			follow.CancelTime = ""
			if err := tx.Updates(&follow).Error; err != nil {
				return err
			}
		} else {
			follow = Follow{
				Id:          RandStringBytes(3),
				FollowerID:  userInfo.UserId,
				FollowingID: following.UserId,
				Status:      1,
				FollowTime:  now,
			}
			if err := tx.Create(&follow).Error; err != nil {
				return err
			}
		}
		incCount := gorm.Expr("COALESCE(follow_count, 0) + ?", 1)
		fansInc := gorm.Expr("COALESCE(fans_count, 0) + ?", 1)
		if err := tx.Table("users").Where("user_id = ?", userInfo.UserId).Update("follow_count", incCount).Error; err != nil {
			return err
		}
		if err := tx.Table("users").Where("user_id = ?", following.UserId).Update("fans_count", fansInc).Error; err != nil {
			return err
		}
		return nil
	})
}

/*取关账号*/
func cancelFollowHandler(ctx iris.Context) {
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
	err = cancelFollow(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*取关账号*/
func cancelFollow(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	var following User
	if err := db.Where("user_id = ?", params.Id).First(&following).Error; err != nil {
		return newError(404, "E_INVALID_USER")
	}
	var follow Follow
	if err := db.Where("follower_id = ? AND following_id = ?", userInfo.UserId, following.UserId).First(&follow).Error; err != nil {
		return newError(400, "E_NOT_FOLLOWING")
	}
	if follow.Status == -1 {
		return newError(400, "E_UNFOLLOWED")
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&follow).Updates(map[string]interface{}{
			"status":       -1,
			"cancel_time":  now,
		}).Error; err != nil {
			return err
		}
		decExpr := gorm.Expr("GREATEST(COALESCE(follow_count, 0) - ?, 0)", 1)
		fansDec := gorm.Expr("GREATEST(COALESCE(fans_count, 0) - ?, 0)", 1)
		if err := tx.Table("users").Where("user_id = ?", userInfo.UserId).Update("follow_count", decExpr).Error; err != nil {
			return err
		}
		if err := tx.Table("users").Where("user_id = ?", following.UserId).Update("fans_count", fansDec).Error; err != nil {
			return err
		}
		return nil
	})
}


/*查询关注列表*/
func userFollowListHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	params.Page = AtoUI(ctx.URLParam("page"), 0)
	params.Limit = AtoUI(ctx.URLParam("limit"), 10)
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} 
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	data, err := userFollowList(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*查询关注列表*/
func userFollowList(params *Params) ([]FollowOut, error) {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("token = ?", params.Token).First(&userInfo).Error
	if err != nil {
		return nil, newError(401, "E_NO_TOKEN")
	}
	db = db.Select("follows.id, follows.following_id, follows.follow_time, users.nickname, users.avatar, users.biography, users.follow_count, users.fans_count, users.post_count").
	Table("follows").
	Joins("left join users on follows.following_id = users.user_id").
	Where("follows.status = ?", 1).
	Where("follows.follower_id = ?", userInfo.UserId)

	err = db.Count(&params.Total).Error
	if err != nil {
		return nil, err
	}
    data := []FollowOut{}
    if params.Page > 0 {
    	db = db.Offset((params.Page - 1) * params.Limit).Limit(params.Limit)
    }
    err = db.Order("follow_time desc").Find(&data).Error
    return data, err
}
