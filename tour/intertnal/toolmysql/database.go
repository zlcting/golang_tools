/**
*FileName: mysql
*golang中mysql的用法
 */

package toolmysql

import (
	"database/sql"
	"fmt"

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
	Id                     int    `db:"id"`
	Name                   string `db:"name"`
	Coordx2                string `db:"coordx2"`
	Coordy2                string `db:"coordy2"`
	Value                  string `db:"value"`
	Is_push_tmall          int    `db:"is_push_tmall"`
	Is_push_lejuflag       int    `db:"is_push_lejuflag"`
	Unified_city_cn_code   string `db:"unified_city_cn_code"`
	City_cn                string `db:"city_cn"`
	Salestate              int    `db:"salestate"`
	Status                 int    `db:"status"`
	Eunified_city_cn_code  string `db:"eunified_city_cn_code"`
	Eunified_city_cn       string `db:"eunified_city_cn"`
	Eunified_district_code string `db:"eunified_district_code"`
	Eunified_district      string `db:"eunified_district"`
	Site                   string `db:"site"`
	Createtime             int    `db:"createtime"`
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

type MyDb struct {
	*sql.DB
}

func NewDb() (*MyDb, error) {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)
	db, err := sql.Open("mysql", dbDSN)

	if err != nil {
		return nil, err
	}
	mydb := &MyDb{DB: db}
	return mydb, nil
}

func (this *MyDb) Query(id string) (map[string]string, error) {
	r, err := this.DB.Query("SELECT e.unified_city_cn_code,c.city_en from house_extension as e left join city as c on e.unified_city_cn_code = c.unified_city_cn_code where house_id = ? ", id)
	var city_en string
	var unified_city_cn_code string
	dbmap := make(map[string]string)
	for r.Next() {
		r.Scan(&unified_city_cn_code, &city_en)
		dbmap["unified_city_cn_code"] = unified_city_cn_code
		dbmap["city_en"] = city_en
	}

	if err != nil {
		return nil, err
	}
	return dbmap, nil
}

func (this *MyDb) HouseQuery(site string, name string) (map[string]string, error) {
	r, _ := this.DB.Query("SELECT site,hid,id from house where site = ? and name = ?", site, name)
	var city_en string
	var hid string
	var id string
	aamap := make(map[string]string)
	for r.Next() {
		r.Scan(&city_en, &hid, &id)
		aamap["city_en"] = city_en
		aamap["hid"] = hid
		aamap["id"] = id
	}

	return aamap, nil
}

// 初始化链接
// func init() {

// 	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)

// 	// 打开连接失败
// 	MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)
// 	//defer MysqlDb.Close();
// 	if MysqlDbErr != nil {
// 		log.Println("dbDSN: " + dbDSN)
// 		panic("数据源配置不正确: " + MysqlDbErr.Error())
// 	}

// 	// 最大连接数
// 	MysqlDb.SetMaxOpenConns(100)
// 	// 闲置连接数
// 	MysqlDb.SetMaxIdleConns(20)
// 	// 最大连接周期
// 	MysqlDb.SetConnMaxLifetime(100 * time.Second)

// 	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
// 		panic("数据库链接失败: " + MysqlDbErr.Error())
// 	}

// }

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
	houses := make([]House_location, 0)
	rows, _ := MysqlDb.Query("SELECT h.id,h.name,h.coordy2,h.coordx2,o.value,e.is_push_lejuflag,e.is_push_tmall,c.unified_city_cn_code,c.city_cn,h.status,h.salestate,e.unified_city_cn_code as Eunified_city_cn_code,e.unified_city_cn as Eunified_city_cn,e.unified_district_code as Eunified_district_code,e.unified_district as Eunified_district,h.site,h.createtime FROM data_house_sina_com_cn.house AS h LEFT JOIN house_extension AS e ON h.id = e.house_id LEFT JOIN house_options AS o ON h.district = o.id AND o.type = 'district' left join city as c on h.site = c.city_en LIMIT ? , ?", limit, offset)
	// 遍历
	var house House_location
	var id int
	var name string
	var coordy2 string
	var coordx2 string
	var is_push_lejuflag int
	var is_push_tmall int
	var unified_city_cn_code string
	var city_cn string
	var value string
	var salestate int
	var status int
	var eunified_city_cn_code string
	var eunified_city_cn string
	var eunified_district_code string
	var eunified_district string
	var site string
	var createtime int
	for rows.Next() {

		rows.Scan(&id, &name, &coordy2, &coordx2, &value, &is_push_lejuflag, &is_push_tmall, &unified_city_cn_code, &city_cn, &status, &salestate, &eunified_city_cn_code, &eunified_city_cn, &eunified_district_code, &eunified_district, &site, &createtime)

		house.Id = id
		house.Name = name
		house.Coordy2 = coordy2
		house.Coordx2 = coordx2
		house.Is_push_lejuflag = is_push_lejuflag
		house.Is_push_tmall = is_push_tmall
		house.Unified_city_cn_code = unified_city_cn_code
		house.City_cn = city_cn
		house.Value = value
		house.Salestate = salestate
		house.Status = status
		house.Eunified_city_cn_code = eunified_city_cn_code
		house.Eunified_city_cn = eunified_city_cn
		house.Eunified_district_code = eunified_district_code
		house.Eunified_district = eunified_district
		house.Site = site
		house.Createtime = createtime

		houses = append(houses, house)
	}

	return houses
}

