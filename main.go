package main

import (
	"fmt"
	"time"
)

func main() {
	pauseCh := make(chan struct{})
	go Waiting(pauseCh)
	ticker := tickFactory(pauseCh)
	for range ticker {
		fmt.Println(".")
	}
}

func Waiting(pause chan struct{}) {
	for range time.Tick(time.Second * 5) {
		pause <- struct{}{}
	}
}

func tickFactory(pause chan struct{}) chan time.Time {
	cht := make(chan time.Time)
	isPause := false

	tick := time.Tick(time.Second)

	go func() {
		for {
			select {
			case <-tick:
				if !isPause {
					cht <- time.Now()
				}

			case <-pause:
				isPause = !isPause
			}
		}
	}()

	return cht
}
