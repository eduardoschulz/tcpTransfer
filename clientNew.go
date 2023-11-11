package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

//create struct for server
type Server struct {
  conn net.Conn
  port string
  ip string
}

type File struct {
  name string
  size int64
  data []byte
}

func createServer(Port string, Ip string) *Server { /* Cria um construtor para o Server */
  
  p := &Server{

    port:  Port,
    ip: Ip,

  }
    
  return p
  
}

func createFile(Name string) (*File, error){

  file, err := os.Open(Name)

  if err != nil {
    log.Println(err)
    return nil, err
  }
  defer file.Close()
  


  fi,_ := file.Stat()

  println(fi.Size())

  data := make([]byte, fi.Size())

  count,_ := file.Read(data)
  println(count)

  p := &File{

    

  }

  return p, nil
}



func handleConnection() net.Conn{

  // Connect to server
  conn, err := net.Dial("tcp", ":9999")
  if err != nil {
    fmt.Println("Error connecting to server")
    os.Exit(1)
  }
  defer conn.Close()


  return conn

}


//create a function that divides a byte array into 4 parts
func divideArray(data []byte) [][]byte {
  
    var dividedData [][]byte
  
    var i int
  
    for i = 0; i < len(data); i += 4 {
  
      dividedData = append(dividedData, data[i:i+4])
  
    }
  
    return dividedData
  
  }




func main() {
  f,_  := createFile("test.e")

  fmt.Println(f.name)
}






