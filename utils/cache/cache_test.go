package cache

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	t.Log("####TestCache test ")

	Set("test", "1")

	resp := Get("test")
	if !resp.Success() {
		t.Error("TestCache set get err")
	}

	if gconv.String(resp.Data) != "1" {
		t.Error("TestCache set value err")
	}

	Del("test")
	resp = Get("test")
	if resp.Success() {
		t.Error("TestCache del get err")
	}

}

func TestCacheTimeout(t *testing.T) {
	t.Log("####TestCacheTimeout test ")

	Setex("test", "1", 1200)

	resp := Get("test")
	if !resp.Success() {
		t.Error("TestCacheTimeout set get err")
	}

	if gconv.String(resp.Data) != "1" {
		t.Error("TestCacheTimeout set value err")
	}

	time.Sleep(time.Second * 1)

	resp = Get("test")
	if !resp.Success() {
		t.Error("TestCacheTimeout set get err")
	}

	time.Sleep(time.Second * 1)

	resp = Get("test")
	if resp.Success() {
		t.Error("TestCacheTimeout set get err")
	}

}

func TestCacheMap(t *testing.T) {
	t.Log("####TestCacheMap test ")

	SetMap("test", g.Map{"a": "1"})

	resp := GetMap("test")
	if !resp.Success() {
		t.Error("TestCacheMap set get err")
	}

	if gconv.Map(resp.Data)["a"] != "1" {
		t.Error("TestCacheMap set value err")
	}

	Del("test")
	resp = GetMap("test")
	if resp.Success() {
		t.Error("TestCacheMap del get err")
	}

}

func TestCacheTimeoutMap(t *testing.T) {
	t.Log("####TestCacheTimeoutMap test ")

	SetexMap("test", g.Map{"a": "1"}, 1200)

	resp := GetMap("test")
	if !resp.Success() {
		t.Error("TestCacheTimeoutMap set get err")
	}

	if gconv.Map(resp.Data)["a"] != "1" {
		t.Error("TestCacheTimeoutMap set value err")
	}

	time.Sleep(time.Second * 1)

	resp = GetMap("test")
	if !resp.Success() {
		t.Error("TestCacheTimeoutMap set get err")
	}

	time.Sleep(time.Second * 1)

	resp = GetMap("test")
	if resp.Success() {
		t.Error("TestCacheTimeoutMap set get err")
	}

}
