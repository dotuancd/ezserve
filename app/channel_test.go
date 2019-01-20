package app_test

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {

	fmt.Print(math.Ceil(float64(11)/ float64(10)))

}

func sleep(seconds int, done chan string) {
	duration := time.Duration(seconds) * time.Second
	time.Sleep(duration)
	done <- fmt.Sprintf("sleep %s\n", duration.String())
}

