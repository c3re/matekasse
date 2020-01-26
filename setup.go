package matekasse

import "fmt"

var listen = ":8080"
var dbfile = "matekasse.sqlite"
var script = ""

func SetDb(s string) {
	dbfile = s
}

func SetIf(s string) {
	listen = s
}

func SetScript(s string) {
	script = s
}

func Start() {
	fmt.Println("Welcome to Matekasse!\n")
	fmt.Println("Listen Interface: ", listen)
	fmt.Println("SQLite-File:      ", dbfile)
	if script == "" {
		fmt.Println("external script:   none")
	} else {
		fmt.Println("external script:  ", script)
	}
	err := connectDB()
	if err != nil {
		panic("Without database, its not worth starting...")
	}
	startServer()
}

func ce(err error) {
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
	}
}
