package test

import (
	"flag"
)

func main(){
	runserver := flag.Bool("runserver",false,"run a server")
	flag.Parse()
	if *runserver {

	}else{

	}
}


