// 
// @Author: jason zhou
// @Description: 
// @File:  base_client_request.go
// @Version: 0.0.1
// @Date: 2022/6/6 17:06
// 

package common

//
//
//
type BaseRequest interface{
    GetReqMsgId() string
}


//
//
//
type DefBaseRequest struct {
    ReqMsgId   string `json:"req_msg_id"`
    Method     string `json:"method"`
    Version    string `json:"version"`
    ReqBizId   string `json:"req_biz_id,omitempty"`
    AuthToken  string `json:"auth_token,omitempty"`
    SdkVersion string `json:"sdk_version"`
    SignType   string `json:"sign_type"`
}

//
//
//
func NewDefBaseRequest() *DefBaseRequest {
    return &DefBaseRequest{}
}

func (_this *DefBaseRequest) GetReqMsgId() string {
    return _this.ReqMsgId
}