package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"

	"github.com/kidoman/embd"
	"github.com/kidoman/embd/controller/pca9685"
	_ "github.com/kidoman/embd/host/all"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 30 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// var (
// 	newline = []byte{'\n'}
// 	space   = []byte{' '}
// )

var WebSocketClientMap map[*WebSocketClient]bool

type WebSocketClient struct {

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

var pca9685Inst *pca9685.PCA9685
var steeringAdcValue uint16
var steeringMax uint16 = 2047
var steeringTargetPoint uint16 = steeringMax / 2

var stopSteeringLoopChan = make(chan struct{}, 1)

func outputFailure(writer http.ResponseWriter) {

	writer.Header().Set("Content-Type", "application/json")
	io.WriteString(writer, `{"success":false}`)

}

func (c *WebSocketClient) webSocketClientReader() {
	defer func() {
		unRegisterClient(c)
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		//message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		//c.hub.broadcast <- message
		processClientMessage(message)
	}
}
func (c *WebSocketClient) webSocketClientWriter() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				//w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			fmt.Println("Sending socket ping")
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
func registerClient(conn *websocket.Conn) {

	client := &WebSocketClient{conn: conn, send: make(chan []byte, 256)}

	WebSocketClientMap[client] = true

	go client.webSocketClientReader()
	go client.webSocketClientWriter()

}
func unRegisterClient(client *WebSocketClient) {

	log.Println("Closed WebSocket Connection UnRegistering")

	if _, ok := WebSocketClientMap[client]; ok {
		delete(WebSocketClientMap, client)
		close(client.send)
	}
}
func webSocketSendJsonToAllClients(jsonData interface{}) {

	//log.Printf("writing to %d clients\n", len(connectedClients))
	for client := range WebSocketClientMap {
		client.conn.WriteJSON(jsonData)
	}
}

func processClientMessage(message []byte) {

	jsonData := make(map[string]interface{})

	json.Unmarshal(message, &jsonData)
	fmt.Println(string(message))
	//fmt.Printf("%+v\n", jsonData)

	if _, ok := jsonData["action"]; !ok {
		log.Println("Invalid Message, no action param")
		return
	}

	switch jsonData["action"] {
	case "updateSteering":
		//fmt.Printf("%+v\n", jsonData["value"])
		pos := jsonData["value"].(float64)
		setSteeringPosition(pos)
		break
	case "stopSteeringMovement":
		stopSteeringMovement()
		break
	case "stopSteeringControlLoop":
		stopSteeringLoopChan <- struct{}{}
		break
	case "startSteeringControlLoop":
		stopSteeringLoopChan <- struct{}{}

		for len(stopSteeringLoopChan) > 0 {
			<-stopSteeringLoopChan
		}
		go startSteeringControlLoop()
		break
	}

}

func stopSteeringMovement() {
	setPwmChanPercent(2, 0)
	setPwmChanPercent(3, 0)
	setPwmChanPercent(4, 0)
}
func setSteeringPosition(pos float64) {
	//fmt.Printf("set pos=%v, ", pos)
	steeringTargetPoint = uint16(float64(steeringMax) * pos / 100.0)
	//fmt.Printf("set steeringTargetPoint=%v\n", steeringTargetPoint)
	//fmt.Printf("steeringMax %v * pos %v / 100\n", steeringMax, pos)
}
func steeringSetPointAdjust() {
	successfulThreshold := .09

	fmt.Printf("set=%v, actual=%v\n", steeringTargetPoint, steeringAdcValue)

	if float64(steeringAdcValue) < float64(steeringTargetPoint)*float64(1+successfulThreshold) {
		if float64(steeringAdcValue) > float64(steeringTargetPoint)*(1-successfulThreshold) {
			stopSteeringMovement()
			return
		}
	}
	if steeringTargetPoint > steeringAdcValue {
		setPwmChanPercent(2, 60)
		setPwmChanPercent(3, 100)
		setPwmChanPercent(4, 0)
	} else {
		setPwmChanPercent(2, 60)
		setPwmChanPercent(3, 0)
		setPwmChanPercent(4, 100)
	}
}
func startSteeringControlLoop() {
	fmt.Println("startSteeringControlLoop")
	checkSteeringPositionTicker := time.Tick(time.Millisecond * 30)
	for {
		select {
		case <-checkSteeringPositionTicker:
			steeringSetPointAdjust()
		case <-stopSteeringLoopChan:
			stopSteeringMovement()
			return
		}
	}
}
func setPwmChanPercent(chanNo int, percent int) {

	//fmt.Printf("setting PWM chan=%d ", chanNo)

	pwmCalc := 4095.0 * percent / 100.0
	//fmt.Printf(", calc=%v ", pwmCalc)
	pwmSet := int(pwmCalc)
	//fmt.Printf(", set=%v \n", pwmSet)

	if err := pca9685Inst.SetPwm(chanNo, 0, pwmSet); err != nil {
		panic(err)
	}
}
func adcTicker(bus embd.I2CBus) {

	initAdc(bus)

	steeringAdcValue = readAdcValue(bus)

	adcValueBroadcastTicker := time.Tick(time.Millisecond * 150)
	adcReadTicker := time.Tick(time.Millisecond * 20)
	var adcTickNumber uint16 = 0
	for {
		select {
		case <-adcValueBroadcastTicker:
			jsonData := make(map[string]interface{})
			jsonData["steeringCurrent"] = fmt.Sprintf("%v", steeringAdcValue)
			jsonData["steeringTargetPoint"] = fmt.Sprintf("%v", steeringTargetPoint)

			webSocketSendJsonToAllClients(jsonData)
		case <-adcReadTicker:
			adcTickNumber++

			//start := time.Now()

			steeringAdcValue = readAdcValue(bus)

			//elapsed := time.Since(start)

			//fmt.Printf("%4.d Read ADC Value %d\n", adcTickNumber, steeringAdcValue)

			//fmt.Printf(" time: %s\n", elapsed)
		}
	}
}
func main() {
	flag.Parse()

	if err := embd.InitI2C(); err != nil {
		panic(err)
	}
	defer embd.CloseI2C()
	i2cBus := embd.NewI2CBus(1)

	//https://github.com/adafruit/Adafruit-Motor-HAT-Python-Library/blob/master/Adafruit_MotorHAT/Adafruit_PWM_Servo_Driver.py
	//Adafruit board is address 0x60
	//Generic PCA9685 address is 0x40
	pca9685Inst = pca9685.New(i2cBus, 0x60)
	pca9685Inst.Freq = 60
	pca9685Inst.Wake()
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

	WebSocketClientMap = make(map[*WebSocketClient]bool)
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
	go startSteeringControlLoop()

	//block waiting for channel
	<-shutdown

	stopSteeringMovement()
	log.Println("Server is shutting down")
	os.Exit(0)

}
