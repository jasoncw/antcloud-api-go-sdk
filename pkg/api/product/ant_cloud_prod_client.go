// 
// @Author: jason zhou
// @Description: 
// @File:  ant_cloud_prod_client.go
// @Version: 0.0.1
// @Date: 2022/6/7 14:00
// 

package product

import (
    "antcloud-api-go-sdk/pkg/api/common"
    "encoding/json"
    "errors"
)

//
//
//
type AntCloudProdClient struct {
    common.BaseGwClient
}

//
//
//
func NewAntCloudProdClient(baseGwClient *common.BaseGwClient) *AntCloudProdClient {
    return &AntCloudProdClient{BaseGwClient: *baseGwClient}
}

//
//
//
func (_this *AntCloudProdClient) ExecuteAntCloudProdClientRequest(antCloudProdClientRequest *AntCloudProdClientRequest) (*AntCloudProdClientResponse, error) {
    baseClientResponse, err := _this.BaseGwClient.Execute(antCloudProdClientRequest)
    if err != nil {
        return nil, err
    }
    return NewAntCloudProdClientResponse(baseClientResponse.GetData()), nil
}

//
//
//
func (_this *AntCloudProdClient) ExecuteAntCloudClientRequest(antCloudProdRequest AntCloudProdRequest) (AntCloudProdResponse, error) {
    antCloudProdClientRequest := NewAntCloudProdClientRequest()
    err := antCloudProdClientRequest.PutParametersFromObject(antCloudProdRequest)
    if err != nil {
        return nil, err
    }
    antCloudProdClientResponse, err := _this.ExecuteAntCloudProdClientRequest(antCloudProdClientRequest)
    if err != nil {
        return nil, err
    }
    antCloudProdResponse := antCloudProdRequest.NewResponse()
    if antCloudProdResponse == nil {
        return nil, errors.New("")
    }
    jsonData, err := common.GwJsons_ToString(antCloudProdClientResponse.GetData())
    if err != nil {
        return nil, err
    }
    err = json.Unmarshal([]byte(jsonData), &antCloudProdResponse)
    if err != nil {
        return nil, err
    }
    return antCloudProdResponse, err
}

type AntCloudProdClientBuilder struct {
    common.BaseGwClientBuilder
}

//
//
//
func NewAntCloudProdClientBuilder() *AntCloudProdClientBuilder {
    return &AntCloudProdClientBuilder{BaseGwClientBuilder: *common.NewBaseGwClientBuilder()}
}

//
//
//
func (_this *AntCloudProdClientBuilder) Build() (*AntCloudProdClient, error) {
    baseGwClient, err := _this.BaseGwClientBuilder.Build()
    if err != nil {
        return nil, err
    }
    return NewAntCloudProdClient(baseGwClient), nil
}
