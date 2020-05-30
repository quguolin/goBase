package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
	"goBase/xtime"
	_ "github.com/go-sql-driver/mysql"
)

const (
	sqlQuery = "select id,`name`,`desc`,`ctime`,`mtime` from `test`"
	sqlInset = "insert into `test` (`name`,`desc`,`ctime`) VALUES(?,?,?)"
)

type Test struct {
	ID    int
	Name  string
	Desc  string
	CTime xtime.Time `json:"ctime"`
	Mtime xtime.Time `json:"mtime"`
}

var (
	rows sql.Rows
	st   []*Test
)

func main() {
	db, err := sql.Open("mysql", "root:test@tcp(127.0.0.1:3306)/test?parseTime=true&loc=Local&charset=utf8,utf8mb4")
	if err != nil {
		panic(err)
	}
	row := db.QueryRow(sqlQuery)
	s := &Test{}
	if err = row.Scan(&s.ID, &s.Name, &s.Desc, &s.CTime, &s.Mtime); err != nil {
		panic(err)
	}
	b, _ := json.Marshal(s)
	fmt.Println(string(b))

	res, err := db.Exec(sqlInset, "name", "desc", xtime.Time(time.Now().Unix()-3600))
	if err != nil {
		panic(err)
	}
	fmt.Println(res.LastInsertId())
}
