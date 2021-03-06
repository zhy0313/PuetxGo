package PtUtil

/**
 * Title:常用函数工具
 * User: iuoon
 * Date: 2016-9-23
 * Version: 1.0
 */
import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//GeneratePIDFile 生成一个pID文件，格式：<name>.<ID>.pID
func GeneratePIDFile(name string, id int) {
	var filename string
	if id == 0 {
		filename = fmt.Sprintf("%s.pid", name)
	} else {
		filename = fmt.Sprintf("%s.%d.pid", name, id)
	}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	pid := os.Getpid()
	fmt.Fprintf(f, "%d", pid)
	Debug("new %s[%d], pID: %d", name, id, pid)
}

var r *rand.Rand

//GetRandomString 生成指定长度的随机字符串
func GetRandomString(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//GetRandomInterval 生成指定区间一个随机数
func GetRandomInterval(low, high int) int {
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return r.Intn(high-low+1) + low
}

//GetRandom 生成一个1-n的随机数
func GetRandom(n int) int {
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return r.Intn(n) + 1
}

//TimeToStr 当前时间转字符串
func TimeToStr(t int64) string {
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}

//StrToTime 字符串转当前时间，格式2006-01-02 15:04:05
func StrToTime(str string) int64 {
	loc, _ := time.LoadLocation("Local")
	the_time, err := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	if err == nil {
		return the_time.Unix()
	} else {
		return 0
	}
}

//NextTime 当前时间下一个指定时间
func NextTime(h, m, s int) int64 {
	now := time.Now()
	year, mon, day := now.Date()
	hour, min, sec := now.Clock()

	ts := StrToTime(fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", year, mon, day, 0, 0, 0))

	if (3600*h + 60*m + s) <= (3600*hour + 60*min + sec) {
		ts += int64(3600*24 + 3600*h + 60*m + s)
	} else {
		ts += int64(3600*h + 60*m + s)
	}
	return ts
}

func getTypeName(obj interface{}) (typestr string) {
	typ := reflect.TypeOf(obj)
	typestr = typ.String()

	lastDotIndex := strings.LastIndex(typestr, ".")
	if lastDotIndex != -1 {
		typestr = typestr[lastDotIndex+1:]
	}

	return
}

//snakeCasedName 小写字符串
func snakeCasedName(name string) string {
	newstr := make([]rune, 0)
	firstTime := true

	for _, chr := range name {
		if isUpper := 'A' <= chr && chr <= 'Z'; isUpper {
			if firstTime == true {
				firstTime = false
			} else {
				newstr = append(newstr, '_')
			}
			chr -= ('A' - 'a')
		}
		newstr = append(newstr, chr)
	}

	return string(newstr)
}

//snakeCasedName 大写字符串
func titleCasedName(name string) string {
	newstr := make([]rune, 0)
	upNextChar := true

	for _, chr := range name {
		switch {
		case upNextChar:
			upNextChar = false
			chr -= ('a' - 'A')
		case chr == '_':
			upNextChar = true
			continue
		}

		newstr = append(newstr, chr)
	}

	return string(newstr)
}

//pluralizeString ...
func pluralizeString(str string) string {
	if strings.HasSuffix(str, "data") {
		return str
	}
	if strings.HasSuffix(str, "y") {
		str = str[:len(str)-1] + "ie"
	}
	return str + "s"
}

//scanMapIntoStruct map转结构
func scanMapIntoStruct(obj interface{}, objMap map[string][]byte) error {
	dataStruct := reflect.Indirect(reflect.ValueOf(obj))
	if dataStruct.Kind() != reflect.Struct {
		return errors.New("expected a pointer to a struct")
	}

	dataStructType := dataStruct.Type()

	for i := 0; i < dataStructType.NumField(); i++ {
		field := dataStructType.Field(i)
		fieldv := dataStruct.Field(i)

		err := scanMapElement(fieldv, field, objMap)
		if err != nil {
			return err
		}
	}

	return nil
}

//scanMapElement ...
func scanMapElement(fieldv reflect.Value, field reflect.StructField, objMap map[string][]byte) error {

	objFieldName := field.Name
	bb := field.Tag
	sqlTag := bb.Get("sql")

	if bb.Get("beedb") == "-" || sqlTag == "-" || reflect.ValueOf(bb).String() == "-" {
		return nil
	}
	sqlTags := strings.Split(sqlTag, ",")
	sqlFieldName := objFieldName
	if len(sqlTags[0]) > 0 {
		sqlFieldName = sqlTags[0]
	}
	inline := false
	//omitempty := false //TODO!
	// CHECK INLINE
	if len(sqlTags) > 1 {
		if stringArrayContains("inline", sqlTags[1:]) {
			inline = true
		}
	}
	if inline {
		if field.Type.Kind() == reflect.Struct && field.Type.String() != "time.Time" {
			for i := 0; i < field.Type.NumField(); i++ {
				err := scanMapElement(fieldv.Field(i), field.Type.Field(i), objMap)
				if err != nil {
					return err
				}
			}
		} else {
			return errors.New("A non struct type can't be inline.")
		}
	}

	// not inline

	data, ok := objMap[sqlFieldName]

	if !ok {
		return nil
	}

	var v interface{}

	switch field.Type.Kind() {

	case reflect.Slice:
		v = data
	case reflect.String:
		v = string(data)
	case reflect.Bool:
		v = string(data) == "1"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		x, err := strconv.Atoi(string(data))
		if err != nil {
			return errors.New("arg " + sqlFieldName + " as int: " + err.Error())
		}
		v = x
	case reflect.Int64:
		x, err := strconv.ParseInt(string(data), 10, 64)
		if err != nil {
			return errors.New("arg " + sqlFieldName + " as int: " + err.Error())
		}
		v = x
	case reflect.Float32, reflect.Float64:
		x, err := strconv.ParseFloat(string(data), 64)
		if err != nil {
			return errors.New("arg " + sqlFieldName + " as float64: " + err.Error())
		}
		v = x
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		x, err := strconv.ParseUint(string(data), 10, 64)
		if err != nil {
			return errors.New("arg " + sqlFieldName + " as int: " + err.Error())
		}
		v = x
	//Supports Time type only (for now)
	case reflect.Struct:
		if fieldv.Type().String() != "time.Time" {
			return errors.New("unsupported struct type in Scan: " + fieldv.Type().String())
		}

		x, err := time.Parse("2006-01-02 15:04:05", string(data))
		if err != nil {
			x, err = time.Parse("2006-01-02 15:04:05.000 -0700", string(data))

			if err != nil {
				return errors.New("unsupported time format: " + string(data))
			}
		}

		v = x
	default:
		return errors.New("unsupported type in Scan: " + reflect.TypeOf(v).String())
	}

	fieldv.Set(reflect.ValueOf(v))

	return nil
}

//scanStructIntoMap ...
func scanStructIntoMap(obj interface{}) (map[string]interface{}, error) {
	dataStruct := reflect.Indirect(reflect.ValueOf(obj))
	if dataStruct.Kind() != reflect.Struct {
		return nil, errors.New("expected a pointer to a struct")
	}

	dataStructType := dataStruct.Type()

	mapped := make(map[string]interface{})

	for i := 0; i < dataStructType.NumField(); i++ {
		field := dataStructType.Field(i)
		fieldv := dataStruct.Field(i)
		fieldName := field.Name
		bb := field.Tag
		sqlTag := bb.Get("sql")
		sqlTags := strings.Split(sqlTag, ",")
		var mapKey string

		inline := false

		if bb.Get("beedb") == "-" || sqlTag == "-" || reflect.ValueOf(bb).String() == "-" {
			continue
		} else if len(sqlTag) > 0 {
			//TODO: support tags that are common in json like omitempty
			if sqlTags[0] == "-" {
				continue
			}
			mapKey = sqlTags[0]
		} else {
			mapKey = fieldName
		}

		if len(sqlTags) > 1 {
			if stringArrayContains("inline", sqlTags[1:]) {
				inline = true
			}
		}

		if inline {
			// get an inner map and then put it insIDe the outer map
			map2, err2 := scanStructIntoMap(fieldv.Interface())
			if err2 != nil {
				return mapped, err2
			}
			for k, v := range map2 {
				mapped[k] = v
			}
		} else {
			value := dataStruct.FieldByName(fieldName).Interface()
			mapped[mapKey] = value
		}
	}

	return mapped, nil
}

//StructName 返回结构名
func StructName(s interface{}) string {
	v := reflect.TypeOf(s)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Name()
}

//GetTableName 返回结构名
func GetTableName(s interface{}) string {
	v := reflect.TypeOf(s)
	if v.Kind() == reflect.String {
		s2, _ := s.(string)
		return snakeCasedName(s2)
	}
	tn := scanTableName(s)
	if len(tn) > 0 {
		return tn
	}
	return GetTableName(StructName(s))
}

//scanTableName ...
func scanTableName(s interface{}) string {
	if reflect.TypeOf(reflect.Indirect(reflect.ValueOf(s)).Interface()).Kind() == reflect.Slice {
		sliceValue := reflect.Indirect(reflect.ValueOf(s))
		sliceElementType := sliceValue.Type().Elem()
		for i := 0; i < sliceElementType.NumField(); i++ {
			bb := sliceElementType.Field(i).Tag
			if len(bb.Get("tname")) > 0 {
				return bb.Get("tname")
			}
		}
	} else {
		tt := reflect.TypeOf(reflect.Indirect(reflect.ValueOf(s)).Interface())
		for i := 0; i < tt.NumField(); i++ {
			bb := tt.Field(i).Tag
			if len(bb.Get("tname")) > 0 {
				return bb.Get("tname")
			}
		}
	}
	return ""

}

func stringArrayContains(needle string, haystack []string) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}
	return false
}

/*
 * byte[] to string
 */
func BytetoString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}