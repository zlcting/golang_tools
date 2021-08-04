/**
*FileName: mysql
*golang中mysql的用法
 */

package toolmysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Unified_city struct {
	Id       int64  `db:"id"`
	Code     int64  `db:"code"`
	Name     string `db:"name"`
	Polyline string `db:"polyline"`
	Level    string `db:"level"`
	Center   string `db:"center"`
}

type House_location struct {
	Id                   int64  `db:"id"`
	Name                 string `db:"name"`
	Coordx2              string `db:"coordx2"`
	Coordy2              string `db:"coordy2"`
	Value                string `db:"value"`
	Is_push_tmall        int    `db:"is_push_tmall"`
	Is_push_lejuflag     int    `db:"is_push_lejuflag"`
	Unified_city_cn_code string `db:"unified_city_cn_code"`
	City_cn              string `db:"city_cn"`
}

var MysqlDb *sql.DB
var MysqlDbErr error

// const (
// 	USER_NAME = "root"
// 	PASS_WORD = "123"
// 	NETWORK   = "tcp"
// 	HOST      = "localhost"
// 	PORT      = 3306
// 	DATABASE  = "bishe"
// 	CHARSET   = "utf8"
// )

// const (
// 	USER_NAME = "loupantest"
// 	PASS_WORD = "v0oMrlie7iV=hs"
// 	NETWORK   = "tcp"
// 	HOST      = "i.yz.mytest.leju.com"
// 	PORT      = "63353"
// 	DATABASE  = "data_house_sina_com_cn"
// 	CHARSET   = "utf8"
// )

const (
	USER_NAME = "datahouseuser"
	PASS_WORD = "daTA@user123"
	NETWORK   = "tcp"
	HOST      = "i.yz.mytest.leju.com"
	PORT      = "53353"
	DATABASE  = "data_house_sina_com_cn"
	CHARSET   = "utf8"
)

// 初始化链接
func init() {

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)

	// 打开连接失败
	MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)
	//defer MysqlDb.Close();
	if MysqlDbErr != nil {
		log.Println("dbDSN: " + dbDSN)
		panic("数据源配置不正确: " + MysqlDbErr.Error())
	}

	// 最大连接数
	MysqlDb.SetMaxOpenConns(100)
	// 闲置连接数
	MysqlDb.SetMaxIdleConns(20)
	// 最大连接周期
	MysqlDb.SetConnMaxLifetime(100 * time.Second)

	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
		panic("数据库链接失败: " + MysqlDbErr.Error())
	}

}

// 查询数据，指定字段名
func StructQueryField() {

	city := new(Unified_city)
	row := MysqlDb.QueryRow("select id, name, age from users where id=?", 1)
	if err := row.Scan(&city.Id, &city.Name, &city.Code); err != nil {
		fmt.Printf("scan failed, err:%v", err)
		return
	}
	fmt.Println(city.Id, city.Name, city.Code)
}

// 查询数据，取所有字段
func StructQueryAllField() []Unified_city {

	// 通过切片存储
	citys := make([]Unified_city, 100)
	rows, _ := MysqlDb.Query("SELECT id,code,name FROM `unified_city` limit ?", 4000)
	// 遍历
	var city Unified_city
	var id int64
	var code int64
	var name string
	for rows.Next() {
		rows.Scan(&id, &code, &name)
		city.Id = id
		city.Code = code
		city.Name = name
		citys = append(citys, city)
	}
	return citys
}

// 插入数据
func StructInsert() {

	ret, _ := MysqlDb.Exec("insert INTO users(name,age) values(?,?)", "小红", 23)

	//插入数据的主键id
	lastInsertID, _ := ret.LastInsertId()
	fmt.Println("LastInsertID:", lastInsertID)

	//影响行数
	rowsaffected, _ := ret.RowsAffected()
	fmt.Println("RowsAffected:", rowsaffected)

}

// 更新数据
func StructUpdate(polyline string, center string, level string, code string) {

	ret, _ := MysqlDb.Exec("UPDATE unified_city set polyline=?,center=?,level=? where code=?", polyline, center, level, code)
	upd_nums, _ := ret.RowsAffected()

	fmt.Println("RowsAffected:", code, upd_nums)
}

// 删除数据
// func StructDel() {

// 	ret, _ := MysqlDb.Exec("delete from users where id=?", 1)
// 	del_nums, _ := ret.RowsAffected()

// 	fmt.Println("RowsAffected:", del_nums)
// }

//查询坐标和天猫状态
func StructHouseflied(page int, offset int) []House_location {
	limit := (page - 1) * offset
	houses := make([]House_location, offset)

	rows, _ := MysqlDb.Query("SELECT h.id,h.name,h.coordy2,h.coordx2,o.value,e.is_push_lejuflag,e.is_push_tmall,c.unified_city_cn_code,c.city_cn FROM data_house_sina_com_cn.house AS h LEFT JOIN house_extension AS e ON h.id = e.house_id LEFT JOIN house_options AS o ON h.district = o.id AND o.type = 'district' left join city as c on h.site = c.city_en WHERE h.status = 1 LIMIT ? , ?", limit, offset)
	// 遍历
	var house House_location
	var id int64
	var name string
	var coordy2 string
	var coordx2 string
	var is_push_lejuflag int
	var is_push_tmall int
	var unified_city_cn_code string
	var city_cn string
	var value string
	for rows.Next() {

		rows.Scan(&id, &name, &coordy2, &coordx2, &value, &is_push_lejuflag, &is_push_tmall, &unified_city_cn_code, &city_cn)

		house.Id = id
		house.Name = name
		house.Coordy2 = coordy2
		house.Coordx2 = coordx2
		house.Is_push_lejuflag = is_push_lejuflag
		house.Is_push_tmall = is_push_tmall
		house.Unified_city_cn_code = unified_city_cn_code
		house.City_cn = city_cn
		house.Value = value

		houses = append(houses, house)
	}

	return houses
}
