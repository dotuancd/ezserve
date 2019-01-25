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

func TestOpenClose(t *testing.T) {
	octo := OctoCat{}
	octo.LegsCount = 5;
	fmt.Println(octo.Legs())
	octo.PrintLegs()
}

func sleep(seconds int, done chan string) {
	duration := time.Duration(seconds) * time.Second
	time.Sleep(duration)
	done <- fmt.Sprintf("sleep %s\n", duration.String())
}

type Cat struct {
	Name string
	LegsCount int
}

func (c Cat) Legs() int {
	return 4
}

func (c Cat) PrintLegs() {
	fmt.Printf("I have %d legs\n", c.Legs())
}

type OctoCat struct {
	Cat
}

func (c OctoCat) Legs() int {
	return 5
}


