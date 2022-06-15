// 
// @Author: jason zhou
// @Description: 
// @File:  http_config.go
// @Version: 0.0.1
// @Date: 2022/6/6 15:48
// 

package acapi

type HttpConfig struct {

    // 连接配置
    MaxIdleConns            int
    IdleConnTimeoutMillis   int
    KeepAliveDurationMillis int

    // 超时设置
    ConnectionTimeoutMillis int
    ReadTimeoutMillis       int
    WriteTimeoutMillis      int

    // 连接池配置
    // enableConnectionPools bool
    // 每个目标主机的最大空闲连接数
    MaxIdleConnsPerHost int
    // 每个目标主机的最大连接数
    MaxConnsPerHost int
}

func NewHttpConfig() *HttpConfig {
    return &HttpConfig{
        MaxIdleConns:            100,
        IdleConnTimeoutMillis:   60 * 1000,
        KeepAliveDurationMillis: 5000,
        ConnectionTimeoutMillis: 20000,
        ReadTimeoutMillis:       20000,
        WriteTimeoutMillis:      20000,
        MaxIdleConnsPerHost:     10,
        MaxConnsPerHost:         10,
    }
}

//
// func (_this *HttpConfig) GetMaxIdleConnections() int {
//     return _this.maxIdleConnections
// }
//
// func (_this *HttpConfig) SetMaxIdleConnections(maxIdleConnections int)   {
//     _this.maxIdleConnections = maxIdleConnections
// }
//
// func (_this *HttpConfig) GetKeepAliveDurationMillis() int {
//     return _this.keepAliveDurationMillis
// }
//
// func (_this *HttpConfig) SetKeepAliveDurationMillis(keepAliveDurationMillis int)   {
//       _this.keepAliveDurationMillis = keepAliveDurationMillis
// }
//
//
// func (_this *HttpConfig) GetConnectionTimeoutMillis() int {
//     return _this.connectionTimeoutMillis
// }
//
// func (_this *HttpConfig) SetConnectionTimeoutMillis(connectionTimeoutMillis int)   {
//     _this.connectionTimeoutMillis = connectionTimeoutMillis
// }
