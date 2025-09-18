package sqlx

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB
var err error

/*
*
Sqlx入门
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
*
*/
type Employees struct {
	Id         int
	Name       string
	Department string
	Salary     int
}

func SelectByDepartment(department string) []Employees {
	employes := make([]Employees, 0)
	err := DB.Select(&employes, "select * from employees where department = ?", department)
	if err != nil {
		fmt.Println("查询失败", err)
	}
	return employes
}

/*
*
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*
*/
func SelectBySslary() Employees {
	employes := Employees{}
	err := DB.Get(&employes, "select * from employees where salary = (select max(salary) from employees)")
	if err != nil {
		fmt.Println("查询失败", err)
	}
	return employes
}

/*
*
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
*
*/
type Book struct {
	Id     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float32 `db:"price"`
}

/*
*
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*
*/
func SelectBooks() []Book {
	ctx := context.Background()
	books := make([]Book, 0)
	sql := "select * from books where price > ?"
	stm, err := DB.PreparexContext(ctx, sql)
	if err != nil {
		fmt.Println("准备查询失败", err)
	}
	defer stm.Close()
	if err := stm.SelectContext(ctx, &books, 50); err != nil {
		fmt.Println("查询失败", err)
	}
	return books
}

func init() {
	dsn := "root:hcrz1234@tcp(192.168.1.123:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接异常", err)
	}
	// defer DB.Close()
}
