package word

import (
	"strings"
	"unicode"
)

//Toupper 转换为大写
func Toupper(s string) string {
	return strings.ToUpper(s)
}

//ToLower 转换为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

//UnderscoreToUpperCamelCase 转换为大写驼峰
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

//UnderscoreToLowerCamelCase 转换为小写驼峰
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]

}

//CamelCaseToUndersCore 驼峰转下划线单词
func CamelCaseToUndersCore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))

	}
	return string(output)
}
