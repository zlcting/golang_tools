package readtowrite

import (
	"fmt"
	"io"
	"os"
)

//Writefile 按行写入文件
func Writefile(writepath string, line string) error {
	var f *os.File
	var err error
	if checkFileIsExist(writepath) { //如果文件存在
		f, err = os.OpenFile(writepath, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err = os.Create(writepath) //创建文件
		fmt.Println("文件不存在")
	}
	fmt.Println(err)
	io.WriteString(f, line)
	//fmt.Printf("写入 %d 个字节n", n)
	//fmt.Println(err)

	if err != nil {
		return err
	}
	return nil
}

//checkFileIsExist 检查路径存不存在
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
