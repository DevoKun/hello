package main

import (
  "fmt"
  "github.com/fatih/color"
  "io"
  "net/http"
  "time"
  "os"

) // import

func helloAnswer(w http.ResponseWriter, r *http.Request) {
  
  fmt.Println(color.YellowString(time.Now().String()),
              color.CyanString("Answering HTTP"),
              color.RedString(r.Method),
              color.CyanString("Request"),
              color.WhiteString(" ("),
              color.GreenString(r.URL.Path),
              color.WhiteString(")") )

  timeNow := time.Now()
  
  w.Header().Set("Date",         timeNow.Format("Mon, 2 Jan 2006 15:04:05 MST"))
  w.Header().Set("Server",       "hello")
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  io.WriteString(w, `{"hello":"world"}` + "\n")

} // func helloAnswer

func main() {
  httpPort     := "8765"
  httpPortTemp := os.Getenv("HTTP_PORT")

  if ( ("x"+httpPortTemp) != "x" ) {
    httpPort = httpPortTemp
  } // if

  fmt.Println(color.YellowString(time.Now().String()), 
              color.GreenString("Starting up http listener"),
              color.WhiteString("... :"),
              color.CyanString(httpPort) )

  http.HandleFunc("/", helloAnswer)

  http.ListenAndServe(":"+httpPort, nil)

} // func main

