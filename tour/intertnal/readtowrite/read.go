package readtowrite

import (
	"bufio"
	"io"
	"os"
)

//Readfile 按行读取文件
func Readfile(readpath string, ch chan string) error {

	f, err := os.Open(readpath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		ch <- line
		//fmt.Println(<-ch)
	}
	ch <- "|end|"

	return nil
}

//PathExists 判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
