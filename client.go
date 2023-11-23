package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

//os args

func main() {

  args := os.Args[1]

  if len(args) < 1 {
    log.Fatal("Usage: go run client.go <client name>")
  }

  conn, err := net.Dial("tcp", "localhost:9999") /*connects socket to server */
  defer conn.Close() /*closes connection when function returns*/

  if err != nil {
    log.Println(err)
    return
  }


  file, err := os.Open("test.jpg") /*file to be sent*/
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

  var s string /* parses the client name from the command line*/
  if "a" == args {
    s = "client a"
    fmt.Println("Client A")
  }else {
    s = "client b"
    fmt.Println("Client B")
  }

  _,err = conn.Write([]byte(s)) /*sends the client name to the server*/

  time.Sleep(100 * time.Millisecond)

  fmt.Println("Sending file...")
  _,err = conn.Write(data[:]) /*sends the file to the server*/
  fmt.Println("File sent!")
  if err != nil {
    log.Fatal(err)
  }
  conn.Write([]byte("\n")) /*sends a newline character to the server*/


  otherfile := make([]byte, 1024)

  var fileC *os.File

  if s == "client a" {
    fileC, err = os.Create("recvB.jpg")
    conn.Write([]byte("is b ready?"))
    if err != nil {
      log.Fatal(err)
    }else{
    fileC, err = os.Create("recvA.jpg")
    conn.Write([]byte("is a ready?"))
    if err != nil {
      log.Fatal(err)
    }
  }

 buf := make([]byte, 1024)

 _,err := conn.Read(buf[:])

 if err != nil {
  log.Fatal(err)
 }

 fmt.Println(string(buf))

  

  defer fileC.Close()





  for {
    _, err = conn.Read(otherfile[:])
    if err != nil {
      if err != io.EOF {
        log.Println("Error reading from connection: ", err)
      }
      break
    }

    _, err = file.Write(otherfile[:])
    if err != nil {
      log.Fatal("Error writing to the file:", err)
    }


  }






}
}

