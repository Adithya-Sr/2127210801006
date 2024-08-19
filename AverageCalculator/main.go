package main

import (
	"fmt"
	"log"
)



func main(){

server:=NewAPIServer(":3000")
fmt.Println("server running...")
if err:=server.Run();err!=nil{
	log.Fatal(err)
}
}

