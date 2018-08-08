package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"time"

	"github.com/RyAndrew/pidctrl"
	"github.com/gorilla/websocket"
	"github.com/tarm/serial"

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

	arduinoLedBlinkRed = "5"
	arduinoLedBlue     = "2"
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

var jawMax int = 1820
var jawMin int = 1510
var jawRange int = jawMax - jawMin

var trexTiltMax int = 1690
var trexTiltMin int = 1000
var trexTiltRange int = trexTiltMax - trexTiltMin

var trexPanMax int = 1000
var trexPanMin int = 2000
var trexPanRange int = trexPanMax - trexPanMin

var steeringMax uint16 = 1130
var steeringMin uint16 = 890
var steeringRange uint16 = steeringMax - steeringMin
var steeringTargetPoint uint16 = (steeringRange / 2) + steeringMin

var throttlePwmFreq float64 = 50.0

//var throttlePwmFreqUsCalc float64 = 1000 / float64(throttlePwmFreq) / 4096 * 1000

var throttlePwmFreq1ms int = 200
var throttlePwmFreq1500us int = 300
var throttlePwmFreqUsCalc float64 = float64(throttlePwmFreq1ms) / 1000

var throttlePwmChannel int = 0
var throttlePwmMax float64 = 1000.0
var throttlePwmOffset float64 = 1000.0

var stopSteeringLoopChan = make(chan struct{}, 1)

var serialPortMessages = make(chan []byte, 25)

func outputFailure(writer http.ResponseWriter) {

	writer.Header().Set("Content-Type", "application/json")
	io.WriteString(writer, `{"success":false}`)

}

