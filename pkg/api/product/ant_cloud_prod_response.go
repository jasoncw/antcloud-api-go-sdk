// 
// @Author: jason zhou
// @Description: 
// @File:  ant_cloud_prod_response.go
// @Version: 0.0.1
// @Date: 2022/6/7 14:19
// 

package product

import "antcloud-api-go-sdk/pkg/api/common"

//
//
//
type AntCloudProdResponse interface {
    GetResponseType() string
}

//
//
//
type DefAntCloudProdResponse struct {
    common.DefBaseResponse
}

//
//
//
func NewDefAntCloudProdResponse(BaseResponse common.BaseResponse) *DefAntCloudProdResponse {
    return &DefAntCloudProdResponse{}
}

//
//
//
func (_this DefAntCloudProdResponse) GetResponseType() string {
    return "DefAntCloudProdResponse"
}
