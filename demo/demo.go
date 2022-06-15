// 
// @Author: jason zhou
// @Description: 
// @File:  API.go
// @Version: 0.0.1
// @Date: 2022/6/6 16:50
// 

package main

import (
    "antcloud-api-go-sdk/pkg/api/log"
    "fmt"
)

//
// 
//
func main() {
    baseUrl := "https://openapi.antchain.antgroup.com"
    accessKey :=  "<your-access-key>"
    accessSecret := "<your-access-secret>"
    regionName := "CN-HANGZHOU-FINANCE"
    appDid := "did:mychain:fsdaf45423fsda3jgh3456jsgsdf3a9fedgfsdgf83d3a41ddvfgdfs867kjhgdf"
    productInstanceId := "appex"

    // 请使用正式的日志实现log.Logger替换StdLogger，否则日志输出到控制台
    log.InitLogger(log.NewStdLogger())

    filenotarySampleService := NewFilenotarySampleService(baseUrl, accessKey, accessSecret, regionName, appDid, productInstanceId)

    originalFileUrl := "原始文件地址"

    //  创建存证 包括初始化文件存证、上传文件和通知文件上传完成三个步骤
    filenotary, err := filenotarySampleService.CreateFilenotary(originalFileUrl)
    if err != nil {
        panic(err.Error())
    }
    // 获取文件存证的状态
    // 此处轮询查询状态接口获得异步存证是否成功的结果
    // 实际代码中请自行处理轮询逻辑和超时逻辑
    filenotary, err = filenotarySampleService.DoGetSolutionFilenotaryStatusRequest(filenotary.FileNotaryId)
    if err != nil {
        panic(err.Error())
    }
    // 查询文件存证  TxHash 存证哈希，存证未完成时为空，必须获取文件存证的状态已全部完成的情况下才能获取到
    if !filenotary.Finished || len(filenotary.TxHash) == 0 {
        panic("存证未完成")
    }
    querySolutionFilenotaryResponse, err := filenotarySampleService.DoQuerySolutionFilenotaryRequest(filenotary.TxHash)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println(filenotary.FileNotaryId + "," + querySolutionFilenotaryResponse.FileContentHash + "," + querySolutionFilenotaryResponse.QrCodeUrl)
}

