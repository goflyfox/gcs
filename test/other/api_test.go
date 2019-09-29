package other_test

import (
	"fmt"
	"gcs/module/system"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"reflect"
	"testing"
)

func TestObject(t *testing.T) {
	arr := []*system.SysConfig{}
	arr = append(arr, &system.SysConfig{Id: 1, Value: "111"})
	arr = append(arr, &system.SysConfig{Id: 2, Value: "222"})
	for _, data := range arr {
		data.Id = 0
	}
	for _, data := range arr {
		glog.Info(data)
	}

}

func TestRun(t *testing.T) {
	foreach("1", "2")
	var paramsStr []string
	paramsStr = append(paramsStr, "3")
	paramsStr = append(paramsStr, "4")
	foreach(paramsStr)

	var params []interface{}
	params = append(params, "5")
	params = append(params, "6")
	foreach(params)
	foreach(params)

}

func TestReflect(t *testing.T) {
	model := system.SysConfig{}
	model.UpdateId = 1
	re := reflect.ValueOf(model).FieldByName("BaseModel")
	updateId := gconv.Int(re.FieldByName("UpdateId").Interface())
	fmt.Println(updateId)
}

func foreach(args ...interface{}) {
	for _, a := range args {
		fmt.Println("#" + gconv.String(a))
	}
}

func TestFile(t *testing.T) {
	gfile.CopyFile("D:\\work\\ddddd.dsdr", "D:\\work\\ddddd.dsdr.bak")
	content := gfile.GetBytes("D:\\work\\ddddd.dsdr.bak")
	gfile.PutBytes("D:\\work\\ddddd.dsdr.bak1", content)
	array := gstr.Split(string(content), "\r\n")
	datas := gstr.Split(array[0], ",")
	fmt.Println(array[0])
	fmt.Println(datas[7])
	fmt.Println(len(content) / (1024 * 1024))
	fmt.Println(len(array))

	//for _, data := range array {
	//	datas := gstr.Split(data, ",")
	//	fmt.Println(datas)
	//}

	list, _ := gfile.DirNames("D:\\test")
	for _, name := range list {
		fmt.Println(name)
	}
}
