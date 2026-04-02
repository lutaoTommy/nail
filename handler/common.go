package handler

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

const ADMIN = "admin"
const TOURIST = "tourist"
const PNG = "image/png"
const JPEG = "image/jpeg"
const MacReg = "^([0-9A-Fa-f]{2}:){5}[0-9A-Fa-f]{2}$"
const PhoneReg = `^1([3456789][0-9]|4[4-9]|5[^4]|6[6-7]|9[189])\d{8}$`
const MailReg = `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

/*验证码随机数*/
func randInt() string {
	rand.Seed(time.Now().Unix())
	return strconv.Itoa(rand.Intn(899998) + 100000)
}

var idMu sync.Mutex
var timeCounter int
var prevTime int64
/*生成随机字符串*/
func RandStringBytes(n int) string {
	idMu.Lock()
	defer idMu.Unlock()

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	/*增加累计数，防止重复*/
	nowTime := time.Now().Unix()
	if prevTime == nowTime {
		timeCounter++
	} else {
		prevTime = nowTime
		timeCounter = 0
	}
	return fmt.Sprintf("%s%d%d", string(b), nowTime, timeCounter)
}

/*从字符串获取正整数，否则返回默认值*/
func AtoUI(str string, dint int) int {
	number, err := strconv.Atoi(str)
	if err == nil && number > 0 {
		return number
	} else {
		return dint
	}
}

/*取最小值*/
func min(args ...int) int {
	var min int
	if len(args) > 0 {
		min = args[0]
		for _, item := range args {
			if item < min {
				min = item
			}
		}
	}
	return min
}

/*是否在数组中*/
func inArr(target string, arr *[]string) int {
	for i, item := range *arr {
		if item == target {
			return i
		}
	}
	return -1
}
