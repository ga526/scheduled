package common

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"math/rand"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

type utils struct{}

func Utils() *utils{
	data := new(utils)
	return data
}

/**
 * 函数生成从 0- n 的随机数
 */
func (this *utils) Rand() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}

/**
 * int64 转 string
 */
func (this *utils) Int64ToString(data int64) string {
	return strconv.FormatInt(data, 10)
}

func (this *utils) StringToInt64(data string) int64{
	int64,_ := strconv.ParseInt(data, 10, 64)
	return int64
}

/**
 * 获取当前可执行文件路径
 */
func (this *utils) GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}

/**
 * []byte 转为整形
 */
func (this *utils) ByteArrToString(source []byte) string {
	return string(source[:])
}

/**
 * 对字符串进行 sha256 编码
 */
func (this *utils) Sha256String(str string, salt ...string) string {
	encryptStruct := sha256.New()
	if salt != nil {
		str = str + "_" + salt[0]
	}

	encryptStruct.Write([]byte(str))
	result := fmt.Sprintf("%x", encryptStruct.Sum(nil))
	return result
}

/**
 * 对字符串进行 md5 编码
 */
func (this *utils) Md5String(str string, salt ...string) string {
	encryptStruct := md5.New()
	if salt != nil {
		str = str + "_" + salt[0]
	}

	encryptStruct.Write([]byte(str))
	result := fmt.Sprintf("%x", encryptStruct.Sum(nil))
	return result
}

/**
 * int64 转 int
 */
func (this *utils) Int64ToInt(data int64) int {
	return int(data)
}

/**
 * int 转 string
 */
func (this *utils) IntToString(data int) string {
	return strconv.Itoa(data)
}

/**
 * string 转 int
 */
func (this *utils) StringToInt(data string) int {
	int, err := strconv.Atoi(data)
	if err != nil {
		fmt.Print(err.Error())
		return 0
	}
	return int
}

func (this *utils) StructToMap(data interface{}) map[string]interface{} {
	myMap := make(map[string]interface{})
	elem := reflect.ValueOf(data).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		myMap[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	return myMap
}
/**
 * 获取当前时间
 */
func (this *utils) GetNowTime() string{
	return time.Now().Format("2006-01-02 15:04:05");
}

func (this *utils) SortByKey(sourceData map[string]interface{}) map[string]interface{}{
	newMap := make(map[string]interface{})
	var keys []string
	for key,_ := range sourceData{
		keys = append(keys,key)
	}

	sort.Strings(keys)
	for _,item := range keys{
		newMap[item] = sourceData[item]
	}
	return newMap
}

func (this *utils) HttpBuildQuery(queryData map[string]string) string{
	urlValue := url.Values{}
	for key,item := range queryData{
		urlValue.Add(key,item)
	}
	return urlValue.Encode()
}

func (this *utils) UTF82GB2312(s []byte)([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.HZGB2312.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
func (this *utils) GB23122GBUTF8(s []byte)([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), unicode.UTF8.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func (this *utils) EncodeSelf(str string) string{
	data := []byte(str)
	for k,val := range data{
		data[k] = val + 1
	}
	return string(data)
}