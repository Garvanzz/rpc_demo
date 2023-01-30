package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"rpc_demo/task"
	"syscall"
)

func main() {
	var module string
	flag.StringVar(&module, "module", "", "run module")
	flag.Parse()
	fmt.Printf("start run %s module", module)
	switch module {
	case "task":
		task.New().Run()
	default:
		log.Fatal("exiting,module param error!")
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	fmt.Println("server exiting")
}
