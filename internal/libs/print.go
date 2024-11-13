package libs

import (
	"github.com/k0kubun/pp"
)

// DebugPrint 封装 pp.Println()，用于打印调试信息
func DebugPrint(v ...interface{}) {
	// 设置 pp 包的颜色输出
	pp.ColoringEnabled = true
	pp.Println(v...)
}
