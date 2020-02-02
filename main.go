package main

import (
  "fmt"
  "os"
  "time"
  "net"
	"strings"
  "./config"
)

func printDetails(){
  printDots()
  targetIP := os.Args[1]
  fmt.Print("Target IP : ",targetIP,"\n")
  fmt.Print("Time started : ",time.Now(),"\n\n")
}

func printDots(){
  fmt.Println()
  fmt.Println()
  for i:=0 ; i <config.LINE_MAX_LENGTH;i++{
    fmt.Print("-")
  }
  fmt.Println()
  fmt.Println()
}

func main() {

  // Retrieve the arguments
  argsWithoutProg := os.Args[1:]

  // Open ports
  var o_ports [] int

  // Exit if no ip is given 
  if len(argsWithoutProg) == 0 {
    fmt.Println("Syntax error: go run main.go <ip>")
    os.Exit(-1)
  }

  // Print out some details
  printDetails()

  targetIP := os.Args[1]

  // Open a connection
  for p:=config.STARTING_PORT; p < config.ENDING_PORT;p++{
    addr := fmt.Sprintf("%s:%d",string(targetIP),p)
    // Check if timeout then the port is closed
    conn, err := net.DialTimeout("tcp",addr,time.Duration(10)*time.Second)
    if err != nil && strings.Contains(err.Error(),"no route to host") { 
      fmt.Println("Hostname could not be resolved")
      printDots()
      os.Exit(1)
    } else if err == nil {
      o_ports = append(o_ports,p)
      conn.Close()
    } 
  }

  fmt.Print("Open Ports: \n\n")
  for i:=0 ;i < len(o_ports);i++{
    conc := fmt.Sprintf("%d/%s",o_ports[i],config.SERVICES[o_ports[i]])
    fmt.Println(conc)
  }
  printDots()
}
