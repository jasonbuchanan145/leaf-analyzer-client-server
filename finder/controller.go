package finder
import(
	"log"
	"github.com/plgd-dev/go-coap/v3/message"
)
func (c *Controller) receiveImage(w coap.ResponseWriter, req *coap.Request) {
    imageData := req.Msg.Payload()
	//send image to service
    if err := c.service.ProcessImage(imageData); err != nil {
        log.Printf("Image service excepiton: %v", err)
        respondWithError(w, "Could not process image")
        return
    }
    coapResp := coap.Message{
        Type: coap.Acknowledgement,
        Code: coap.Content,
        Payload: nil,
    }
    coapResp.SetOption(coap.ContentFormat, coap.TextPlain)
}