func (client *WebSocketClient) webSocketClientReader() {
	defer func() {
		unRegisterClient(client)
		client.conn.Close()
	}()
	client.conn.SetReadLimit(maxMessageSize)
	client.conn.SetReadDeadline(time.Now().Add(pongWait))
	client.conn.SetPongHandler(func(string) error { client.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		//message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		//client.hub.broadcast <- message
		processClientMessage(client, message)
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
		jsonText, _ := json.Marshal(jsonData)
		client.send <- jsonText
	}
}

func processClientMessage(client *WebSocketClient, message []byte) {

	jsonData := make(map[string]interface{})

	json.Unmarshal(message, &jsonData)
	fmt.Println(string(message))
	//fmt.Printf("%+v\n", jsonData)

	if _, ok := jsonData["action"]; !ok {
		log.Println("Invalid Message, no action param")
		return
	}

	switch jsonData["action"] {
	case "updateTrexPan":
		if fmt.Sprintf("%T", jsonData["value"]) != "float64" {
			fmt.Printf("Invalid updateTrexPan value %v\n", jsonData["value"])
			break
		}
		pos := jsonData["value"].(float64)
		setTrexPanPos(pos)
		break
	case "updateTrexTilt":
		if fmt.Sprintf("%T", jsonData["value"]) != "float64" {
			fmt.Printf("Invalid updateTrexTilt value %v\n", jsonData["value"])
			break
		}
		pos := jsonData["value"].(float64)
		setTrexTiltPos(pos)
		break
	case "updateTrexJaw":
		if fmt.Sprintf("%T", jsonData["value"]) != "float64" {
			fmt.Printf("Invalid updateTrexJaw value %v\n", jsonData["value"])
			break
		}
		throttleFloat := jsonData["value"].(float64)
		setTrexJawPos(throttleFloat)
		break
	case "updateThrottle":
		if fmt.Sprintf("%T", jsonData["value"]) != "float64" {
			fmt.Printf("Invalid updateThrottle value %v\n", jsonData["value"])
			break
		}
		throttleFloat := jsonData["value"].(float64)
		setThrottle(throttleFloat)
		break
	case "updateSteering":
		//fmt.Printf("%+v\n", jsonData["value"])
		if fmt.Sprintf("%T", jsonData["value"]) != "float64" {
			fmt.Printf("Invalid updateSteering value %v\n", jsonData["value"])
			break
		}
		posFloat := jsonData["value"].(float64)
		setSteeringPosition(posFloat)
		break
	case "arduinoCommand":
		//fmt.Printf("%v\n", jsonData["command"])
		fmt.Println("arduinoCommand")
		fmt.Printf("%T\n", jsonData["command"])
		fmt.Printf("%v\n", jsonData["command"])
		if fmt.Sprintf("%T", jsonData["command"]) != "string" {
			fmt.Printf("Invalid updateSteering value %v\n", jsonData["value"])
			break
		}
		command := jsonData["command"].(string)

		serialPortMessages <- []byte(command)
		break
	case "playnextsound":
		go playnextsound()
		break
	case "trexscream":
		go trexScreamMp3()
		serialPortMessages <- []byte(arduinoLedBlinkRed)
		break
	case "updatePidConstants":
		//fmt.Printf("%+v\n", jsonData["value"])
		invalidData := false
		if fmt.Sprintf("%T", jsonData["p"]) != "float64" {
			fmt.Printf("Invalid updatePidConstants p value %v\n", jsonData["p"])
			invalidData = true
		}
		if fmt.Sprintf("%T", jsonData["i"]) != "float64" {
			fmt.Printf("Invalid updatePidConstants i value %v\n", jsonData["i"])
			invalidData = true
		}
		if fmt.Sprintf("%T", jsonData["d"]) != "float64" {
			fmt.Printf("Invalid updatePidConstants d value %v\n", jsonData["d"])
			invalidData = true
		}
		if invalidData {
			break
		}
		constP := jsonData["p"].(float64)
		constI := jsonData["i"].(float64)
		constD := jsonData["d"].(float64)
		pidControl.SetPID(constP, constI, constD)
		break
	case "stopThrottle":
		stopThrottle()
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
	case "getPidConstants":
		p, i, d := pidControl.PID()

		jsonData := make(map[string]interface{})
		jsonData["msgType"] = "pidConstants"
		jsonData["p"] = fmt.Sprintf("%.2f", p)
		jsonData["i"] = fmt.Sprintf("%.2f", i)
		jsonData["d"] = fmt.Sprintf("%.2f", d)

		jsonText, _ := json.Marshal(jsonData)
		client.send <- jsonText
		break
	}

}

func stopSteeringMovement() {
	setPwmChanPercent(2, 0)
	setPwmChanPercent(3, 0)
	setPwmChanPercent(4, 0)
}
func setTrexJawPos(pos float64) {
	pos = pos / 1000

	updateTrexJawPulse := float64(jawRange)*pos + float64(jawMin)
	//940 or less = brake

	fmt.Printf("updateTrexJaw pulse %v\n", updateTrexJawPulse)

	//pulseWidth := int(math.Round(1200 * throttlePwmFreqUsCalc))
	//	fmt.Printf("set pulseLengthUs=%v, throttlePwmFreqUsCalc=%v, pulseStart=%v\n", pulseLengthUs, throttlePwmFreqUsCalc, pulseStart)
	if err := pca9685Inst.SetPwm(15, 0, int(throttlePwmFreqUsCalc*updateTrexJawPulse)); err != nil {
		panic(err)
	}
}
func setTrexPanPos(pos float64) {
	pos = pos / 1000

	updatePulse := float64(trexPanRange)*pos + float64(trexPanMin)

	fmt.Printf("updateTrexPan pulse %v\n", updatePulse)

	//pulseWidth := int(math.Round(1200 * throttlePwmFreqUsCalc))
	//	fmt.Printf("set pulseLengthUs=%v, throttlePwmFreqUsCalc=%v, pulseStart=%v\n", pulseLengthUs, throttlePwmFreqUsCalc, pulseStart)
	if err := pca9685Inst.SetPwm(1, 0, int(throttlePwmFreqUsCalc*updatePulse)); err != nil {
		panic(err)
	}
}
func setTrexTiltPos(pos float64) {
	pos = 1 - (pos / 1000)

	updatePulse := float64(trexTiltRange)*pos + float64(trexTiltMin)

	fmt.Printf("updateTrexTilt pulse %v\n", updatePulse)

	//pulseWidth := int(math.Round(1200 * throttlePwmFreqUsCalc))
	//	fmt.Printf("set pulseLengthUs=%v, throttlePwmFreqUsCalc=%v, pulseStart=%v\n", pulseLengthUs, throttlePwmFreqUsCalc, pulseStart)
	if err := pca9685Inst.SetPwm(14, 0, int(throttlePwmFreqUsCalc*updatePulse)); err != nil {
		panic(err)
	}
}
func stopThrottle() {
	setThrottleMicroSeconds(0)
}
func setThrottle(pos float64) {
	//subtract too to start at 800ms
	throttleMax := 1000.0 //only operate within the first 3ms
	//940 or less = brake

	microSecondSetValue := int(pos * throttleMax / 1000)
	//+ (throttlePwmOffset - 200) // add 800ms?
	// + (throttlePwmMax / 2)
	setThrottleMicroSeconds(microSecondSetValue + 1000)
}

func setThrottleFullRange(pulseEnd int) {

	fmt.Printf("set pulseEnd=%v\n", pulseEnd)
	if err := pca9685Inst.SetPwm(throttlePwmChannel, 0, pulseEnd); err != nil {
		panic(err)

	}
}
func setThrottleArm() {
	setThrottleMicroSeconds(1500)
}
func setThrottleCalibration() {

	setThrottleMicroSeconds(2000) // 2000us
	time.Sleep(time.Second * 8)
	setThrottleMicroSeconds(1000) // 1000us
	time.Sleep(time.Second * 8)
	setThrottleMicroSeconds(1500) // 1500us
}
func setThrottleMicroSeconds(pulseLengthUs int) {

	pulseStart := int(math.Round(float64(pulseLengthUs) * throttlePwmFreqUsCalc))
	//pulseWidth := int(math.Round(1200 * throttlePwmFreqUsCalc))
	//	fmt.Printf("set pulseLengthUs=%v, throttlePwmFreqUsCalc=%v, pulseStart=%v\n", pulseLengthUs, throttlePwmFreqUsCalc, pulseStart)
	if err := pca9685Inst.SetPwm(throttlePwmChannel, 0, pulseStart); err != nil {
		panic(err)
	}
}
func setSteeringPosition(pos float64) {

	//fmt.Printf("set pos=%v, ", pos)

	pos = 1 - (pos / 1000)

	//fmt.Printf("set pos%%=%v, ", pos)

	steeringTargetPoint = uint16(float64(steeringRange)*pos) + steeringMin
	//fmt.Printf("set steeringTargetPoint=%v\n", steeringTargetPoint)

	pos = pos * 100
	//fmt.Printf("set pos=%v, ", pos)
	pidSet = pos
	pidControl.Set(pos)
}

var pidControl *pidctrl.PIDController
var pidOutput float64 = 0
var pidError float64 = 0
var pidSet float64 = 0
var pidAdc float64 = 0

func steeringSetPointAdjust() {

	successfulThreshold := .009 //1.8 % total range

	if steeringAdcValue > steeringMax {
		steeringAdcValue = steeringMax
	}
	if steeringAdcValue < steeringMin {
		steeringAdcValue = steeringMin
	}

	pidAdc = (float64(steeringAdcValue-steeringMin) / float64(steeringRange)) * 100
	pidError = pidSet - pidAdc
	pidOutput = pidControl.Update(pidAdc)

	//fmt.Printf("pidSet=%.2f, pidAdc=%.2f, pidError=%.2f, pidOutput=%.2f\n", pidSet, pidAdc, pidError, pidOutput)

	//fmt.Printf("set=%v, actual=%v\n", steeringTargetPoint, steeringAdcValue)

	if float64(steeringAdcValue) < float64(steeringTargetPoint)*float64(1+successfulThreshold) {
		if float64(steeringAdcValue) > float64(steeringTargetPoint)*(1-successfulThreshold) {
			stopSteeringMovement()
			return
		}
	}
	//fmt.Printf("pidOutput=%.2f\n", pidOutput)
	if math.Abs(pidOutput) < 10 {
		stopSteeringMovement()
		return
	}
	if pidOutput > 0 {
		//pidOutput = pidOutput * .5
		//fmt.Printf("pidOutput 50%%=%.2f\n", pidOutput)

		setPwmChanPercent(2, int(pidOutput))
		setPwmChanPercent(3, 100)
		setPwmChanPercent(4, 0)
	} else {
		pidOutput = pidOutput * -1
		//pidOutput = pidOutput * .5
		//fmt.Printf("pidOutput 50%%=%.2f\n", pidOutput)

		setPwmChanPercent(2, int(pidOutput))
		setPwmChanPercent(3, 0)
		setPwmChanPercent(4, 100)
	}
}
func startSteeringControlLoop() {
	fmt.Println("startSteeringControlLoop")
	checkSteeringPositionTicker := time.Tick(time.Millisecond * 20)
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
func serialPortReader(serialPortToRead *serial.Port) {
	fmt.Println("Starting serial port read loop")
	buf := make([]byte, 1)
	minimumMessageSeparator := 30 * time.Millisecond
	lastMessage := time.Now()
	for {
		readMessage, err := serialPortToRead.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		now := time.Now()
		if time.Since(lastMessage) > minimumMessageSeparator {
			fmt.Print("Serial Message:\n")
			lastMessage = now
		}
		fmt.Printf("%s", buf[:readMessage])
	}
}
func serialPortWriter() {

	comConfig := &serial.Config{Name: "/dev/ttyS0", Baud: 115200}
	serialPort, err := serial.OpenPort(comConfig)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Open serial port %v\n", comConfig.Name)

	go serialPortReader(serialPort)
	redEyesTicker := time.Tick(time.Second * 10)
	redEyesTickerCurrent := 1
	//var adcTickNumber uint16 = 0
	for {
		select {
		case message := <-serialPortMessages:
			fmt.Printf("sending serial command: %s\n", message)
			_, err = serialPort.Write([]byte(fmt.Sprintf("%s", message)))
			if err != nil {
				log.Println(err)
			}
			break
		case <-redEyesTicker:

			//serialPortMessages <- []byte(fmt.Sprintf("%v", redEyesTickerCurrent))

			//fmt.Printf("writing to serial port: %v\n", redEyesTickerCurrent)
			// _, err = serialPort.Write([]byte())
			// if err != nil {
			// 	log.Println(err)
			// }
			redEyesTickerCurrent++
			if redEyesTickerCurrent > 8 {
				redEyesTickerCurrent = 1
			}
			break
		}
	}
}
func adcTicker(bus embd.I2CBus) {

	initAdc(bus)

	steeringAdcValue = readAdcValue(bus)

	adcValueBroadcastTicker := time.Tick(time.Millisecond * 50)
	adcReadTicker := time.Tick(time.Millisecond * 16)
	//var adcTickNumber uint16 = 0
	for {
		select {
		case <-adcValueBroadcastTicker:
			jsonData := make(map[string]interface{})
			jsonData["msgType"] = "status"
			jsonData["steeringCurrent"] = fmt.Sprintf("%v", steeringAdcValue)
			jsonData["steeringTargetPoint"] = fmt.Sprintf("%v", steeringTargetPoint)

			//jsonData["pidOutput"] = fmt.Sprintf("%v.2", pidOutput)
			jsonData["pidError"] = fmt.Sprintf("%.2v", pidError)
			//jsonData["pidSet"] = fmt.Sprintf("%v", pidSet)
			//jsonData["pidAdc"] = fmt.Sprintf("%.2v", pidAdc)

			webSocketSendJsonToAllClients(jsonData)
		case <-adcReadTicker:
			//adcTickNumber++

			//start := time.Now()

			steeringAdcValue = readAdcValue(bus)

			//elapsed := time.Since(start)

			//fmt.Printf("%4.d Read ADC Value %d\n", adcTickNumber, steeringAdcValue)

			//fmt.Printf(" time: %s\n", elapsed)
		}
	}
}

var mp3filesCurrent int = 0

var mp3files = [8]string{
	"Welcome... to Jurassic park.mp3",
	//"Theme song.mp3",
	"trex hold on to your butts.mp3",
	"toast.mp3",
	"trex you didnt say the magic word.mp3",
	"trex clever girl.mp3",
	"TRex growls.mp3",
	"Malcom Laugh.mp3",
	"meat.mp3",
	//"God creates dinosaurs.mp3"
}

var mp3filesLen = len(mp3files) - 1

func playnextsound() {
	playMp3(mp3files[mp3filesCurrent])
	fmt.Printf("playing %v\n", mp3files[mp3filesCurrent])
	mp3filesCurrent++
	if mp3filesCurrent > mp3filesLen {
		mp3filesCurrent = 0
	}
}
func trexScreamMp3() {
	playMp3("TRex screams.mp3")
	time.Sleep(150 * time.Millisecond)
	setTrexJawPos(500)
	time.Sleep(40 * time.Millisecond)
	setTrexJawPos(600)
	time.Sleep(120 * time.Millisecond)
	setTrexJawPos(800)
	time.Sleep(180 * time.Millisecond)
	setTrexJawPos(400)
	time.Sleep(120 * time.Millisecond)
	setTrexJawPos(700)
	time.Sleep(180 * time.Millisecond)
	setTrexJawPos(800)
	time.Sleep(120 * time.Millisecond)
	setTrexJawPos(700)
	time.Sleep(180 * time.Millisecond)
	setTrexJawPos(600)
	time.Sleep(140 * time.Millisecond)
	setTrexJawPos(500)
	time.Sleep(100 * time.Millisecond)
	setTrexJawPos(800)
	time.Sleep(120 * time.Millisecond)
	setTrexJawPos(500)
	time.Sleep(100 * time.Millisecond)
	setTrexJawPos(600)
	time.Sleep(100 * time.Millisecond)
	setTrexJawPos(800)
	time.Sleep(120 * time.Millisecond)
	setTrexJawPos(500)
	time.Sleep(100 * time.Millisecond)
	setTrexJawPos(20)
}
func playMp3(fileName string) {
	//"God creates dinosaurs.mp3"
	//"Theme song.mp3"
	//"TRex growls.mp3"
	//"TRex screams.mp3"
	//"Welcome... to Jurassic park.mp3"
	//"Malcom Laugh.mp3"

	cmd := exec.Command("mpg123", "/home/pi/trex-sounds/"+fileName)
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
}
func setVolume() {
	cmd := exec.Command("amixer", "sset", "'PCM'", "200%")
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
}
func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
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
	pca9685Inst.Freq = throttlePwmFreq
	pca9685Inst.Wake()
	defer pca9685Inst.Close()

	// fmt.Println("setting 1ms on port 15")
	// if err := pca9685Inst.SetPwm(15, 0, throttlePwmFreq1500us); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("setting 1ms on port 14")
	// if err := pca9685Inst.SetPwm(14, 0, throttlePwmFreq1500us); err != nil {
	// 	panic(err)
	// }
	fmt.Println("setting 1ms on port 1")
	if err := pca9685Inst.SetPwm(1, 0, int(throttlePwmFreqUsCalc*float64(jawMin))); err != nil {
		panic(err)
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	srv := &http.Server{
		Addr: ":8090",
		//Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fs := http.FileServer(http.Dir("/home/pi/rc-steering-test-pid-gamepad/webroot"))
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

	pidControl = pidctrl.NewPIDController(1.5, .03, .03)
	pidControl.SetOutputLimits(-100, 100)
	pidSet = 50
	pidControl.Set(pidSet)

	//setThrottleCalibration()
	setThrottleArm()

	fmt.Printf("Setting intial steering position to %v\n", steeringTargetPoint)

	go adcTicker(i2cBus)
	go startSteeringControlLoop()

	setVolume()
	playMp3("Theme song.mp3")

	go serialPortWriter()

	ip := getOutboundIP()
	fmt.Printf("outbound ip: %v\n", ip)
	serialPortMessages <- []byte(fmt.Sprintf("t1%v", ip))

	//block waiting for channel
	<-shutdown

	stopSteeringMovement()
	stopThrottle()

	setTrexPanPos(500)
	setTrexTiltPos(500)
	setTrexJawPos(50)

	log.Println("Server is shutting down")
	os.Exit(0)

}
