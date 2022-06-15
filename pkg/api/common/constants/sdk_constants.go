// 
// @Author: jason zhou
// @Description: 
// @File:  sdk_constants.go
// @Version: 0.0.1
// @Date: 2022/6/7 14:37
// 

package constants

const (
    EMPTY  string = ""
    HYPHEN        = "-"

    BASE_SDK_VERSION_VALUE string = "3.4.0"
    DEFAULT_CHARSET        string = "UTF-8"
    // public static final Charset SIGN_CHARSET    string   = Charset.forName("UTF-8")
    DEFAULT_SIGN_TYPE string = "HmacSHA1"
    SIGN_TYPE_SHA1    string = "HmacSHA1"
    SIGN_TYPE_SHA256  string = "HmacSHA256"
    BASE64URL         string = "antcloud-base64://"

    // ParamKeys
    ParamKeys_RESPONSE               string = "response"
    ParamKeys_REQ_MSG_ID             string = "req_msg_id"
    ParamKeys_RESULT_CODE            string = "result_code"
    ParamKeys_RESULT_MSG             string = "result_msg"
    ParamKeys_RESULT_MSG_PLACEHOLDER string = "result_msg_placeholder"
    ParamKeys_RESULT_MSG_ARGS        string = "result_msg_args"
    ParamKeys_SIGN_TYPE              string = "sign_type"
    ParamKeys_SIGN                   string = "sign"
    ParamKeys_REQ_TIME               string = "req_time"
    ParamKeys_BASE_SDK_VERSION       string = "base_sdk_version"
    ParamKeys_METHOD                 string = "method"
    ParamKeys_VERSION                string = "version"
    ParamKeys_ACCESS_KEY             string = "access_key"
    ParamKeys_SECURITY_TOKEN         string = "security_token"
    ParamKeys_PRODUCT_INSTANCE_ID    string = "product_instance_id"
    ParamKeys_REGION_NAME            string = "region_name"
    ParamKeys_INVOKER_USER           string = "invoker.user"
    ParamKeys_INTERNAL_API           string = "internal_api"

    // ResultCodes
    ResultCodes_OK                    string = "OK"
    ResultCodes_MISSING_PARAMETER     string = "MISSING_PARAMETER"
    ResultCodes_INVALID_PARAMETER     string = "INVALID_PARAMETER"
    ResultCodes_TRANSPORT_ERROR       string = "TRANSPORT_ERROR"
    ResultCodes_PARASE_URL_ERROR      string = "PARASE_URL_ERROR"
    ResultCodes_RESPONSE_FORMAT_ERROR string = "RESPONSE_FORMAT_ERROR"
    ResultCodes_BAD_SIGNATURE         string = "INVALID_RESPONSE_SIGNATURE"
    ResultCodes_UNKNOWN_ERROR         string = "UNKNOWN_ERROR"
    ResultCodes_ACCESS_DENIED         string = "ACCESS_DENIED"
    ResultCodes_METHOD_NOT_FOUND      string = "METHOD_NOT_FOUND"
    ResultCodes_CALL_HTTP_FAILED      string = "CALL_HTTP_FAILED"

    // ResultMsgPlaceholders
    ResultMsgPlaceholders_PROVIDER_UNKNOWN_ERROR string = "PROVIDER_UNKNOWN_ERROR"
)
