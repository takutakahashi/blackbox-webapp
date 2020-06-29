package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/prometheus/common/log"
)

var globalCnt int

func main() {
	globalCnt = 10
	port := os.Getenv("PORT")
	if port == "" {
		panic("environment variable PORT is empty")
	}
	mysqlHost := os.Getenv("MYSQL_HOST")
	if mysqlHost == "" {
		panic("environment variable MYSQL_HOST is empty")
	}
	mysqlPort := os.Getenv("MYSQL_PORT")
	if mysqlPort == "" {
		panic("environment variable MYSQL_PORT is empty")
	}
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", mysqlHost, mysqlPort), 3*time.Second)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.Write([]byte("zebra\n"))

	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		panic("environment variable REDIS_HOST is empty")
	}

	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		panic("environment variable REDIS_PORT is empty")
	}
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err = rdb.Set(ctx, "animal", "elephant", 0).Err()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
func hello(c echo.Context) error {
	if globalCnt < 0 {
		log.Info("gorilla")
	} else {
		log.Info(globalCnt)
	}
	globalCnt--
	return c.String(http.StatusOK, "Lion")
}
