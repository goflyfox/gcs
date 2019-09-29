package test

import (
	"encoding/json"
	"gcs/utils/resp"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"testing"
)

const (
	TestURL string = "http://127.0.0.1:80"
)

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
		defer r.Close()

		content := string(r.ReadAll())
		var respData resp.Resp
		err := json.Unmarshal([]byte(content), &respData)
		if err != nil {
			t.Error(err)
		}

		if !respData.Success() {
			t.Error("resp fail:" + respData.Json())
		}
		t.Log(content)
	}

}

func TestApiVersionFail(t *testing.T) {
	name := "test"
	secretKey := "12345678XX"
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
		defer r.Close()

		content := string(r.ReadAll())
		var respData resp.Resp
		err := json.Unmarshal([]byte(content), &respData)
		if err != nil {
			t.Error(err)
		}

		if respData.Success() {
			t.Error("resp fail:" + respData.Json())
		}
		t.Log(content)
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
		defer r.Close()

		content := string(r.ReadAll())
		var respData resp.Resp
		err := json.Unmarshal([]byte(content), &respData)
		if err != nil {
			t.Error(err)
		}

		if !respData.Success() {
			t.Error("resp fail:" + respData.Json())
		}
		t.Log(content)
	}
}
