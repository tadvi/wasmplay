package main

import (
	"github.com/gorilla/websocket"
)

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
)

var (
	listen = flag.String("listen", ":8080", "listen address")
	dir    = flag.String("dir", "static/", "directory to serve")
)

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)

	http.HandleFunc("/ajax", ajaxHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/websocket", websocketHandler)

	http.Handle("/", http.FileServer(http.Dir(*dir)))
	log.Fatal(http.ListenAndServe(*listen, nil))
}

// ajaxHandler dumps entire request into stdout.
func ajaxHandler(w http.ResponseWriter, req *http.Request) {
	b, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
	fmt.Fprintf(w, "ok")
}

// formHandler dumps form values into stdout.
func formHandler(w http.ResponseWriter, req *http.Request) {
	err := req.ParseMultipartForm(1 << 20)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(req.Form)
	fmt.Fprintf(w, "ok")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// websocketHandler echo back limited number of websocket messages and then close connection.
// Limit is accepted via 'limit' parameter.
func websocketHandler(w http.ResponseWriter, req *http.Request) {

	limit := 1 << 30
	s := req.FormValue("limit")
	if len(s) > 0 {
		var err error
		limit, err = strconv.Atoi(s)
		if err != nil {
			log.Println(err)
			return
		}
	}

	// Perform Websocket upgrade.
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("Upgrade Error: ", err)
		return
	}
	defer conn.Close()

	log.Println("WebSocket connection initiated.")

	for i := 0; i < limit; i++ {
		msgType, bytes, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read Error: ", err)
			break
		}

		// We don't recognize any message that is not text.
		if msg := string(bytes); msgType != websocket.TextMessage {
			log.Println("Non-text message received, ignoring.")
			continue
		} else {
			log.Println("Received:", msg)
		}

		log.Println("Sending: same message as echo back.")
		err = conn.WriteMessage(websocket.TextMessage, bytes)
		if err != nil {
			log.Println("Write Error: ", err)
			break
		}
	}
	log.Println("WebSocket connection terminated.")
}
