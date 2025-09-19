golang练习作业

个人博客belogs

1. 运行环境
golang + gin + gorm

2. SQL脚本文件
/db/init.sql

3. 配置文件
路径：/config/config.yml
+ 修改数据库配置db
+ key作为加密参数使用，可以不用修改
```
db:
  url: 127.0.0.1:3306
  username: root
  password: 123456
  database: belogs
key: 123456789abc
```
4. 启动项目
+ 方式一
```
go run main.go
```
+ 方式二
```
使用fresh
go install github.com/gravityblast/fresh@latest
fresh
```