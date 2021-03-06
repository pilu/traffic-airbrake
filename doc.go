/*
Package airbrake implements a Traffic Middleware for Airbrake.

This is a Middleware for Traffic (https://github.com/pilu/traffic).
It is base on @tobi's Airbrake library (https://github.com/tobi/airbrake-go).

Example:

  package main

  import (
    "os"
    "fmt"
    "time"
    "net/http"
    "github.com/pilu/traffic"
    "github.com/pilu/traffic-airbrake"
  )

  func rootHandler(w traffic.ResponseWriter, r *traffic.Request) {
    err := fmt.Sprintf("Error at %v", time.Now())
    panic(err)
  }

  func main() {
    traffic.SetVar("env", "production")
    router := traffic.New()
    router.Use(airbrake.New(os.Getenv("AIRBRAKE_API_KEY")))

    // Routes
    router.Get("/", rootHandler)
    router.Run()
  }
*/
package airbrake
