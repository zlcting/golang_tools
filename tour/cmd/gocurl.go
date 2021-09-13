package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"tools/tour/intertnal/gocurl"
	"tools/tour/intertnal/help"
	"tools/tour/intertnal/readtowrite"
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

var locationCmd = &cobra.Command{
	Use:   "location",
	Short: "location 坐标",
	Long:  "location 坐标",
	Run: func(cmd *cobra.Command, args []string) {
		var url string
		var str string
		dictStatus := map[int]string{
			1:  "正常",
			-1: "删除",
			-2: "隐藏",
		}

		dictsale := map[int]string{
			0:  "",
			1:  "新盘",
			2:  "在售",
			3:  "尾盘",
			10: "待售",
			4:  "已售完",
			11: "停售",
			5:  "出租",
			6:  "新盘-商业",
			7:  "在售-商业",
			8:  "租售-商业",
			9:  "满租",
		}
		dictTmall := map[int]string{
			0: "无状态",
			1: "已推送",
			2: "上架",
			3: "下架",
		}
		var Eunified_district string
		var Eunified_district_code string
		var Eunified_city_cn string

		i := 0
		for {
			i++
			//houses := toolmysql.StructHouseflied(i, 5000)
			houses := toolmysql.StructHousefliedId()
			for _, v := range houses {
				url = "https://restapi.amap.com/v3/geocode/regeo?key=c7894840f8de303c5b556d509f395cfb&location=" + v.Coordx2 + "," + v.Coordy2
				fmt.Println(url)
				res, err := gocurl.Get(url)
				if err == nil {
					a := selfjson.Json2Gaodecitybylocation(res)
					Eunified_city_cn = a.Regeocode.AddressComponent.City
					Eunified_district_code = a.Regeocode.AddressComponent.Adcode
					Eunified_district = a.Regeocode.AddressComponent.District

					fmt.Println(a)
				}

				flag := help.Citycompare(v.Site, v.Unified_city_cn_code, Eunified_district_code)

				if !flag {
					// fmt.Println(v)
					str = strconv.Itoa(i) + "," + v.City_cn + "," + v.Value + "," + strconv.Itoa(v.Id) + "," + v.Name + "," + time.Unix(int64(v.Createtime), 0).Format("2006-01-02 15:04:05") + "," + dictStatus[v.Status] + "," + dictsale[v.Salestate] + "," + dictTmall[v.Is_push_tmall] + "," + dictTmall[v.Is_push_lejuflag] + "," + Eunified_city_cn + "," + Eunified_district + "\n"
					fmt.Println(str)
					readtowrite.Writefile("/home/zlc/goProject/src/tools/houseyidi2.txt", str)
				}

			}
			//time.Sleep(time.Microsecond * 200)
			//break
			if len(houses) == 0 {
				break
			}
			break
		}
	},
}

var yidihouseCmd = &cobra.Command{
	Use:   "yidhouse",
	Short: "yidhouse",
	Long:  "yidhouse",
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan string)
		//var str string
		db, _ := toolmysql.NewDb()

		go readtowrite.Readfile("/home/zlc/goProject/src/tools/houseyidi.txt", ch)
		//go readtowrite.Readfile("/home/zlc/goProject/src/tools/houseyidi_new.csv", ch)

		for {
			tmp, ok := <-ch
			fmt.Println(tmp)

			if !ok {
				break
			}

			countSplit := strings.Split(tmp, ",")
			if len(countSplit) > 3 {
				dbmap, _ := db.Query(countSplit[2])
				aa, _ := db.HouseQuery(dbmap["city_en"], countSplit[3])

				str = strings.Replace(tmp, "\n", "", -1) + "," + dbmap["unified_city_cn_code"] + "," + aa["city_en"] + "," + aa["hid"] + "," + aa["id"] + "\n"
				fmt.Println(str)

				readtowrite.Writefile("/home/zlc/goProject/src/tools/houseyidi_new2.csv", str)
			}

		}

	},
}

//go run main.go gocurlCmd locationtoaddress
var locationtoaddressCmd = &cobra.Command{
	Use:   "locationtoaddress",
	Short: "locationtoaddress",
	Long:  "locationtoaddress",
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan string)
		//var str string
		go readtowrite.Readfile("/home/zlc/goProject/src/tools/1112.csv", ch)

		for {
			tmp, ok := <-ch
			//fmt.Println(tmp)
			if !ok {
				break
			}
			countSplit := strings.Split(tmp, ",")
			if len(countSplit) < 2 {
				continue
			}

			url = "https://restapi.amap.com/v3/geocode/regeo?key=c7894840f8de303c5b556d509f395cfb&location=" + countSplit[1] + "," + strings.Replace(countSplit[2], "\r\n", "", -1) + "&poitype=&radius=100&extensions=all&batch=false&roadlevel=0"
			// url = "https://restapi.amap.com/v3/geocode/regeo?key=c7894840f8de303c5b556d509f395cfb&location=114.5276645,30.9807585&poitype=&radius=100&extensions=all&batch=false&roadlevel=0"
			res, _ := gocurl.Get(url)
			a := selfjson.Json2Gaodecitybylocation(res)
			str = strings.Replace(tmp, "\r\n", "", -1) + "," + a.Regeocode.FormattedAddress + "," + a.Regeocode.AddressComponent.Adcode + "," + a.Regeocode.AddressComponent.District + ","
			if len(a.Regeocode.Pois) > 0 {
				str = str + a.Regeocode.Pois[0].ID + "," + a.Regeocode.Pois[0].Name
			}
			str = str + "\r\n"

			fmt.Println(str)
			readtowrite.Writefile("/home/zlc/goProject/src/tools/1112end2.csv", str)
		}

	},
}

//注册子命令
func init() {
	gocurlCmd.AddCommand(getdistrictCmd)
	gocurlCmd.AddCommand(locationCmd)
	gocurlCmd.AddCommand(yidihouseCmd)
	gocurlCmd.AddCommand(locationtoaddressCmd)

}
