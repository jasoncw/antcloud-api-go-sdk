// 
// @Author: jason zhou
// @Description: 
// @File:  init_solution_filenotary_rest_request.go
// @Version: 0.0.1
// @Date: 2022/6/6 17:03
// 

package request

import (
    "antcloud-api-go-sdk/pkg/api/product"
    "antcloud-api-go-sdk/pkg/appex/v1_0_0/response"
)

//
// 初始化文件存证
//
type InitSolutionFilenotaryRequest struct {
    product.DefAntCloudProdRequest
    AppDid string `json:"app_did"`
}

//
//
//
func NewInitSolutionFilenotaryRequest() *InitSolutionFilenotaryRequest {
    defAntCloudProdRequest := product.NewDefAntCloudProdRequest("blockchain.appex.solution.filenotary.init", "1.0", "")
    defAntCloudProdRequest.SdkVersion = "Go-SDK-20220601"
    return &InitSolutionFilenotaryRequest{DefAntCloudProdRequest: *defAntCloudProdRequest}
}

//
//
//
func (_this InitSolutionFilenotaryRequest) NewResponse() product.AntCloudProdResponse {
    return response.NewInitSolutionFilenotaryResponse()
}
