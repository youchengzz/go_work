package crud

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "root:hcrz1234@tcp(192.168.1.123:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接异常", err)
	}
}
/*
*
题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
*
*/
type Student struct {
	Id    int
	Name  string
	Age   int
	Grade string
}

/*
*
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
*
*/
func AddStudent() {
	student := Student{
		Name:  "李四",
		Age:   16,
		Grade: "二年级",
	}
	tx := DB.Create(&student)
	fmt.Println(tx)
}

/*
*
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
*
*/
func SelectByAge(age int) {
	student := Student{}
	DB.Where("age > ?", age).Find(&student)
	fmt.Println(student)
}

/*
*
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
*
*/
func Update(name string) {
	student := Student{}
	tx := DB.Where("name = ?", name).Find(&student)
	if tx.Error != nil {
		fmt.Println("sql执行异常", tx.Error)
		return
	}
	student.Grade = "四年级"
	DB.Save(&student)

}

/*
*
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*
*/
func Delete(age int) {
	student := Student{}
	tx := DB.Where("age < ?", age).Delete(&student)
	if tx.Error != nil {
		fmt.Println("sql执行异常", tx.Error)
		return
	}
}

/*
*
题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*
*/
type Accounts struct {
	Id      int
	Balance int
}
type Transactions struct {
	Id            int
	FromAccountId int
	ToAccountId   int
	Amount        int
}

func TransferAccounts(balance int) {
	// 账户A id 1
	A := Accounts{
		Id: 1,
	}
	B := Accounts{
		Id: 2,
	}
	// 账户B id 2
	err := DB.Transaction(func(tx *gorm.DB) error {
		// 查询账户A余额
		if err := tx.First(&A).Error; err != nil {
			return fmt.Errorf("查询账户A失败 %s", err)
		}
		if err := tx.First(&B).Error; err != nil {
			return fmt.Errorf("查询账户B失败 %s", err)
		}
		fmt.Println("A账户信息:", A)
		if A.Balance > balance {
			A.Balance -= balance
			B.Balance += balance
		} else {
			return fmt.Errorf("账户余额不足")
		}
		if err := tx.Save(&A).Error; err != nil {
			return fmt.Errorf("转出余额失败: %s", err)
		}
		if err := tx.Save(&B).Error; err != nil {
			return fmt.Errorf("转入余额失败: %s", err)
		}
		transaction := Transactions{
			FromAccountId: A.Id,
			ToAccountId:   B.Id,
			Amount:        balance,
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return fmt.Errorf("创建转账记录失败 %s", err)
		}
		return nil
	})
	if err != nil {
		fmt.Println("转账失败", err)
		return
	}
	fmt.Println("转账成功")
}
