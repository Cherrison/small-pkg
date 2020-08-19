package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Cherrison/jstime"

	game "github.com/Cherrison/small-pkg/life-game"
)

func main() {
	stopCh := make(chan struct{})
	l := game.NewLife(100, 100)
	stop := jstime.SetInterval(time.Millisecond*300, func() {
		l.Step()
		fmt.Print("\x0c", l, "\n\n 按下Enter 停止") // Clear screen and print field.
	})
	go func() {
		scaner := bufio.NewScanner(os.Stdin)
		for scaner.Scan() {
			stop.Stop()
			stopCh <- struct{}{}
		}
	}()
	<-stopCh
}
