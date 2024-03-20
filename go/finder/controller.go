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
    payload, error := io.ReadAll(r.Body())
	if error!=nil{
        log.Printf("error %v",error)
    }
    ProcessImage(payload)
    customResp := w.Conn().AcquireMessage(r.Context())
	defer w.Conn().ReleaseMessage(customResp)
	customResp.SetCode(codes.Content)
	customResp.SetToken(r.Token())
	customResp.SetContentFormat(message.TextPlain)
	customResp.SetBody(bytes.NewReader([]byte("Processed and placed image on the queue")))
	error = w.Conn().WriteMessage(customResp)
	if error != nil {
		log.Printf("cannot set response: %v", error)
	}
}
func HelloWorld(w mux.ResponseWriter, r *mux.Message){
    customResp := w.Conn().AcquireMessage(r.Context())
	defer w.Conn().ReleaseMessage(customResp)
	customResp.SetCode(codes.Content)
	customResp.SetToken(r.Token())
	customResp.SetContentFormat(message.TextPlain)
	customResp.SetBody(bytes.NewReader([]byte("hello world")))
	w.Conn().WriteMessage(customResp)
}
