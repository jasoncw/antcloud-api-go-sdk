// 
// @Author: jason zhou
// @Description: 
// @File:  base_gw_client.go
// @Version: 0.0.1
// @Date: 2022/6/6 16:09
// 

package common

import (
    "antcloud-api-go-sdk/pkg/api/acapi"
    "antcloud-api-go-sdk/pkg/api/common/constants"
    "antcloud-api-go-sdk/pkg/api/log"
    "encoding/json"
    "errors"
    "net/http"
    "net/url"
    "strconv"
    "strings"
    "time"
)

//
//
//
type BaseGwClient struct {
    endpoint        string
    accessKey       string
    accessSecret    string
    checkSign       bool
    enableAutoRetry bool
    autoRetryLimit  int

    httpClient    *acapi.AntCloudHttpClient
    securityToken string
}

//
//
//
func NewBaseGwClient(endpoint, accessKey, accessSecret string, checkSign, enableAutoRetry bool, autoRetryLimit int,
    httpClient *acapi.AntCloudHttpClient, securityToken string) (*BaseGwClient, error) {
    if len(endpoint) == 0 {
        return nil, errors.New("endpoint is empty")
    }
    if len(accessKey) == 0 {
        return nil, errors.New("accessKey is empty")
    }
    if len(accessSecret) == 0 {
        return nil, errors.New("accessSecret is empty")
    }
    return &BaseGwClient{
        endpoint:        endpoint,
        accessKey:       accessKey,
        accessSecret:    accessSecret,
        checkSign:       checkSign,
        enableAutoRetry: enableAutoRetry,
        autoRetryLimit:  autoRetryLimit,
        httpClient:      httpClient,
        securityToken:   securityToken,
    }, nil
}

