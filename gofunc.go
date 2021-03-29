//Time    : 2021/3/20 9:37 下午
//Author  : xushiyin
//contact : yuqingxushiyin@gmail.com
package testmodule

import (
	"fmt"
	"runtime"
)

// GoGetVersion 获取当前Go版本，也可通过 go env 查看
func GoGetVersion() string {
	return fmt.Sprintf("we are using Go: <%s>, Go GOARCH: <%s>", runtime.Version(), runtime.GOARCH)
}
