module guilhem-mateo.fr/mod.v0

require (
	github.com/foolin/gin-template v0.0.0-20190415034731-41efedfb393b // indirect
	github.com/gin-gonic/gin v1.6.3
	guilhem-mateo.fr/git/Wariie/modbase.git v0.0.0-20200709160058-145b8bece1b2
	guilhem-mateo.fr/testgo/app/rand v0.0.0 // indirect
)

go 1.14

//replace guilhem-mateo.fr/git/Wariie/modbase.git => guilhem-mateo.fr/git/Wariie/modbase

replace guilhem-mateo.fr/testgo/app/rand => ../testgo/app/rand

replace guilhem-mateo.fr/testgo/app/com => ../testgo/app/com
