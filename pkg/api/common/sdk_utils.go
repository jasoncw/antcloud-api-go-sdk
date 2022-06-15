// 
// @Author: jason zhou
// @Description: 
// @File:  sdk_utils.go
// @Version: 0.0.1
// @Date: 2022/6/7 14:40
// 

package common

import (
    "antcloud-api-go-sdk/pkg/api/log"
    "bytes"
    "encoding/hex"
    "errors"
    uuid "github.com/satori/go.uuid"
    "strconv"
    "time"
    "unicode"
)

//
//
//
func SdkUtils_CheckNotNull(msg string, arg interface{}) error {
    if arg == nil {
        return errors.New(msg)
    }
    return nil
}

//
//
//
func SdkUtils_GenerateReqMsgId() string {
    u := [uuid.Size]byte(uuid.NewV4())
    buf := make([]byte, 32)
    hex.Encode(buf[0:8], u[0:4])
    hex.Encode(buf[8:12], u[4:6])
    hex.Encode(buf[12:16], u[6:8])
    hex.Encode(buf[16:20], u[8:10])
    hex.Encode(buf[20:], u[10:])
    return string(buf)
}

func SdkUtils_FormatDate(time0 time.Time) string {
    return time0.UTC().Format("2006-01-02T15:04:05Z")
}



//
// 驼峰式写法转为下划线写法
//
func SdkUtils_Camel2Case(name string) string {
    buffer := NewBuffer()
    for i, r := range name {
        if unicode.IsUpper(r) {
            if i != 0 {
                buffer.Append('_')
            }
            buffer.Append(unicode.ToLower(r))
        } else {
            buffer.Append(r)
        }
    }
    return buffer.String()
}

// 内嵌bytes.Buffer，支持连写
type Buffer struct {
    *bytes.Buffer
}

func NewBuffer() *Buffer {
    return &Buffer{Buffer: new(bytes.Buffer)}
}

func (b *Buffer) Append(i interface{}) *Buffer {
    switch val := i.(type) {
    case int:
        b.append(strconv.Itoa(val))
    case int64:
        b.append(strconv.FormatInt(val, 10))
    case uint:
        b.append(strconv.FormatUint(uint64(val), 10))
    case uint64:
        b.append(strconv.FormatUint(val, 10))
    case string:
        b.append(val)
    case []byte:
        b.Write(val)
    case rune:
        b.WriteRune(val)
    }

    return b
}

func (b *Buffer) append(s string) *Buffer {
    defer func() {
        if err := recover(); err != nil {
            log.GetLogger().Error("*****内存不够了！******")
        }
    }()
    b.WriteString(s)
    return b
}
