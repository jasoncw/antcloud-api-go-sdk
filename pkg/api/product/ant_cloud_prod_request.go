// 
// @Author: jason zhou
// @Description: 
// @File:  ant_cloud_prod_request.go
// @Version: 0.0.1
// @Date: 2022/6/6 16:39
// 

package product

import "antcloud-api-go-sdk/pkg/api/common"

//
//
//
type AntCloudProdRequest interface {
    NewResponse() AntCloudProdResponse
}

//
//
//
type DefAntCloudProdRequest struct {
    common.DefBaseRequest
    ProductInstanceId string `json:"product_instance_id"`
    RegionName string `json:"region_name"`
}

//
//
//
func NewDefAntCloudProdRequest(method, version, productInstanceId string) *DefAntCloudProdRequest {
    defBaseRequest := common.NewDefBaseRequest()
    defBaseRequest.Method = method
    defBaseRequest.Version = version
    defBaseRequest.SdkVersion = "UnSet-SDK-Version"
    return &DefAntCloudProdRequest{DefBaseRequest: *defBaseRequest, ProductInstanceId: productInstanceId}
}

//
//
//
func (_this DefAntCloudProdRequest) NewResponse() AntCloudProdResponse {
    return nil
}
