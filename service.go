package main

import (
  "bitbucket.org/kardianos/service"
  "fmt"
  "os"
)

var srvLog service.Logger

func main() {
  var name = "GoSensuKeepalive"
  var displayName = "Sensu keepalive client in Go"
  var desc = "Sends keepalive messages to Sensu server."

  var s, err = service.NewService(name, displayName, desc)
  srvLog = s

  if err != nil {
    fmt.Printf("%s unable to start: %s", displayName, err)
    return
  }

  if len(os.Args) > 1 {
    var err error
    verb := os.Args[1]
    switch verb {
    case "install":
      err = s.Install()
      if err != nil {
        fmt.Printf("Failed to install: %s\n", err)
        return
      }
      fmt.Printf("Service \"%s\" installed.\n", displayName)
    case "remove":
      err = s.Remove()
      if err != nil {
        fmt.Printf("Failed to remove: %s\n", err)
        return
      }
      fmt.Printf("Service \"%s\" removed.\n", displayName)
    case "run":
      doWork()
    case "start":
      err = s.Start()
      if err != nil {
        fmt.Printf("Failed to start: %s\n", err)
        return
      }
      fmt.Printf("Service \"%s\" started.\n", displayName)
    case "stop":
      err = s.Stop()
      if err != nil {
        fmt.Printf("Failed to stop: %s\n", err)
        return
      }
      fmt.Printf("Service \"%s\" stopped.\n", displayName)
    }
    return
  }
  err = s.Run(func() error {
    // start
    go doWork()
    return nil
  }, func() error {
    // stop
    stopWork()
    return nil
  })
  if err != nil {
    s.Error(err.Error())
  }
}

func doWork() {
  srvLog.Info("Service started")
  sensu()
}
func stopWork() {
  srvLog.Info("Service stopped")
}
