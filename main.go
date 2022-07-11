package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/", fs)
	http.HandleFunc("/call", handler)

	fmt.Println("listening...")
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

const (
	content = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Gather>
        <Say>Hello Monkey, please press 1</Say>
    </Gather>
</Response>`
	echoMsg = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
	<Say>Hello, World!</Say>
    <Play>https://api.twilio.com/Cowbell.mp3</Play>
</Response>`
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/xml")

	rec := r.FormValue("Digits")

	if rec == "" {
		fmt.Fprint(w, content)
		return
	}

	fmt.Fprint(w, echoMsg)
}

func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
