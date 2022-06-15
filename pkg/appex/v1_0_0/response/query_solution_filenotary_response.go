// 
// @Author: jason zhou
// @Description: 
// @File:  query_solution_filenotary_response.go
// @Version: 0.0.1
// @Date: 2022/6/15 16:35
// 

package response

import "github.com/jasoncw/antcloud-api-go-sdk/pkg/api/product"

//
//
//
type QuerySolutionFilenotaryResponse struct {
    product.DefAntCloudProdResponse
    FileContentHash string `json:"file_content_hash"`
    QrCodeUrl       string `json:"qr_code_url"`
    Url             string `json:"url"`
}

//
//
//
func NewQuerySolutionFilenotaryResponse() *QuerySolutionFilenotaryResponse {
    return &QuerySolutionFilenotaryResponse{}
}

//
//
//
func (_this QuerySolutionFilenotaryResponse) GetResponseType() string {
    return "QuerySolutionFilenotaryResponse"
}
