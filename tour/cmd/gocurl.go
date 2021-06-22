package cmd

import (
	"fmt"
	"strconv"
	"time"
	"tools/tour/intertnal/gocurl"
	"tools/tour/intertnal/selfjson"
	"tools/tour/intertnal/toolmysql"

	"github.com/spf13/cobra"
)

var url string

var gocurlCmd = &cobra.Command{
	Use:   "curl",
	Short: "curl请求",
	Long:  "curl请求",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("输入url", str)
		url := "https://restapi.amap.com/v3/config/district?keywords=110108&subdistrict=0&key=2caa5a3c8b92ad1e1c2dcd5437975a01&extensions=all"
		res, err := gocurl.Get(url)
		if err == nil {
			res = res
		}

	},
}

//go run main.go curl district
var getdistrictCmd = &cobra.Command{
	Use:   "district",
	Short: "district 获取高德地图行政区坐标 围栏 等级等数据",
	Long:  "district 获取高德地图行政区坐标 围栏 等级等数据",
	Run: func(cmd *cobra.Command, args []string) {
		citys := toolmysql.StructQueryAllField()
		var url string
		for _, v := range citys {
			//fmt.Println(v)
			url = "https://restapi.amap.com/v3/config/district?keywords=" + strconv.FormatInt(v.Code, 10) + "&subdistrict=0&key=2caa5a3c8b92ad1e1c2dcd5437975a01&extensions=all"

			res, err := gocurl.Get(url)
			if err == nil {
				a := selfjson.Json2struct(res)

				for _, v := range a.Districts {
					// fmt.Println(v)
					toolmysql.StructUpdate(v.Polyline, v.Center, v.Level, v.Adcode)
				}
			}
			time.Sleep(time.Microsecond * 200)
		}
	},
}

//注册子命令
func init() {
	gocurlCmd.AddCommand(getdistrictCmd)
}
