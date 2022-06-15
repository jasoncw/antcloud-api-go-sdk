// 
// @Author: jason zhou
// @Description: 
// @File:  ant_fin_tech_api_client.go
// @Version: 0.0.1
// @Date: 2022/6/6 16:05
// 

package api

import (
    "antcloud-api-go-sdk/pkg/api/acapi"
    "antcloud-api-go-sdk/pkg/api/common/constants"
    "antcloud-api-go-sdk/pkg/api/product"
    "strings"
)

//
//
type AntFinTechApiClient struct {
    antCloudProdClient *product.AntCloudProdClient
}

//
//
//
func NewAntFinTechApiClient(profile *AntFinTechProfile) *AntFinTechApiClient {
    httpConfig := profile.HttpConfig
    if httpConfig == nil {
        httpConfig = acapi.NewHttpConfig()
    }
    httpClient := acapi.NewAntCloudHttpClient(httpConfig)
    securityToken := constants.EMPTY
    if profile.STSConfig != nil {
        securityToken = profile.STSConfig.SecurityToken
    }
    baseUrl := profile.BaseUrl
    if !strings.HasSuffix(baseUrl, "/") {
        baseUrl += "/"
    }
    antCloudProdClientBuilder := product.NewAntCloudProdClientBuilder()
    antCloudProdClientBuilder.SetEndpoint(baseUrl + "gateway.do")
    antCloudProdClientBuilder.SetAccess(profile.AccessKey, profile.AccessSecret)
    antCloudProdClientBuilder.SetCheckSign(profile.CheckSign)
    antCloudProdClientBuilder.SetEnableAutoRetry(profile.EnableAutoRetry)
    antCloudProdClientBuilder.SetAutoRetryLimit(profile.AutoRetryLimit)
    antCloudProdClientBuilder.SetHttpClient(httpClient)
    antCloudProdClientBuilder.SetSecurityToken(securityToken)
    antCloudProdClient, err := antCloudProdClientBuilder.Build()
    if err != nil {
        panic(err.Error())
    }
    return &AntFinTechApiClient{antCloudProdClient: antCloudProdClient}
}

//
//
//
func (_this *AntFinTechApiClient) Execute(request product.AntCloudProdRequest) (product.AntCloudProdResponse, error) {
    return _this.antCloudProdClient.ExecuteAntCloudClientRequest(request)
}