//
//
//
func (_this *BaseGwClient) buildRequest(endpoint string, request map[string]string) (*http.Request, error) {
    url0, err := url.Parse(endpoint)
    if err != nil {
        return nil, err
    }
    values := make(url.Values)
    for k, v := range request {
        values.Set(k, v)
    }
    body := values.Encode()
    req, err := http.NewRequest("POST", url0.String(), strings.NewReader(body))
    if err != nil {
        return nil, err
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
    return req, nil
}

//
//
//
func (_this *BaseGwClient) Execute(request BaseClientRequest) (BaseClientResponse, error) {
    if request == nil {
        return nil, errors.New("request cannot be nil")
    }
    if len(request.GetMethod()) == 0 {
        return nil, errors.New("method cannot be empty")
    }
    if len(request.GetVersion()) == 0 {
        return nil, errors.New("version cannot be empty")
    }
    _this.prepareParameters(request)

    return _this.sendRequest(request)
}

//
//
//
func (_this BaseGwClient) sendRequest(request BaseClientRequest) (BaseClientResponse, error) {
    retried := 0
    for {
        httpUriRequest, err := _this.buildRequest(_this.endpoint, request.GetParameters())
        if err != nil {
            log.GetLogger().Error(err.Error())
            return nil, err
        }
        httpResponse, err := _this.httpClient.Invoke(httpUriRequest)
        if err != nil {
            log.GetLogger().Error(err.Error())
            return nil, err
        }
        if httpResponse.StatusCode != 200 {
            log.GetLogger().Error(httpResponse.Status)
            return nil, errors.New("Invoke " + _this.endpoint + ", httpResponse.Status:" + httpResponse.Status)
        }
        wholeJson := make(map[string]interface{})
        err = json.Unmarshal(httpResponse.Body, &wholeJson)
        if err != nil {
            if _this.enableAutoRetry && retried < _this.autoRetryLimit {
                // 具备重试条件
                retried += 1
                log.GetLogger().Error("retry send request, retried count = " + strconv.Itoa(retried))
                continue
            } else {
                return nil, err
            }
        }
        if len(wholeJson) == 0 {
            log.GetLogger().Error(err.Error())
            return nil, errors.New("")
        }
        responseNode := wholeJson["response"]
        if responseNode == nil {
            log.GetLogger().Error(err.Error())
            return nil, errors.New("")
        }
        // 驼峰统一修正为下划线
        responseNode = keyCamel2Case(responseNode)

        jsonStr, err := GwJsons_ToString(responseNode)
        if err != nil {
            log.GetLogger().Error(err.Error())
            return nil, errors.New("")
        }
        data, err := GwKeyValues_ToObjMapFromJsonStr(jsonStr)
        response := NewDefBaseClientResponse(data)
        if response.IsSuccess() && _this.checkSign {
            // sign := wholeJson[constants.ParamKeys_SIGN]

        }
        return response, nil
    }
    return nil, nil
}

//
//
//
func keyCamel2Case(responseNode interface{}) interface{} {
    responseNodeCastMap, ok := responseNode.(map[string]interface{})
    if ok {
        tempMap := make(map[string]interface{})
        for k, v := range responseNodeCastMap {
            nk := SdkUtils_Camel2Case(k)
            tempMap[nk] = v
        }
        responseNode = tempMap
    }
    return responseNode
}

//
//
//
func (_this *BaseGwClient) putParameterIfAbsent(request BaseClientRequest, key, value string) {
    if len(request.GetParameter(key)) == 0 {
        request.PutParameter(key, value)
    }
}

//
//
//
func (_this *BaseGwClient) prepareParameters(request BaseClientRequest) {
    request.PutParameter(constants.ParamKeys_ACCESS_KEY, _this.accessKey)

    _this.putParameterIfAbsent(request, constants.ParamKeys_SIGN_TYPE, constants.DEFAULT_SIGN_TYPE)
    _this.putParameterIfAbsent(request, constants.ParamKeys_REQ_MSG_ID, SdkUtils_GenerateReqMsgId())
    _this.putParameterIfAbsent(request, constants.ParamKeys_REQ_TIME, SdkUtils_FormatDate(time.Now()))

    signType := request.GetParameter(constants.ParamKeys_SIGN_TYPE)
    if signType != constants.DEFAULT_SIGN_TYPE && signType != constants.DEFAULT_SIGN_TYPE {
        // signType非法值时, 默认使用HmacSHA1
        signType = constants.DEFAULT_SIGN_TYPE
        request.PutParameter(constants.ParamKeys_SIGN_TYPE, signType)
    }

    // STS token
    if len(_this.securityToken) > 0 {
        _this.putParameterIfAbsent(request, constants.ParamKeys_SECURITY_TOKEN, _this.securityToken)
    }

    // 基础包版本信息
    _this.putParameterIfAbsent(request, constants.ParamKeys_BASE_SDK_VERSION,
        constants.BASE_SDK_VERSION_VALUE)

    sign := GwSigns_Sign(request.GetParameters(), signType, _this.accessSecret)
    request.PutParameter(constants.ParamKeys_SIGN, sign)
}

type BaseGwClientBuilder struct {
    endpoint        string
    accessKey       string
    accessSecret    string
    checkSign       bool
    enableAutoRetry bool
    autoRetryLimit  int

    httpClient    *acapi.AntCloudHttpClient
    securityToken string
}

//
//
//
func NewBaseGwClientBuilder() *BaseGwClientBuilder {
    return &BaseGwClientBuilder{
        checkSign:       true,
        enableAutoRetry: false,
        autoRetryLimit:  3,
    }
}

//
//
//
func (_this *BaseGwClientBuilder) Build() (*BaseGwClient, error) {
    if _this.httpClient == nil {
        _this.httpClient = acapi.NewAntCloudHttpClient(acapi.NewHttpConfig())
    }
    return NewBaseGwClient(
        _this.endpoint,
        _this.accessKey,
        _this.accessSecret,
        _this.checkSign,
        _this.enableAutoRetry,
        _this.autoRetryLimit,
        _this.httpClient,
        _this.securityToken)
}

//
//
//
func (_this *BaseGwClientBuilder) SetHttpClient(httpClient *acapi.AntCloudHttpClient) *BaseGwClientBuilder {
    _this.httpClient = httpClient
    return _this
}

//
//
//
func (_this *BaseGwClientBuilder) SetEndpoint(endpoint string) *BaseGwClientBuilder {
    _this.endpoint = endpoint
    return _this
}

//
//
//
func (_this *BaseGwClientBuilder) SetAccess(accessKey, accessSecret string) *BaseGwClientBuilder {
    _this.accessKey = accessKey
    _this.accessSecret = accessSecret
    return _this
}

//
//
//
func (_this *BaseGwClientBuilder) SetCheckSign(checkSign bool) *BaseGwClientBuilder {
    _this.checkSign = checkSign
    return _this
}

//
//
//
func (_this *BaseGwClientBuilder) SetEnableAutoRetry(enableAutoRetry bool) *BaseGwClientBuilder {
    _this.enableAutoRetry = enableAutoRetry
    return _this
}

//
//
//
func (_this *BaseGwClientBuilder) SetAutoRetryLimit(autoRetryLimit int) *BaseGwClientBuilder {
    _this.autoRetryLimit = autoRetryLimit
    return _this
}

//
//
//
func (_this *BaseGwClientBuilder) SetSecurityToken(securityToken string) *BaseGwClientBuilder {
    _this.securityToken = securityToken
    return _this
}
