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
type AntCloudProdClientResponse struct {
    common.DefBaseClientResponse
}

//
//
//
func NewAntCloudProdClientResponse(responseData map[string]interface{}) *AntCloudProdClientResponse {
    return &AntCloudProdClientResponse{
        DefBaseClientResponse: *common.NewDefBaseClientResponse(responseData),
    }
}
