package initialize

import "github.com/yitter/idgenerator-go/idgen"

// SnowFlakeInitialize 雪花算法初始化
func SnowFlakeInitialize() {
	// 创建 IdGeneratorOptions 对象，可在构造函数中输入 WorkerId：
	var options = idgen.NewIdGeneratorOptions(12)
	idgen.SetIdGenerator(options)
}
