package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var clientA bool = false
var clienteB bool = false



type Server struct {
}

type Client struct {
  id string
  conn net.Conn
}

func (s *Server) run() {
  //clientA = false
  //clientB = false

  listener, err := net.Listen("tcp", ":9999")
  log.Print("Listening on port 9999")
  if err != nil {
    log.Fatal(err)
  }

  for {

    conn, err := listener.Accept()
    if err != nil {
      log.Fatal(err)
    }

    go s.handleConnection(conn)
  }
}


func (s *Server) handleConnection(conn net.Conn){
  //defer conn.Close()

  clientID := make([]byte, 1024) /*stores a id for the client*/
  n,err := conn.Read(clientID)

  if err != nil {
    log.Fatal(err)
  }

  log.Printf("Client %s connected...", clientID[:n])
  filename := string(clientID[:n]) + "recv.jpg" /*defines the name of the file to be saved, hardcoded to jpg TODO*/ 
  fmt.Printf("Filename: %s\n", filename)
 

  file,err := os.Create(filename) /*creates the file to be saved*/
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  buffer := make([]byte, 1024)
  for {
    n, err = conn.Read(buffer)
    
    if err != nil {
      if err != io.EOF {
        log.Println("Error reading from connection: ", err)
      }
    }

    lenv := len(buffer[:n]) /*gets the length of the buffer*/


    if lenv < 1024 {
        break
    }

    _, err = file.Write(buffer[:n])
    if err != nil {
      log.Fatal("Error writing to the file:", err)
    }
  }


  if strings.Contains(string(clientID[:n]),"a"){
    clientA = true
  }else if strings.Contains(string(clientID[:n]),"b"){
    clienteB = true
  }

 

  fmt.Printf("Finished writing file...\n")
  fmt.Printf("File saved...\n")
  err = file.Close()
  if err != nil {
    log.Fatal(err)
  }

  time.Sleep(100 * time.Millisecond)
  for { 
    //fmt.Printf("Waiting for other client...\n")
    buff := make([]byte, 1024) 
    n, err = conn.Read(buff)
    fmt.Printf("%s: %s\n", string(clientID[:n]), string(buff[:n]))

    if strings.Contains(string(clientID[:n]),"a") && clienteB == true{
      time.Sleep(100 * time.Millisecond)
      log.Printf("entered condition client a and b ready")

      conn.Write([]byte("yes"))
      sendFile(conn, "client brecv.jpg")
      break

    }else if strings.Contains(string(clientID[:n]),"b") && clientA == true{
      time.Sleep(100 * time.Millisecond)
      conn.Write([]byte("yes"))

      log.Printf("entered condition client a and b ready")
     // fmt.Printf("is a ready... yes\n")
      sendFile(conn, "client arecv.jpg")
      break
    }

    time.Sleep(100 * time.Millisecond)
    conn.Write([]byte("no"))
    fmt.Printf("%b %b\n", clientA, clienteB)
  }

  time.Sleep(2 * time.Second)
  if strings.Contains(string(clientID[:n]),"a"){
    clientA = false
  }else if strings.Contains(string(clientID[:n]),"b"){
    clienteB = false
  }
  conn.Close()

}

func sendFile(conn net.Conn, filename string) {

  file, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  }

  fileInfo, err := file.Stat()
  if err != nil {
    log.Fatal(err)
  }

  buffer := make([]byte, fileInfo.Size())
  _, err = file.Read(buffer)
  conn.Write(buffer)
  time.Sleep(100 * time.Millisecond)
  buffer = make([]byte, 1)
  conn.Write(buffer)

  file.Close()

}


func main() {
  server := Server{
  }

  server.run()
}

