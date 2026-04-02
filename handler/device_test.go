package handler

import (
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
)

func newDeviceApp() *iris.Application {
	app := iris.New()
	app.PartyFunc("/device", DeviceHandler)
	return app
}

func newUserAndDeviceApp() *iris.Application {
	app := iris.New()
	app.PartyFunc("/user", UserHandler)
	app.PartyFunc("/device", DeviceHandler)
	return app
}

// /device/list：缺少 token 时应返回 401
func TestDeviceListHandler_NoToken(t *testing.T) {
	app := newDeviceApp()
	e := httptest.New(t, app)

	resp := e.GET("/device/list").
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	resp.Value("result_code").Number().Equal(401)
	resp.Value("result_msg").String().NotEmpty()
}

// 设备上报 -> 列表 -> 该名 -> 删除
func TestDeviceFlow_WithDBUser(t *testing.T) {
	u, plainPwd := prepareTestUser(t)

	app := newUserAndDeviceApp()
	e := httptest.New(t, app)

	// 1. 登录
	login := e.POST("/user/login").
		WithJSON(iris.Map{
			"phone":  u.Phone,
			"passwd": plainPwd,
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	token := login.Value("token").String().Raw()

	// 2. 设备上报 (Upload)
	mac := "00:11:22:33:44:55"
	upload := e.POST("/device/upload").
		WithHeader("token", token).
		WithJSON(iris.Map{
			"devices": []map[string]interface{}{
				{
					"mac":     mac,
					"bat":     "100",
					"name":    "Test Device",
					"rssi":    -50,
					"version": "1.0.0",
				},
			},
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	upload.Value("result_code").Number().Equal(200)

	// 3. 设备列表 (List)
	list := e.GET("/device/list").
		WithHeader("token", token).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	list.Value("result_code").Number().Equal(200)
	devices := list.Value("data").Array()
	devices.NotEmpty()
	// 验证刚才上报的设备
	found := false
	for _, val := range devices.Iter() {
		obj := val.Object()
		if obj.Value("mac").String().Raw() == mac {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("uploaded device mac=%s not found in list", mac)
	}

	// 4. 设备改名 (Rename)
	newName := "Renamed Device"
	rename := e.POST("/device/rename").
		WithHeader("token", token).
		WithJSON(iris.Map{
			"mac":  mac,
			"name": newName,
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	rename.Value("result_code").Number().Equal(200)

	// 验证改名是否生效 (Name List)
	nameList := e.GET("/device/name/list").
		WithHeader("token", token).
		WithQuery("mac", mac).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	nameList.Value("result_code").Number().Equal(200)
	nameData := nameList.Value("data").Array()
	nameData.NotEmpty()
	nameData.Element(0).Object().Value("name").String().Equal(newName)

	// 5. 投送记录 (Post) -> 依赖 Color 表，这里需要模拟一些 Color 数据或使用 try-catch
	// 为了简化，我们先跳过复杂的 Post 流程，或者先确保 Color 存在。
	// 我们可以调用 colorHandler 的逻辑或者直接插库，但这里尽量保持黑盒。
	// 由于 devicePost 需要 params.Id (color ids)，我们使用 getAnyColorID 类似逻辑
	// 但 devicePost 需要具体的颜色 ID 列表。这里暂不深度测试 Post 逻辑，
	// 除非我们引入 Color 的创建逻辑。

	// 6. 设备删除 (Remove)
	remove := e.POST("/device/remove").
		WithHeader("token", token).
		WithJSON(iris.Map{
			"mac": mac,
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	remove.Value("result_code").Number().Equal(200)

	// 再次查询列表，确认已删除
	listAfter := e.GET("/device/list").
		WithHeader("token", token).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	devicesAfter := listAfter.Value("data").Array()
	for _, val := range devicesAfter.Iter() {
		obj := val.Object()
		if obj.Value("mac").String().Raw() == mac {
			t.Errorf("device mac=%s should be removed", mac)
		}
	}
}

// 测试设备投送记录相关的逻辑 (需要 Color 数据)
func TestDevicePostFlow_WithDBUser(t *testing.T) {
	u, plainPwd := prepareTestUser(t)
	colorID := getAnyColorID(t) // 复用 color_test.go 中的 helper

	app := newUserAndDeviceApp()
	e := httptest.New(t, app)

	// Login
	login := e.POST("/user/login").
		WithJSON(iris.Map{
			"phone":  u.Phone,
			"passwd": plainPwd,
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()
	token := login.Value("token").String().Raw()

	// Upload Device first
	mac := "AA:BB:CC:DD:EE:FF"
	e.POST("/device/upload").
		WithHeader("token", token).
		WithJSON(iris.Map{
			"devices": []map[string]interface{}{{"mac": mac, "name": "PostTest"}},
		}).
		Expect().
		Status(iris.StatusOK)

	// Post
	post := e.POST("/device/post").
		WithHeader("token", token).
		WithJSON(iris.Map{
			"mac": mac,
			"id":  []string{colorID},
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	post.Value("result_code").Number().Equal(200)

	// Post History
	history := e.GET("/device/post/history").
		WithHeader("token", token).
		WithQuery("mac", mac).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	history.Value("result_code").Number().Equal(200)
	data := history.Value("data").Array()
	data.NotEmpty()
	// 验证第一条记录
	firstHistory := data.Element(0).Object()
	firstHistory.Value("mac").String().Equal(mac)

	// Remove History
	historyId := firstHistory.Value("id").String().Raw()
	removeHistory := e.POST("/device/post/history/remove").
		WithHeader("token", token).
		WithJSON(iris.Map{
			"id": historyId,
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	removeHistory.Value("result_code").Number().Equal(200)
}
