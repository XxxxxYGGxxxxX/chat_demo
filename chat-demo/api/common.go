package api

import (
	"chat-demo/conf"
	"chat-demo/serializer"
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// validator 库做参数校验是否实用，包括错误翻译等提示
// 返回错误信息 ErrorResponse
func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		// 通过for range 遍历错误 validator.ValidationErrors断言错误信息
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return serializer.Response{
				Status: 400,
				Msg:    fmt.Sprintf("%s%s", field, tag),
				Error:  fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 400,
			Msg:    "JSON类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}

	return serializer.Response{
		Status: 400,
		Msg:    "参数错误",
		Error:  fmt.Sprint(err),
	}
}