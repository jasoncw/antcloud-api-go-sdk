// 
// @Author: jason zhou
// @Description: 
// @File:  query_solution_filenotary_request.go
// @Version: 0.0.1
// @Date: 2022/6/15 16:34
// 

package request

import (
    "antcloud-api-go-sdk/pkg/api/product"
    "antcloud-api-go-sdk/pkg/appex/v1_0_0/response"
)

//
//
//
type QuerySolutionFilenotaryRequest struct {
    product.DefAntCloudProdRequest
    AppDid string `json:"app_did"`
    TxHash string `json:"tx_hash"`
}

//
//
//
func NewQuerySolutionFilenotaryRequest() *QuerySolutionFilenotaryRequest {
    defAntCloudProdRequest := product.NewDefAntCloudProdRequest("blockchain.appex.solution.filenotary.query", "1.0", "")
    defAntCloudProdRequest.SdkVersion = "Go-SDK-20220601"
    return &QuerySolutionFilenotaryRequest{DefAntCloudProdRequest: *defAntCloudProdRequest}
}

//
//
//
func (_this QuerySolutionFilenotaryRequest) NewResponse() product.AntCloudProdResponse {
    return response.NewQuerySolutionFilenotaryResponse()
}
