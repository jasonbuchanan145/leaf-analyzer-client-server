import time
import stomp
import modelRunner

class ImageProcessorListener(stomp.ConnectionListener):
    def on_message(self, headers, message):
        print(f'Received file name: {message}')
        modelRunner.process_image(message)
#connect to the activemq dockercontainer on the stomp port
conn = stomp.Connection([('activemq', 61613)])

conn.set_listener('', ImageProcessorListener())
conn.connect('admin', 'admin', wait=True)

conn.subscribe(destination='/queue/insect-finding-queue', id=1, ack='auto')
#Busy loop to keep the connection active and script running
try:
    while True:
        time.sleep(10)
except KeyboardInterrupt:
    print("Exiting")