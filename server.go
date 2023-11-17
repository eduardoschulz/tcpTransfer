package main

import (
	"fmt"
	"log"
	"net"
	"os"
)


func main() {
  handleConnection()

}


func handleConnection() {

  listener, err := net.Listen("tcp", ":9999")

  if err != nil {
    log.Println(err)
    return
  }

  defer listener.Close()
  fmt.Println("Listening on port 9999")

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Println(err)
      continue
    }
    
    go handleClient(conn)



}
}

func handleClient(conn net.Conn){
  //defer conn.Close()
  fmt.Println("Client connected from " + conn.RemoteAddr().String())
 // for {
   /* 
    var buf [2701662]byte cria buffer que vai armazenar a mensagem recebida pelo socket 
    n, err := conn.Read(buf[0:])

    if err != nil {    verifica se ocorreu algum erro na leitura da mensagem 
      log.Println(err) 
      return
    }
    s := string(buf[0:n])  converte o buffer em string 
    fmt.Println("Received:", s)
    _, err2 := conn.Write([]byte("Received: "+ conn.LocalAddr().String()))
    if err2 != nil {
      log.Println(err2)
      return
    }
    */
    f, err := os.Create("received.jpg") /* cria um arquivo para armazenar a mensagem recebida */
    if err != nil { /* verifica se ocorreu algum erro na criação do arquivo */
      log.Fatal(err) 
    }
    var buf []byte

    go conn.Read(buf[0:]) /* lê a mensagem recebida pelo socket */

    f.Write(buf)/* escreve o stream de dados recebida no arquivo */
    f.Close() /* fecha o arquivo */
 // }

}
