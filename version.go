package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
)

var (
	VERSION = "0.1poc"
)

func LogVersion() {
	log.Fatalf("Version: %s", VERSION)
}

func PrintVersion() {
	fmt.Printf("Version: %s\n", VERSION)
}
