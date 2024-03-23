package finder
import (
    "time"
    "os"
    "github.com/go-stomp/stomp/v3"
    "github.com/google/uuid"
    "log"
)

//keep the connection with activemq open
var activeMqConnection *stomp.Conn

func InitActiveMQConnection() {
    options := []func(*stomp.Conn) error{
        stomp.ConnOpt.Login("artemis", "artemis"),
    }
    conn, err := stomp.Dial("tcp", "activemq:61613", options...)
    if err != nil {
        log.Fatalf("Failed to connect to ActiveMQ: %v", err)
    }
    activeMqConnection = conn
}



func ProcessImage(imageData []byte) error {
    // Save the image data to a temporary file
    //outputs the date in the format of yyyymmddhhmmss_UUID.jpg where UUID is just the first 8 characters of the uuid
    pseudo_unique_file  := "image-" + time.Now().Format("20060102150405") +uuid.New().String()[:8]+".jpg"
    log.Printf("Created file name "+pseudo_unique_file)
    filePath := "/app/temp/"+pseudo_unique_file
    // Write the image data to the file in the specified directory
    err := os.WriteFile(filePath, imageData, 0644)
    if err != nil {
        log.Printf(err.Error())
    }
    return notifyActiveMq(filePath) 
}

func notifyActiveMq(filePath string)error{
    if activeMqConnection == nil {
        log.Printf("ActiveMQ connection is not initialized")
    }

    // Use the existing connection to send the message
    err := activeMqConnection.Send("queue.insect-finding-queue", "text/plain", []byte(filePath), nil)
    if err != nil {
        log.Printf(err.Error())
    }
    return nil
}
