package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Wariie/go-woxy/com"
	modbase "github.com/Wariie/go-woxy/modbase"
)

func main() {
	var m modbase.ModuleImpl

	m.Name = "mod.v0"
	m.InstanceName = "mod test v0"
	m.SetServer("", "", "2985", "")
	m.SetCommand("msg", msg)
	m.Init()
	m.Register("GET", "/", index, "WEB")
	m.Run()
}

func index(ctx *gin.Context) {
	ctx.HTML(http.StatusAccepted, "index.html", gin.H{
		"title":  "Guilhem MATEO",
		"secret": modbase.GetModManager().GetSecret(),
		"hash":   modbase.GetModManager().GetMod().Hash,
	})
	log.Println("GET / mod.v0", ctx.Request.RemoteAddr)
}

func msg(r *com.Request, c *gin.Context, mod *modbase.ModuleImpl) (string, error) {
	cr := (*r).(*com.CommandRequest)
	log.Println("MESSAGE :", cr.Content)
	return "OK", nil
}
