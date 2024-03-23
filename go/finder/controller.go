package finder
import(
	"log"
    "io"
    "bytes"
	"github.com/plgd-dev/go-coap/v3/message"
	"github.com/plgd-dev/go-coap/v3/message/codes"
	"github.com/plgd-dev/go-coap/v3/mux"

)
// @Summary recieve the image to save and then place on the queue for processing by the pytorch script
// @Accept json
// @Produce json
// @Router /image
func ReceiveImage(w mux.ResponseWriter, r *mux.Message) {
	log.Printf("hit the recieve")
	test:=r.Body()
	if test == nil{
		log.Printf("It's null")
	}
    payload, error := io.ReadAll(r.Body())
	if error!=nil{
        log.Printf("error reading data %v",error)
    }
    ProcessImage(payload)
	w.SetResponse(codes.Changed,message.TextPlain,nil)
}
func HelloWorld(w mux.ResponseWriter, r *mux.Message){
	err:=w.SetResponse(codes.GET, message.TextPlain, bytes.NewReader([]byte("hello world")))
	if err!=nil{
		log.Printf(err.Error())
	}else{
		log.Printf("printed")
	}
}
