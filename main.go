package main

import (
	"log"

	"github.com/dialogbox/mpipego/mpipe"

	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	mpipe.Execute()
}
