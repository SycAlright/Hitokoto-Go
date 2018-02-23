/*
 * Hitokoto-Go
 * Version: 1.02
 * Author: Syc <syc@bilibili.de>
 * GNU General Public License v3.0
*/

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"net/http"
	"math/rand"
	"strconv"
	"time"
	"syscall"
	"github.com/garyburd/redigo/redis"
)

var (
	//Service_Conf
	Service_ip		=	""
	Service_port	=	"8080"

	//Mysql_Conf
	mysql_host  	=	"127.0.0.1"
	mysql_port		=	"3306"
    mysql_user		=	"root"
    mysql_pass		=	""
	mysql_name     	=	"hitokoto"

	//Redis_Conf
	redis_host  	=	"127.0.0.1"
	redis_port		=	"6379"
    redis_pass		=	""
	redis_dbns     	=	0
	redis_pool		=	10

	//Init
	RedisClient *redis.Pool
	VERSION = 1.12

	db = MysqlClient()
	client = RedisPool()
	kid int64 = int64(Count())
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU());
	log.Printf("[Hitokoto] CPU Number %d", runtime.NumCPU())
    c := make(chan os.Signal)
    //监听指定信号 ctrl+c kill
    signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
    go func() {
        for s := range c {
            switch s {
            case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
                Exit()
            default:
                log.Print(s)
            }
        }
    }()
	log.Printf("[Hitokoto] Service Version %.2f \n", VERSION)
	Run_Pid()
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(Handler_Index))
	log.Fatal(http.ListenAndServe(Service_ip + ":" + Service_port, mux))
}

func Exit()  {
	log.Print("[Hitokoto] Mysql-Conn Close")
	db.Close()
	log.Print("[Hitokoto] Redis-Conn Close")
	client.Close()
    log.Print("[Hitokoto] Service Stop")
    os.Exit(0)
}

func Run_Pid() {
	pid := os.Getpid()
	log.Print("[Hitokoto] PID ", pid, "\n")
	log.Print("[Hitokoto] Service Start\n")
}

func Handler_Index(w http.ResponseWriter, r *http.Request) {
	cid := Rand_Id(0, kid)
	wid := strconv.Itoa(int(cid))
	User(w, r)
	conn := client.Get()
	defer conn.Close()
	val, err := redis.String(conn.Do("GET", wid))
    if err != nil {
		log.Printf("[Hitokoto] Redis Get Unknow (%s)\n", wid)
		text := Hitokoto(cid)
		fmt.Fprintf(w, "%s\n", text)  
	}else{
		log.Printf("[Hitokoto] Redis Get Success (%s)\n", wid)
		fmt.Fprintf(w, "%s\n", val)
	}
	conn.Close()
}

func Rand_Id(min, max int64) int64 {  
    rand.Seed(time.Now().UnixNano())
    return min + rand.Int63n(max-min+1)  
}  

func User(w http.ResponseWriter, r *http.Request) {
	log.Print("[User] " + r.RequestURI + " - " + r.RemoteAddr + "\n")
}

func CheckErr(err error) {  
    if err != nil {  
		panic(err)
    }  
}  