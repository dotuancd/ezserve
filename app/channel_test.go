package app_test

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	c := make(chan string)

	begin := time.Now()
	go sleep(3, c)
	go sleep(2, c)
	go sleep(1, c)

	x, y, z := <- c, <-c, <- c
	finish := time.Now().Sub(begin)
	fmt.Printf("Finish after: %s\n", finish.String())
	fmt.Println(x, y, z)
}

func sleep(seconds int, done chan string) {
	duration := time.Duration(seconds) * time.Second
	time.Sleep(duration)
	done <- fmt.Sprintf("sleep %s\n", duration.String())
}

