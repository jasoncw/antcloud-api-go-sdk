// 
// @Author: jason zhou
// @Description: 
// @File:  sync_solution_filenotary_request.go
// @Version: 0.0.1
// @Date: 2022/6/15 14:06
// 

package request

import (
    "antcloud-api-go-sdk/pkg/api/product"
    "antcloud-api-go-sdk/pkg/appex/v1_0_0/response"
)

//
//
//
type SyncSolutionFilenotaryRequest struct {
    product.DefAntCloudProdRequest
    AppDid       string `json:"app_did"`
    FileNotaryId string `json:"file_notary_id"`
}

//
//
//
func NewSyncSolutionFilenotaryRequest() *SyncSolutionFilenotaryRequest {
    defAntCloudProdRequest := product.NewDefAntCloudProdRequest("blockchain.appex.solution.filenotary.sync", "1.0", "")
    defAntCloudProdRequest.SdkVersion = "Go-SDK-20220601"
    return &SyncSolutionFilenotaryRequest{DefAntCloudProdRequest: *defAntCloudProdRequest}
}

//
//
//
func (_this SyncSolutionFilenotaryRequest) NewResponse() product.AntCloudProdResponse {
    return response.NewSyncSolutionFilenotaryResponse()
}
