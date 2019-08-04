package Dao

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"strings"
	"transaction/entity"
)

const (
	userName = "root"
	password = ""
	ip = "localhost"
	port = "3306"
	dbName = "mercer"
)

func IsErr(err error,msg string)  {
	if err != nil{
		fmt.Printf("error:%s,msg:%s",err,msg)
	}
}

var DB *sql.DB

//func Connection()  {
//	建立连接
	//path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//DB,err := sql.Open("mysql",path)
	//IsErr(err,"failed to connect")
	//DB.SetConnMaxLifetime(100)
	//DB.SetMaxIdleConns(10)
	//if err := DB.Ping(); err != nil{
	//	fmt.Println("opon database fail")
	//	return
	//}
	//fmt.Println("connnect success")
//}

//更改商铺库存表
func Update(shop *entity.Shop) {
	//Connection()
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	DB,err := sql.Open("mysql",path)
	IsErr(err,"failed to connect")

	stmt,err := DB.Prepare("update shop set gNum = gNum - ? where sID = ? and gID = ?")
	IsErr(err,"failed to prepare a sql")
	result,err := stmt.Exec(shop.GNum,shop.SID,shop.GID)
	IsErr(err,"failed to excute a sql")
	fmt.Print(result.RowsAffected())

}

//记录进订单表
func Insert(order *entity.Order){
	//Connection()
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	DB,err := sql.Open("mysql",path)
	IsErr(err,"failed to connect")

	row,err := DB.Prepare("INSERT INTO test (cID,gID,sID,number,time) VALUES (?,?,?,?,?)")
	IsErr(err,"failed to prepare a sql")
	res,err := row.Exec(order.CId,order.GID,order.SID,order.Number,order.Time)
	IsErr(err,"failed to excute a sql")
	fmt.Println(res.LastInsertId())
}

//查询物品数量
func Select()  {
	//Connection()
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	DB,err := sql.Open("mysql",path)
	IsErr(err,"failed to connect")

	rows,err := DB.Query("select gNum from shop where gID = ? and sID = ?")
	IsErr(err,"failed to excute a sql")
	defer rows.Close()
	var gNum string
	for rows.Next(){
		rows.Scan(&gNum)
		fmt.Print(gNum)
	}
}