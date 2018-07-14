package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/kidoman/embd"
	"github.com/kidoman/embd/controller/pca9685"
	_ "github.com/kidoman/embd/host/all"
)

var pca9685Inst *pca9685.PCA9685

func outputFailure(writer http.ResponseWriter) {

	writer.Header().Set("Content-Type", "application/json")
	io.WriteString(writer, `{"success":false}`)

}
func udpateServo(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	servo := request.FormValue("servo")
	if servo == "" {
		outputFailure(writer)
		return
	}
	value := request.FormValue("value")
	if value == "" {
		outputFailure(writer)
		return
	}

	servoNum := servo[len(servo)-1:]
	fmt.Print("Servo #")
	fmt.Print(servoNum)
	fmt.Print(" = ")
	fmt.Println(value)

	servoSetNum, err := strconv.Atoi(servoNum)

	if err != nil {
		outputFailure(writer)
		return
	}
	servoSetValue, err := strconv.Atoi(value)
	if err != nil {
		outputFailure(writer)
		return
	}

	fmt.Println("setting MOTOR scaling range between 0 and 4096")
	servoSetValue = (servoSetValue * 4095) / 100

	fmt.Printf("setting servo %d to %d\r\n", servoSetNum, servoSetValue)

	if err := pca9685Inst.SetPwm(servoSetNum, 0, servoSetValue); err != nil {
		panic(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	io.WriteString(writer, `{"success":true}`)
}

func main() {

	if err := embd.InitI2C(); err != nil {
		panic(err)
	}
	defer embd.CloseI2C()
	bus := embd.NewI2CBus(1)

	//Adafruit board is address 0x60
	//Generic PCA9685 address is 0x40
	pca9685Inst = pca9685.New(bus, 0x60)

	pca9685Inst.Freq = 60
	defer pca9685Inst.Close()

	if err := pca9685Inst.SetPwm(1, 0, 150); err != nil {
		panic(err)
	}
	if err := pca9685Inst.SetPwm(2, 0, 150); err != nil {
		panic(err)
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	srv := &http.Server{
		Addr:           ":8090",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fs := http.FileServer(http.Dir("/home/pi/pca9685-servo-webtest/webroot"))
	http.Handle("/", fs)
	http.Handle("/updateServo", http.HandlerFunc(udpateServo))

	log.Println("Listening on 8090")

	go func() {
		err :=
			srv.ListenAndServe()
		if err != nil {
			log.Printf("Httpserver: ListenAndServe() quitting: %s", err)
			shutdown <- nil
		}
	}()

	//blocking call, wait for shutdown message
	<-shutdown

	log.Println("Server shutting down")
	os.Exit(0)

}
