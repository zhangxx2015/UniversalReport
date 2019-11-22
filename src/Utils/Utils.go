package Utils

import (
	"bufio"
	"crypto/md5"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	crypto_rand "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"math"

	"github.com/satori/go.uuid"
)

func Typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func ToNumberSystem26(n int) string { //1,26,27>A,Z,AA
	s := ""
	for n > 0 {
		m := n % 26 // 求余数
		if 0 == m { // 最后一位
			m = 26
		}
		s = string(m+64) + s // int数值转ascii字符后拼接字符串
		n = (n - m) / 26     // 减余数后降位
	}
	return s
}

func FromNumberSystem26(s string) int { //A,Z,AA>1,26,27
	r := 0
	for i := len(s); i > 0; i-- {
		//ch := s[i-1]
		//fmt.Println("从低位顺序", len(s)-i, "asc字符", string(ch), "int数值", ch, "十进制数", ch-64, "所在位基数", int(math.Pow(float64(26), float64(len(s)-i))))
		r += int(s[i-1]-64) * // 从低位顺序取字符(二十六进制),转为十进制数
			int(math.Pow(float64(26), float64(len(s)-i))) // 乘所在位基数
	}
	return r
}

func ToString(inter interface{}) string {
	switch inter.(type) {
	case string:
		return fmt.Sprint(inter.(string))
		break
	case int64:
		return fmt.Sprint(inter.(int64))
		break
	case int:
		return fmt.Sprint(inter.(int))
		break
	case float32:
		return fmt.Sprint(inter.(float32))
		break
	case float64:
		return fmt.Sprint(inter.(float64))
		break
	case bool:
		return fmt.Sprint(inter.(bool))
		break
	}
	// var err error = errors.New("unknow type")
	panic("unknow type")
	return ""
}
func ParseInt64Default(s string, defaultValue int64) int64 {
	val := defaultValue
	ret, err := strconv.ParseInt(s, 10, 64)
	if nil == err {
		val = int64(ret)
	}
	return val
}
func ParseIntDefault(s string, defaultValue int) int {
	val := defaultValue
	ret, err := strconv.ParseInt(s, 10, 32)
	if nil == err {
		val = int(ret)
	}
	return val
}
func ParseInt(s string) int {
	val := -1
	ret, err := strconv.ParseInt(s, 10, 32)
	if nil != err {
		panic(err.Error())
	}
	val = int(ret)
	return val
}

func HttpBody(r *http.Request) []byte {
	httpBody, err := ioutil.ReadAll(r.Body)
	if nil != err {
		return []byte{}
	}
	return httpBody
}
func Payload(r *http.Request) string {
	return string(HttpBody(r))
}
func PayloadMap(r *http.Request) map[string]interface{} {
	dict := make(map[string]interface{}, 0)
	httpBody := HttpBody(r)
	if len(httpBody) > 0 {
		if err := json.Unmarshal(httpBody, &dict); err == nil {
			return dict
		}
		panic("error")
	}
	return dict
}

func PayloadMapNotContainskeys(r *http.Request, handler ContainskeyCallback, key ...string) map[string]interface{} {
	//fmt.Println("PayloadMapNotContainskeys!")
	payloadMap := PayloadMap(r)
	//fmt.Println("PayloadMapNotContainskeys:", "payloadMap:", payloadMap)
	for _, k := range key {
		exist := false
		for p, _ := range payloadMap {
			//fmt.Println("PayloadMapNotContainskeys:", k, p)
			if p == k {
				exist = true
				break
			}
		}
		if !exist {
			handler(k, false)
		}
	}
	return payloadMap
}

func MapNotContainskeys(maps map[string]interface{}, handler ContainskeyCallback, key ...string) map[string]interface{} {
	for _, k := range key {
		exist := false
		for p, _ := range maps {
			if p == k {
				exist = true
				break
			}
		}
		if !exist {
			handler(k, false)
		}
	}
	return maps
}

// func MapContainskeys(payMap map[string]interface{}, key string) interface{} {
// 	if val, ok := payMap[key]; ok {
// 		return val
// 	}
// 	return nil
// }

