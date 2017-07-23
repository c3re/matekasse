package main

import (
	"fmt"
	mk "github.com/c3re/matekasse"
	"os"
)

const helptext string = "matekasse usage: \n\n -l IF:port default: *:80 \n -f DB-File default: matekasse.sqlite \n -h print help and exit\n"

func main() {
	conf := true
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-l":
			i++
			if i < len(os.Args) {
				mk.SetIf(os.Args[i])
			} else {
				conf = false
			}
		case "-f":
			i++
			if i < len(os.Args) {
				mk.SetDb(os.Args[i])
			} else {
				conf = false
			}
		default:
			conf = false
			i = len(os.Args)
		}

	}
	if conf {
		mk.Start()
	} else {
		printHelp()
	}
}

func printHelp() {
	fmt.Printf(helptext)
}
