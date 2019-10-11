package task

import (
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/glog"
)

func init() {
	glog.Info("task init...")

	gcron.Add("0 * * * * *", checkVersion)

	glog.Info("task finish")
}
