package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

type Person struct {
	Pid     int    `db:"id"`
	Pname   string `db:"name"`
	Pmobile string `db:"mobile"`
}

func init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql error: ", err)
	}

	Db = database

}

func main() {
	//delete1()
	//update1()
	insert1()
	select1()
	defer Db.Close()
}

func delete1()  {
	exec, err := Db.Exec("delete from person where name=?", "hhh11")
	if err != nil {
		fmt.Println("db ex error: ", err)
	}
	affected, err := exec.RowsAffected()
	if err != nil {
		fmt.Println("rows error: ", err)
	}
	fmt.Println("succ",affected)

}

func update1() {
	exec, err := Db.Exec("update person set mobile=? where name=?", "111122222", "zzz")
	if err != nil {
		fmt.Println("exec error: ", err)
	}
	affected, err := exec.RowsAffected()
	if err != nil {
		fmt.Println("rows error: ", err)
	}
	fmt.Println("success",affected)
}

func insert1() {
	exec, err := Db.Exec("insert into person(name,mobile) values(?,?)", "hhh11", "123123")
	if err != nil {
		fmt.Println("db exec error: ", err)
	}
	id, err := exec.LastInsertId()
	if err != nil {
		fmt.Println("lsat insert error: ", err)
	}
	fmt.Println("success", id)
}

func select1() {
	var person []Person
	err := Db.Select(&person, "select * from person")
	if err != nil {
		fmt.Println("ecex  error: ", err)
	}
	fmt.Println(person)

}
