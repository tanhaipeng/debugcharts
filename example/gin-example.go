package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mkevac/debugcharts"
	"log"
	_ "net/http/pprof"
	"time"
)

func ginDummyAllocations() {
	type t struct {
		a uint64
		b map[uint64][]byte
	}
	d := make([]t, 0, 0)

	for {
		select {
		case <-time.Tick(time.Second * 20):
			d = make([]t, 0, 0)
			log.Println(len(d))
		case <-time.Tick(time.Second * 8):
			d = make([]t, 500000, 500000)
			log.Println(len(d))
		}
	}
}

func main() {
	go ginDummyAllocations()

	router := gin.Default()
	debugcharts.Register(router)

	log.Fatal(router.Run(":8080"))
}
