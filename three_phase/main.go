package main

import (
	"Go-Project/go_work/three_phase/blogs"
)

func main() {
	// AddStudent()
	// SelectByAge(18)
	// Update("张三")
	// Delete(15)

	// crud.TransferAccounts(100)

	// res:=sqlx.SelectByDepartment("技术部")
	// res := sqlx.SelectBySslary()
	// res := sqlx.SelectBooks()
	// res := blogs.SelectUserPost(1)
	// res := blogs.SelectCountComment()
	blogs.DeleteComment(8)
	// fmt.Println(res)
}
