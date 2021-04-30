package cmd

import (
	"log"
	"tools/tour/intertnal/timer"

	"github.com/spf13/cobra"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

//获取当前时间
var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("当前时间: %s,%d,", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

//注册子命令
func init() {
	timeCmd.AddCommand(nowTimeCmd)
}
