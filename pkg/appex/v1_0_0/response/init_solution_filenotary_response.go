// 
// @Author: jason zhou
// @Description: 
// @File:  init_solution_filenotary_response.go.go
// @Version: 0.0.1
// @Date: 2022/6/6 17:14
// 

package response

import (
    "github.com/jasoncw/antcloud-api-go-sdk/pkg/api/product"
)

//
// 初始化文件存证
//
type InitSolutionFilenotaryResponse struct {
    product.DefAntCloudProdResponse
    FileNotaryId string `json:"file_notary_id"`
    Url          string `json:"url"`
}

//
//
//
func NewInitSolutionFilenotaryResponse() *InitSolutionFilenotaryResponse {
    return &InitSolutionFilenotaryResponse{}
}

//
//
//
func (_this InitSolutionFilenotaryResponse) GetResponseType() string {
    return "InitSolutionFilenotaryResponse"
}
