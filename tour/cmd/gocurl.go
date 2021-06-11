package cmd

import (
	"fmt"
	"tools/tour/intertnal/gocurl"

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

//注册子命令
func init() {

}
