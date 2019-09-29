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
	BenchmarkURL = "http://127.0.0.1:80"
)

func BenchmarkTmp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := gmd5.Encrypt("123123123")
		if err != nil {
			glog.Error("mac encrypt error", err)
		}
	}
}

func BenchmarkApiVersion(b *testing.B) {
	name := "test"
	secretKey := "12345678"
	no := gtime.Now().Format("YmdHis")

	for i := 0; i < b.N; i++ {
		keyStr, err := gmd5.Encrypt(name + no + secretKey)
		if err != nil {
			glog.Error("mac encrypt error", err)
		}
		mac := gstr.ToLower(keyStr)
		params := "?name=" + name + "&no=" + no + "&mac=" + mac
		r, e := ghttp.Get(TestURL + "/config/api/version" + params)
		defer r.Close()
		if e != nil {
			b.Error(e)
		} else {
			// b.Log(string(r.ReadAll()))
		}
	}
}

func BenchmarkLogin(b *testing.B) {
	passwd, _ := gmd5.Encrypt("123456")

	for i := 0; i < b.N; i++ {
		if r, e := ghttp.Post(TestURL+"/login/submit", "username="+Username+"&passwd="+passwd); e != nil {
			b.Error(e)
		} else {
			defer r.Close()

			content := string(r.ReadAll())
			var respData resp.Resp
			err := json.Unmarshal([]byte(content), &respData)
			if err != nil {
				b.Error(err)
			}

			if !respData.Success() {
				b.Error("resp fail:" + respData.Json())
			}

		}
	}
}
