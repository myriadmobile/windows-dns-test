package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

// Config - holds values for configuration
type Config struct {
	SleepTime int
	Hostname  string
}

// Initializes a global Config struct
var config Config

// Send log output to stdout as well as logger (which writes to log.txt)
func doLog(logLine string) {
	fmt.Println(logLine)
	// add carriage return for Windows
	log.Println(logLine + "\r")
}

func main() {
	// Parse flags/defaults
	flag.IntVar(&config.SleepTime, "sleep-time", 1, "run checks against hostname every X seconds")
	flag.StringVar(&config.Hostname, "hostname", "kube-dns.kube-system.svc.cluster.local", "override the DNS hostname to resolve")
	flag.Parse()

	// initialize logging
	f, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)

	fmt.Println("Saving output to log.txt in current directory")
	doLog(fmt.Sprintf("Starting DNS lookups for %s (%d second(s) sleep between checks)", config.Hostname, config.SleepTime))
	for {
		_, err := net.LookupHost(config.Hostname)
		if err != nil {
			doLog(err.Error())
		} else {
			doLog("DNS lookup succeeded")
		}
		time.Sleep(time.Duration(config.SleepTime) * time.Second)
	}

}
