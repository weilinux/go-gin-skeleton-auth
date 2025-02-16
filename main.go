package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gookit/color"
	"github.com/weilinux/go-gin-skeleton-auth/web"

	// boot and init some services(log, cache, eureka)
	"github.com/weilinux/go-gin-skeleton-auth/app"
	"github.com/weilinux/go-gin-skeleton-auth/model/mongo"
	"github.com/weilinux/go-gin-skeleton-auth/model/myrds"
	"github.com/weilinux/go-gin-skeleton-auth/model/mysql"
)

func init() {
	var err error
	app.Bootstrap("./config")

	// - redis, mongo, mysql connection
	err = myrds.InitRedis()
	checkError("Rds init error:", err)

	err = mysql.InitMysql()
	checkError("Db init error:", err)

	err = mongo.InitMongo()
	checkError("Mgo init error:", err)
	// initEurekaService()

	web.InitServer()
}

func main() {
	listenSignals()

	// init services
	log.Printf("======================== Begin Running(PID: %d) ========================", os.Getpid())

	web.Run()
}

func checkError(prefix string, err error) {
	if err != nil {
		color.Error.Println(prefix, err.Error())
		os.Exit(2)
	}
}

// listenSignals Graceful start/stop server
func listenSignals() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go handleSignals(sigChan)
}

// handleSignals handle process signal
func handleSignals(c chan os.Signal) {
	log.Print("Notice: System signal monitoring is enabled(watch: SIGINT,SIGTERM,SIGQUIT)\n")

	switch <-c {
	case syscall.SIGINT:
		fmt.Println("\nShutdown by Ctrl+C")
	case syscall.SIGTERM: // by kill
		fmt.Println("\nShutdown quickly")
	case syscall.SIGQUIT:
		fmt.Println("\nShutdown gracefully")
		// do graceful shutdown
	}

	// sync logs
	_ = app.Logger.Sync()
	_ = mysql.Close()
	_ = myrds.ClosePool()
	mongo.Close()

	// unregister from eureka
	// erkServer.Unregister()

	// 等待一秒
	time.Sleep(1e9 / 2)
	color.Info.Println("\n  GoodBye...")

	os.Exit(0)
}
