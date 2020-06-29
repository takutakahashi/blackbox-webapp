package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	mysqlHost := os.Getenv("MYSQL_HOST")
	if mysqlHost == "" {
		panic("MYSQL_HOST is empty")
	}
	mysqlPort := os.Getenv("MYSQL_PORT")
	if mysqlPort == "" {
		panic("MYSQL_PORT is empty")
	}
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", mysqlHost, mysqlPort), 3*time.Second)
	if err != nil {
		panic(err)
	}
	conn.Close()
}
