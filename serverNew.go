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

  listener, err := net.Listen("tcp", ":9999") /* cria um socket TCP na porta 9999 */

  if err != nil {
    log.Println(err)
    return
  }

  fmt.Println("Escutando na porta 9999")

    conn, err := listener.Accept()
    if err != nil {
      log.Println(err)
    }


    var buf [1024]byte /* cria buffer que vai armazenar a mensagem recebida pelo socket*/
    
    n,err := conn.Read(buf[0:])

    switch (string(buf[0:n])) {
      case "cliente A":
        fmt.Println("Cliente A conectado")
        handleClientc(conn, "A") /*TODO*/

      case "cliente B":
        fmt.Println("Cliente B conectado")
        handleClientc(conn, "B") /*TODO*/

      case "END":
          fmt.Println("Encerrando conexão")
        conn.Close()
        log.Println("Conexão encerrada")

      default:
        fmt.Println("Cliente desconhecido")
        conn.Close()
        log.Println("Conexão encerrada")
    }

  }

  func handleClientc(conn net.Conn, nome string){

    fmt.Println("Cliente conectado de " + conn.RemoteAddr().String())

    var totalSize [1024]byte /* cria buffer que vai armazenar o tamanho do arquivo a ser recebido */
    
    _,err := conn.Read(totalSize[0:]) /* lê a mensagem recebida pelo socket */
    fmt.Println("Tamanho do arquivo: " + string(totalSize[0:]))
    if err != nil {
      log.Println(err)
      return
    }


    

    
    f,err := os.Create("file"+ nome + ".jpg") /* cria um arquivo para armazenar a mensagem recebida */
    if err != nil {
      log.Println(err)
      return
    }



    //transform the byte array into a int
    
    var buf []byte /* cria buffer que vai armazenar a mensagem recebida pelo socket */
    var buffTemp []byte
    for string(buffTemp) != "END" {
        fmt.Println("Recebendo arquivo...")
      _,err := conn.Read(buffTemp[0:]) /* lê a mensagem recebida pelo socket */
      buf = append(buf, buffTemp...)
       

      if err != nil {
        log.Println(err)
      }
    }
    
    f.Write(buf)/* escreve o stream de dados recebida no arquivo */

  }
