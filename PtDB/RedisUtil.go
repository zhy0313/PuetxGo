package PtDB

/**
 * Title:Redis工具类
 * User: iuoon
 * Date: 2016-9-23
 * Version: 1.0
 */

import (
	"errors"
	"gopkg.in/redis.v3"
	"reflect"
	"strconv"
	"time"
	"github.com/iuoon/PuetxGo/PtUtil"
)

//目前支持整数浮点数字符串，时间戳用int64的unix时间戳，不支持time.TIme
//主键使用PK表示，以便支持beedb
// type TestStruct struct {
// 	UID        int    `PK`
// 	Username   string `PK`
// 	Departname string
// 	Created    int64
// }

// func main() {
// 	LoadConfig("config.json")
// 	InitLogger("NewServer")

// 	if !connect_redis() {
// 		Debug("connect redis fail")
// 		return
// 	}

// 	var saveone TestStruct
// 	saveone.UID = 1
// 	saveone.Username = "name"
// 	saveone.Departname = "Test Add Departname"
// 	saveone.Created = time.Now().Unix()

// 	SaveToRedis(rdClient, &saveone)

// 	var saveone1 TestStruct
// 	saveone1.UID = 1
// 	saveone1.Username = "name"
// 	LoadFromRedis(rdClient, &saveone1)
// 	fmt.Println(saveone1)
// 	fmt.Println("ok")
// }
func getRedisTableName(info interface{}) string {
	tableName := "h_" + PtUtil.GetTableName(info)

	//sacn key
	dataStruct := reflect.Indirect(reflect.ValueOf(info))
	dataStructType := dataStruct.Type()
	for i := 0; i < dataStructType.NumField(); i++ {
		fieldType := dataStructType.Field(i)
		fieldValue := dataStruct.Field(i)
		fieldTag := fieldType.Tag
		if fieldTag.Get("pk") == "true" {
			switch fieldType.Type.Kind() {
			case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
				str := strconv.FormatInt(fieldValue.Int(), 10)
				tableName += ":" + str
			case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				str := strconv.FormatUint(fieldValue.Uint(), 10)
				tableName += ":" + str
			case reflect.Float32, reflect.Float64:
				str := strconv.FormatFloat(fieldValue.Float(), 'f', -1, 64)
				tableName += ":" + str
			case reflect.String:
				tableName += ":" + fieldValue.String()
			//时间类型
			case reflect.Struct:
				str := strconv.FormatInt(fieldValue.Interface().(time.Time).Unix(), 10)
				tableName += ":" + str
			case reflect.Bool:
				if fieldValue.Bool() {
					tableName += ":1"
				} else {
					tableName += ":0"
				}
			case reflect.Slice:
				if fieldType.Type.Elem().Kind() == reflect.Uint8 {
					tableName += ":" + string(fieldValue.Interface().([]byte))
				}
			}
		}
	}

	return tableName
}

//SaveToRedis 将一个结构保存到redis中
func SaveToRedis(client *redis.Client, info interface{}) {
	tableName := getRedisTableName(info)

	dataStruct := reflect.Indirect(reflect.ValueOf(info))
	dataStructType := dataStruct.Type()
	for i := 0; i < dataStructType.NumField(); i++ {
		fieldType := dataStructType.Field(i)
		fieldValue := dataStruct.Field(i)

		switch fieldType.Type.Kind() {
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
			str := strconv.FormatInt(fieldValue.Int(), 10)
			client.HSet(tableName, fieldType.Name, str)
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			str := strconv.FormatUint(fieldValue.Uint(), 10)
			client.HSet(tableName, fieldType.Name, str)
		case reflect.Float32, reflect.Float64:
			str := strconv.FormatFloat(fieldValue.Float(), 'f', -1, 64)
			client.HSet(tableName, fieldType.Name, str)
		case reflect.String:
			client.HSet(tableName, fieldType.Name, fieldValue.String())
		//时间类型
		case reflect.Struct:
			str := strconv.FormatInt(fieldValue.Interface().(time.Time).Unix(), 10)
			client.HSet(tableName, fieldType.Name, str)
		case reflect.Bool:
			if fieldValue.Bool() {
				client.HSet(tableName, fieldType.Name, "1")
			} else {
				client.HSet(tableName, fieldType.Name, "0")
			}
		case reflect.Slice:
			if fieldType.Type.Elem().Kind() == reflect.Uint8 {
				client.HSet(tableName, fieldType.Name, string(fieldValue.Interface().([]byte)))
			}
		}
	}
}

