package cmd

import (
	"log"
	"tools/tour/intertnal/readtowrite"

	"github.com/spf13/cobra"
)

var readpath string
var writepath string

var readtowriteCmd = &cobra.Command{
	Use:   "readtowrite",
	Short: "文件读取和写入",
	Long:  "文件读取和写入",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("请输入自命令调用,读取文件 -r 写入文件 -w")
	},
}

//go run main.go readtowrite  readtoarray -r='/home/zlc/goProject/src/tools/ceshi.txt'
var readtoarrayCmd = &cobra.Command{
	Use:   "readtoarray",
	Short: "文件读取和写入",
	Long:  "文件读取和写入",
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan string, 2)
		go readtowrite.Readfile(readpath, ch)
		for {
			//fmt.Println(<-ch)
			tmp := <-ch
			if "|end|" == tmp { //结束标识位
				break
			}
			readtowrite.Writefile(writepath, tmp)
		}

	},
}

//注册子命令
func init() {
	readtowriteCmd.AddCommand(readtoarrayCmd)
	readtoarrayCmd.Flags().StringVarP(&readpath, "readpath", "r", "", "读取文件的路径")
	readtoarrayCmd.Flags().StringVarP(&writepath, "writepath", "w", "", "写入文件的路径")
}
