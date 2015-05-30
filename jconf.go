package jconf

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

//package jconf is a simple json based config module in Go.
//It exports those Interfaces:
//func Init(path string) error
//func GetInt(key string, def int) int
//func GetIntArray(key string, def[]int) []int
//func GetStr(key, def string) string
//func GetStrArray(key string, def []string) []string

const (
	CONF_TYPE_SIMPLE = iota
	CONF_TYPE_JSON
)

type ConfReader interface {
	Init() error
	GetInt(key string, def int) (val int)
	GetIntArray(key string, def []int) (val []int)
	GetStr(key string, def string) (val string)
	GetStrArray(key string, def []string) (val []string)
}

type JsonConf struct {
	ConfName string
	ConfType int
}

type SimpleConf struct {
	ConfName string
	ConfType int
}

//Internal map
var confMap map[string]interface{}
var sConfMap map[string]string
var confType int

func NewConfReader(path string) (r ConfReader, err error) {
	if strings.HasSuffix(path, ".json") {
		return &JsonConf{ConfName: path, ConfType: CONF_TYPE_JSON}, nil
	}
	if strings.HasSuffix(path, ".conf") {
		return &SimpleConf{ConfName: path, ConfType: CONF_TYPE_SIMPLE}, nil
	}

	return nil, errors.New("jconf: unsupported conf file!")
}

func (c *JsonConf) Init() (err error) {
	str, err := ioutil.ReadFile(c.ConfName)
	if err != nil {
		err = errors.New(fmt.Sprintf("jconf: read %s file fail!", c.ConfName))
		return
	}
	if err = json.Unmarshal(str, &confMap); err != nil {
		err = errors.New(fmt.Sprintf("jconf: parse json file error(%s)", err.Error()))
		return
	}
	return
}

func (c *JsonConf) GetInt(key string, def int) (val int) {
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

func (c *JsonConf) GetIntArray(key string, def []int) (val []int) {
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

func (c *JsonConf) GetStr(key string, def string) (val string) {
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

func (c *JsonConf) GetStrArray(key string, def []string) (val []string) {
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

func (c *SimpleConf) Init() (err error) {
	str, err := ioutil.ReadFile(c.ConfName)
	if err != nil {
		err = errors.New(fmt.Sprintf("jconf: read %s file fail!", c.ConfName))
		return
	}
	sConfMap = make(map[string]string)
	if e := parseSimpleConf(string(str)); e != nil {
		err = errors.New(fmt.Sprintf("jconf: parse conf file error(%s)", err.Error()))
		return
	}
	return
}

//parse simple conf file, store each element in map
func parseSimpleConf(str string) (err error) {
	re := regexp.MustCompile(`\[(.*)\]`)
	reLine := regexp.MustCompile(`\s*(\S+)\s*=\s*(.*)`)
	r := bufio.NewReader(strings.NewReader(str))
	prefix := ""
	for {
		line, e := r.ReadString('\n')
		if e == io.EOF {
			break
		}
		if strings.Trim(line, " ") == "\n" { //skip empty line
			continue
		}
		s := re.FindStringSubmatch(line)
		if len(s) == 2 {
			prefix = s[1]
			continue
		}
		l := reLine.FindStringSubmatch(line)
		if len(l) != 3 {
			continue
		}
		var k string
		if prefix != "" {
			k = prefix + "." + l[1]
		} else {
			k = l[1]
		}
		v := l[2]
		sConfMap[k] = v
	}
	return
}

func (c *SimpleConf) GetInt(key string, def int) (val int) {
	if _, ok := sConfMap[key]; !ok {
		val = def
		return
	}
	val, err := strconv.Atoi(sConfMap[key])
	if err != nil {
		val = def
		return
	}
	return
}

func (c *SimpleConf) GetIntArray(key string, def []int) (val []int) {
	if _, ok := sConfMap[key]; !ok {
		val = def
		return
	}
	list := strings.Split(sConfMap[key], ",")
	for _, v := range list {
		v, err := strconv.Atoi(strings.Trim(v, " "))
		if err != nil {
			val = def
			return
		}
		val = append(val, v)
	}

	return
}

func (c *SimpleConf) GetStr(key string, def string) (val string) {
	if _, ok := sConfMap[key]; !ok {
		val = def
		return
	}
	val = sConfMap[key]

	return
}

func (c *SimpleConf) GetStrArray(key string, def []string) (val []string) {
	if _, ok := sConfMap[key]; !ok {
		val = def
		return
	}
	list := strings.Split(sConfMap[key], ",")
	for _, v := range list {
		val = append(val, strings.Trim(v, " "))
	}

	return
}
