// 
// @Author: jason zhou
// @Description: 
// @File:  get_solution_filenotary_status_response.go
// @Version: 0.0.1
// @Date: 2022/6/15 14:33
// 

package response

import "github.com/jasoncw/antcloud-api-go-sdk/pkg/api/product"

//
//
//
type GetSolutionFilenotarystatusResponse struct {
    product.DefAntCloudProdResponse
    Accepted     bool   `json:"accepted"`
    DeniedReason string `json:"denied_reason"`
    FileNotaryId string `json:"file_notary_id"`
    Finished     bool `json:"finished"`
    TxHash       string `json:"tx_hash"`
}

//
//
//
func NewGetSolutionFilenotarystatusResponse() *GetSolutionFilenotarystatusResponse {
    return &GetSolutionFilenotarystatusResponse{}
}

//
//
//
func (_this GetSolutionFilenotarystatusResponse) GetResponseType() string {
    return "GetSolutionFilenotarystatusResponse"
}
