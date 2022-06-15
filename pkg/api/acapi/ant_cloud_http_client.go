// 
// @Author: jason zhou
// @Description: 
// @File:  ant_cloud_http_client.go
// @Version: 0.0.1
// @Date: 2022/6/7 16:50
// 

package acapi

import (
    "crypto/tls"
    "io/ioutil"
    "net"
    "net/http"
    "time"
)

//
//
//
type AntCloudHttpClient struct {
    httpConfig *HttpConfig
    httpClient *http.Client
}

// export http_proxy='http://user:password@prox-server:3128'
// export https_proxy='http://user:password@prox-server:3128'
// export HTTP_PROXY='http://user:password@prox-server:3128'
// export HTTPS_PROXY='http://user:password@prox-server:3128'
//
//
//
func NewAntCloudHttpClient(httpConfig *HttpConfig) *AntCloudHttpClient {
    dialer := &net.Dialer{
        Timeout:   time.Duration(httpConfig.ConnectionTimeoutMillis) * time.Millisecond,
        KeepAlive: time.Duration(httpConfig.KeepAliveDurationMillis) * time.Millisecond,
        DualStack: false,
    }

    transport := &http.Transport{
        TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
        Proxy:                 http.ProxyFromEnvironment, // http.ProxyURL(ProxyURL),
        DialContext:           dialer.DialContext,
        ForceAttemptHTTP2:     false,
        MaxIdleConns:          httpConfig.MaxIdleConns,
        MaxConnsPerHost:       httpConfig.MaxConnsPerHost,
        MaxIdleConnsPerHost:   httpConfig.MaxIdleConnsPerHost,
        IdleConnTimeout:       time.Duration(httpConfig.IdleConnTimeoutMillis) * time.Millisecond,
        TLSHandshakeTimeout:   10 * time.Second, // 10 * time.Second
        ExpectContinueTimeout: 1 * time.Second,  // 1 * time.Second
        // ResponseHeaderTimeout:
        DisableKeepAlives: false,
    }
    httpClient := &http.Client{
        Transport: transport,
        Timeout:   time.Duration(httpConfig.ReadTimeoutMillis) * time.Millisecond,
        CheckRedirect: func(req *http.Request, via []*http.Request) error {
            return http.ErrUseLastResponse
        },
    }
    return &AntCloudHttpClient{httpConfig: httpConfig, httpClient: httpClient}
}

//
//
//
func (_this *AntCloudHttpClient) Invoke(request *http.Request) (*HttpResponse, error) {
    resp, err := _this.httpClient.Do(request)
    if err != nil {
        return nil, err
    }
    readCloser := resp.Body
    if readCloser != nil {
        defer readCloser.Close()
    }
    body, err := ioutil.ReadAll(readCloser)
    if err != nil {
        return nil, err
    }
    return &HttpResponse{Status: resp.Status, StatusCode: resp.StatusCode, Body: body}, nil
}

//
//
//
type HttpResponse struct {
    Status     string
    StatusCode int
    Body       []byte
}
