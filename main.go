package main

import (
	"github.com/opinionated/scrapeClient/client"
	"github.com/opinionated/utils/log"
)

func main() {
	// initialize logs to write to standard out
	log.InitStd()

	// see doc for server's IP, we don't want to upload it
	// TODO: store things like IP in hidden config files
	c := client.Client{IP: "http://107.170.148.224:8080/"}

	c.Run()
}
