package main

import (
	"fmt"
	"log"
	"net"
	"os"
)


func main() {
  conn, err := net.Dial("tcp", "localhost:9999")
  
  if err != nil {
    log.Println(err)
    return
  }

  //for {
  
       // pid := os.Getppid()
       // pidStr := strconv.Itoa(pid)
  
       // conn.Write([]byte(pidStr))
       // write a sleep for 10 seconds

        file, err := os.Open("test.jpg")
        if err != nil {
          log.Fatal(err)
        }

        fi,_ := file.Stat()
        defer file.Close()

        // get the file size
        size := fi.Size()
        println("File size:", size)


        data := make([]byte, size) 
        

        count, err := file.Read(data)
        

        if err != nil {
          log.Fatal(err)
        }
        //divide the file into 600 parts
       // for i := 0; i < 600; i++ {
        //  conn.Write(data[i*count/600:(i+1)*count/600])
        //  time.Sleep(100 * time.Millisecond)
        //}




        n,err := conn.Write(data[:count])

        if err != nil {
          log.Println(err)
          return
        }
        fmt.Println(n)

}

