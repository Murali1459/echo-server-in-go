package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
    "strings"
)

const(
    HOST = "localhost"
    PORT = "8080"
    TYPE = "tcp"
)


func main(){
    listen,err := net.Listen(TYPE,HOST+":"+PORT)
    if err!= nil{
        log.Print(err)
        return
    }
    debug_msg("listening on " + HOST + ":" + PORT + "\n")
    for{
        conn,err := listen.Accept()
        if err != nil{
            return
        }
        go handlerequest(conn) 
    }

}
func debug_msg(msg string){
    fmt.Print("[DEBUG]: ",msg)
}
func recv_data(con net.Conn)(string){
    message,_ := bufio.NewReader(con).ReadString('\n')
    return string(message)
}
func handlerequest(con net.Conn){
    remoteip := con.RemoteAddr()
    debug_msg("Incoming "  + remoteip.Network() + " Connection from " + remoteip.String() + "\n" )
    for {
        clientmessage := recv_data(con)
        debug_msg("[RECEIVED DATA] -> " + clientmessage)
        sendmessage := []byte("We Received your message: ") 
        sendmessage = append(sendmessage,clientmessage...)
        clientmessage = strings.ToUpper(clientmessage)
        _,err := con.Write([]byte(sendmessage))
        if err != nil{
            log.Fatal(err)
            con.Close()
        }
        if clientmessage == "EXIT\n"{
            break
        } 
    }
    debug_msg("Closing "  + remoteip.Network() + " Connection " + remoteip.String() + "\n" )
    con.Write([]byte("Closing Connection\n"))
    con.Close()
}