//LoadFromRedis 从redis中读取一个结构
func LoadFromRedis(client *redis.Client, info interface{}) error {
	tableName := getRedisTableName(info)

	if !client.Exists(tableName).Val() {
		return errors.New("key not existst")
	}

	dataStruct := reflect.Indirect(reflect.ValueOf(info))
	dataStructType := dataStruct.Type()
	m := client.HGetAllMap(tableName)
	for key, value := range m.Val() {
		for i := 0; i < dataStructType.NumField(); i++ {
			fieldType := dataStructType.Field(i)
			fieldValue := dataStruct.Field(i)
			if fieldType.Name == key {
				switch fieldType.Type.Kind() {
				case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
					n, _ := strconv.Atoi(value)
					fieldValue.SetInt(int64(n))
				case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					n, _ := strconv.Atoi(value)
					fieldValue.SetUint(uint64(n))
				case reflect.Float32:
					n, _ := strconv.ParseFloat(value, 32)
					fieldValue.SetFloat(n)
				case reflect.Float64:
					n, _ := strconv.ParseFloat(value, 64)
					fieldValue.SetFloat(n)
				case reflect.String:
					fieldValue.SetString(value)
				case reflect.Bool:
					fieldValue.SetBool(value == "1")
				case reflect.Slice:
					fieldValue.SetBytes([]byte(value))
				}
				break
			}
		}
	}
	return nil
}

//DelFromRedis 从redis中删除一个结构对应的key
func DelFromRedis(client *redis.Client, info interface{}) error {
	tableName := getRedisTableName(info)
	client.Del(tableName)
	return nil
}

//GetFieldFromRedis 从redis中读取一个结构指定的字段
func GetFieldFromRedis(client *redis.Client, info interface{}, field string) error {
	tableName := getRedisTableName(info)
	if !client.Exists(tableName).Val() {
		return errors.New("key not existst")
	}
	dataStruct := reflect.Indirect(reflect.ValueOf(info))
	dataStructType := dataStruct.Type()

	for i := 0; i < dataStructType.NumField(); i++ {
		fieldType := dataStructType.Field(i)
		if fieldType.Name == field {
			fieldValue := dataStruct.Field(i)
			v := client.HGet(tableName, field).Val()
			switch fieldType.Type.Kind() {
			case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
				n, _ := strconv.Atoi(v)
				fieldValue.SetInt(int64(n))
			case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				n, _ := strconv.Atoi(v)
				fieldValue.SetUint(uint64(n))
			case reflect.Float32:
				n, _ := strconv.ParseFloat(v, 32)
				fieldValue.SetFloat(n)
			case reflect.Float64:
				n, _ := strconv.ParseFloat(v, 64)
				fieldValue.SetFloat(n)
			case reflect.String:
				fieldValue.SetString(v)
			case reflect.Bool:
				fieldValue.SetBool(v == "1")
			case reflect.Slice:
				fieldValue.SetBytes([]byte(v))
			}
			return nil
		}
	}

	return errors.New("field not existst")
}

//SetFieldFromRedis 从redis中保存一个结构指定的字段
func SetFieldFromRedis(client *redis.Client, info interface{}, field string) error {
	tableName := getRedisTableName(info)
	if !client.Exists(tableName).Val() {
		return errors.New("key not existst")
	}
	dataStruct := reflect.Indirect(reflect.ValueOf(info))
	dataStructType := dataStruct.Type()

	for i := 0; i < dataStructType.NumField(); i++ {
		fieldType := dataStructType.Field(i)
		if fieldType.Name == field {
			fieldValue := dataStruct.Field(i)
			switch fieldType.Type.Kind() {
			case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
				str := strconv.FormatInt(fieldValue.Int(), 10)
				client.HSet(tableName, fieldType.Name, str)
			case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				str := strconv.FormatUint(fieldValue.Uint(), 10)
				client.HSet(tableName, fieldType.Name, str)
			case reflect.Float32, reflect.Float64:
				str := strconv.FormatFloat(fieldValue.Float(), 'f', -1, 64)
				client.HSet(tableName, fieldType.Name, str)
			case reflect.String:
				client.HSet(tableName, fieldType.Name, fieldValue.String())
			//时间类型
			case reflect.Struct:
				str := strconv.FormatInt(fieldValue.Interface().(time.Time).Unix(), 10)
				client.HSet(tableName, fieldType.Name, str)
			case reflect.Bool:
				if fieldValue.Bool() {
					client.HSet(tableName, fieldType.Name, "1")
				} else {
					client.HSet(tableName, fieldType.Name, "0")
				}
			case reflect.Slice:
				if fieldType.Type.Elem().Kind() == reflect.Uint8 {
					client.HSet(tableName, fieldType.Name, string(fieldValue.Interface().([]byte)))
				}
			}
			return nil
		}
	}
	return errors.New("field not existst")
}

