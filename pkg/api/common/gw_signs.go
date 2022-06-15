// 
// @Author: jason zhou
// @Description: 
// @File:  gw_signs.go
// @Version: 0.0.1
// @Date: 2022/6/6 16:50
// 

package common

import (
    "antcloud-api-go-sdk/pkg/api/common/constants"
    "crypto/hmac"
    "crypto/sha1"
    "crypto/sha256"
    "encoding/base64"
    "hash"
    "net/url"
    "regexp"
    "sort"
    "strings"
)

const patternStr = `\Q+\E|\Q*\E|\Q%7E\E|\Q%2F\E`

var EncodedCharactersPattern, _ = regexp.Compile(patternStr)

//
//
//
func GwSigns_Sign(params map[string]string, algorithm, secret string) string {
    hmac := getHash(algorithm, secret)
    keys := signKeyFilter(params)
    sort.Strings(keys)
    for i, key := range keys {
        if i != 0 {
            hmac.Write([]byte("&"))
        }
        hmac.Write([]byte(urlEncode(key)))
        hmac.Write([]byte("="))
        hmac.Write([]byte(urlEncode(params[key])))
    }
    signData := hmac.Sum(nil)
    return base64.StdEncoding.EncodeToString(signData)
}

//
//
//
func signKeyFilter(params map[string]string) []string {
    keys := make([]string, 0)
    for k, v := range params {
        if !strings.HasPrefix(v, constants.BASE64URL) {
            keys = append(keys, k)
        }
    }
    return keys
}

//
//
//
func urlEncode(value string) string {
    return url.QueryEscape(value)
}

//
//
//
func getHash(algorithm, secret string) hash.Hash {
    key := []byte(secret)
    if algorithm == constants.SIGN_TYPE_SHA256 {
        return hmac.New(sha256.New, key)
    }
    return hmac.New(sha1.New, key)
}
