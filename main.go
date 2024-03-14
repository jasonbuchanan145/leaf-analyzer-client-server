package main
//see https://github.com/plgd-dev/go-coap/blob/master/examples/simple/server/main.go the definition for the server is really similar
import (
    "log"
    coap "github.com/plgd-dev/go-coap/v3"
    "github.com/plgd-dev/go-coap/v3/mux"
    controller "github.com/jasonbuchanan145/leaf-analyzer-server/finder"
)

func loggingMiddleware(next mux.Handler) mux.Handler {
		return mux.HandlerFunc(func(w mux.ResponseWriter, r *mux.Message) {
					log.Printf("ClientAddress %v, %v\n", w.Conn().RemoteAddr(), r.String())
							next.ServeCOAP(w, r)
					})
}

func main() {
    server := mux.NewRouter()
    server.Use(loggingMiddleware)
    server.Handle("/image", mux.HandlerFunc(controller.ReceiveImage))
    log.Fatal(coap.ListenAndServe("udp",":5688",server))
}
