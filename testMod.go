package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	modbase "guilhem-mateo.fr/git/Wariie/modbase.git"
)

func main() {
	var m modbase.ModuleImpl
	m.Init()
	m.Name = "mod.v0"
	m.InstanceName = "mod test v0"
	modbase.HubAddress = "localhost"
	modbase.ModulePort = "2985"

	//TODO ADD REGISTER
	//m.Register("GET", "/", index, "")
	m.Register("GET", "/", index, "WEB")
	m.Run()
}

func index(ctx *gin.Context) {
	ctx.HTML(http.StatusAccepted, "index.html", gin.H{
		"title": "Main website", //IGNORE THIS
	})
	log.Println("GET / mod.v0", ctx.Request.RemoteAddr)
}
