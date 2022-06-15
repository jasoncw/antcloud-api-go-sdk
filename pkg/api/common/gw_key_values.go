// 
// @Author: jason zhou
// @Description: 
// @File:  gw_key_values.go
// @Version: 0.0.1
// @Date: 2022/6/7 14:50
// 

package common

import (
    "encoding/json"
    "errors"
)

//
//
//
func GwKeyValues_ToMapFromObj(obj interface{}) (map[string]string, error) {
    jsonStr, err := GwJsons_ToString(obj)
    if err != nil {
        return nil, err
    }
    return GwKeyValues_ToMapFromJsonStr(jsonStr)
}

//
//
//
func GwKeyValues_ToMapFromJsonStr(jsonStr string) (map[string]string, error) {
    if len(jsonStr) == 0 {
        return nil, errors.New("jsonStr is empty")
    }
    mp := make(map[string]string)
    err := json.Unmarshal([]byte(jsonStr), &mp)
    if err != nil {
        return nil, err
    }
    return mp, nil
}

//
//
//
func GwKeyValues_ToObjMapFromJsonStr(jsonStr string) (map[string]interface{}, error) {
    if len(jsonStr) == 0 {
        return nil, errors.New("jsonStr is empty")
    }
    mp := make(map[string]interface{})
    err := json.Unmarshal([]byte(jsonStr), &mp)
    if err != nil {
        return nil, err
    }
    return mp, nil
}
