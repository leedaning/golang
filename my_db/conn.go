package my_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type StdResp struct {
	Id   int
	Name string
	Age  int
}

func Conn() {
	fmt.Println("Connect database.")
	Db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/study")
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
		fmt.Println("连接数据库失败")
	}

	//监测数据库连接
	/*errs := Db.Ping()
	if errs != nil {
		fmt.Println("连接数据库失败")
		return
	} else {
		fmt.Println("连接成功")
	}*/
	//关闭连接，使用defer保证一定会关闭
	defer Db.Close()

	var stds []StdResp
	fmt.Printf("stds的类型为：%T\n", stds)

	//var Id, Age int
	var Name string
	/*rows, _ := Db.Query("select id, name, age from student")
	for rows.Next() {
		rows.Scan(&Id, &Name, &Age)
		fmt.Println("id=", Id, " name=", Name, " age=", Age)
	}*/

	err = Db.QueryRow("select name from student").Scan(&Name)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("没有符合条件的数据")
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println("row", Name)
	/*rows.Scan(&stds)
	fmt.Println("学生信息：", stds)*/
}
