package main

import (
  "fmt"
  "os"
  "time"
  "net"
  "strings"
)

var LINE_MAX_LENGTH = 60
var STARTING_PORT = 50
var ENDING_PORT = 85

func printDots(){
  fmt.Println()
  fmt.Println()
  for i:=0 ; i <LINE_MAX_LENGTH;i++{
    fmt.Print("-")
  }
  fmt.Println()
  fmt.Println()
}

func printTarget(target string){
  fmt.Print("Target IP : ",target,"\n")
  fmt.Print("Time started : ",time.Now(),"\n\n")
}

func printDetails(){
  // Print out some details
  printDots()
  targetIP := os.Args[1]
  printTarget(targetIP)
}

func main() {

  // Retrieve the arguments
  argsWithoutProg := os.Args[1:]

  // Exit if no ip is given 
  if len(argsWithoutProg) == 0 {
    fmt.Println("Syntax error: go run main.go <ip>")
    os.Exit(-1)
  }

  // Print out some details
  printDetails()

  targetIP := os.Args[1]


  // Open a connection
  for p:=STARTING_PORT; p < ENDING_PORT;p++{
    fmt.Println("Checking port : ",p)
    addr := fmt.Sprintf("%s:%d",string(targetIP),p)

    // Check if timeout then the port is closed
    conn, err := net.DialTimeout("tcp",addr,time.Duration(10)*time.Second)
    if err != nil && strings.Contains(err.Error(),"refused") {
      fmt.Println("Port ",p," is close.")
    } else {
      fmt.Println("Port ",p," is open.")
      conn.Close()
    }
  }

 printDots()
}
