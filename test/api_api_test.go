package main

import (
	"testing"
)

func TestApiVersion(t *testing.T) {
	// 登录，访问用户信息
	//name := "test"
	//secretKey := "12345678"
	//no := gtime.str
	//
	//keyStr, err := gmd5.Encrypt(name + no + secretKey)
	//if err != nil {
	//	glog.Error("mac encrypt error", err)
	//	return resp.Fail("mac encrypt error")
	//}
	//mac :=
	//params := ""
	//if r, e := ghttp.Get(TestURL + "/config/api/version" + params); e != nil {
	//	t.Error(e)
	//} else {
	//	t.Log(string(r.ReadAll()))
	//	r.Close()
	//}

}

func TestApiData(t *testing.T) {
	// 登录，访问用户信息
	params := ""
	data := Post(t, "/system/user/get"+params, "id=1")
	if data.Success() {
		t.Log(data.Json())
	} else {
		t.Error(data.Json())
	}
}
