package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
  "strings"
)

//os args

func main() {

  args := os.Args[1]

  if len(args) < 1 {
    log.Fatal("Usage: go run client.go <client name>")
  }

  conn, err := net.Dial("tcp", "localhost:9999") /*connects socket to server */

  if err != nil {
    log.Println(err)
    return
  }

  var file *os.File
  var s string /* parses the client name from the command line*/
  if "a" == args {
    s = "client a"
    fmt.Println("Client A")
    file, err = os.Open("sendA.png") 
  }else {
    s = "client b"
    fmt.Println("Client B")
    file, err = os.Open("sendB.png") 
  }



  /*file to be sent*/
  if err != nil {
    log.Fatal(err)
  }

  fi,_ := file.Stat()
  defer file.Close()

  size := fi.Size()
  fmt.Printf("File size: %d\n", size)


  data := make([]byte, size) 

  _,err = file.Read(data)


  if err != nil {
    log.Fatal(err)
  }

    _,err = conn.Write([]byte(s)) /*sends the client name to the server*/

  time.Sleep(100 * time.Millisecond)

  fmt.Println("Sending file...")
  _,err = conn.Write(data[:]) /*sends the file to the server*/
  fmt.Println("File sent!")
  if err != nil {
    log.Fatal(err)
  }
  time.Sleep(100 * time.Millisecond)

  conn.Write([]byte("\n")) /*sends a newline character to the server*/
  conn.Write(make([]byte, 0)) /*sends a empty byte array to the server*/

  var fileC *os.File

  fmt.Println("Waiting for response...")
  clientReady := false
  for !clientReady {

    if s == "client a" {
      fileC, err = os.Create("recvB.png")
      conn.Write([]byte("is b ready?"))
      fmt.Printf("Is B ready?\n")
      if err != nil {
        log.Fatal(err)
      }
    }else if s == "client b"{
      fileC, err = os.Create("recvA.png")
      conn.Write([]byte("is a ready?"))
      fmt.Printf("Is A ready?\n")
      if err != nil {
        log.Fatal(err)
      }
    }
    
    time.Sleep(100 * time.Millisecond)
    tmpBuf := make([]byte, 1024)
    n, err := conn.Read(tmpBuf)
    //fmt.Printf("Read %s bytes\n", string(tmpBuf[:n]))
    if err != nil {
      if err == io.EOF {
        fmt.Println("EOF")
        return
      }
    }
    if strings.Contains(string(tmpBuf[:n]), "yes") {
      clientReady = true
      break // test
    }else if strings.Contains(string(tmpBuf[:n]), "no") {
      time.Sleep(2 * time.Second)
    }


    time.Sleep(200 * time.Millisecond)
    fmt.Printf("Waiting for client...\n")
  }
  
  finished := false

  tmpBuf := make([]byte, 1024)
  for !finished {
    n, err := conn.Read(tmpBuf)
    
    if err != nil {
      if err != io.EOF {
        log.Println("Error reading from connection: ", err)
      }
    }

    lenv := len(tmpBuf[:n])

    //fmt.Printf("Size of Buffer: %d\n", lenv)

    if lenv < 1024 {
      finished = true
    }

    _, err = fileC.Write(tmpBuf[:n])
    if err != nil {
      log.Fatal("Error writing to the file:", err)
    }
  }
  log.Print("File received!")


  conn.Close()

}

