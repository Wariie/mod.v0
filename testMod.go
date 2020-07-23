package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	modbase "github.com/Wariie/go-woxy/modbase"
)

func main() {
	var m modbase.ModuleImpl
	
	m.Name = "mod.v0"
	m.InstanceName = "mod test v0"
	modbase.HubAddress = "localhost"
	modbase.ModulePort = "2985"

	m.Init()
	m.Register("GET", "/", index, "WEB")
	m.Run()
}

func index(ctx *gin.Context) {
	ctx.HTML(http.StatusAccepted, "index.html", gin.H{
		"title": "Guilhem MATEO", //IGNORE THIS
	})
	log.Println("GET / mod.v0", ctx.Request.RemoteAddr)
}
