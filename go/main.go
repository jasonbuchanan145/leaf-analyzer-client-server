package main
//see https://github.com/plgd-dev/go-coap/blob/master/examples/simple/server/main.go the definition for the server is really similar
import (
    "log"
    coap "github.com/plgd-dev/go-coap/v3"
    "github.com/plgd-dev/go-coap/v3/mux"
    controller "github.com/jasonbuchanan145/leaf-analyzer-server/finder"
)

func main() {
    server := mux.NewRouter()
    server.Handle("/image", mux.HandlerFunc(controller.ReceiveImage))
    server.Handle("/hello", mux.HandlerFunc(controller.HelloWorld))
    log.Fatal(coap.ListenAndServe("tcp",":5688",server))
}
