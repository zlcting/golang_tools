package cmd

import (
	"log"
	"strings"
	"tools/tour/intertnal/word"

	"github.com/spf13/cobra"
)

const (
	MODE_UPPER                        = iota + 1 //全部单词转为大写
	MODE_LOWER                                   //全部转为小写
	MODE_UNDERSCORE_TOUPPER_CAMELCASE            //下划线转为大写驼峰
	MODE_UNDERSCORE_TOLOWER_CAMELCASE            //下划线单词转为小写驼峰
	MODE_CAMELCASE_TOUNDERSCORE                  //驼峰单词转为下划线单词
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1:全部单词转为大写",
	"2:全部转为小写",
	"3:下划线转为大写驼峰",
	"4:下划线单词转为小写驼峰",
	"5:驼峰单词转为下划线单词",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case MODE_UPPER:
			content = word.Toupper(str)
		case MODE_LOWER:
			content = word.ToLower(str)
		case MODE_UNDERSCORE_TOUPPER_CAMELCASE:
			content = word.UnderscoreToUpperCamelCase(str)
		case MODE_UNDERSCORE_TOLOWER_CAMELCASE:
			content = word.UnderscoreToLowerCamelCase(str)
		case MODE_CAMELCASE_TOUNDERSCORE:
			content = word.CamelCaseToUndersCore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行help word 查看帮助文档")
		}
		log.Printf("输出结果：%s", content)
	},
}
