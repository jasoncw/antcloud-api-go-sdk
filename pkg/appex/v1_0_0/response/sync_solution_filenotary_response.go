// 
// @Author: jason zhou
// @Description: 
// @File:  sync_solution_filenotary_response.go
// @Version: 0.0.1
// @Date: 2022/6/15 14:09
// 

package response

import "antcloud-api-go-sdk/pkg/api/product"

//
//
//
type SyncSolutionFilenotaryResponse struct {
    product.DefAntCloudProdResponse
    Accepted     bool   `json:"accepted"`
    DeniedReason string `json:"denied_reason"`
    FileNotaryId string `json:"file_notary_id"`
}

//
//
//
func NewSyncSolutionFilenotaryResponse() *SyncSolutionFilenotaryResponse {
    return &SyncSolutionFilenotaryResponse{}
}

//
//
//
func (_this SyncSolutionFilenotaryResponse) GetResponseType() string {
    return "SyncSolutionFilenotaryResponse"
}
