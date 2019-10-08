package cache

import (
	"github.com/gogf/gf/util/gconv"
	"testing"
)

func TestCache(t *testing.T) {
	t.Log("cache test ")

	Set("test", "1")

	resp := Get("test")
	if !resp.Success() {
		t.Error("cache set get err")
	}

	if gconv.String(resp.Data) != "1" {
		t.Error("cache set value err")
	}

}
