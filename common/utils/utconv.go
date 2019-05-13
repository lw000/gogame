package ggutils

import (
	"errors"
	"reflect"
)

func GetFloat64(v interface{}) (float64, error) {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Float64 {
		return v.(float64), nil
	}
	return 0, errors.New("数据类型错误")
}

func GetInt(v interface{}) (int, error) {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Int {
		return v.(int), nil
	}
	return 0, errors.New("数据类型错误")
}

func GetInt32(v interface{}) (int32, error) {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Int32 {
		return v.(int32), nil
	}
	return 0, errors.New("数据类型错误")
}

func GetInt64(v interface{}) (int64, error) {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Int64 {
		return v.(int64), nil
	}
	return 0, errors.New("数据类型错误")
}

func GetString(v interface{}) (string, error) {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.String {
		return v.(string), nil
	}
	return "", errors.New("数据类型错误")
}
