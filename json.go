package rockjson

import (
	"errors"
	"encoding/json"
	"reflect"
)

type JsonObject = map[string]interface{}
type JsonArray = []interface{}

func (jsobj JsonObject) GetObj(key string) (interface{}, error) {
	obj := jsobj[key]
	if obj == nil {
		return nil, errors.New("key不存在")
	}
	return obj, nil
}

func (jsonobj JsonObject) GetJsonObject(key string) (JsonObject, error) {
	obj := jsonobj[key]
	if obj == nil {
		return nil, errors.New("key不存在")
	}
	res, ok := obj.(JsonObject)
	if !ok {
		return nil, errors.New("转为JsonObject失败")
	}
	return res, nil
}

func (jsonobj JsonObject) GetJsonArray(key string) (JsonArray, error) {
	obj := jsonobj[key]
	if obj == nil {
		return nil, errors.New("key不存在")
	}
	resArray, ok := obj.(JsonArray)

	if !ok {
		return nil, errors.New("转为JsonArray失败")
	}
	return resArray, nil
}

func (jsonobj JsonObject) GetString(key string, defaultstr string) (string, error) {
	obj := jsonobj[key]
	if obj == nil {
		return defaultstr, errors.New("key不存在")
	}
	res, ok := obj.(string)

	if !ok {
		return defaultstr, errors.New(key + "的内容转为string失败")
	}
	return res, nil
}
func (jsonobj JsonObject) GetInt64(key string, defaultarg int64) (int64, error) {
	obj := jsonobj[key]
	if obj == nil {
		return defaultarg, errors.New("key不存在")
	}

	switch obj.(type) {
	case json.Number:
		return obj.(json.Number).Int64()
	case float32, float64:
		return int64(reflect.ValueOf(obj).Float()), nil
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(obj).Int(), nil
	case uint, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(obj).Uint()), nil
	}
	return defaultarg, errors.New("数据类型错误")
}
