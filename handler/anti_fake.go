package handler

import (
	crypto_rand "crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

/*防伪码字符集（大写字母+数字）*/
const antiFakeCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

/*防伪码路由注册*/
func AntiFakeHandler(antiFake iris.Party) {
	/*生成防伪码*/
	antiFake.Post("/generate", antiFakeGenerateHandler)
	/*激活防伪码*/
	antiFake.Post("/activate", antiFakeActivateHandler)
	/*验证防伪码*/
	antiFake.Post("/verify", antiFakeVerifyHandler)
	/*查询防伪码列表*/
	antiFake.Get("/list", antiFakeListHandler)
	/*删除防伪码*/
	antiFake.Post("/remove", antiFakeRemoveHandler)
}

/*生成防伪码（入库）*/
func antiFakeGenerateHandler(ctx iris.Context) {
	var err error
	var params AntiFakeGenerateReq

	/*1. 解析并校验参数*/
	err = ctx.ReadJSON(&params)
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	if err = params.checkParams(); err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}

	/*2. 校验token和管理员权限*/
	db := getMysqlConn()
	token := ctx.GetHeader("token")
	if token == "" {
		err = newError(401, "E_NO_TOKEN")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	var userInfo User
	if err = db.Where("token = ?", token).First(&userInfo).Error; err != nil {
		err = newError(401, "E_NO_TOKEN")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	if userInfo.UserId != ADMIN {
		err = newError(403, "E_NO_AUTH")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}

	/*3. 生成防伪码并入库（重试机制确保唯一性）*/
	var code, randomPart string
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		code, randomPart, err = generateAntiFakeCode(params.ProductModel, params.Batch, params.Number)
		if err != nil {
			ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
			return
		}

		antiFakeCode := AntiFakeCode{
			Code:         code,
			ProductModel: params.ProductModel,
			Name:         params.Name,
			Batch:        params.Batch,
			Number:       params.Number,
			RandomPart:   randomPart,
			Status:       AntiFakeStatusInactive,
			CreatedAt:    time.Now(),
		}

		if err = db.Create(&antiFakeCode).Error; err != nil {
			if isDuplicateKeyError(err) {
				continue
			}
			ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
			return
		}
		break
	}

	if err != nil {
		err = newError(500, "E_GENERATE_CODE_FAILED")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}

	/*4. 返回结果*/
	ctx.JSON(iris.Map{
		"result_code": 200,
		"result_msg":  "success",
		"data": iris.Map{
			"anti_fake_code": code,
			"product_model":  params.ProductModel,
			"name":           params.Name,
			"batch":          params.Batch,
			"number":         params.Number,
			"random_part":    randomPart,
		},
	})
}

/*判断是否为唯一索引冲突错误*/
func isDuplicateKeyError(err error) bool {
	errMsg := err.Error()
	return strings.Contains(errMsg, "Duplicate entry") || strings.Contains(errMsg, "unique constraint") || strings.Contains(errMsg, "UNIQUE")
}

/*激活防伪码*/
func antiFakeActivateHandler(ctx iris.Context) {
	var err error
	var params AntiFakeActivateReq

	/*1. 解析并校验参数*/
	err = ctx.ReadJSON(&params)
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	if err = params.checkParams(); err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}

	/*2. 校验token和管理员权限*/
	db := getMysqlConn()
	token := ctx.GetHeader("token")
	if token == "" {
		err = newError(401, "E_NO_TOKEN")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	var userInfo User
	if err = db.Where("token = ?", token).First(&userInfo).Error; err != nil {
		err = newError(401, "E_NO_TOKEN")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	if userInfo.UserId != ADMIN {
		err = newError(403, "E_NO_AUTH")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}

	/*3. 查询防伪码*/
	var antiFakeCode AntiFakeCode
	if err = db.Where("code = ?", params.Code).First(&antiFakeCode).Error; err != nil {
		err = newError(400, "E_NO_ANTI_FAKE_CODE")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}

	/*4. 检查是否已激活*/
	if antiFakeCode.Status == AntiFakeStatusActive {
		err = newError(400, "E_ANTI_FAKE_CODE_ACTIVATED")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}

	/*5. 更新为已激活*/
	now := time.Now()
	if err = db.Model(&antiFakeCode).Updates(map[string]interface{}{
		"status":       AntiFakeStatusActive,
		"activated_by": userInfo.UserId,
		"activated_at": now,
	}).Error; err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}

	/*6. 返回结果*/
	ctx.JSON(iris.Map{
		"result_code": 200,
		"result_msg":  "success",
		"data": iris.Map{
			"anti_fake_code": params.Code,
			"status":         "activated",
		},
	})
}

/*验证防伪码*/
func antiFakeVerifyHandler(ctx iris.Context) {
	var err error
	var params AntiFakeVerifyReq

	/*1. 解析并校验参数*/
	err = ctx.ReadJSON(&params)
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	if err = params.checkParams(); err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}

	/*2. 校验token和用户*/
	db := getMysqlConn()
	token := ctx.GetHeader("token")
	if token == "" {
		err = newError(401, "E_NO_TOKEN")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	var verifier User
	if err = db.Where("token = ?", token).First(&verifier).Error; err != nil {
		err = newError(401, "E_NO_TOKEN")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}

	/*3. 查询防伪码*/
	var antiFakeCode AntiFakeCode
	if err = db.Where("code = ?", params.Code).First(&antiFakeCode).Error; err != nil {
		ctx.JSON(iris.Map{
			"result_code": 200,
			"result_msg":  "success",
			"data": iris.Map{
				"valid":          false,
				"anti_fake_code": params.Code,
				"message":        "anti_fake_code_not_found",
			},
		})
		return
	}

	/*4. 检查防伪码是否已激活*/
	if antiFakeCode.Status != AntiFakeStatusActive {
		err = newError(400, "E_ANTI_FAKE_CODE_NOT_ACTIVATED")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}

	/*5. 检查是否已有别人验证过*/
	if antiFakeCode.VerifiedBy != "" && antiFakeCode.VerifiedBy != verifier.UserId {
		err = newError(400, "E_ANTI_FAKE_CODE_VERIFIED_BY_OTHER")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}

	/*6. 记录验证信息并递增查询次数（仅第一次验证时保存验证者信息和位置）*/
	now := time.Now()
	if antiFakeCode.VerifiedBy == "" {
		db.Model(&antiFakeCode).Updates(map[string]interface{}{
			"verified_by":  verifier.UserId,
			"verified_at":  now,
			"province":     params.Province,
			"city":         params.City,
			"district":     params.District,
			"address":      params.Address,
			"latitude":     params.Latitude,
			"longitude":    params.Longitude,
			"verify_count": gorm.Expr("verify_count + 1"),
		})
	} else {
		db.Model(&antiFakeCode).Update("verify_count", gorm.Expr("verify_count + 1"))
	}

	/*7. 查询当前验证的顺序（按第一次验证时间排序，保持排名稳定）*/
	var verifyOrder int64
	verifyBaseTime := now // 首次验证用当前时间
	if antiFakeCode.VerifiedBy != "" {
		verifyBaseTime = *antiFakeCode.VerifiedAt // 非首次验证用该码的首次验证时间
	}
	db.Model(&AntiFakeCode{}).
		Where("product_model = ? AND batch = ? AND number = ? AND verified_at <= ?",
			antiFakeCode.ProductModel, antiFakeCode.Batch, antiFakeCode.Number, verifyBaseTime).
		Count(&verifyOrder)

	/*8. 统计同产品型号和批次下已激活的设备总数*/
	var activatedCount int64
	db.Model(&AntiFakeCode{}).
		Where("product_model = ? AND batch = ? AND status = ?",
			antiFakeCode.ProductModel, antiFakeCode.Batch, AntiFakeStatusActive).
		Count(&activatedCount)

	/*8. 返回结果*/
	activatedAt := ""
	if antiFakeCode.ActivatedAt != nil {
		activatedAt = antiFakeCode.ActivatedAt.Format("2006-01-02 15:04:05")
	}
	verifiedAt := ""
	if antiFakeCode.VerifiedAt != nil {
		verifiedAt = antiFakeCode.VerifiedAt.Format("2006-01-02 15:04:05")
	}

	ctx.JSON(iris.Map{
		"result_code": 200,
		"result_msg":  "success",
		"data": iris.Map{
			"valid":           true,
			"anti_fake_code":  antiFakeCode.Code,
			"name":            antiFakeCode.Name,
			"product_model":   antiFakeCode.ProductModel,
			"batch":           antiFakeCode.Batch,
			"number":          antiFakeCode.Number,
			"status":          antiFakeCode.Status,
			"activated_at":    activatedAt,
			"verified_at":     verifiedAt,
			"province":        antiFakeCode.Province,
			"city":            antiFakeCode.City,
			"district":        antiFakeCode.District,
			"address":         antiFakeCode.Address,
			"latitude":        antiFakeCode.Latitude,
			"longitude":       antiFakeCode.Longitude,
			"activated_count": activatedCount,
			"verify_count":    antiFakeCode.VerifyCount + 1,
			"verify_order":    verifyOrder,
		},
	})
}

/*查询防伪码列表*/
func antiFakeListHandler(ctx iris.Context) {
	var err error

	/*1. 校验token和管理员权限*/
	token := ctx.GetHeader("token")
	if token == "" {
		err = newError(401, "E_NO_TOKEN")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	db := getMysqlConn()
	var userInfo User
	if err = db.Where("token = ?", token).First(&userInfo).Error; err != nil {
		err = newError(401, "E_NO_TOKEN")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	if userInfo.UserId != ADMIN {
		err = newError(403, "E_NO_AUTH")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}

	/*2. 获取并校验查询参数*/
	page := AtoUI(ctx.URLParam("page"), 1)
	limit := AtoUI(ctx.URLParam("limit"), 20)
	productModel := ctx.URLParam("product_model")
	batch := ctx.URLParam("batch")
	status := ctx.URLParam("status")

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	/*3. 构建查询条件*/
	query := db.Model(&AntiFakeCode{})
	if productModel != "" {
		query = query.Where("product_model = ?", productModel)
	}
	if batch != "" {
		query = query.Where("batch = ?", batch)
	}
	if status != "" {
		query = query.Where("status = ?", AtoUI(status, 0))
	}

	/*4. 统计总数*/
	var total int64
	if err = query.Count(&total).Error; err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}

	/*5. 计算分页*/
	totalPages := int(total) / limit
	if int(total)%limit > 0 {
		totalPages++
	}
	if page > totalPages && totalPages > 0 {
		page = totalPages
	}

	/*6. 分页查询*/
	var list []AntiFakeCode
	offset := (page - 1) * limit
	if err = query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}

	/*7. 格式化返回数据*/
	dataList := make([]iris.Map, 0, len(list))
	for _, item := range list {
		activatedAt := ""
		if item.ActivatedAt != nil {
			activatedAt = item.ActivatedAt.Format("2006-01-02 15:04:05")
		}
		verifiedAt := ""
		if item.VerifiedAt != nil {
			verifiedAt = item.VerifiedAt.Format("2006-01-02 15:04:05")
		}
		dataList = append(dataList, iris.Map{
			"id":            item.Id,
			"code":          item.Code,
			"product_model": item.ProductModel,
			"name":          item.Name,
			"batch":         item.Batch,
			"number":        item.Number,
			"random_part":   item.RandomPart,
			"status":        item.Status,
			"activated_by":  item.ActivatedBy,
			"activated_at":  activatedAt,
			"verified_by":   item.VerifiedBy,
			"verified_at":   verifiedAt,
			"verify_count":  item.VerifyCount,
			"created_at":    item.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	/*8. 返回结果*/
	ctx.JSON(iris.Map{
		"result_code": 200,
		"result_msg":  "success",
		"data": iris.Map{
			"total":       total,
			"page":        page,
			"limit":       limit,
			"total_pages": totalPages,
			"list":        dataList,
		},
	})
}

/*删除防伪码（仅未激活的可删除）*/
func antiFakeRemoveHandler(ctx iris.Context) {
	var err error
	var params struct {
		Token string `json:"token"`
		Id    uint   `json:"id"`
	}

	/*1. 解析参数*/
	err = ctx.ReadJSON(&params)
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}

	/*2. 校验token和管理员权限*/
	db := getMysqlConn()
	token := ctx.GetHeader("token")
	if token == "" {
		err = newError(401, "E_NO_TOKEN")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	var userInfo User
	if err = db.Where("token = ?", token).First(&userInfo).Error; err != nil {
		err = newError(401, "E_NO_TOKEN")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	if userInfo.UserId != ADMIN {
		err = newError(403, "E_NO_AUTH")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}

	/*3. 查询防伪码*/
	var antiFakeCode AntiFakeCode
	if err = db.Where("id = ?", params.Id).First(&antiFakeCode).Error; err != nil {
		err = newError(400, "E_NO_ANTI_FAKE_CODE")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}

	/*4. 检查是否已激活*/
	if antiFakeCode.Status == AntiFakeStatusActive {
		err = newError(400, "E_ANTI_FAKE_CODE_ACTIVATED")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}

	/*5. 执行删除*/
	if err = db.Delete(&antiFakeCode).Error; err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}

	/*6. 返回结果*/
	ctx.JSON(iris.Map{
		"result_code": 200,
		"result_msg":  "success",
	})
}

/*生成24位防伪码：产品型号(2) + 规格(2) + 生产批次(6) + 随机字符串(14) = 24位*/
func generateAntiFakeCode(productModel, batch, number string) (code string, randomPart string, err error) {
	/*固定部分：10位*/
	fixedPart := productModel + number + batch

	/*随机部分：14位，使用密码学安全随机数*/
	randomPart, err = generateSecureRandomCode(14)
	if err != nil {
		return "", "", newError(500, "E_GENERATE_CODE_FAILED")
	}

	/*拼接完整防伪码：24位*/
	code = fixedPart + randomPart

	return code, randomPart, nil
}

/*生成密码学安全的随机字符串（大写字母+数字）*/
func generateSecureRandomCode(length int) (string, error) {
	result := make([]byte, length)
	charsetLen := big.NewInt(int64(len(antiFakeCharset)))

	for i := 0; i < length; i++ {
		/*使用crypto/rand生成密码学安全的随机数*/
		randomIndex, err := crypto_rand.Int(crypto_rand.Reader, charsetLen)
		if err != nil {
			return "", fmt.Errorf("failed to generate secure random: %w", err)
		}
		result[i] = antiFakeCharset[randomIndex.Int64()]
	}

	return string(result), nil
}
