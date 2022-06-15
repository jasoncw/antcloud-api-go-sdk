// 
// @Author: jason zhou
// @Description: 
// @File:  base_client_request.go
// @Version: 0.0.1
// @Date: 2022/6/7 14:38
// 

package common

import (
    "errors"
    "github.com/jasoncw/antcloud-api-go-sdk/pkg/api/common/constants"
)

//
//
//
type BaseClientRequest interface {
    PutParameter(key, value string) error
    RemoveParameter(key string)
    GetParameter(key string) string
    PutParameters(params map[string]string) error
    PutParametersFromObject(obj interface{}) error
    GetParameters() map[string]string
    GetMethod() string
    SetMethod(method string)
    GetVersion() string
    SetVersion(version string)
    GetReqMsgId() string
    SetReqMsgId(reqMsgId string)
}

//
//
//
type BaseClientRequestImpl struct {
    parameters map[string]string
}

//
//
//
func NewBaseClientRequestImpl() *BaseClientRequestImpl {
    return &BaseClientRequestImpl{parameters: make(map[string]string)}
}

//
//
//
func (_this *BaseClientRequestImpl) PutParameter(key, value string) error {
    if len(key) == 0 {
        return errors.New("key is empty")
    }
    _this.parameters[key] = value
    return nil
}

//
//
//
func (_this *BaseClientRequestImpl) RemoveParameter(key string) {
    delete(_this.parameters, key)
}

//
//
//
func (_this *BaseClientRequestImpl) GetParameter(key string) string {
    return _this.parameters[key]
}

//
//
//
func (_this *BaseClientRequestImpl) PutParameters(params map[string]string) error {
    if params == nil {
        return errors.New("params is nil")
    }
    for k, v := range params {
        if len(k) == 0 {
            return errors.New("empty key is not allowed")
        }
        _this.parameters[k] = v
    }
    return nil
}

//
//
//
func (_this *BaseClientRequestImpl) PutParametersFromObject(obj interface{}) error {
    params, err := GwKeyValues_ToMapFromObj(obj)
    if err != nil {
        return err
    }
    return _this.PutParameters(params)
}

//
//
//
func (_this *BaseClientRequestImpl) GetParameters() map[string]string {
    return _this.parameters
}

//
//
//
func (_this *BaseClientRequestImpl) GetMethod() string {
    return _this.GetParameter(constants.ParamKeys_METHOD)
}

//
//
//
func (_this *BaseClientRequestImpl) SetMethod(method string) {
    _this.PutParameter(constants.ParamKeys_METHOD, method)
}

//
//
//
func (_this *BaseClientRequestImpl) GetVersion() string {
    return _this.GetParameter(constants.ParamKeys_VERSION)
}

//
//
//
func (_this *BaseClientRequestImpl) SetVersion(version string) {
    _this.PutParameter(constants.ParamKeys_VERSION, version)
}

//
//
//
func (_this *BaseClientRequestImpl) GetReqMsgId() string {
    return _this.GetParameter(constants.ParamKeys_REQ_MSG_ID)
}

//
//
//
func (_this *BaseClientRequestImpl) SetReqMsgId(reqMsgId string) {
    _this.PutParameter(constants.ParamKeys_REQ_MSG_ID, reqMsgId)
}
