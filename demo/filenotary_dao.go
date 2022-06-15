// 
// @Author: jason zhou
// @Description: 
// @File:  filenotary_sample_dao.go
// @Version: 0.0.1
// @Date: 2022/6/6 18:21
// 

package main

//
// Filenotary数据访问接口示例
//
type FilenotarySampleDao struct {
    store map[string]*Filenotary
}

//
//
//
func NewFilenotarySampleDao() *FilenotarySampleDao {
    return &FilenotarySampleDao{store: make(map[string]*Filenotary, 0)}
}

//
//
//
func (_this *FilenotarySampleDao) FindFilenotary(fileNotaryId string) (*Filenotary, error) {
    return _this.store[fileNotaryId], nil
}

//
//
//
func (_this *FilenotarySampleDao) SaveOrUpdate(filenotary *Filenotary) error {
    // 文件存证id 唯一
    _this.store[filenotary.FileNotaryId] = filenotary
    return nil
}
