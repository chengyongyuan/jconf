package jconf

import (
	"fmt"
	"testing"
)

var basicTest = []struct {
	Path       string
	ServerName string
	IPLIST     []string
	Port       []int
	ID         int
}{
	{
		Path:       "testdata/basic.json",
		ServerName: "testserver",
		IPLIST:     []string{"127.0.0.1", "192.168.0.1", "192.168.0.3"},
		Port:       []int{80, 443, 14000},
		ID:         8888,
	},
	{
		Path:       "testdata/simple.conf",
		ServerName: "colin't test machine",
		IPLIST:     []string{"10.137.2.221", "127.0.0.1"},
		Port:       []int{80, 443, 14000},
		ID:         9999,
	},
}

var simpleConfTest = []struct {
	Path       string
	ServerName string
	IPLIST     []string
	Port       []int
	ID         int
}{
	{
		Path:       "testdata/simple_sect.conf",
		ServerName: "15T",
		IPLIST:     []string{"8.8.8.8", "1.1.1.1"},
		Port:       []int{80, 443},
		ID:         1,
	},
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

func TestSimpleConfSect(tt *testing.T) {
	for _, t := range simpleConfTest {
		fmt.Printf("Testing conf file: %s\n", t.Path)
		r, err := NewConfReader(t.Path)
		if err != nil {
			tt.Error("Init fail" + err.Error())
		}
		if err := r.Init(); err != nil {
			tt.Error("Init fail" + err.Error())
		}
		//test misc sector
		s := r.GetStr("misc.ServerName", "default")
		if s != t.ServerName {
			tt.Errorf("GetStr Fail! Extpect:%v, Real:%v", t.ServerName, s)
		}
		sa := r.GetStrArray("main.IPLIST", []string{})
		if !eqStrArray(sa, t.IPLIST) {
			tt.Errorf("GetStrArray Fail! Extpect:%v, Real:%v", t.IPLIST, sa)
		}
		//test missing sector
		ia := r.GetIntArray("Port", []int{})
		if !eqIntArray(ia, []int{}) {
			tt.Errorf("GetIntArray Fail! Extpect:%v, Real:%v", t.Port, ia)
		}
		//test main sector
		ia = r.GetIntArray("main.Port", []int{})
		if !eqIntArray(ia, t.Port) {
			tt.Errorf("GetIntArray Fail! Extpect:%v, Real:%v", t.Port, ia)
		}
		i := r.GetInt("misc.ID", 0)
		if i != t.ID {
			tt.Errorf("GetInt Fail! Extpect:%v, Real:%v", t.ID, i)
		}
		//test int default
		i = r.GetInt("intkey", -1)
		if i != -1 {
			tt.Errorf("GetInt Fail! Extpect:%v, Real:%v", -1, i)
		}

		//test string default
		s = r.GetStr("strkey", "default")
		if s != "default" {
			tt.Errorf("GetStr Fail! Extpect:%v, Real:%v", "default", s)
		}
	}
}

func TestBasic(tt *testing.T) {
	for _, t := range basicTest {
		fmt.Printf("Testing conf file: %s\n", t.Path)
		r, err := NewConfReader(t.Path)
		if err != nil {
			tt.Error("Init fail!" + err.Error())
		}
		if err := r.Init(); err != nil {
			tt.Error("Init fail!" + err.Error())
		}
		//if err := Init(t.Path); err != nil {
		//	tt.Error("Init fail!" + err.Error())
		//}
		//test get string
		s := r.GetStr("ServerName", "default")
		if s != t.ServerName {
			tt.Errorf("GetStr Fail! Extpect:%v, Real:%v", t.ServerName, s)
		}
		//test get string Array
		sa := r.GetStrArray("IPLIST", []string{})
		if !eqStrArray(sa, t.IPLIST) {
			tt.Errorf("GetStrArray Fail! Extpect:%v, Real:%v", t.IPLIST, sa)
		}

		//test get int array
		ia := r.GetIntArray("Port", []int{})
		if !eqIntArray(ia, t.Port) {
			tt.Errorf("GetIntArray Fail! Extpect:%v, Real:%v", t.Port, ia)
		}

		//test get int
		i := r.GetInt("ID", 0)
		if i != t.ID {
			tt.Errorf("GetInt Fail! Extpect:%v, Real:%v", t.ID, i)
		}

		//test int default
		i = r.GetInt("intkey", -1)
		if i != -1 {
			tt.Errorf("GetInt Fail! Extpect:%v, Real:%v", -1, i)
		}

		//test string default
		s = r.GetStr("strkey", "default")
		if s != "default" {
			tt.Errorf("GetStr Fail! Extpect:%v, Real:%v", "default", s)
		}
	}
}
