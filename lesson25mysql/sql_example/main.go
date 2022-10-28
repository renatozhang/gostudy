package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func initdb() error {
	var err error
	dsn := "root:123456@(localhost:3306)/golang_db"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	DB.SetMaxOpenConns(100) // 最大打开连接数
	DB.SetMaxIdleConns(16)  // 空闲连接数
	return nil
}

type User struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func testQueryMutilRow() {
	sqlstr := "select id,name,age from user where id>?"
	rows, err := DB.Query(sqlstr, 0)
	// 重点关注，row对象一定要close掉
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		fmt.Printf("query failed,err:%v\n", err)
		return
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", user.Id, user.Name, user.Age)
	}
}

func testQueryData() {
	for i := 0; i < 101; i++ {
		fmt.Printf("query %d times\n", i)
		sqlstr := "select id,name,age from user where id=?"
		row := DB.QueryRow(sqlstr, 2)
		// if row != nil {
		// 	continue
		// }
		var user User
		err := row.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", user.Id, user.Name, user.Age)
	}

}

func testInsertData() {
	sqlstr := "insert into user(name,age)values(?,?)"
	result, err := DB.Exec(sqlstr, "tom", 18)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get last insert id failed, err:%v\n", err)
		return
	}
	fmt.Printf("id is %d\n", id)
}

func testUpdateData() {
	sqlstr := "update user set name=? where id=?"
	result, err := DB.Exec(sqlstr, "jim", 4)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get affected row failed, err:%v\n", err)
	}
	fmt.Printf("update db succ, affected rows:%d\n", affected)
}

func testDeleteData() {
	sqlstr := "delete from user where id=?"
	result, err := DB.Exec(sqlstr, 4)
	if err != nil {
		fmt.Printf("delete data failed, err:%v\n", err)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get affected row failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete db succ, affected rows:%d\n", affected)
}

func testPrepareData() {
	sqlstr := "select id,name,age from user where id > ?"
	stmt, err := DB.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	rows, err := stmt.Query(0)
	// 重点关注，row对象一定要close掉
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Name)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s, age:%d\n", user.Id, user.Name, user.Age)
	}

}

func testPrepareRow() {
	sqlstr := "select id, name, age from user where id=?"
	stmt, err := DB.Prepare(sqlstr)
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	if err != nil {
		fmt.Printf("prepare dailed, err:%v\n", err)
		return
	}

	row := stmt.QueryRow(7)
	var user User
	err = row.Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", user.Id, user.Name, user.Age)
}

func testPrepareUpdate() {
	sqlstr := "update user set name=? where id=?"
	stmt, err := DB.Prepare(sqlstr)
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	result, err := stmt.Exec("tom", 3)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}

	id, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get affected row failed, err:%v\n", err)
		return
	}
	fmt.Printf("update db succ, affected rows:%d\n", id)
}

func testPrepareInsertData() {
	sqlstr := "insert into user(name,age) values(?,?)"
	stmt, err := DB.Prepare(sqlstr)
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	result, err := stmt.Exec("zhang", 32)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get last row id failed,err:%v\n", err)
		return
	}
	fmt.Printf("id is %d\n", id)
}

func testTrans() {
	conn, err := DB.Begin()
	if err != nil {
		if conn != nil {
			conn.Rollback()
		}
		fmt.Printf("begin failed, err:%v\n", err)
		return
	}
	sqlstr := "update user set age=1 where id=?"
	_, err = conn.Exec(sqlstr, 1)
	if err != nil {
		conn.Rollback()
		fmt.Printf("exec sql:%s failed, err:%v\n", sqlstr, err)
		return
	}

	sqlstr = "update user set age=2; where id=?"
	_, err = conn.Exec(sqlstr, 2)
	if err != nil {
		conn.Rollback()
		fmt.Printf("exec second sql:%s failed, err:%v\n", sqlstr, err)
		return
	}

	err = conn.Commit()
	if err != nil {
		fmt.Printf("commit failed, err:%v\n", err)
		conn.Rollback()
		return
	}

}

func main() {
	err := initdb()
	if err != nil {
		fmt.Printf("init db failed., err:%v\n", err)
		return
	}
	// testQueryData()
	// testQueryMutilRow()
	// testInsertData()
	// testUpdateData()
	// testDeleteData()

	// testPrepareData()
	// testPrepareUpdate()
	// testPrepareInsertData()
	// testPrepareRow()

	//事务
	testTrans()
}
