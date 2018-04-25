package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

// Config - holds values for configuration
type Config struct {
	Hostname    string
	RandomSleep bool
	RandomStart int
	RandomEnd   int
	SleepTime   int
}

// Initializes a global Config struct
var config Config

// Send log output to stdout as well as logger (which writes to log.txt)
func doLog(logLine string) {
	fmt.Println(logLine)
	// add carriage return for Windows
	log.Println(logLine + "\r")
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	// Parse flags/defaults
	flag.StringVar(&config.Hostname, "hostname", "kube-dns.kube-system.svc.cluster.local", "override the DNS hostname to resolve")
	flag.BoolVar(&config.RandomSleep, "random-sleep", false, "sleep a random amount of time (between 1 and 180 seconds) between checks")
	flag.IntVar(&config.RandomStart, "random-start", 1, "Start of range for random sleep")
	flag.IntVar(&config.RandomEnd, "random-end", 180, "End of range for random sleep")
	flag.IntVar(&config.SleepTime, "sleep-time", 1, "run checks against hostname every X seconds")

	flag.Parse()

	// initialize logging
	f, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)

	fmt.Println("Saving output to log.txt in current directory")

	if config.RandomSleep {
		rand.Seed(time.Now().UnixNano())
		doLog(fmt.Sprintf("Starting DNS lookups for %s (%d-%d second(s) random sleep between checks)", config.Hostname, config.RandomStart, config.RandomEnd))
	} else {
		doLog(fmt.Sprintf("Starting DNS lookups for %s (%d second(s) sleep between checks)", config.Hostname, config.SleepTime))
	}

	for {
		_, err := net.LookupHost(config.Hostname)
		if err != nil {
			doLog(err.Error())
		} else {
			doLog("DNS lookup succeeded")
		}

		var sleepTime int
		if config.RandomSleep {
			sleepTime = random(config.RandomStart, config.RandomEnd)
		} else {
			sleepTime = config.SleepTime
		}
		time.Sleep(time.Duration(sleepTime) * time.Second)

	}

}