func PayloadMapString(payMap map[string]interface{}, key string, defaultValue string) string {
	strVal := defaultValue
	if val, ok := payMap[key]; ok {
		strVal = fmt.Sprint(val)
	}
	return strVal
}
func PayloadMapInt(payMap map[string]interface{}, key string, defaultValue int) int {
	intVal := defaultValue
	if val, ok := payMap[key]; ok {
		intVal = val.(int)
	}
	return intVal
}

func QueryBool(r *http.Request, key string, defaultValue bool) bool {
	result := defaultValue
	val, err := strconv.ParseBool(QueryString(r, key, ""))
	if nil == err {
		result = val
	}
	return result
}

func QueryFloat(r *http.Request, key string, defaultValue float32) float32 {
	result := defaultValue
	val, err := strconv.ParseFloat(QueryString(r, key, ""), 64)
	if nil == err {
		result = float32(val)
	}
	return result
}
func QueryInt(r *http.Request, key string, defaultValue int) int {
	result := defaultValue
	s := QueryString(r, key, "")
	val, err := strconv.ParseInt(s, 10, 32)
	if nil == err {
		result = int(val)
	}
	return result

}
func QueryString(r *http.Request, key string, defaultValue string) string {
	result := defaultValue
	u, err := url.Parse(r.URL.String())
	if nil == err {
		for k, v := range u.Query() {
			if key == k {
				result = strings.Join(v, "")
				break
			}
		}
	}
	return result
}

func QueryContainskey(r *http.Request, key string) bool {
	u, _ := url.Parse(r.URL.String())
	for k, _ := range u.Query() {
		if key == k {
			return true
			break
		}
	}
	return false
}

type ContainskeyCallback func(string, bool)

func QueryContainskeys(r *http.Request, handler ContainskeyCallback, key ...string) {
	for _, k := range key {
		exist := QueryContainskey(r, k)
		handler(k, exist)
	}
}
func QueryNotContainskeys(r *http.Request, handler ContainskeyCallback, key ...string) {
	for _, k := range key {
		if !QueryContainskey(r, k) {
			handler(k, false)
		}
	}
}

func JobNumber() string {
	// 67000000:11GB zhangxx @ 2017-07-08
	return base64.URLEncoding.EncodeToString([]byte(uuid.NewV4().String()))[0:8] + uuid.NewV4().String()[0:8]
}

func VCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	vcode := ""
	for i := 0; i < length; i++ {
		vcode += strconv.Itoa(rand.Intn(9))
	}
	return vcode
}

func CurrDate() string {
	return time.Now().Format("2006-01-02")
}

func CurrTime() string {
	return time.Now().Format("15:04:05")
}

func CurrUtc() int {
	return int(time.Now().Unix())
}

func CurrUtcTo(offsetSeconds int) int {
	return int(time.Now().Add(time.Duration(offsetSeconds) * time.Second).Unix())
}

func Sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func GetGuid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(crypto_rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err.Error())
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func Get_internal() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		os.Exit(1)
	}
	ip := ""
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//os.Stdout.WriteString(ipnet.IP.String() + "\n")
				ip = ipnet.IP.String()
				//break
			}
		}
	}
	//os.Exit(0)
	return ip
}

func Md5Sum(file string) string {
	f, err := os.Open(file)
	if err != nil {
		return ""
	}
	defer f.Close()
	r := bufio.NewReader(f)
	h := md5.New()
	_, err = io.Copy(h, r)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// func ToMap(obj interface{}) map[string]interface{} {
// 	types := reflect.TypeOf(obj)
// 	value := reflect.ValueOf(obj)
// 	var data = make(map[string]interface{})

// 	for i := 0; i < types.NumField(); i++ {
// 		key := types.Field(i).Name
// 		val := value.Field(i).Interface()
// 		typ := reflect.TypeOf(val)

// 		if typ == reflect.TypeOf(BaseModel{}) {
// 			childs := ToMap(val)
// 			for key, value := range childs {
// 				//typ := reflect.TypeOf(value)
// 				//fmt.Println(typ, "@", key, ":", value)
// 				data[key] = value
// 			}
// 			continue
// 		}
// 		//fmt.Println(typ, "@", key, ":", val)
// 		data[key] = val
// 	}
// 	return data
// }
