package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"
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

//推算时间
var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTmer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTmer = timer.GetNowTime()
		} else {
			var err error

			_, err1 := strconv.ParseFloat(calculateTime, 64) //判断不是时间戳(纯数字),时间戳格式不变

			if !strings.Contains(calculateTime, " ") && err1 != nil {
				layout = "2006-01-02"
			}

			currentTmer, err = time.Parse(layout, calculateTime)

			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTmer = time.Unix(int64(t), 0)
			}

		}

		if duration == "" { //没有值的时候默认为0h
			duration = "0h"
		}

		calculateTime, err := timer.GetCaulculateTime(currentTmer, duration)

		if err != nil {
			log.Fatalf("time.GetCaulculateTime err : %v", err)
		}

		log.Printf("输出结果: %s, %d", calculateTime.Format(layout), calculateTime.Unix())

	},
}

//注册子命令
func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要计算的时间,有效果格式为时间戳或已格式化后的时间")

	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间有效时间为 "ns","us","ms","s","m","h"`)
}
