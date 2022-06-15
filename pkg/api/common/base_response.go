// 
// @Author: jason zhou
// @Description: 
// @File:  base_client_request.go
// @Version: 0.0.1
// @Date: 2022/6/6 17:06
// 

package common

import "antcloud-api-go-sdk/pkg/api/common/constants"

//
//
//
type BaseResponse interface {
    IsSuccess() bool
}

//
//
//
type DefBaseResponse struct {
    Secrets []string `json:"secrets"`
    // API调用结果码，成功为OK，失败的结果码参考文档里的"结果码详情"表格
    ResultCode string `json:"result_code"`
    //  API调用结果描述，比如调用失败的时候会显示具体的错误信息
    ResultMsg string `json:"result_msg"`
    ReqMsgId  string `json:"req_msg_id"`
}

//
//
//
func NewDefBaseResponse() *DefBaseResponse {
    return &DefBaseResponse{}
}

//
//
//
func (_this *DefBaseResponse) IsSuccess() bool {
    return constants.ResultCodes_OK == _this.ResultCode
}
