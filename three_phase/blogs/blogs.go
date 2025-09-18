package blogs

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
进阶gorm
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
*
*/
type User struct {
	Id        int
	Name      string
	PostCount int
	Posts     []Post
}

func (User) TableName() string {
	return "user"
}

type Post struct {
	Id       int
	UserId   int
	Title    string
	Status   string
	Comments []Comment
}

type Comment struct {
	Id      int
	PostId  int
	Message string
}

func (Comment) TableName() string {
	return "comment"
}

/*
*
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
*
*/
func SelectUserPost(id int) User {
	user := User{}
	err := DB.Preload("Posts.Comments").Where("id = ?", id).Find(&user).Error
	if err != nil {
		fmt.Println("查询失败", id)
	}
	return user
}

/*
*
编写Go代码，使用Gorm查询评论数量最多的文章信息。
*
*/
func SelectCountComment() Post {
	post := Post{}
	sub := DB.Model(&Comment{}).Select("post_id, count(post_id) as num").Group("post_id").Order("num desc")
	err := DB.Table("posts p").Joins("JOIN (?) as tmp on p.id = tmp.post_id", sub).Preload("Comments").Find(&post)
	if err.Error != nil {
		fmt.Println("查询失败", err)
	}
	return post
}

/*
*
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
*
*/
func (c *Post) AfterSave(tx *gorm.DB) error {
	user := User{}
	if err := tx.Where("id = ?", c.UserId).Find(&user).Error; err != nil {
		fmt.Println("查询用户失败", err)
		return err
	}
	user.PostCount += 1
	err := tx.Save(&user).Error
	if err != nil {
		fmt.Println("更新用户文章失败", err)
		return err
	}
	return nil
}

func AddPost() Post {
	post := Post{
		UserId: 1,
		Title:  "测试钩子",
	}
	err := DB.Create(&post).Error
	if err != nil {
		fmt.Println("创建文章失败", err)
	}
	return post
}

/*
*
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*
*/
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	var num int64
	if err := tx.Where("post_id = ?", c.PostId).Find(c).Count(&num).Error; err != nil {
		fmt.Println("查询文章评论数量异常", err)
		return err
	}
	fmt.Println("文章评论数:", num, c)
	if num == 0 {
		if res := tx.Model(&Post{}).Where("id = ?", c.PostId).Update("status", "无评论"); res.Error != nil {
			fmt.Println("更新文章评论状态异常", err)
			return err
		}
	}
	return nil
}

func DeleteComment(id int) {
	comment := Comment{}
	if err := DB.Where("id = ?", id).Find(&comment).Error; err != nil {
		fmt.Println("查询评论异常", err)
	}
	DB.Delete(&comment)
}
