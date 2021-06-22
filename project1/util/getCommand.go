package util

import (
	"strings"

	flag "github.com/spf13/pflag"
)

// 定义命令行参数对应的变量
var CliJsonPath = flag.StringP("json-path", "p", "./config/config.army.model.json", "Input Json Path")

func wordSepNormalizeFunc(f *flag.FlagSet, name string) flag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return flag.NormalizedName(name)
}

func GetJsonPath() string {
	// 设置标准化参数名称的函数
	flag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)
	flag.Parse()
	return *CliJsonPath
}
