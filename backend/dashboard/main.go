package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var servers = make(map[string]string)
var rdb *redis.Client

func main() {

	// 读取环境变量的服务器列表
	prefix := "SERVER_"
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		key, value := pair[0], pair[1]
		if strings.HasPrefix(key, prefix) {
			envKey := strings.ToLower(strings.TrimPrefix(key, prefix))
			servers[envKey] = value
		}
	}

	// 读取环境变量的Redis连接信息
	redisHost, exists := os.LookupEnv("REDIS_HOST")
	if !exists {
		redisHost = "localhost"
	}
	redisPort, exists := os.LookupEnv("REDIS_PORT")
	if !exists {
		redisPort = "6379"
	}
	rawRedisDb, exists := os.LookupEnv("REDIS_DB")
	redisDb, err := strconv.Atoi(rawRedisDb)
	if err != nil || !exists {
		log.Fatalln("无法读取数据库编号：" + err.Error())
	}
	redisPassword, exists := os.LookupEnv("REDIS_PASSWORD")

	// 读取环境变量的程序启动信息
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8080"
	}

	// 初始化Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		DB:       redisDb,
		Password: redisPassword,
	})
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		log.Fatalln("连接Redis失败：", err.Error())
	}

	defaultServer, exists := os.LookupEnv("DEFAULT_SERVER")
	if !exists {
		log.Fatalln("请提供一个默认查询服务器")
	}

	router := gin.Default()
	//允许CORS跨域
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/player/:name", func(ctx *gin.Context) {
		playerName := ctx.Param("name")
		server := ctx.DefaultQuery("server", defaultServer)
		result := rdb.HGet(
			context.Background(),
			server+":name_uuid",
			playerName,
		)
		if err := result.Err(); err != nil {
			ctx.AbortWithStatusJSON(200, gin.H{"msg": "没有这个玩家"})
			return
		}
		filename := result.Val() + ".json"
		path := filepath.Join(servers[server], "world/stats/", filename)
		fmt.Println(path)
		file, err := os.Open(path)
		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"msg": "打开文件失败"})
			fmt.Println(err.Error())
			return
		}
		data, err := io.ReadAll(file)
		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"msg": "读取数据失败"})
			return
		}
		ctx.Data(200, "application/json", data)
	})

	router.GET("/:stat", func(ctx *gin.Context) {
		server := ctx.DefaultQuery("server", defaultServer)
		limit := 10
		if rawUserLimit := ctx.Query("limit"); rawUserLimit != "" {
			if userLimit, err := strconv.Atoi(rawUserLimit); err != nil {
				ctx.AbortWithStatusJSON(400, gin.H{"msg": "限制值不合法"})
				return
			} else {
				limit = userLimit
			}
		}
		key := server + ":" + ctx.Param("stat")
		var result []redis.Z
		ctxt := context.Background()
		if limit > 0 {
			result, err = rdb.ZRevRangeWithScores(
				ctxt, key, 0, int64(limit-1),
			).Result()
		} else {
			result, err = rdb.ZRangeWithScores(
				ctxt, key, 0, int64(-limit-1),
			).Result()
		}
		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"msg": "查询失败"})
			return
		}
		ctx.AbortWithStatusJSON(200, gin.H{
			"msg":  "查询成功",
			"data": result,
		})
	})
	router.Run(":" + port)
}
