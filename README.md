# Hitokoto-Go

一言Hitokoto（GoLang）

Version 1.02

#### Blog: 
[《一言（Hitokoto-Go）》- SycBlog][1]

#### 开源协议
GNU General Public License v3.0

#### 运行环境
 - GoLang 1.9+
 - Mysql
 - Redis

#### Pkg
    go get -u github.com/go-sql-driver/mysql
    go get github.com/garyburd/redigo/redis

#### Mysql：
> 导入 ‘hitokoto.sql’ (附10条测试数据)

修改 main.go 中的 Mysql_Config

    mysql_host = "127.0.0.1" 
	mysql_port = "3306"      
    mysql_user = "root"
    mysql_pass = ""
	mysql_name = "hitokoto"

#### Redis：
> 修改 main.go 中的 Redis_Config

    redis_host = "127.0.0.1"
	redis_port = "6379"
    redis_pass = ""
	redis_dbns = 0  //存储库
	redis_pool = 10 //连接池最大数量

#### 编译  
    cd ./src/
    go build ./

#### 运行
    测试: ./hitokoto
    后台: nohup ./hitokoto

#### 大功告成
访问 `127.0.0.1：8080` 显示如下：
> 我相信十年后的八月，我们还能再相遇。

[1]: https://www.mfeng.cc/archives/2018/02/23/Hitokoto-Go.html