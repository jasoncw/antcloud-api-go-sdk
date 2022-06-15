// 
// @Author: jason zhou
// @Description: 
// @File:  ant_cloud_prod_client_request.go
// @Version: 0.0.1
// @Date: 2022/6/7 16:00
// 

package product

import "antcloud-api-go-sdk/pkg/api/common"

//
//
//
type AntCloudProdClientRequest struct {
    common.BaseClientRequestImpl
}

//
//
//
func NewAntCloudProdClientRequest() *AntCloudProdClientRequest {
    baseClientRequestImpl := common.NewBaseClientRequestImpl()
    return &AntCloudProdClientRequest{BaseClientRequestImpl: *baseClientRequestImpl}
}



