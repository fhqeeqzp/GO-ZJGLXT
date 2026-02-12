package utils

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"context"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	SuccessCode = 200
	ErrorCode   = 500
)

func Success(data interface{}) Response {
	return Response{
		Code:    SuccessCode,
		Message: "success",
		Data:    data,
	}
}

func Error(message string) Response {
	return Response{
		Code:    ErrorCode,
		Message: message,
		Data:    nil,
	}
}

func EmitEvent(ctx context.Context, eventName string, data interface{}) {
	runtime.EventsEmit(ctx, eventName, data)
}
