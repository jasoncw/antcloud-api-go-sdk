// 
// @Author: jason zhou
// @Description: 
// @File:  gw_sign_type.go
// @Version: 0.0.1
// @Date: 2022/6/6 16:40
// 

package common

import "github.com/jasoncw/antcloud-api-go-sdk/pkg/api/common/constants"

type GwSignType struct {
    code string
}

func newGwSignType(code string) *GwSignType {
    return &GwSignType{code: code}
}

func (_this *GwSignType) GetCode() string {
    return _this.code
}

var (
    GwSignType_HmacSHA1   = newGwSignType(constants.SIGN_TYPE_SHA1)
    GwSignType_HmacSHA256 = newGwSignType(constants.SIGN_TYPE_SHA256)
)
