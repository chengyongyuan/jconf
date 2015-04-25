package jconf

import (
	"testing"
)

var basicTest = struct {
	Path       string
	ServerName string
	IPLIST     []string
	Port       []int
	ID         int
}{
	Path:       "basic.json",
	ServerName: "testserver",
	IPLIST:     []string{"127.0.0.1", "192.168.0.1", "192.168.0.3"},
	Port:       []int{80, 443, 14000},
	ID:         8888,
}

func eqIntArray(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func eqStrArray(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestBasic(t *testing.T) {
	if err := Init(basicTest.Path); err != nil {
		t.Error("Init fail!" + err.Error())
	}
	//test get string
	s := GetStr("ServerName", "default")
	if s != basicTest.ServerName {
		t.Error("GetStr Fail! Extpect:%v, Real:%v", basicTest.ServerName, s)
	}
	//test get string Array
	sa := GetStrArray("IPLIST", []string{})
	if !eqStrArray(sa, basicTest.IPLIST) {
		t.Error("GetStrArray Fail! Extpect:%v, Real:%v", basicTest.IPLIST, sa)
	}

	//test get int array
	ia := GetIntArray("Port", []int{})
	if !eqIntArray(ia, basicTest.Port) {
		t.Error("GetIntArray Fail! Extpect:%v, Real:%v", basicTest.Port, ia)
	}

	//test get int
	i := GetInt("ID", 0)
	if i != basicTest.ID {
		t.Error("GetInt Fail! Extpect:%v, Real:%v", basicTest.ID, i)
	}

	//test int default
	i = GetInt("intkey", 8888)
	if i != 8888 {
		t.Error("GetInt Fail! Extpect:%v, Real:%v", 8888, i)
	}

	//test string default
	s = GetStr("strkey", "default")
	if s != "default" {
		t.Error("GetStr Fail! Extpect:%v, Real:%v", "default", s)
	}
}
