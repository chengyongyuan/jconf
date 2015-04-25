package jconf

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

//package jconf is a simple json based config module in Go.
//It exports those Interfaces:
//func Init(path string) error
//func GetInt(key string, def int) int
//func GetIntArray(key string, def[]int) []int
//func GetStr(key, def string) string
//func GetStrArray(key string, def []string) []string

//Internal map
var confMap map[string]interface{}

func Init(path string) (err error) {
	jstr, e := ioutil.ReadFile(path)
	if e != nil {
		err = errors.New(fmt.Sprintf("jconf: read %s file fail!", path))
		return
	}
	if e := json.Unmarshal(jstr, &confMap); e != nil {
		err = errors.New(fmt.Sprintf("jconf: parse json file error(%s)", err.Error()))
		return
	}
	return
}

func GetInt(key string, def int) (val int) {
	if _, ok := confMap[key]; !ok {
		val = def
		return
	}
	//go take json number as float64
	if _, ok := confMap[key].(float64); !ok {
		val = def
		return
	}
	val = int(confMap[key].(float64))
	return
}

func GetIntArray(key string, def []int) (val []int) {
	if _, ok := confMap[key]; !ok {
		val = def
		return
	}
	if _, ok := confMap[key].([]interface{}); !ok {
		val = def
		return
	}
	for _, s := range confMap[key].([]interface{}) {
		v, ok := s.(float64)
		if !ok {
			val = def
			return
		}
		val = append(val, int(v))
	}
	return
}

func GetStr(key string, def string) (val string) {
	if _, ok := confMap[key]; !ok {
		val = def
		return
	}
	if _, ok := confMap[key].(string); !ok {
		val = def
		return
	}
	val = confMap[key].(string)
	return
}

func GetStrArray(key string, def []string) (val []string) {
	if _, ok := confMap[key]; !ok {
		val = def
		return
	}
	if _, ok := confMap[key].([]interface{}); !ok {
		val = def
		return
	}
	for _, s := range confMap[key].([]interface{}) {
		v, ok := s.(string)
		if !ok {
			val = def
			return
		}
		val = append(val, v)
	}
	return
}
