# Hitokoto-Go: 一言

访问 `127.0.0.1:8080` 显示如下：
  
> 我相信十年后的八月，我们还能再相遇。

修正`main.go`的`Mysql_Config`部分

    mysql_host = "127.0.0.1" 
    mysql_port = "3306"      
    mysql_user = "root"
    mysql_pass = ""
    mysql_name = "hitokoto"

修正`main.go`的`Redis_Config`部分

    redis_host = "127.0.0.1"
    redis_port = "6379"
    redis_pass = ""
    redis_dbns = 0
    redis_pool = 10

Blog: [php.wf/archives/hitokoto-go.html][1]

  [1]: https://php.wf/archives/hitokoto-go.html
