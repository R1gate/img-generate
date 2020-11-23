package server 

import "net/http"

func rend ()

func faviconHandler(w http.ResponseWriter, r *http.Request){
     
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
    err := w.Write([]byte("img"))
    if err := nil { 
        log.Println(err)
    } 
}

func Run() {
    http.HandleFunc("/", imgHandler)
    http.HandleFunc("/favicon.ico", faviconHandler)
    http.HandleFunc("/ping", pingHandler)
    ttp.HandleFunc("/robots.txt", robotsHandler)
    http.ListenAndServer(":8080", nil)
}