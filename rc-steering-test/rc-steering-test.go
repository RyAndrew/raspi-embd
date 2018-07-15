package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"

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

	fmt.Printf("setting servo %v to %v\r\n", servoSetNum, servoSetValue)

	if err := pca9685Inst.SetPwm(servoSetNum, 0, servoSetValue); err != nil {
		panic(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	io.WriteString(writer, `{"success":true}`)
}

var connectedClients []*webSocketClient
var mutex = &sync.Mutex{}

type webSocketClient struct {

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

func webSocketClientReader() {

}
func webSocketClientWriter(conn *websocket.Conn, socketData []byte) {
	if err := conn.WriteMessage(websocket.TextMessage, socketData); err != nil {
		log.Printf("Error writing to socket: %s", err)
		unRegisterClient(conn)
		return
	}
}
func registerClient(conn *websocket.Conn) {

	client := &webSocketClient{conn: conn, send: make(chan []byte, 256)}

	mutex.Lock()
	connectedClients = append(connectedClients, client)
	mutex.Unlock()

	//_, message, err := conn.ReadMessage()
}
func unRegisterClient(conn *websocket.Conn) {
	log.Println("Closed WebSocket Connection UnRegistering")

	mutex.Lock()
	var connectedClientsNew []*webSocketClient

	//remove the disconnected client from the connectedClients slice
	for _, client := range connectedClients {
		if client.conn == conn {
			continue
		} else {
			connectedClientsNew = append(connectedClientsNew, client)
		}
	}

	connectedClients = connectedClientsNew
	mutex.Unlock()
}
func webSocketSendJsonToAllClients(jsonData interface{}) {

	//log.Printf("writing to %d clients\n", len(connectedClients))

	for _, client := range connectedClients {
		client.conn.WriteJSON(jsonData)
	}
}
func adcTicker(bus embd.I2CBus) {
	tick := time.Tick(time.Millisecond * 2000)
	for {
		select {
		case <-tick:

			jsonData := make(map[string]interface{})
			jsonData["steeringCurrent"] = fmt.Sprintf("%v", readAdcValue(bus))

			webSocketSendJsonToAllClients(jsonData)
		}
	}
}
func main() {

	if err := embd.InitI2C(); err != nil {
		panic(err)
	}
	defer embd.CloseI2C()
	i2cBus := embd.NewI2CBus(1)

	//Adafruit board is address 0x60
	//Generic PCA9685 address is 0x40
	pca9685Inst = pca9685.New(i2cBus, 0x60)
	pca9685Inst.Freq = 60
	defer pca9685Inst.Close()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	srv := &http.Server{
		Addr: ":8090",
		//Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fs := http.FileServer(http.Dir("/home/pi/rc-steering-test/webroot"))
	http.Handle("/", fs)
	http.Handle("/updateServo", http.HandlerFunc(udpateServo))

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	http.HandleFunc("/wsapi", func(w http.ResponseWriter, request *http.Request) {

		log.Println("New WebSocket Connection")
		log.Println(request.RemoteAddr)
		conn, err := upgrader.Upgrade(w, request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		registerClient(conn)
	})

	log.Println("Listening on 8090")

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Printf("Httpserver: ListenAndServe() quitting: %s", err)
			shutdown <- nil
		}
	}()

	go adcTicker(i2cBus)

	//block waiting for channel
	<-shutdown

	log.Println("Server is shutting down")
	os.Exit(0)

}
