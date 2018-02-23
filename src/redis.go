/*
 * Hitokoto-Go
 * Version: 1.02
 * Author: Syc <syc@bilibili.de>
 * GNU General Public License v3.0
*/

package main

import (
    "log"
    "github.com/garyburd/redigo/redis"
    "time"
)

func RedisPool() *redis.Pool {
	RedisClient = &redis.Pool{
		MaxIdle: 1,  
		MaxActive: redis_pool,
        IdleTimeout: 180 * time.Second,
        Dial: func() (redis.Conn, error) {
            c, err := redis.Dial("tcp", redis_host + ":" + redis_port)
            if err != nil {
                return nil, err
            }
            if _, err := c.Do("AUTH", redis_pass); err != nil {
                c.Close()
                return nil, err
            }
            return c, err
        },
        TestOnBorrow: func(c redis.Conn, t time.Time) error {
            if time.Since(t) < time.Minute {
                return nil
            }
            _, err := c.Do("PING")
            return err
        },
    }
    return RedisClient
}

func SetRedis(key, value string) {
    conn := client.Get()
    defer conn.Close()
    _, err := conn.Do("SET", key, value)
    CheckErr(err)
    log.Printf("[Hitokoto] Redis Set Success (%s)\n", key)
    conn.Close()
}