//查询坐标和天猫状态
func StructHousefliedId() []House_location {
	houses := make([]House_location, 0)
	rows, _ := MysqlDb.Query("SELECT h.id,h.name,h.coordy2,h.coordx2,o.value,e.is_push_lejuflag,e.is_push_tmall,c.unified_city_cn_code,c.city_cn,h.status,h.salestate,e.unified_city_cn_code as Eunified_city_cn_code,e.unified_city_cn as Eunified_city_cn,e.unified_district_code as Eunified_district_code,e.unified_district as Eunified_district,h.site,h.createtime FROM data_house_sina_com_cn.house AS h LEFT JOIN house_extension AS e ON h.id = e.house_id LEFT JOIN house_options AS o ON h.district = o.id AND o.type = 'district' left join city as c on h.site = c.city_en where h.id in (10276,35588,97463,113401,135814,140609,142476,145443,148008,148548,149291,151586,152862,155973,156326,157791,157900,157958,159079,160043,160881,161772,161774,161776,161964,162156,162482,162716,163459,163767,163793,163880,163910,164311,164709,165134,165333,165828,166318,167109,167408,167904,168413,168635,168960,169176,169203,169279,169647,169667,169788,169864,170525,171517,171701,171818,171852,172078,172207,172584,172586,172847,173994,174675,176046,176497,177658,178012,179754,179956,180370,180475,180494,180523,180525,180535,180542,180544,180546,180553,180562,180569,180572,180575,181254,181336,181398,181667,181725,181771,182394,182412,182689,182849,182874,182878,183377,183387,183388,183390,183404,183406,183409,183415,184217,184686,184813,185339,185538,185953,186356,186794,186861,187192,187598,187628,187652,187723,187909,187938,188243,188245,188347,188444,188552,188752,188843,189049,189323,189413,189461,189537,189729,190344,190456,191061,191127,191544,191631,191636,191844,192175,192256,192264,192351,192523,192668,192673,192838,192855,193088,193112,193186,193353,193403,193474,193548,193978,194004,194028,194102,194136,194236,194300,194305,194311,194315,194316,195165,196528)")
	// 遍历
	var house House_location
	var id int
	var name string
	var coordy2 string
	var coordx2 string
	var is_push_lejuflag int
	var is_push_tmall int
	var unified_city_cn_code string
	var city_cn string
	var value string
	var salestate int
	var status int
	var eunified_city_cn_code string
	var eunified_city_cn string
	var eunified_district_code string
	var eunified_district string
	var site string
	var createtime int
	for rows.Next() {

		rows.Scan(&id, &name, &coordy2, &coordx2, &value, &is_push_lejuflag, &is_push_tmall, &unified_city_cn_code, &city_cn, &status, &salestate, &eunified_city_cn_code, &eunified_city_cn, &eunified_district_code, &eunified_district, &site, &createtime)

		house.Id = id
		house.Name = name
		house.Coordy2 = coordy2
		house.Coordx2 = coordx2
		house.Is_push_lejuflag = is_push_lejuflag
		house.Is_push_tmall = is_push_tmall
		house.Unified_city_cn_code = unified_city_cn_code
		house.City_cn = city_cn
		house.Value = value
		house.Salestate = salestate
		house.Status = status
		house.Eunified_city_cn_code = eunified_city_cn_code
		house.Eunified_city_cn = eunified_city_cn
		house.Eunified_district_code = eunified_district_code
		house.Eunified_district = eunified_district
		house.Site = site
		house.Createtime = createtime

		houses = append(houses, house)
	}

	return houses
}
