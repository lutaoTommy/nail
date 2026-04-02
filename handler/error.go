package handler

import (
    "nail/language"
)

/*自定义错误*/
type HttpError struct {
    Code     int
    Temp     string
    Message  string
}

/*自定义错误*/
func (e HttpError) Error() string {
    return e.Message
}

/*自定义错误*/
func newErr(temp string) HttpError {
    return HttpError{
        Code: 500,
        Temp: temp,
        Message: language.GetRawMessage(temp),
    }
}

/*自定义错误*/
func newError(code int, temp string) HttpError {
    return HttpError{
        Code: code,
        Temp: temp,
        Message: language.GetRawMessage(temp),
    }
}


/*自定义错误*/
func getErrCode(err error) int {
    e, ok := err.(HttpError)
    if ok {
        return e.Code
    } else {
        return 500
    }
}





