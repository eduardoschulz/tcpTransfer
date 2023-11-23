package main

import (
  "log"
  "net"
  "fmt"
  "os"
  "io"
  "strings"
)

type Server struct {
  clients []Client
}

type Client struct {
  id string
  conn net.Conn
}

func (s *Server) run() {
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
 

  file,err := os.CreateTemp("./temp", filename)
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

    lenv := len(buffer[:n])

    fmt.Printf("Size of Buffer: %d\n", lenv)
    

    if strings.Contains(string(buffer[:n]), "\n") {
      break
    }

    _, err = file.Write(buffer[:n])
    if err != nil {
      log.Fatal("Error writing to the file:", err)
    }
  }

  fmt.Printf("Finished writing file...\n")

  err = file.Close()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("File closed...\n")
  
  buffer = make([]byte, 1024)
  n, err = conn.Read(buffer)  
  if err != nil {
    log.Fatal(err)
  }

  if string(buffer[:n]) == "is b ready?" {
    conn.Write([]byte("yes"))
  }

}





func main() {
  server := Server{}
  server.run()
}

