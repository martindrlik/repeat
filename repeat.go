package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

var (
	n        = flag.Int("n", 1000, "repeat command n times, negative n for infinity")
	delay    = flag.Duration("delay", time.Duration(0), "delay first command start")
	distance = flag.Duration("distance", time.Duration(2*time.Second), "time distance between command starts")
)

func main() {
	flag.Parse()
	args := flag.Args()
	time.Sleep(*delay)
	for {
		ch := time.After(*distance)
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "repeat: %v\n", err)
		}
		if *n > 0 {
			*n--
		}
		if *n == 0 {
			break
		}
		<-ch
	}
}
