package main

import (
	"embed"
	"log"
	"net/http"
	"text/template"

	"golang.org/x/net/websocket"
)

//go:embed index.html
var indexTmpl embed.FS

var conns []*websocket.Conn

func main() {
	http.HandleFunc("/", index)
	http.Handle("/ws", websocket.Handler(msgHandler))

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(indexTmpl, "index.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, "")
	if err != nil {
		log.Fatal(err)
	}
}

func msgHandler(ws *websocket.Conn) {
	conns = append(conns, ws)
	msgReceiver(ws)
}

func msgReceiver(ws *websocket.Conn) {
	for {
		// メッセージを受信する
		msg := ""
		var err = websocket.Message.Receive(ws, &msg)
		if err != nil {
			log.Fatalln(err)
		}

		for _, conn := range conns {
			dist := conn
			go func() {
				// メッセージを返信する
				err = websocket.Message.Send(dist, msg)
				if err != nil {
					log.Fatalln(err)
				}
			}()
		}
	}
}
