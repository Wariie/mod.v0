module guilhem-mateo.fr/mod.v0

require (
	github.com/foolin/gin-template v0.0.0-20190415034731-41efedfb393b // indirect
	github.com/gin-gonic/gin v1.6.3
	guilhem-mateo.fr/git/Wariie/modbase.git v0.0.0-20200709152053-458cf801fe55 // indirect
	guilhem-mateo.fr/testgo/app/com v0.0.0 // indirect
	guilhem-mateo.fr/testgo/app/rand v0.0.0 // indirect
)

go 1.14

//replace guilhem-mateo.fr/git/Wariie/modbase.git => guilhem-mateo.fr/git/Wariie/modbase

replace guilhem-mateo.fr/testgo/app/rand => ../testgo/app/rand

replace guilhem-mateo.fr/testgo/app/com => ../testgo/app/com
