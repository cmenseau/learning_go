package main

import (
	"flag"
	"log"
	"main/internal/runner"
	"os"

	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = pprof.StartCPUProfile(f)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer pprof.StopCPUProfile()
	}

	for range 50 {
		runner.Run([]string{"-iwr", `defer\|func`, "/home/menseau/Documents/Go/learning_go/0_go_class_matt_holiday"})
	}
}
