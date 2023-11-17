package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

/*Cria struct Server */
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

  return conn

}


/*
func divideArray(data []byte) [][]byte {
  
    var dividedData [][]byte
  
    var i int
  
    for i = 0; i < len(data); i += 4 {
  
      dividedData = append(dividedData, data[i:i+4])
  
    }
  
    return dividedData
  
  }

*/



func main() {

  file,_ := os.ReadFile("fileA.jpg")
  



  //f,_  := createFile("fileA.jpg")


  conn := handleConnection()
  conn.Write([]byte("cliente A"))
  

  conn.Write([]byte("2000000"))
  n,err := conn.Write(file)


  for i := 0; i < 600; i++ {
    conn.Write(file[i*n/600:(i+1)*n/600])
  }


  if err != nil {
    log.Println(err)
    return
  }

  conn.Write([]byte("END"))

  fmt.Println(n)
  

  defer conn.Close()

}






