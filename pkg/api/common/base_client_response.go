// 
// @Author: jason zhou
// @Description: 
// @File:  base_client_response.go
// @Version: 0.0.1
// @Date: 2022/6/14 16:00
// 

package common

import "github.com/jasoncw/antcloud-api-go-sdk/pkg/api/common/constants"

//
//
//
type BaseClientResponse interface {
    GetData() map[string]interface{}
    IsSuccess() bool
}

//
//
//
type DefBaseClientResponse struct {
    data map[string]interface{}
}

//
//
//
func NewDefBaseClientResponse(data map[string]interface{}) *DefBaseClientResponse {
    return &DefBaseClientResponse{data: data}
}

//
//
//
func (_this DefBaseClientResponse) GetData() map[string]interface{} {
    return _this.data
}

//
//
//
func (_this DefBaseClientResponse) IsSuccess() bool {
    return _this.data[constants.ParamKeys_RESULT_CODE] == constants.ResultCodes_OK
}
