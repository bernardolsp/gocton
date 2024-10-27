package main

import (
	"sync"
	"time"
)

var retryAmount = 5
var retryDelay = 2 * time.Second
var initialDelay = retryDelay
var wg sync.WaitGroup

func main() {
	messager := bootstrap_messager()
	bootstrap_parser(messager)
}
