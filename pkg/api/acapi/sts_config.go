// 
// @Author: jason zhou
// @Description: 
// @File:  sts_config.go
// @Version: 0.0.1
// @Date: 2022/6/15 10:29
// 

package acapi

//
//
//
type STSConfig struct {
    SecurityToken string
}

//
//
//
func NewStsConfig() *STSConfig  {
    return &STSConfig{}
}