// 
// @Author: jason zhou
// @Description: 
// @File:  Filenotary.go
// @Version: 0.0.1
// @Date: 2022/6/6 18:19
// 

package main

const (
    StatusInitSuccess         = "初始化文件存证成功"
    StatusUploadSuccess       = "文件存证上传成功"
    StatusUploadFail          = "文件存证上传失败"
    StatusSyncSolutionSuccess = "通知文件上传完成成功"
    StatusSyncSolutionFail    = "通知文件上传完成失败"
    StatusFinish              = "存证接受并且完成"
)

//
// 存证记录
//
type Filenotary struct {
    // 主键
    Id int64
    // 原文件存储url
    OriginalFileUrl string
    // 应用Id
    AppDid string
    // 文件存证Id 唯一
    FileNotaryId string
    // 上传url
    UploadUrl string
    // 存证状态
    Status string
    // 存证是否接受
    Accepted bool
    // 存证拒绝原因
    DeniedReason string
    // 存证是否完成
    Finished bool
    // 存证哈希，存证未完成时为空
    TxHash string
    // 调用结果码
    ResultCode string
    // 调用结果描述
    ResultMsg string
}

//
//
//
func NewFilenotary() *Filenotary {
    return &Filenotary{}
}
