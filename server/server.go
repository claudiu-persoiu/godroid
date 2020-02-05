package server

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/claudiu-persoiu/godroid/gpio"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

var motorLeft gpio.Motor
var motorRight gpio.Motor

type message struct {
	Action string `json:"action"`
	Data   string `json:"data,omitempty"`
}

// StartServer start https communcation
func StartServer(address string, mLeft gpio.Motor, mRight gpio.Motor) error {
	motorLeft = mLeft
	motorRight = mRight
	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.Handle("/ws", http.HandlerFunc(webSocket))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("public/index.html"))
		err := t.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	return http.ListenAndServe(address, router)
}

func webSocket(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	go waitForMessage(c)
}

func waitForMessage(c *websocket.Conn) {

	defer c.Close()

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("recv: %s", msg)

		var obj message
		if err := json.Unmarshal(msg, &obj); err == nil {
			processMessage(obj)
		} else {
			log.Println("Error parcing message:")
			log.Println(err)
		}
	}
}

func processMessage(msg message) error {

	switch msg.Action {
	case "left":
		switch msg.Data {
		case "up":
			motorLeft.Forward(1)
		case "down":
			motorLeft.Backword(1)
		case "stop":
			motorLeft.Stop()
		}

	case "right":
		switch msg.Data {
		case "up":
			motorRight.Forward(1)
		case "down":
			motorRight.Backword(1)
		case "stop":
			motorRight.Stop()
		}
	}

	return nil
}
