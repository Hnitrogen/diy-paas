package dto

import (
	"encoding/json"
	"github.com/emicklei/go-restful/v3"
)

type BaseResponse struct {
	Code    int         `json:"status_code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(response *restful.Response, code int, message string, data interface{}) {
	response.WriteHeader(code)
	res := BaseResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
	json.NewEncoder(response).Encode(res)
}
