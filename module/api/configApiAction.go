package api

type SysConfig struct {
	// columns START
	No    string `json:"no" gconv:"no,omitempty"`       // 标识
	Name  string `json:"name" gconv:"name,omitempty"`   // 名称
	Key   string `json:"key" gconv:"key,omitempty"`     // 键
	Value string `json:"value" gconv:"value,omitempty"` // 值
	Code  string `json:"code" gconv:"code,omitempty"`   // 编码
	Mac   string `json:"mac" gconv:"mac,omitempty"`     // MAC
	// columns END
}
