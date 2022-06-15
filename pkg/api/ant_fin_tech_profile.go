// 
// @Author: jason zhou
// @Description: 
// @File:  AntFinTechProfile.go
// @Version: 0.0.1
// @Date: 2022/6/6 15:44
// 

package api

import (
    "antcloud-api-go-sdk/pkg/api/acapi"
)

//
//
//
type AntFinTechProfile struct {

    //
    // access key
    //
    AccessKey string

    //
    // access secret
    //
    AccessSecret string

    //
    // 基础URL，例如：https://apigw.cloud.alipay.com
    //
    BaseUrl string

    /**
     * 超时时间
     */
    TimeOutInMillis int

    /**
     * 是否开启自动重试
     */
    EnableAutoRetry bool

    /**
     * 自动重试上限
     */
    AutoRetryLimit int

    /**
     * http 配置
     */
    HttpConfig *acapi.HttpConfig

    //
    //
    //
    STSConfig *acapi.STSConfig

    //
    //
    //
    CheckSign bool
}

//
//
//
func GetAntFinTechProfile(baseUrl, accessKey, accessSecret string) *AntFinTechProfile {
    return &AntFinTechProfile{
        BaseUrl:         baseUrl,
        AccessKey:       accessKey,
        AccessSecret:    accessSecret,
        TimeOutInMillis: DEFAULT_TIMEOUT,
        HttpConfig:      acapi.NewHttpConfig(),
        STSConfig:       acapi.NewStsConfig(),
        CheckSign:       true,
    }
}
