// 
// @Author: jason zhou
// @Description: 
// @File:  filenotary_sample_service.go
// @Version: 0.0.1
// @Date: 2022/6/15 15:16
// 

package main

import (
    "errors"
    "github.com/jasoncw/antcloud-api-go-sdk/pkg/api"
    "github.com/jasoncw/antcloud-api-go-sdk/pkg/api/common"
    "github.com/jasoncw/antcloud-api-go-sdk/pkg/appex/v1_0_0/request"
    "github.com/jasoncw/antcloud-api-go-sdk/pkg/appex/v1_0_0/response"
)

//
//
//
type FilenotarySampleService struct {
    regionName          string
    appDid              string
    productInstanceId   string
    filenotarySampleDao *FilenotarySampleDao
    antFinTechApiClient *api.AntFinTechApiClient
}

//
//
//
func NewFilenotarySampleService(baseUrl, accessKey, accessSecret, regionName, appDid, productInstanceId string) *FilenotarySampleService {
    // 初始化AntFinTechApi客户端
    profile := api.GetAntFinTechProfile(baseUrl, accessKey, accessSecret)
    antFinTechApiClient := api.NewAntFinTechApiClient(profile)

    return &FilenotarySampleService{
        regionName:          regionName,
        appDid:              appDid,
        productInstanceId:   productInstanceId,
        filenotarySampleDao: NewFilenotarySampleDao(),
        antFinTechApiClient: antFinTechApiClient,
    }
}

//
// 创建存证 包括初始化文件存证、上传文件和通知文件上传完成三个步骤
//
func (_this FilenotarySampleService) CreateFilenotary(originalFileUrl string) (*Filenotary, error) {
    // 初始化文件存证
    filenotary, err := _this.doInitSolutionFilenotaryRequest(originalFileUrl)
    if err != nil {
        return nil, err
    }
    // 上传文件
    err = _this.doUpload(filenotary)
    if err != nil {
        return nil, err
    }
    // 通知文件上传完成
    err = _this.doSyncSolutionFilenotaryRequest(filenotary)
    if err != nil {
        return nil, err
    }
    return filenotary, nil
}

//
// 初始化文件存证
//
func (_this FilenotarySampleService) doInitSolutionFilenotaryRequest(originalFileUrl string) (*Filenotary, error) {
    request := request.NewInitSolutionFilenotaryRequest()
    request.AppDid = _this.appDid
    request.ProductInstanceId = _this.productInstanceId
    request.RegionName = _this.regionName
    request.SignType = common.GwSignType_HmacSHA256.GetCode()
    antCloudRestClientResponse, err := _this.antFinTechApiClient.Execute(request)
    if err != nil {
        return nil, err
    }
    response, ok := antCloudRestClientResponse.(*response.InitSolutionFilenotaryResponse)
    if !ok {
        return nil, errors.New("转型失败," + antCloudRestClientResponse.GetResponseType())
    }
    if !response.IsSuccess() {
        return nil, errors.New(response.ResultMsg)
    }
    // 创建数据库存证记录
    filenotary := NewFilenotary()
    filenotary.AppDid = _this.appDid
    filenotary.FileNotaryId = response.FileNotaryId
    filenotary.OriginalFileUrl = originalFileUrl
    filenotary.UploadUrl = response.Url
    filenotary.Status = StatusInitSuccess
    err = _this.filenotarySampleDao.SaveOrUpdate(filenotary)
    if err != nil {
        return nil, err
    }
    return filenotary, nil
}

//
// 执行文件上传
//
func (_this FilenotarySampleService) doUpload(filenotary *Filenotary) error {
    // 执行http上传 TODO
    // 文件下载 filenotary.OriginalFileUrl
    // 文件上传 filenotary.UploadUrl
    var err error = nil
    if err != nil {
        filenotary.Status = StatusUploadFail
        filenotary.ResultMsg = err.Error()
        _this.filenotarySampleDao.SaveOrUpdate(filenotary)
        return err
    }
    filenotary.Status = StatusUploadSuccess
    filenotary.ResultMsg = ""
    return _this.filenotarySampleDao.SaveOrUpdate(filenotary)
}

