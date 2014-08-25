package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Init(infoHandle, warningHandle, errorHandle io.Writer) {
	Info = log.New(infoHandle, "INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle, "WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle, "ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", r.URL.Path[1:])
}

func main() {
	Init(os.Stdout, os.Stdout, os.Stdout)
	Info.Println("WebServer initialized")

	http.HandleFunc("/", handler)
	Info.Println("Starting webserver talkback on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		Info.Println("WebServer talkback finished on port 8080")
	}
}
