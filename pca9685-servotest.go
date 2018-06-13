package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/kidoman/embd"
	"github.com/kidoman/embd/controller/pca9685"

	_ "github.com/kidoman/embd/host/all"
)

func main() {
	flag.Parse()

	if err := embd.InitI2C(); err != nil {
		panic(err)
	}
	defer embd.CloseI2C()

	bus := embd.NewI2CBus(1)

	pca9685 := pca9685.New(bus, 0x40)
	pca9685.Freq = 60
	defer pca9685.Close()

	if err := pca9685.SetPwm(0, 0, 150); err != nil {
		panic(err)
	}
	if err := pca9685.SetPwm(1, 0, 600); err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	tick := time.Tick(time.Millisecond * 1000)
	servoStatus := false

	for {
		select {
		case <-c:
			return
		case <-tick:
			if servoStatus {

				if err := pca9685.SetPwm(0, 0, 150); err != nil {
					panic(err)
				}
				if err := pca9685.SetPwm(1, 0, 600); err != nil {
					panic(err)
				}

				servoStatus = false
			} else {

				if err := pca9685.SetPwm(0, 0, 600); err != nil {
					panic(err)
				}
				if err := pca9685.SetPwm(1, 0, 150); err != nil {
					panic(err)
				}

				servoStatus = true
			}
			fmt.Println("tick!")
		}
	}
}
