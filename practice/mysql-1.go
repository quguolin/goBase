package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	id   int
	name string
)

func main() {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/member?charset=utf8")

	if err != nil {
		fmt.Println("error---------------")
	}

	rows, err := db.Query("select id,name,create_at from crm_member limit 3")
	checkErr(err)
	defer rows.Close()

	row_values := make(map[int]map[string]string) //must creat a new row_value every time
	index := 0

	for rows.Next() {
		var id string
		var name string
		var create_at string

		row_value := make(map[string]string)

		if err = rows.Scan(&id, &name, &create_at); err != nil {
			fmt.Println(err)
		}

		row_value["id"] = id
		row_value["name"] = name
		row_value["create_at"] = create_at

		row_values[index] = row_value

		index++

	}

	for _, value := range row_values {
		fmt.Println(value)
	}

}

func printResult(query *sql.Rows) {
	column, _ := query.Columns() //所有的字段名称
	values := make([][]byte, len(column))
	scans := make([]interface{}, len(column))

	for i := range values {
		scans[i] = &values[i]
	}

	results := make(map[int]map[string]string)
	i := 0
	for query.Next() {
		if err := query.Scan(scans...); err != nil {
			scans[i] = &values[i]
			fmt.Println(err)
			return
		}

		row := make(map[string]string) //每一行的数据
		for k, v := range values {
			key := column[k]
			row[key] = string(v)
		}
		results[i] = row
		i++
	}

	for k, v := range results {
		fmt.Println(k, v)
	}
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
