// 
// @Author: jason zhou
// @Description: 
// @File:  get_solution_filenotary_status_request.go
// @Version: 0.0.1
// @Date: 2022/6/15 14:30
// 

package request

import (
    "github.com/jasoncw/antcloud-api-go-sdk/pkg/api/product"
    "github.com/jasoncw/antcloud-api-go-sdk/pkg/appex/v1_0_0/response"
)

//
//
//
type GetSolutionFilenotaryStatusRequest struct {
    product.DefAntCloudProdRequest
    AppDid       string `json:"app_did"`
    FileNotaryId string `json:"file_notary_id"`
}

//
//
//
func NewGetSolutionFilenotaryStatusRequest() *GetSolutionFilenotaryStatusRequest {
    defAntCloudProdRequest := product.NewDefAntCloudProdRequest("blockchain.appex.solution.filenotarystatus.get", "1.0", "")
    defAntCloudProdRequest.SdkVersion = "Go-SDK-20220601"
    return &GetSolutionFilenotaryStatusRequest{DefAntCloudProdRequest: *defAntCloudProdRequest}
}

//
//
//
func (_this GetSolutionFilenotaryStatusRequest) NewResponse() product.AntCloudProdResponse {
    return response.NewGetSolutionFilenotarystatusResponse()
}
