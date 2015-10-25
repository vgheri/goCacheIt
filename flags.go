package main

import (
	"flag"
)

// Maximum amount of memory to be used by the system, as specified by the user
var maxMemory uint64

// HTTP server port to listen to
var webServerPort int

func init() {
	flag.Uint64Var(&maxMemory, "maxMem", defaultMaxMemory, "Maximum amount of"+
		" allocated memory in MB. Default value is 512 MB")
	flag.IntVar(&webServerPort, "port", defaultHTTPPort, "Port to listen to."+
		" Default value is 3000")
	flag.Parse()
}
