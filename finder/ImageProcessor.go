package finder
import (
    "io/ioutil"
    "os"
    "github.com/go-stomp/stomp/v3"
)
func (s *Service) ProcessImage(imageData []byte) error {
    // Save the image data to a temporary file
	// ioutil.TempFile writes unique naming for each file 
    tempFile, err := ioutil.TempFile("", "image-*.jpg")
    if err != nil {
        return err
    }
    defer tempFile.Close()
	//write returns the number of bytes written and any errors.
	//since I don't care about the number of bytes use _, to discard it
    _, err = tempFile.Write(imageData)
    if err != nil {
        return err
    }
	return notifyActiveMq(tempFile.Name())
    
}

func(s *Service) notifyActiveMq(filePath string)error{
	conn, err := stomp.Dial("tcp", "localhost:61613")
    if err != nil {
        return err
    }
    defer conn.Disconnect()

    // Send the file path to the queue
    err = conn.Send("/queue/insect-finding-queue", "text/plain", []byte(filePath), nil)
    if err != nil {
        return err
    }

    return nil
}
