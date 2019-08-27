package main

import (
	"github.com/gogf/gf/g/crypto/gmd5"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/os/gtime"
	"github.com/gogf/gf/g/text/gstr"
	"strings"
	"testing"
)

func TestCaa(t *testing.T) {
	strings.Repeat(",?", -1)
}

func TestApiVersion(t *testing.T) {
	name := "test"
	secretKey := "12345678"
	no := gtime.Now().Format("YmdHis")

	keyStr, err := gmd5.Encrypt(name + no + secretKey)
	if err != nil {
		glog.Error("mac encrypt error", err)
	}
	mac := gstr.ToLower(keyStr)
	params := "?name=" + name + "&no=" + no + "&mac=" + mac
	glog.Info(TestURL + "/config/api/version" + params)
	if r, e := ghttp.Get(TestURL + "/config/api/version" + params); e != nil {
		t.Error(e)
	} else {
		t.Log(string(r.ReadAll()))
		r.Close()
	}

}

func TestApiData(t *testing.T) {
	name := "test"
	secretKey := "12345678"
	no := gtime.Now().Format("YmdHis")

	keyStr, err := gmd5.Encrypt(name + no + secretKey)
	if err != nil {
		glog.Error("mac encrypt error", err)
	}
	mac := gstr.ToLower(keyStr)
	params := "?name=" + name + "&no=" + no + "&mac=" + mac
	glog.Info(TestURL + "/config/api/data" + params)
	if r, e := ghttp.Get(TestURL + "/config/api/data" + params); e != nil {
		t.Error(e)
	} else {
		t.Log(string(r.ReadAll()))
		r.Close()
	}
}
