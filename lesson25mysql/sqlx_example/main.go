package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

var DB *sqlx.DB

func initDb() error {
	var err error
	dsn := "root:123456@tcp(localhost:3306)/golang_db"
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return err
	}
	DB.SetMaxOpenConns(100)
	DB.SetConnMaxIdleTime(16)
	return nil
}

func testQuery() {
	sqlstr := "select id, name, age from user where id=?"
	var user User
	err := DB.Get(&user, sqlstr, 6)
	if err != nil {
		fmt.Printf("get data failed, err:%v\n", err)
		return
	}
	fmt.Printf("user:%v\n", user)
}

func testQueryMulti() {
	sqlstr := "select id, name, age from user where id > ?"
	var users []User
	err := DB.Select(&users, sqlstr, 0)
	if err != nil {
		fmt.Printf("get data failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}

func testUpdate() {
	sqlstr := "update user set name=? where id=?"
	result, err := DB.Exec(sqlstr, "renato", 7)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	count, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("affected failed, err:%v\n", err)
		return
	}
	fmt.Printf("affected rows %d", count)

}

func testInsert() {
	sqlstr := "insert into user(name,age) values(?,?)"
	result, err := DB.Exec(sqlstr, "zeng", 18)
	if err != nil {
		fmt.Printf("exec failed, err:%v", err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get last id failed, err:%v\n", err)
		return
	}
	fmt.Printf("id is %d\n", id)
}

func testDelete() {
	sqlstr := "delete from user where name=?"
	result, err := DB.Exec(sqlstr, "zeng")
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	count, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("affected failed, err:%v\n", err)
		return
	}
	fmt.Printf("affected rows %d\n", count)
}

func testTrans() {
	sqlstr := "update user set age=? where id=?"
	conn, err := DB.Begin()
	if err != nil {
		if conn != nil {
			conn.Rollback()
		}
		fmt.Printf("begin failed, err:%v\n", err)
		return
	}
	_, err = conn.Exec(sqlstr, 1, 1)
	if err != nil {
		conn.Rollback()
		fmt.Printf("exec sql:%s failed, err:%v\n", sqlstr, err)
		return
	}

	sqlstr = "update user set age=?; where id=?"
	_, err = conn.Exec(sqlstr, 2, 2)
	if err != nil {
		conn.Rollback()
		fmt.Printf("exec sql:%s failed, err:%v\n", sqlstr, err)
		return
	}

	err = conn.Commit()
	if err != nil {
		conn.Rollback()
		fmt.Printf("commit failed, err:%s\n", err)
		return
	}

}

func queryDB(name string) {
	sqlstr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("sql:%s\n", sqlstr)
	var user []User
	err := DB.Select(&user, sqlstr)
	if err != nil {
		fmt.Printf("select failed, err:%v\n", err)
		return
	}

	for _, v := range user {
		fmt.Printf("user:%#v\n", v)
	}
}

func queryDBBySqlx(name string) {
	sqlstr := "select id, name, age from user where name=?"
	fmt.Printf("sql:%s\n", sqlstr)
	var user []User
	err := DB.Select(&user, sqlstr, name)
	if err != nil {
		fmt.Printf("select failed, err:%v\n", err)
		return
	}

	for _, v := range user {
		fmt.Printf("user:%#v\n", v)
	}
}

func SqlInject() {
	// queryDB("abc' or 1 = 1 #")
	// queryDB("name=tom' and (select count(*) from user ) < 10 #")
	// queryDB("name=123' union select * from user #")

	queryDBBySqlx("abc' or 1 = 1 #")
}

func main() {
	err := initDb()
	if err != nil {
		fmt.Printf("init db failed. err:%v\n", err)
		return
	}

	// testQuery()
	// testQueryMulti()
	// testUpdate()
	// testInsert()
	// testDelete()
	// testTrans()

	SqlInject()
}
