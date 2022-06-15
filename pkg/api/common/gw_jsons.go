// 
// @Author: jason zhou
// @Description: 
// @File:  gw_jsons.go
// @Version: 0.0.1
// @Date: 2022/6/7 15:00
// 

package common

import (
    "encoding/json"
    "errors"
    "github.com/jasoncw/antcloud-api-go-sdk/pkg/api/common/constants"
)

func GwJsons_ToString(obj interface{}) (string, error) {
    if obj == nil {
        return constants.EMPTY, errors.New("obj is nil")
    }
    jsonBuf, err := json.Marshal(obj)
    if err != nil {
        return constants.EMPTY, err
    }
    return string(jsonBuf), nil
}
