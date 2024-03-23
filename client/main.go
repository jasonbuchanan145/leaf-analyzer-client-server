package main

import (
	"bytes"
	"context"
	"log"
	"os"
	"time"

	"github.com/plgd-dev/go-coap/v3/message"
	"github.com/plgd-dev/go-coap/v3/message/codes"
	"github.com/plgd-dev/go-coap/v3/tcp"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func main(){
	path := "./payloads/"
	if len(os.Args) > 1 {
		path+=os.Args[1]
	}else{
		path+="small.jpg"
	}
	file,error :=os.ReadFile(path)
	if(file ==nil || len(file)==0){
		log.Printf("Could not open file")
	}
	check(error)
	co, error :=tcp.Dial("localhost:5688")
	check(error)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()
	print("hit")
	resp, error := co.Get(ctx,"/hello")
	log.Printf(resp.String())
	print("ran")
	check(error)
	request,error := co.NewPostRequest(ctx,"/image",message.AppOctets,bytes.NewReader(file))
	
	check(error)
	response,error := co.Do(request)
	check(error)
	if response.Code() != codes.Changed && response.Code() != codes.Created {
		log.Fatalf("Unexpected response code: %v", response.Code())
	}
	log.Print("sent")
}
