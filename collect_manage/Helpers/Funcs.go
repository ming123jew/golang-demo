package Helpers

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
)

type Funcs struct {
}

/*
元素是否存在一个切片中
*/
func (f Funcs) InSlice(value interface{}, list []string) bool {
	switch value.(type) {
	case int:
		value = strconv.Itoa(value.(int)) // value.(int)表示断言：value是一个int
	case uint:
		value = strconv.Itoa(int(value.(uint))) // int(type)表示强类型转换：将uint转为int
	case float64:
		value = strconv.FormatFloat(value.(float64), 'f', -1, 64) // 将float64转为string
	case float32:
		value = strconv.FormatFloat(float64(value.(float32)), 'f', -1, 32) // 将float32先转为float64再转为string
	}

	for _, el := range list {
		if el == value {
			return true
		}
	}
	return false
}

/*
去除切片中的空值
*/
func (f Funcs) FilterSlice(list []string) (result []string) {
	for _, value := range list {
		if value != "" {
			result = append(result, value)
		}
	}
	return
}

/*
去除切片中的重复值
*/
func (f Funcs) UniqueSlice(list []string) (result []string) {
	// 存储每个元素的数量
	countEle := make(map[string]int)
	for _, value := range list {
		countEle[value] += 1
		if countEle[value] > 1 {
			continue
		}
		result = append(result, value)
	}
	return
}

/*
将多个同类型切片合并成一个切片
*/
func (f Funcs) MergeSlice(list ...[]string) (result []string) {
	for _, value := range list {
		for _, v := range value {
			result = append(result, v)
		}
	}
	return
}

/*
将多个map合并成一个map
*/
func (f Funcs) MergeMap(mapList ...map[string]string) map[string]string {
	result := make(map[string]string)
	for _, value := range mapList {
		for k, v := range value {
			result[k] = v
		}
	}
	return result
}

func (f Funcs) Hello() {
	fmt.Println("hello world.")
}

/*
md5函数
*/
func (f Funcs) Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

/*
base64加密
*/
func (f Funcs) Base64Encode(str string, usedUrl bool) string {
	if usedUrl {
		// 使用在url地址中
		return base64.URLEncoding.EncodeToString([]byte(str))
	} else {
		// 普通使用
		return base64.StdEncoding.EncodeToString([]byte(str))
	}
}

/*
Base64解密
*/
func (f Funcs) Base64Decode(str string, usedUrl bool) string {
	// 普通使用
	decodeBytes, err := base64.StdEncoding.DecodeString(str)
	// 使用在url地址中
	if usedUrl {
		decodeBytes, err = base64.URLEncoding.DecodeString(str)
	}

	if err != nil {
		return ""
	}
	return string(decodeBytes)
}

/*
url编码
*/
func (f Funcs) UrlEncode(str string) string {
	return url.QueryEscape(str)
}

/*
url解码
*/
func (f Funcs) UrlDecode(str string) string {
	decodeStr, _ := url.QueryUnescape(str)
	return decodeStr
}
