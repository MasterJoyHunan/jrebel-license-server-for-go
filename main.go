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
	flag.StringVar(&host, "h", "0.0.0.0", "host")
	flag.IntVar(&port, "p", 8877, "port")
}

func main() {
	flag.Parse()
	engine := gin.Default()
	engine.GET("/", handle.Index)
	engine.GET("/jrebel/leases", handle.JrebelLeasesHandler)
	engine.GET("/jrebel/leases/1", handle.JrebelLeases1Handler)
	engine.GET("/agent/leases", handle.JrebelLeasesHandler)
	engine.GET("/agent/leases/1", handle.JrebelLeases1Handler)
	engine.GET("/jrebel/validate-connection", handle.JrebelValidateHandler)
	engine.GET("/rpc/ping.action", handle.PingHandler)
	engine.GET("/rpc/obtainTicket.action", handle.ObtainTicketHandler)
	engine.GET("/rpc/releaseTicket.action", handle.ReleaseTicketHandler)
	panic(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), engine))
}
