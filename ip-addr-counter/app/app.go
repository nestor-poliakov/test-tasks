package app

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

func Run() {
	t := time.Now()
	if len(os.Args) < 2 {
		log.Fatalf("no filename in command line args")
	}
	threads := runtime.NumCPU()
	if len(os.Args) >= 3 {
		var err error
		threads, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("parse number of threads error: %s", err)
		}
		if threads < 1 {
			log.Fatalf("error: number or threads less then 1")
		}
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	p := NewProcessor()
	p.Process(f, threads)
	result := p.Result()

	fmt.Printf("result: %d\nprocessing time: %s\n", result, time.Since(t))
}
