package controllers

import (
	"encoding/json"
	"github.com/goBack/pkg/serializer"
)

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	// 处理 Validator 产生的错误
	//if ve, ok := err.(validator.ValidationErrors); ok {
	//	for _, e := range ve {
	//		return serializer.ParamErr(
	//			ParamErrorMsg(e.Field(), e.Tag()),
	//			err,
	//		)
	//	}
	//}
	//Type assertion is it a json marshall error?
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.ParamErr("JSON marshall error", err)
	}

	return serializer.ParamErr("Parameter error", err)
}
