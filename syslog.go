package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	//currentProgramName := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, "BING")
	if err != nil {
		log.Fatal(err)
	}
	sysLog.Crit("Crit: Logging in Go!")
	sysLog.Info("Info: Bing Bing BO")
	fmt.Fprintf(sysLog, "log.Print: Logging in Go!")
}
