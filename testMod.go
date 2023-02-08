package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/Wariie/go-woxy/com"
	"github.com/Wariie/go-woxy/modbase"
)

func main() {
	var m modbase.ModuleImpl

	m.Name = "mod.v0"
	m.InstanceName = "mod test v0"

	m.SetHubAddress("guilhem-mateo.fr")
	//m.SetHubPort("2000")
	m.SetHubProtocol("https")
	m.SetPort("2985")
	m.SetCommand("msg", msg)
	m.Init()
	m.Register("/", index(), "WEB")
	m.Run()
}

func index() com.HandlerFunc {
	return com.HandlerFunc(func(ctx *com.Context) {

		data := IndexPage{
			Title:  "Guilhem MATEO",
			Secret: modbase.GetModManager().GetSecret(),
			Hash:   modbase.GetModManager().GetMod().Hash,
		}

		tmpl := template.Must(template.ParseFiles("./views/index.html"))
		err := tmpl.ExecuteTemplate(ctx.ResponseWriter, "index", data)
		if err != nil {
			log.Fatalln("Error : ", err)
		}

		log.Println("GET / mod.v0", ctx.RemoteAddr)
	})
}

type IndexPage struct {
	Title  string
	Secret string
	Hash   string
}

func msg(r *com.Request, w http.ResponseWriter, re *http.Request, mod *modbase.ModuleImpl) (string, error) {
	cr := (*r).(*com.CommandRequest)
	log.Println("MESSAGE :", cr.Content)
	return "OK", nil
}
