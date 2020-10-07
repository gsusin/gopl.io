// Giancarlo Susin
// Exercício 9.4

package main

import (
	"fmt"
	"time"
)

func main() {
	var nStages int = 2000000
	channels := make([]chan struct{}, nStages)
	done := make(chan struct{})
	t0 := time.Now()
	for n := 0; n < nStages; n++ {
		channels[n] = make(chan struct{})
		if n == 0 {
			go func() {
				channels[0] <- struct{}{}
				<-done
			}()
			continue
		}
		go func(chPrev, chNext chan struct{}) {
			v := <-chPrev
			chNext <- v
			<-done
		}(channels[n-1], channels[n])
	}
	<-channels[nStages-1]
	t1 := time.Now()
	fmt.Printf("terminou %d estágios em %f s\n", nStages, t1.Sub(t0).Seconds())
}
