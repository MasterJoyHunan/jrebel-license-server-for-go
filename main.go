package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-jrebel-license/handle"
	"net/http"
)

var (
	host string
	port int
)

func init() {
	flag.StringVar(&host, "h", "127.0.0.1", "host")
	flag.IntVar(&port, "p", 8877, "port")
}

func main() {
	flag.Parse()
	engine := gin.Default()
	engine.GET("/", handle.Index)
	engine.POST("/jrebel/leases", handle.JrebelLeasesHandler)
	engine.POST("/jrebel/leases/1", handle.JrebelLeases1Handler)
	engine.POST("/agent/leases", handle.JrebelLeasesHandler)
	engine.POST("/agent/leases/1", handle.JrebelLeases1Handler)
	engine.POST("/jrebel/validate-connection", handle.JrebelValidateHandler)
	engine.POST("/rpc/ping.action", handle.PingHandler)
	engine.POST("/rpc/obtainTicket.action", handle.ObtainTicketHandler)
	engine.POST("/rpc/releaseTicket.action", handle.ReleaseTicketHandler)
	panic(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), engine))
}
