/*
 * Hitokoto-Go
 * Version: 1.02
 * Author: Syc <syc@bilibili.de>
 * GNU General Public License v3.0
*/

package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func MysqlClient() *sql.DB {
	db, err := sql.Open("mysql", mysql_user + ":" + mysql_pass + "@tcp(" + mysql_host + ":" + mysql_port + ")/" + mysql_name)
	if err != nil {
		log.Fatal(err, "\n")
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err, "\n")
	}else{
		log.Print("[Hitokoto] Connect Mysql Success")
	}
	return db
}

func Hitokoto(cid int64) string {
	rows, err := db.Query("SELECT * FROM hitokoto WHERE id = ?", cid)
	CheckErr(err)
	log.Printf("[Hitokoto] Mysql Query Success (%d)\n", cid)
	var id string
	var hitokoto string 
	for rows.Next() {
		err = rows.Scan(&id, &hitokoto)  
		CheckErr(err)  
		SetRedis(id, hitokoto)
	}
	return hitokoto
}

func Count() int {
	rows, err := db.Query("SELECT * FROM hitokoto")
	CheckErr(err)
	var kid int = 0
	for rows.Next() {
		kid++
	}
	log.Printf("[Hitokoto] Total Number %d\n", kid)
	return kid
}