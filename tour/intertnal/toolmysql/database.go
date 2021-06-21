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

const (
	USER_NAME = "loupantest"
	PASS_WORD = "v0oMrlie7iV=hs"
	NETWORK   = "tcp"
	HOST      = "i.yz.mytest.leju.com"
	PORT      = "63353"
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
func StructQueryAllField() {

	// 通过切片存储
	citys := make([]Unified_city, 0)
	rows, _ := MysqlDb.Query("SELECT * FROM `unified_city` limit ?", 100)
	// 遍历
	var city Unified_city
	for rows.Next() {
		rows.Scan(&city.Id, &city.Name, &city.Code)
		citys = append(citys, city)
	}
	fmt.Println(citys)
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
func StructUpdate() {

	ret, _ := MysqlDb.Exec("UPDATE users set age=? where id=?", "100", 1)
	upd_nums, _ := ret.RowsAffected()

	fmt.Println("RowsAffected:", upd_nums)
}

// 删除数据
func StructDel() {

	ret, _ := MysqlDb.Exec("delete from users where id=?", 1)
	del_nums, _ := ret.RowsAffected()

	fmt.Println("RowsAffected:", del_nums)
}