//
// 通知文件上传完成
//
func (_this FilenotarySampleService) doSyncSolutionFilenotaryRequest(filenotary *Filenotary) error {
    request := request.NewSyncSolutionFilenotaryRequest()
    request.AppDid = _this.appDid
    request.ProductInstanceId = _this.productInstanceId
    request.RegionName = _this.regionName
    request.FileNotaryId = filenotary.FileNotaryId
    request.SignType = common.GwSignType_HmacSHA256.GetCode()
    antCloudRestClientResponse, err := _this.antFinTechApiClient.Execute(request)
    if err != nil {
        filenotary.Status = "通知文件上传完成失败"
        filenotary.ResultMsg = err.Error()
        _this.filenotarySampleDao.SaveOrUpdate(filenotary)
        return err
    }
    response, ok := antCloudRestClientResponse.(*response.SyncSolutionFilenotaryResponse)
    if !ok {
        filenotary.Status = StatusSyncSolutionFail
        filenotary.ResultMsg = "转型失败," + antCloudRestClientResponse.GetResponseType()
        _this.filenotarySampleDao.SaveOrUpdate(filenotary)
        return errors.New(filenotary.ResultMsg)
    }
    if !response.IsSuccess() {
        filenotary.Status = StatusSyncSolutionFail
        filenotary.ResultMsg = response.ResultMsg
        _this.filenotarySampleDao.SaveOrUpdate(filenotary)
        return errors.New(response.ResultMsg)
    }
    filenotary.Accepted = response.Accepted
    filenotary.DeniedReason = response.DeniedReason
    filenotary.ResultMsg = ""
    filenotary.Status = StatusSyncSolutionSuccess
    return _this.filenotarySampleDao.SaveOrUpdate(filenotary)
}

//
// 获取文件存证的状态
//
func (_this FilenotarySampleService) DoGetSolutionFilenotaryStatusRequest(fileNotaryId string) (*Filenotary, error) {
    request := request.NewGetSolutionFilenotaryStatusRequest()
    request.AppDid = _this.appDid
    request.ProductInstanceId = _this.productInstanceId
    request.RegionName = _this.regionName
    request.FileNotaryId = fileNotaryId
    request.SignType = common.GwSignType_HmacSHA256.GetCode()

    antCloudRestClientResponse, err := _this.antFinTechApiClient.Execute(request)
    if err != nil {
        return nil, err
    }
    response, ok := antCloudRestClientResponse.(*response.GetSolutionFilenotarystatusResponse)
    if !ok {
        return nil, errors.New("转型失败," + antCloudRestClientResponse.GetResponseType())
    }
    filenotary, err := _this.filenotarySampleDao.FindFilenotary(fileNotaryId)
    if err != nil {
        return nil, err
    }
    if filenotary == nil {
        filenotary = NewFilenotary()
        filenotary.AppDid = _this.appDid
        filenotary.FileNotaryId = fileNotaryId
    }
    if len(filenotary.TxHash) > 0 {
        return filenotary, nil
    }
    filenotary.Accepted = response.Accepted
    filenotary.DeniedReason = response.DeniedReason
    filenotary.Finished = response.Finished
    filenotary.TxHash = response.TxHash
    filenotary.ResultMsg = ""
    if response.Finished {
        filenotary.Status = StatusFinish
    }
    return filenotary, _this.filenotarySampleDao.SaveOrUpdate(filenotary)
}

//
// 查询文件存证
//
func (_this FilenotarySampleService) DoQuerySolutionFilenotaryRequest(txHash string) (*response.QuerySolutionFilenotaryResponse, error) {
    request := request.NewQuerySolutionFilenotaryRequest()
    request.AppDid = _this.appDid
    request.ProductInstanceId = _this.productInstanceId
    request.RegionName = _this.regionName
    request.TxHash = txHash
    request.SignType = common.GwSignType_HmacSHA256.GetCode()

    antCloudRestClientResponse, err := _this.antFinTechApiClient.Execute(request)
    if err != nil {
        return nil, err
    }
    response, ok := antCloudRestClientResponse.(*response.QuerySolutionFilenotaryResponse)
    if !ok {
        return nil, errors.New("转型失败," + antCloudRestClientResponse.GetResponseType())
    }
    return response, nil
}
